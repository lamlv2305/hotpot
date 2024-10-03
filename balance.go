package hotpot

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lamlv2305/hotpot/contract"
	"github.com/pkg/errors"
	"math/big"
)

type BalanceRequest struct {
	Wallet common.Address
	Token  common.Address
}

type BalanceResponse struct {
	Wallet  common.Address
	Token   common.Address
	Balance *big.Int
	Success bool
}

func (h Hotpot) Balance(ctx context.Context, requests []BalanceRequest) ([]BalanceResponse, error) {
	client := h.client(ctx)
	if client == nil {
		return nil, ErrNoClient
	}

	var calls []contract.Multicall3Call3

	for _, req := range requests {
		var calldata []byte
		var err error

		if req.Token == zeroAddress {
			calldata, err = multicallABI.Pack("getEthBalance", req.Wallet)
		} else {
			calldata, err = erc20ABI.Pack("balanceOf", req.Wallet)
		}

		if err != nil {
			return nil, err
		}

		calls = append(calls, contract.Multicall3Call3{
			Target:       req.Token,
			AllowFailure: true,
			CallData:     calldata,
		})
	}

	aggregateCallData, err := multicallABI.Pack("aggregate3", calls)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to encode Multicall aggregate3")
	}

	// Create the call message
	msg := ethereum.CallMsg{
		To:   &h.multicallAddress,
		Data: aggregateCallData,
	}

	// Execute the call
	resp, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to execute Multicall aggregate3")
	}

	var unpackedResult []contract.Multicall3Result
	err = multicallABI.UnpackIntoInterface(&unpackedResult, "aggregate3", resp)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to unpack Multicall aggregate3 result")
	}

	var result []BalanceResponse
	for i, data := range unpackedResult {
		balance := new(big.Int).SetBytes(data.ReturnData)

		result = append(result, BalanceResponse{
			Wallet:  requests[i].Wallet,
			Token:   requests[i].Token,
			Balance: balance,
			Success: data.Success,
		})
	}

	return result, nil
}
