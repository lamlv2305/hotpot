package hotpot

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lamlv2305/hotpot/contract"
	"log"
	"math/big"
	"testing"
)

func TestEncodeData(t *testing.T) {
	tx := MetaTransaction{
		To:        common.HexToAddress("0x105554FF86200a8d133eb783D9E0A92200ED8d72"),
		Value:     big.NewInt(100),
		Data:      []byte{},
		Operation: 0,
	}

	bytes, err := tx.Pack()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Encoded data:", hexutil.Encode(bytes))
}

func TestMultiSendApproveAllowance(t *testing.T) {
	// pk := "5248df84fa5588b2a2eb8ee8a9094f628031cabb0589108c847eae2204c545fd"
	// publicKey := "0x65E2919189990246ab4934B8953cd4bcaBB9fb9D"

	privateKeyBytes, err := hexutil.Decode("5248df84fa5588b2a2eb8ee8a9094f628031cabb0589108c847eae2204c545fd")
	if err != nil {
		t.Fatal(err)
	}

	// Convert bytes to ecdsa.PrivateKey
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		t.Fatal(err)
	}

	abi, err := contract.ERC20MetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	approveBytes, err := abi.Pack("approve", common.HexToAddress("0x105554FF86200a8d133eb783D9E0A92200ED8d72"), big.NewInt(100))
	if err != nil {
		t.Fatal(err)
	}

	signature, err := crypto.Sign(approveBytes, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	if err != nil {
		t.Fatal(err)
	}

	ms := NewMultiSend(client, big.NewInt(421614))
	ms.Send(rootPk(), MetaTransaction{
		To:        common.HexToAddress("0x3520c884c6211fdc12a27fe73696c79d45e11334"),
		Value:     big.NewInt(0),
		Data:      signature,
		Operation: OperationTypeDelegateCall,
	})
}

func rootPk() *ecdsa.PrivateKey {

}
