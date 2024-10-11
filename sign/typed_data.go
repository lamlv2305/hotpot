package sign

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func TypedData(privateKey *ecdsa.PrivateKey, typedData apitypes.TypedData) ([]byte, error) {
	// EIP-712 typed data marshalling
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}

	// add magic string prefix
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	dataHash := crypto.Keccak256(rawData)
	signature, err := crypto.Sign(dataHash, privateKey)
	if err != nil {
		return nil, err
	}

	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}
