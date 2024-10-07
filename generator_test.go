package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateCode(t *testing.T) {
	jsonABI, err := os.ReadFile("./test_abi/usdc-implementation.json")
	require.NoError(t, err, "failed to load test ABI")
	_, err = GenerateCodeForJSON("foo", "Foo", "0x0000000000000000000000000000000000000000", string(jsonABI))
	require.NoError(t, err)
}
