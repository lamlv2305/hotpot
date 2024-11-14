package safe

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/lamlv2305/hotpot/contract"
)

type OperationType uint8

const (
	OperationCall         OperationType = 0
	OperationDelegateCall OperationType = 1
)

var (
	multiSend141         = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
	multiSendCallOnly141 = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
)

type MetaTransaction struct {
	Operation OperationType
	To        common.Address
	Value     *big.Int
	Data      []byte
}

func (m MetaTransaction) Encode() ([]byte, error) {
	input := [][]byte{
		{byte(m.Operation)},
		m.To.Bytes(),
		math.U256Bytes(m.Value),
		math.U256Bytes(big.NewInt(int64(len(m.Data)))),
		m.Data,
	}

	return bytes.Join(input, nil), nil
}

func encodeMulti(txs []MetaTransaction, to common.Address) (*MetaTransaction, error) {
	var transactionsEncoded []byte

	for idx := range txs {
		encodedData, err := txs[idx].Encode()
		if err != nil {
			return nil, err
		}

		transactionsEncoded = append(transactionsEncoded, encodedData...)
	}

	multisendAbi, _ := contract.MultiSendMetaData.GetAbi()
	data, err := multisendAbi.Pack("multiSend", transactionsEncoded)
	if err != nil {
		return nil, err
	}

	mtx := MetaTransaction{
		To:    to,
		Value: big.NewInt(0),
		Data:  data,
	}

	return &mtx, nil
}
