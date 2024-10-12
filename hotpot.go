package hotpot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lamlv2305/hotpot/contract"
	"math/rand"
	"time"
)

var (
	source      = rand.NewSource(time.Now().UnixNano())
	rng         = rand.New(source)
	zeroAddress = common.Address{}

	multicallABI, _ = contract.MultiCallMetaData.GetAbi()
	erc20ABI, _     = contract.ERC20MetaData.GetAbi()
)

func WithMultiCallAddress(address common.Address) func(*Hotpot) {
	return func(h *Hotpot) {
		h.multicallAddress = address
	}
}

func NewHotpot(rpc []string, options ...func(*Hotpot)) *Hotpot {
	hp := &Hotpot{
		// https://www.multicall3.com/
		multicallAddress: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
	}

	for idx := range rpc {
		client, err := ethclient.DialContext(context.Background(), rpc[idx])
		if err != nil {
			continue
		}

		hp.clients = append(hp.clients, client)
	}

	for idx := range options {
		options[idx](hp)
	}

	return hp
}

type Hotpot struct {
	clients          []*ethclient.Client
	multicallAddress common.Address
}

func (h Hotpot) randomClient() *ethclient.Client {
	if len(h.clients) == 0 {
		return nil
	}

	return h.clients[rng.Intn(len(h.clients))]
}
