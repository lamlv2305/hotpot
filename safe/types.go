package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	zero = common.Address{}
)

type TransferParams struct {
	From   common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
}
