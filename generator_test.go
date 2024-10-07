package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateCode(t *testing.T) {
	jsonABI, err := os.ReadFile("./test_abi/usdc-implementation.json")
	require.NoError(t, err, "failed to load test ABI")
	require.NoError(t, GenerateCodeForJSON("foo", "Foo", "0x0000000000000000000000000000000000000000", string(jsonABI), io.Discard))
}
