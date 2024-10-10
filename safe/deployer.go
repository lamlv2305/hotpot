package safe

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lamlv2305/hotpot/contract"
	"math/big"
	"time"
)

func WithProxyFactory(proxyFactory common.Address) func(*Deployer) {
	return func(wallet *Deployer) {
		wallet.proxyFactory = proxyFactory
	}
}

func NewDeployer(client *ethclient.Client, pk *ecdsa.PrivateKey, options ...func(*Deployer)) (*Deployer, error) {
	publicKeyECDSA, ok := pk.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	chainId, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}

	w := &Deployer{
		client:       client,
		pk:           pk,
		pub:          crypto.PubkeyToAddress(*publicKeyECDSA),
		proxyFactory: common.HexToAddress("0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2"),
		chainId:      chainId,
	}

	for idx := range options {
		options[idx](w)
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
		delay:   time.Second * 1,
		timeout: time.Second * 30,
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

	timeoutCtx, cancelTimeoutCtx := context.WithTimeout(ctx, opt.timeout)
	defer cancelTimeoutCtx()

	for {
		if err := timeoutCtx.Err(); err != nil {
			return tx, zero, err
		}

		receipt, err := w.client.TransactionReceipt(timeoutCtx, tx.Hash())
		if err != nil {
			time.Sleep(time.Second * 1)
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
	Spender common.Address
	Amount  *big.Int
	To      common.Address

	ProxyWallet common.Address
}

func (w *Deployer) Approve(ctx context.Context, opts ApproveOptions) (*types.Transaction, error) {
	safeContract, err := contract.NewGnosisSafe(opts.ProxyWallet, w.client)
	if err != nil {
		return nil, err
	}

	//proxyNonce, err := safeContract.Nonce(&bind.CallOpts{Pending: true})
	//if err != nil {
	//	return nil, err
	//}

	tokenAbi, _ := contract.ERC20MetaData.GetAbi()
	approveData, err := tokenAbi.Pack("approve", opts.Spender, opts.Amount)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256Hash(approveData)
	signature, err := signHash(hash.Bytes(), w.pk)
	if err != nil {
		return nil, err
	}

	//fmt.Println("Signature", hexutil.Encode(signature))

	multiSendTx, err := encodeMulti(
		[]MetaTransaction{
			{
				Operation: OperationCall,
				// USDC
				To:    usdc,
				Value: big.NewInt(0),
				Data:  approveData,
			},
		},
		opts.To)
	if err != nil {
		return nil, err
	}

	safeTx := Transaction{
		To:             opts.To,
		Value:          big.NewInt(0),
		Data:           multiSendTx.Data,
		Operation:      OperationDelegateCall,
		SafeTxGas:      big.NewInt(0),
		BaseGas:        big.NewInt(0),
		GasPrice:       big.NewInt(0),
		GasToken:       zero,
		RefundReceiver: zero,
	}

	//rawSafeTx, err := safeContract.EncodeTransactionData(
	//	&bind.CallOpts{Context: ctx, Pending: true},
	//	safeTx.To,
	//	safeTx.Value,
	//	safeTx.Data,
	//	uint8(safeTx.Operation),
	//	safeTx.SafeTxGas,
	//	safeTx.BaseGas,
	//	safeTx.GasPrice,
	//	safeTx.GasToken,
	//	safeTx.RefundReceiver,
	//	proxyNonce,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//safeTxHashBytes := crypto.Keccak256(rawSafeTx)
	//safeTxSignature, err := signHash(safeTxHashBytes, w.pk)
	//if err != nil {
	//	return nil, err
	//}
	//
	//fmt.Sprintln("safe tx sig", safeTxSignature)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(w.pk, w.chainId)
	if err != nil {
		return nil, err
	}

	// to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int,
	// baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte
	tx, err := safeContract.ExecTransaction(
		transactOpts,
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		uint8(safeTx.Operation),
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		signature,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
