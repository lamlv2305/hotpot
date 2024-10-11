package safe

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/lamlv2305/hotpot/contract"
	"github.com/lamlv2305/hotpot/sign"
	"math/big"
	"time"
)

func WithProxyFactory(proxyFactory common.Address) func(*Deployer) {
	return func(deployer *Deployer) {
		deployer.proxyFactory = proxyFactory
	}
}

func WithChainId(chainId *big.Int) func(*Deployer) {
	return func(deployer *Deployer) {
		if chainId != nil && chainId.Cmp(big.NewInt(0)) != 0 {
			deployer.chainId = chainId
		}
	}
}

func NewDeployer(client *ethclient.Client, pk *ecdsa.PrivateKey, options ...func(*Deployer)) (*Deployer, error) {
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	w := &Deployer{
		client: client,
		pk:     pk,
		pub:    crypto.PubkeyToAddress(*publicKeyECDSA),

		// https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#factory-contracts-3
		proxyFactory: common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
	}

	for idx := range options {
		options[idx](w)
	}

	if w.chainId == nil {
		chainId, err := client.ChainID(context.TODO())
		if err != nil {
			return nil, err
		}

		w.chainId = chainId
	}

	return w, nil
}

type Deployer struct {
	client *ethclient.Client

	// the wallet will be paid for the gas
	pk  *ecdsa.PrivateKey
	pub common.Address

	proxyFactory common.Address
	chainId      *big.Int
}

func (w *Deployer) CreateProxy(ctx context.Context, index *big.Int, options ...func(*RequestOptions)) (*types.Transaction, common.Address, error) {
	opt := RequestOptions{
		waitDelay:   time.Second * 1,
		waitTimeout: time.Second * 30,
	}

	for idx := range options {
		options[idx](&opt)
	}

	deployer, err := bind.NewKeyedTransactorWithChainID(w.pk, w.chainId)
	if err != nil {
		return nil, zero, err
	}
	deployer.Context = ctx

	proxyFactory, err := contract.NewProxyFactory(w.proxyFactory, w.client)
	if err != nil {
		return nil, zero, err
	}

	var salt [32]byte
	copy(salt[:], index.Bytes())

	tx, err := proxyFactory.CreateProxy(deployer, zero, big.NewInt(0), zero, salt)
	if err != nil {
		return nil, zero, err
	}

	timeoutCtx, cancelTimeoutCtx := context.WithTimeout(ctx, opt.waitTimeout)
	defer cancelTimeoutCtx()

	for {
		if err := timeoutCtx.Err(); err != nil {
			return tx, zero, err
		}

		receipt, err := w.client.TransactionReceipt(timeoutCtx, tx.Hash())
		if err != nil {
			time.Sleep(opt.waitDelay)
			continue
		}

		if receipt.Status == 0 {
			return tx, zero, fmt.Errorf("submit failed with tx: %s", receipt.TxHash.String())
		}

		// Fetch address
		address, err := w.ComputeProxyAddress(ctx, index)
		if err != nil {
			return tx, zero, err
		}

		return tx, address, nil
	}
}

func (w *Deployer) ComputeProxyAddress(ctx context.Context, index *big.Int) (common.Address, error) {
	proxyFactory, err := contract.NewProxyFactory(w.proxyFactory, w.client)
	if err != nil {
		return zero, err
	}

	var salt [32]byte
	copy(salt[:], index.Bytes())

	return proxyFactory.ComputeProxyAddress(&bind.CallOpts{Context: ctx}, w.pub, salt)
}

type ApproveOptions struct {
	Spender     common.Address
	Amount      *big.Int
	ProxyWallet common.Address
}

func (w *Deployer) Approve(ctx context.Context, opts ApproveOptions) (*types.Transaction, error) {
	tokenAbi, _ := contract.ERC20MetaData.GetAbi()
	approveData, err := tokenAbi.Pack("approve", opts.Spender, opts.Amount)
	if err != nil {
		return nil, err
	}

	return w.Exec(ctx, opts.ProxyWallet, []MetaTransaction{
		{
			To:    usdc,
			Value: big.NewInt(0),
			Data:  approveData,
		},
	})
}

func (w *Deployer) Exec(ctx context.Context, proxyWallet common.Address, txs []MetaTransaction) (*types.Transaction, error) {
	destination := multiSendCallOnly141

	for idx := range txs {
		if txs[idx].Operation == OperationCall {
			destination = multiSend141
			break
		}
	}

	multiSendTx, err := encodeMulti(txs, destination)
	if err != nil {
		return nil, err
	}

	safeContract, err := contract.NewGnosisSafe(proxyWallet, w.client)
	if err != nil {
		return nil, err
	}

	proxyWalletNonce, err := safeContract.Nonce(&bind.CallOpts{
		Pending: true,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	signature, err := sign.TypedData(w.pk, apitypes.TypedData{
		Types: map[string][]apitypes.Type{
			"SafeTx": {
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "operation", Type: "uint8"},
				{Name: "safeTxGas", Type: "uint256"},
				{Name: "baseGas", Type: "uint256"},
				{Name: "gasPrice", Type: "uint256"},
				{Name: "gasToken", Type: "address"},
				{Name: "refundReceiver", Type: "address"},
				{Name: "nonce", Type: "uint256"},
			},
			"EIP712Domain": {
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "SafeTx",
		Domain: apitypes.TypedDataDomain{
			ChainId:           math.NewHexOrDecimal256(w.chainId.Int64()),
			VerifyingContract: proxyWallet.String(),
		},
		Message: map[string]any{
			"to":             destination.String(),
			"value":          big.NewInt(0),
			"operation":      big.NewInt(int64(OperationDelegateCall)), // delegate call
			"data":           multiSendTx.Data,                         // multisend data
			"safeTxGas":      big.NewInt(0),
			"baseGas":        big.NewInt(0),
			"gasPrice":       big.NewInt(0),
			"gasToken":       zero.String(),
			"refundReceiver": zero.String(),
			"nonce":          proxyWalletNonce,
		},
	})
	if err != nil {
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(w.pk, w.chainId)
	if err != nil {
		return nil, err
	}
	transactOpts.Context = ctx

	tx, err := safeContract.ExecTransaction(
		transactOpts,
		destination,
		big.NewInt(0),
		multiSendTx.Data,
		uint8(OperationDelegateCall),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		zero,
		zero,
		signature,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
