package main

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestEtherscan_GetContractABI(t *testing.T) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		t.Skip("env: ETHERSCAN_API_KEY is not set")
	}
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	client := NewEtherscan(apiKey)

	actual, err := client.GetContractABI(context.Background(), usdc)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	expected, _ := os.ReadFile("./test_abi/usdc-proxy.json")
	require.Equal(t, strings.TrimSpace(string(expected)), strings.TrimSpace(string(actual)))
}
