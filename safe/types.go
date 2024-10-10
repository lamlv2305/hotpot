package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	zero = common.Address{}
)

type Transaction struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      OperationType
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signature      []byte
}
