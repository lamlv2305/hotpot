package hotpot

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lamlv2305/hotpot/contract"
	"math/big"
)

type OperationType int

const (
	OperationTypeCall         OperationType = 0
	OperationTypeDelegateCall OperationType = 1

	metaTransactionTypes = `[
		{"name": "operation", "type": "uint8"},
		{"name": "to", "type": "address"},
		{"name": "value", "type": "uint256"},
		{"name": "dataLength", "type": "uint256"},
		{"name": "data", "type": "bytes"}
	]`
)

var (
	uint8Type, _   = abi.NewType("uint8", "", nil)
	addressType, _ = abi.NewType("address", "", nil)
	uint256Type, _ = abi.NewType("uint256", "", nil)
	bytesType, _   = abi.NewType("bytes", "", nil)

	multiSend141 = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
)

type MetaTransaction struct {
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	Data      []byte         `json:"data"`
	Operation OperationType  `json:"operation"`
}

func (m MetaTransaction) Pack() ([]byte, error) {
	// Define the ABI arguments
	arguments := abi.Arguments{
		abi.Argument{Type: uint8Type},   // uint8 (operation)
		abi.Argument{Type: addressType}, // address (to)
		abi.Argument{Type: uint256Type}, // uint256 (value)
		abi.Argument{Type: uint256Type}, // uint256 (data length)
		abi.Argument{Type: bytesType},   // bytes (data)
	}

	operation := uint8(m.Operation)
	dataLength := big.NewInt(int64(len(m.Data)))

	// Pack the arguments into bytes
	encodedData, err := arguments.Pack(operation, m.To, m.Value, dataLength, m.Data)
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}

func NewMultiSend(client *ethclient.Client, chainId *big.Int) *MultiSend {
	return &MultiSend{
		client:  client,
		chainId: chainId,
	}
}

type MultiSend struct {
	client  *ethclient.Client
	chainId *big.Int
}

func (m MultiSend) Send(pk *ecdsa.PrivateKey, tx MetaTransaction) (*types.Transaction, error) {
	// Encode the data
	data, err := tx.Pack()
	if err != nil {
		return nil, err
	}

	ms, err := contract.NewMultiSend(multiSend141, m.client)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(pk, m.chainId)
	if err != nil {
		return nil, err
	}

	return ms.MultiSend(transactOpts, data)
}
