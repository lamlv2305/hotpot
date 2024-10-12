package hotpot

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lamlv2305/hotpot/contract"
	"github.com/pkg/errors"
)

type QueryRequest struct {
	To   common.Address
	Data []byte
}

type QueryResponse struct {
	Result  []byte
	Success bool
}

func (h Hotpot) Query(ctx context.Context, requests []QueryRequest) ([]QueryResponse, error) {
	client := h.randomClient()

	var calls []contract.Multicall3Call3

	for _, req := range requests {
		calls = append(calls, contract.Multicall3Call3{
			Target:       req.To,
			AllowFailure: true,
			CallData:     req.Data,
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

	var result []QueryResponse
	for _, data := range unpackedResult {
		result = append(result, QueryResponse{
			Result:  data.ReturnData,
			Success: data.Success,
		})
	}

	return result, nil
}
