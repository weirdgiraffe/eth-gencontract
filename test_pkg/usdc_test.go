package foo

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestUSDC(t *testing.T) {
	rpc := RPCClient(t)
	// USDC contract is tricky because it is the proxy contract, so should be instantiated
	// using the proxy address and not the implementation address.
	usdc := NewUSDCForAddress(rpc, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))

	name, err := usdc.Name(context.Background())
	require.NoError(t, err)
	fmt.Println("Name:", name)
}

func RPCClient(t *testing.T) *ethclient.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("env: RPC_URL is not set")
	}

	c, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)
	return c
}

func init() {
	envFile := "../.env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Fprintf(os.Stderr, "[WARN] failed to load %s: %v\n", envFile, err)
	}
}
