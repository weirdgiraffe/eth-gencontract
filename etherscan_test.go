package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEtherscan_GetContractABI(t *testing.T) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		t.Skip("env: ETHERSCAN_API_KEY is not set")
	}
	usdc := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	client := NewEtherscan(apiKey)

	abiJSON, err := client.GetContractABI(context.Background(), usdc)
	require.NoError(t, err)
	require.NotEmpty(t, abiJSON)
}
