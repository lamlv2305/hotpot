package safe

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func signHash(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	if sig[64] < 2 {
		sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	return sig, nil
}
