package dm1

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"testing"
)

var (
	client, _    = ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	pk, _        = crypto.HexToECDSA("")
	proxyFactory = common.HexToAddress("0x615e66a51AAC47247f37E0E94DAEA9779a72dE88")
	//deployer, _  = NewDeployer(client, pk, WithProxyFactory(proxyFactory))

	multiSendAddr = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
	//usdc          = common.HexToAddress("0x3520C884C6211FDC12A27fe73696c79d45E11334")
	exchange = common.HexToAddress("0x105554FF86200a8d133eb783D9E0A92200ED8d72")
	zero     = common.Address{}
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
			VerifyingContract: "0x500d83C4c1Bbdbd14528F68e81791Da95d3Cae01",
		},
		Message: map[string]any{
			"to":             multiSendAddr.String(),
			"value":          big.NewInt(0),
			"data":           []byte{},      // multisend data
			"operation":      big.NewInt(1), // delegate call
			"safeTxGas":      big.NewInt(0),
			"baseGas":        big.NewInt(0),
			"gasPrice":       big.NewInt(0),
			"gasToken":       zero.String(),
			"refundReceiver": zero.String(),
			"nonce":          big.NewInt(0),
		},
	}

	// EIP-712 typed data marshalling
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		t.Fatal(err)
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		t.Fatal(err)
	}

	// add magic string prefix
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	dataHash := crypto.Keccak256(rawData)
	fmt.Println("SIG HASH:", hexutil.Encode(dataHash))

	signature, err := crypto.Sign(dataHash, pk)
	if err != nil {
		t.Fatal(err)
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	t.Log(hexutil.Encode(signature))
}
