package safe

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lamlv2305/hotpot/contract"
	"math/big"
)

type OperationType uint8

const (
	OperationCall         OperationType = 0
	OperationDelegateCall OperationType = 1
)

var (
	uint8Type, _   = abi.NewType("uint8", "", nil)
	addressType, _ = abi.NewType("address", "", nil)
	uint256Type, _ = abi.NewType("uint256", "", nil)
	bytesType, _   = abi.NewType("bytes", "", nil)

	usdc              = common.HexToAddress("0x3520C884C6211FDC12A27fe73696c79d45E11334")
	multiSend141      = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
	multiSendCallOnly = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
)

type MetaTransaction struct {
	Operation OperationType
	To        common.Address
	Value     *big.Int
	Data      []byte
}

func (m MetaTransaction) Encode() ([]byte, error) {
	arguments := abi.Arguments{
		abi.Argument{Type: uint8Type},   // uint8 (operation)
		abi.Argument{Type: addressType}, // address (to)
		abi.Argument{Type: uint256Type}, // uint256 (value)
		abi.Argument{Type: uint256Type}, // uint256 (data length)
		abi.Argument{Type: bytesType},   // bytes (data)
	}

	operation := uint8(m.Operation)
	dataLength := big.NewInt(int64(len(m.Data)))

	encodedData, err := arguments.Pack(operation, m.To, m.Value, dataLength, m.Data)
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}

func encodeMulti(txs []MetaTransaction, to common.Address) (*MetaTransaction, error) {
	var data []byte

	for idx := range txs {
		encodedData, err := txs[idx].Encode()
		if err != nil {
			return nil, err
		}

		data = append(data, encodedData...)
	}

	multisendAbi, _ := contract.MultiSendMetaData.GetAbi()
	encodedBytes, err := multisendAbi.Pack("multiSend", data)
	if err != nil {
		return nil, err
	}

	mtx := MetaTransaction{
		Operation: OperationDelegateCall,
		To:        to,
		Value:     big.NewInt(0),
		Data:      encodedBytes,
	}

	return &mtx, nil
}
