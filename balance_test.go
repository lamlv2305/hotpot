package hotpot

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lamlv2305/hotpot/contract"
	"testing"
)

func TestHotpot_Balance(t *testing.T) {
	balances, err := h.Balance(
		context.Background(),
		[]BalanceRequest{
			{
				Wallet: common.HexToAddress("0x31392A5b12445F5792250fa857Bb34a114d224d1"),
				Token:  common.HexToAddress("0x9D0685F859782B4bC3f0E4403bEeF11eDA8AC2E8"),
			},
			{
				Wallet: common.HexToAddress("0x31392A5b12445F5792250fa857Bb34a114d224d1"),
			},
		})
	if err != nil {
		t.Fatalf("Failed to get balances: %v", err)
	}

	for _, b := range balances {
		t.Logf("Token %s balance: %s", b.Token, b.Balance.String())
	}
}

func TestNativeBalance(t *testing.T) {
	addr := common.HexToAddress("0x31392A5b12445F5792250fa857Bb34a114d224d1")

	client, _ := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
	b, err := client.PendingBalanceAt(context.TODO(), addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("B1", b.String())

	m, err := contract.NewMultiCall(addr, client)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := m.GetEthBalance(&bind.CallOpts{}, addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("B2", b2.String())
	if b.String() != b2.String() {
		t.Fatal("balance not match")
	}

	t.Log("Done")
}
