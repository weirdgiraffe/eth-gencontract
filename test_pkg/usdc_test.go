package foo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestUSDC(t *testing.T) {
	rpc := RPCClient(t)
	usdc := NewUSDC(rpc)

	name, err := usdc.Name(context.Background())
	require.NoError(t, err)
	fmt.Println("Name:", name)
}

func RPCClient(t *testing.T) *ethclient.Client {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		t.Skip("env: RPC_URL is not set")
	}

	c, err := ethclient.Dial(rpcURL)
	require.NoError(t, err)
	return c
}

func init() {
	envFile := "../.env"
	if err := godotenv.Load(envFile); err != nil {
		fmt.Fprintf(os.Stderr, "[WARN] failed to load %s: %v\n", envFile, err)
	}
}
