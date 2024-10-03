package hotpot

import (
	"os"
	"testing"
)

var (
	h = NewHotpot([]string{"https://sepolia-rollup.arbitrum.io/rpc"})
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
