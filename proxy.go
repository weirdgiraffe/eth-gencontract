package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var addressSlotProxies = []common.Hash{
	common.HexToHash("0x863df6bfa4469f3ead0be8f9f2aae51c91a907b4"),
}

type ProxyResolver struct {
	rpc       *ethclient.Client
	etherscan *Etherscan
}

func NewProxyResolver(rpc *ethclient.Client, etherscan *Etherscan) *ProxyResolver {
	return &ProxyResolver{
		etherscan: etherscan,
		rpc:       rpc,
	}
}

func (r *ProxyResolver) GetContractABI(ctx context.Context, address common.Address) ([]byte, error) {
	implAddress, err := r.findImplementationAddress(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("failed to get proxy implementation: %w", err)
	}
	if implAddress != zeroAddress {
		return r.etherscan.GetContractABI(ctx, implAddress)
	}
	return r.etherscan.GetContractABI(ctx, address)
}

var (
	zeroAddress = common.Address{}

	// references:
	//   - https://eips.ethereum.org/EIPS/eip-1967
	//   - https://github.com/circlefin/stablecoin-evm/blob/master/contracts/upgradeability/UpgradeabilityProxy.sol
	knownProxies = []common.Hash{
		// bytes32(keccak256("org.zeppelinos.proxy.implementation"));
		common.HexToHash("0x7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3"),
		// bytes32(uint256(keccak256('eip1967.proxy.implementation')) - 1)
		common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"),
		// bytes32(uint256(keccak256('eip1967.proxy.beacon')) - 1)
		common.HexToHash("0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50"),
	}
)

func (r *ProxyResolver) findImplementationAddress(ctx context.Context, address common.Address) (common.Address, error) {
	for _, slot := range knownProxies {
		v, err := r.rpc.StorageAt(ctx, address, slot, nil)
		if err != nil {
			return zeroAddress, fmt.Errorf("failed to get openzeppelin logicAddress slot: %w", err)
		}
		implAddress := common.BytesToAddress(v[len(v)-20:])
		if implAddress != zeroAddress {
			return implAddress, nil
		}
	}
	return zeroAddress, nil
}
