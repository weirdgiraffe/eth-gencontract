# eth-gencontract

Code generator which will generate code for the provided ethereum contract. If
supplied address is a proxy, it will generate contract using the implementation
ABI.

# Usage

command line invocation

```bash
eth-gencontract -pkg=example -name=USDC -addr=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48 -o=usdc.go
```

or go-generate comment

```go
//go:generate eth-gencontract -pkg=example -name=USDC -addr=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48 -o=usdc.go
```

Will generate go structure will all methods for USDC contract like (full code
could be found in [test_pkg/usdc.go](./test_pkg/usdc.go)):

```go
type USDC struct {
	abi     abi.ABI
	address common.Address
	caller  ethereum.ContractCaller
}

func NewUSDC(caller ethereum.ContractCaller) *USDC {
	return NewUSDCForAddress(caller, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))
}


func (c *USDC) Allowance(ctx context.Context, owner common.Address, spender common.Address) (*big.Int, error) {
    // implementation 
}

func (c *USDC) Approve(ctx context.Context, spender common.Address, value *big.Int) (bool, error) {
    // implementation 
}

func (c *USDC) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {
    // implementation 
}

func (c *USDC) Transfer(ctx context.Context, to common.Address, value *big.Int) (bool, error) {
    // implementation 
}

func (c *USDC) TransferFrom(ctx context.Context, from common.Address, to common.Address, value *big.Int) (bool, error) {
    // implementation 
}
// etc.
```

