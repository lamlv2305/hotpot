package safe

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

var (
	client, _    = ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	pk, _        = crypto.HexToECDSA("")
	proxyFactory = common.HexToAddress("0x615e66a51AAC47247f37E0E94DAEA9779a72dE88")
	deployer, _  = NewDeployer(client, pk, WithProxyFactory(proxyFactory))

	multiSendAddr = common.HexToAddress("0x38869bf66a61cF6bDB996A6aE40D5853Fd43B526")
	//usdc          = common.HexToAddress("0x3520C884C6211FDC12A27fe73696c79d45E11334")
	exchange = common.HexToAddress("0x105554FF86200a8d133eb783D9E0A92200ED8d72")
)

func TestCreateProxy(t *testing.T) {
	tx, addr, err := deployer.CreateProxy(context.TODO(), big.NewInt(1))
	if err != nil {
		t.Fatal(err)
	}

	if addr == zero {
		t.Fatal("address is zero")
	}

	t.Log("tx:", tx.Hash().Hex())
	t.Log("addr:", addr.Hex())
}

func TestApproveAllowance(t *testing.T) {
	//salt: big.NewInt(1)
	//deployer_test.go:29: tx: 0x92596dff8fe3d45e11f78765d7f2ebfd79256fd7b95a290b577c471660e80f35
	//deployer_test.go:30: addr: 0x500d83C4c1Bbdbd14528F68e81791Da95d3Cae01

	tx, err := deployer.Approve(context.TODO(), ApproveOptions{
		Spender:     exchange,
		Amount:      abi.MaxUint256,
		To:          multiSendAddr,
		ProxyWallet: common.HexToAddress("0x500d83C4c1Bbdbd14528F68e81791Da95d3Cae01"),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx.Hash().String())
}
