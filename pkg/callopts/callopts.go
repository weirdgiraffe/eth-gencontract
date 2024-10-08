package callopts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type Opt func(cfg *CallConfig)

func From(addr common.Address) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.From = addr
	}
}

func To(addr common.Address) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.To = &addr
	}
}

func Data(data []byte) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.Data = data
	}
}

func Gas(n uint64) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.Gas = n
	}
}

func GasPrice(wei uint64) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.GasPrice = new(big.Int).SetUint64(wei)
	}
}

func GasFeeCap(wei uint64) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.GasFeeCap = new(big.Int).SetUint64(wei)
	}
}

func GasTipCap(wei uint64) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.GasTipCap = new(big.Int).SetUint64(wei)
	}
}

func Value(wei *big.Int) Opt {
	return func(cfg *CallConfig) {
		cfg.CallMsg.Value = wei
	}
}

func AtBlock(blockNumber uint64) Opt {
	return func(cfg *CallConfig) {
		cfg.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
}

type ContractCaller struct {
	impl ethereum.ContractCaller
}

func WrapCaller(impl ethereum.ContractCaller) *ContractCaller {
	return &ContractCaller{impl: impl}
}

type CallConfig struct {
	CallMsg     ethereum.CallMsg
	BlockNumber *big.Int
}

func (c *ContractCaller) Call(ctx context.Context, opts ...Opt) ([]byte, error) {
	var cfg CallConfig
	for _, opt := range opts {
		opt(&cfg)
	}
	return c.impl.CallContract(ctx, cfg.CallMsg, cfg.BlockNumber)
}
