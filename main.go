package main

import (
	"context"
	"flag"
	"fmt"
	"os"
)

var (
	pkg     = flag.String("pkg", "", "package name")
	name    = flag.String("name", "", "go contract name")
	address = flag.String("addr", "", "contract address")
	output  = flag.String("o", "", "where to output")
)

func main() {
	flag.Parse()

	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "ETHERSCAN_API_KEY is required")
		os.Exit(1)
	}
	etherscan := NewEtherscan(apiKey)

	if *pkg == "" {
		fmt.Fprintf(os.Stderr, "pkg is required")
		os.Exit(1)
	}

	if *address == "" {
		fmt.Fprintf(os.Stderr, "address is required")
		os.Exit(1)
	}

	if *name == "" {
		fmt.Fprintf(os.Stderr, "name is required")
		os.Exit(1)
	}

	out := os.Stdout
	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create output file: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		out = f
	}

	jsonABI, err := etherscan.GetContractABI(context.Background(), *address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to init etherscan: %v", err)
		os.Exit(1)
	}

	err = GenerateCodeForJSON(*pkg, *name, *address, string(jsonABI), out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate code for contract: %v", err)
		os.Exit(1)
	}
}
