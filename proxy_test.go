package main

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestProxyResolver(t *testing.T) {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		t.Skip("env: ETHERSCAN_API_KEY is not set")
	}
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("env: RPC_URL is not set")
	}
	rpc, err := ethclient.Dial(rpcURL)
	require.NoError(t, err, "failed to init rpc client")

	resolver := NewProxyResolver(rpc, NewEtherscan(apiKey))
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	actual, err := resolver.GetContractABI(context.Background(), usdc)
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	expected, _ := os.ReadFile("./test_abi/usdc-implementation.json")
	require.Equal(t, strings.TrimSpace(string(expected)), strings.TrimSpace(string(actual)))
}
