package sign

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"testing"
)

var (
	pk, _         = crypto.HexToECDSA("")
	multiSendAddr = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
	zero          = common.Address{}
)

func TestSignTypedData(t *testing.T) {
	typedData := apitypes.TypedData{
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
			ChainId:           math.NewHexOrDecimal256(1),
			VerifyingContract: "0x000000.....",
		},
		Message: map[string]any{
			"to":             multiSendAddr.String(),
			"value":          big.NewInt(0),
			"data":           []byte{},
			"operation":      big.NewInt(1),
			"safeTxGas":      big.NewInt(0),
			"baseGas":        big.NewInt(0),
			"gasPrice":       big.NewInt(0),
			"gasToken":       zero.String(),
			"refundReceiver": zero.String(),
			"nonce":          big.NewInt(0),
		},
	}

	sig, err := TypedData(pk, typedData)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hexutil.Encode(sig))
}
