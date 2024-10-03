package hotpot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestHotpot_Balance(t *testing.T) {
	balances, err := h.Balance(
		context.Background(),
		[]BalanceRequest{
			//{
			//	Wallet: common.HexToAddress("0x31392A5b12445F5792250fa857Bb34a114d224d1"),
			//	Token:  common.HexToAddress("0x9D0685F859782B4bC3f0E4403bEeF11eDA8AC2E8"),
			//},
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
