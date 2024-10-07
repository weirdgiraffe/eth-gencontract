# eth-gencontract

Code generator which will generate code for the provided ethereum contract. If
supplied address is a proxy, it will generate contract using the implementation
ABI.

Example (see `test_pkg`):

The following code

```go
package example
//go:generate eth-gencontract -pkg=example -name=USDC -addr=0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48 -o=usdc.go
```

Will generate go structure will all methods for USDC contract

