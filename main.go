package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/tools/imports"
)

var (
	version = "dev"
	pkg     = flag.String("pkg", "", "package name")
	name    = flag.String("name", "", "go contract name")
	address = flag.String("addr", "", "contract address")
	output  = flag.String("o", "", "where to output")
)

func main() {
	flag.Parse()

	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "ETHERSCAN_API_KEY is required\n")
		os.Exit(1)
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		fmt.Fprintf(os.Stderr, "RPC_URL is required\n")
		os.Exit(1)
	}

	if *address == "" {
		fmt.Fprintf(os.Stderr, "address is required\n")
		os.Exit(1)
	}

	if !common.IsHexAddress(*address) {
		fmt.Fprintf(os.Stderr, "invalid address: %s\n", *address)
		os.Exit(1)
	}

	if *pkg == "" {
		fmt.Fprintf(os.Stderr, "pkg is required\n")
		os.Exit(1)
	}

	if *name == "" {
		fmt.Fprintf(os.Stderr, "name is required\n")
		os.Exit(1)
	}

	rpc, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init rpc client: %v\n", err)
		os.Exit(1)
	}
	etherscan := NewEtherscan(apiKey)
	resolver := NewProxyResolver(rpc, etherscan)

	jsonABI, err := resolver.GetContractABI(context.Background(), common.HexToAddress(*address))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init etherscan: %v\n", err)
		os.Exit(1)
	}

	content, err := GenerateCodeForJSON(*pkg, *name, *address, string(jsonABI))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate code for contract address=%s: %v\n", *address, err)
		os.Exit(1)
	}

	cleaned, err := imports.Process("", content, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run go imports for generated contract code for address=%s: %v\n", *address, err)
		os.Exit(1)
	}

	if *output == "" {
		*output = *name + ".go"
	}
	err = os.WriteFile(*output, cleaned, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output file: %v\n", err)
		os.Exit(1)
	}
}
