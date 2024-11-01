//go:build tools

package hotpot

import (
	_ "github.com/ethereum/go-ethereum/console/prompt"
	_ "github.com/ethereum/go-ethereum/graphql"

	// _ "github.com/ethereum/go-ethereum/internal/era"
	_ "github.com/ethereum/go-ethereum/metrics/influxdb"
)
