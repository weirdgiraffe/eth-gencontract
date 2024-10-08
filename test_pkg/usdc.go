// This file was automatically generated using
//
// eth-gencontract (version: dev)
//
// repository: https://github.com/weirdgiraffe/eth-gencontract
//
// PLEASE DO NOT EDIT MANUALLY
package foo

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type USDC struct {
	caller  ethereum.ContractCaller
	Address common.Address
}

func NewUSDC(caller ethereum.ContractCaller) *USDC {
	return NewUSDCForAddress(caller, common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"))
}

func NewUSDCForAddress(caller ethereum.ContractCaller, address common.Address) *USDC {
	return &USDC{
		caller:  caller,
		Address: address,
	}
}

func (c *USDC) CancelAuthorizationTypehash(ctx context.Context, blockNumber *big.Int) ([32]uint8, error) {
	var (
		out0 [32]uint8
	)

	input, err := UsdcABI.Pack("CANCEL_AUTHORIZATION_TYPEHASH")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("CANCEL_AUTHORIZATION_TYPEHASH", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].([32]uint8)

	return out0, nil
}

func (c *USDC) DomainSeparator(ctx context.Context, blockNumber *big.Int) ([32]uint8, error) {
	var (
		out0 [32]uint8
	)

	input, err := UsdcABI.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("DOMAIN_SEPARATOR", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].([32]uint8)

	return out0, nil
}

func (c *USDC) PermitTypehash(ctx context.Context, blockNumber *big.Int) ([32]uint8, error) {
	var (
		out0 [32]uint8
	)

	input, err := UsdcABI.Pack("PERMIT_TYPEHASH")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("PERMIT_TYPEHASH", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].([32]uint8)

	return out0, nil
}

func (c *USDC) ReceiveWithAuthorizationTypehash(ctx context.Context, blockNumber *big.Int) ([32]uint8, error) {
	var (
		out0 [32]uint8
	)

	input, err := UsdcABI.Pack("RECEIVE_WITH_AUTHORIZATION_TYPEHASH")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("RECEIVE_WITH_AUTHORIZATION_TYPEHASH", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].([32]uint8)

	return out0, nil
}

func (c *USDC) TransferWithAuthorizationTypehash(ctx context.Context, blockNumber *big.Int) ([32]uint8, error) {
	var (
		out0 [32]uint8
	)

	input, err := UsdcABI.Pack("TRANSFER_WITH_AUTHORIZATION_TYPEHASH")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("TRANSFER_WITH_AUTHORIZATION_TYPEHASH", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].([32]uint8)

	return out0, nil
}

func (c *USDC) Allowance(ctx context.Context, owner common.Address, spender common.Address, blockNumber *big.Int) (*big.Int, error) {
	var (
		out0 *big.Int
	)

	input, err := UsdcABI.Pack("allowance", owner, spender)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("allowance", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(*big.Int)

	return out0, nil
}

func (c *USDC) Approve(ctx context.Context, spender common.Address, value *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("approve", spender, value)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("approve", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) AuthorizationState(ctx context.Context, authorizer common.Address, nonce [32]uint8, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("authorizationState", authorizer, nonce)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("authorizationState", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) BalanceOf(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var (
		out0 *big.Int
	)

	input, err := UsdcABI.Pack("balanceOf", account)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("balanceOf", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(*big.Int)

	return out0, nil
}

func (c *USDC) Blacklist(ctx context.Context, _account common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("blacklist", _account)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Blacklister(ctx context.Context, blockNumber *big.Int) (common.Address, error) {
	var (
		out0 common.Address
	)

	input, err := UsdcABI.Pack("blacklister")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("blacklister", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(common.Address)

	return out0, nil
}

func (c *USDC) Burn(ctx context.Context, _amount *big.Int, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("burn", _amount)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) CancelAuthorization(ctx context.Context, authorizer common.Address, nonce [32]uint8, v uint8, r [32]uint8, s [32]uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("cancelAuthorization", authorizer, nonce, v, r, s)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) CancelAuthorization0(ctx context.Context, authorizer common.Address, nonce [32]uint8, signature []uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("cancelAuthorization0", authorizer, nonce, signature)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) ConfigureMinter(ctx context.Context, minter common.Address, minterAllowedAmount *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("configureMinter", minter, minterAllowedAmount)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("configureMinter", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) Currency(ctx context.Context, blockNumber *big.Int) (string, error) {
	var (
		out0 string
	)

	input, err := UsdcABI.Pack("currency")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("currency", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(string)

	return out0, nil
}

func (c *USDC) Decimals(ctx context.Context, blockNumber *big.Int) (uint8, error) {
	var (
		out0 uint8
	)

	input, err := UsdcABI.Pack("decimals")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("decimals", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(uint8)

	return out0, nil
}

func (c *USDC) DecreaseAllowance(ctx context.Context, spender common.Address, decrement *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("decreaseAllowance", spender, decrement)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("decreaseAllowance", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) IncreaseAllowance(ctx context.Context, spender common.Address, increment *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("increaseAllowance", spender, increment)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("increaseAllowance", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) Initialize(ctx context.Context, tokenName string, tokenSymbol string, tokenCurrency string, tokenDecimals uint8, newMasterMinter common.Address, newPauser common.Address, newBlacklister common.Address, newOwner common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("initialize", tokenName, tokenSymbol, tokenCurrency, tokenDecimals, newMasterMinter, newPauser, newBlacklister, newOwner)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) InitializeV2(ctx context.Context, newName string, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("initializeV2", newName)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) InitializeV21(ctx context.Context, lostAndFound common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("initializeV2_1", lostAndFound)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) InitializeV22(ctx context.Context, accountsToBlacklist []common.Address, newSymbol string, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("initializeV2_2", accountsToBlacklist, newSymbol)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) IsBlacklisted(ctx context.Context, _account common.Address, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("isBlacklisted", _account)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("isBlacklisted", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) IsMinter(ctx context.Context, account common.Address, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("isMinter", account)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("isMinter", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) MasterMinter(ctx context.Context, blockNumber *big.Int) (common.Address, error) {
	var (
		out0 common.Address
	)

	input, err := UsdcABI.Pack("masterMinter")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("masterMinter", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(common.Address)

	return out0, nil
}

func (c *USDC) Mint(ctx context.Context, _to common.Address, _amount *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("mint", _to, _amount)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("mint", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) MinterAllowance(ctx context.Context, minter common.Address, blockNumber *big.Int) (*big.Int, error) {
	var (
		out0 *big.Int
	)

	input, err := UsdcABI.Pack("minterAllowance", minter)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("minterAllowance", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(*big.Int)

	return out0, nil
}

func (c *USDC) Name(ctx context.Context, blockNumber *big.Int) (string, error) {
	var (
		out0 string
	)

	input, err := UsdcABI.Pack("name")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("name", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(string)

	return out0, nil
}

func (c *USDC) Nonces(ctx context.Context, owner common.Address, blockNumber *big.Int) (*big.Int, error) {
	var (
		out0 *big.Int
	)

	input, err := UsdcABI.Pack("nonces", owner)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("nonces", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(*big.Int)

	return out0, nil
}

func (c *USDC) Owner(ctx context.Context, blockNumber *big.Int) (common.Address, error) {
	var (
		out0 common.Address
	)

	input, err := UsdcABI.Pack("owner")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("owner", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(common.Address)

	return out0, nil
}

func (c *USDC) Pause(ctx context.Context, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("pause")
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Paused(ctx context.Context, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("paused")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("paused", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) Pauser(ctx context.Context, blockNumber *big.Int) (common.Address, error) {
	var (
		out0 common.Address
	)

	input, err := UsdcABI.Pack("pauser")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("pauser", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(common.Address)

	return out0, nil
}

func (c *USDC) Permit(ctx context.Context, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, signature []uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("permit", owner, spender, value, deadline, signature)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Permit0(ctx context.Context, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]uint8, s [32]uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("permit0", owner, spender, value, deadline, v, r, s)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) ReceiveWithAuthorization(ctx context.Context, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]uint8, signature []uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("receiveWithAuthorization", from, to, value, validAfter, validBefore, nonce, signature)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) ReceiveWithAuthorization0(ctx context.Context, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]uint8, v uint8, r [32]uint8, s [32]uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("receiveWithAuthorization0", from, to, value, validAfter, validBefore, nonce, v, r, s)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) RemoveMinter(ctx context.Context, minter common.Address, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("removeMinter", minter)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("removeMinter", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) RescueErc20(ctx context.Context, tokenContract common.Address, to common.Address, amount *big.Int, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("rescueERC20", tokenContract, to, amount)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Rescuer(ctx context.Context, blockNumber *big.Int) (common.Address, error) {
	var (
		out0 common.Address
	)

	input, err := UsdcABI.Pack("rescuer")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("rescuer", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(common.Address)

	return out0, nil
}

func (c *USDC) Symbol(ctx context.Context, blockNumber *big.Int) (string, error) {
	var (
		out0 string
	)

	input, err := UsdcABI.Pack("symbol")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("symbol", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(string)

	return out0, nil
}

func (c *USDC) TotalSupply(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	var (
		out0 *big.Int
	)

	input, err := UsdcABI.Pack("totalSupply")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("totalSupply", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(*big.Int)

	return out0, nil
}

func (c *USDC) Transfer(ctx context.Context, to common.Address, value *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("transfer", to, value)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("transfer", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) TransferFrom(ctx context.Context, from common.Address, to common.Address, value *big.Int, blockNumber *big.Int) (bool, error) {
	var (
		out0 bool
	)

	input, err := UsdcABI.Pack("transferFrom", from, to, value)
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("transferFrom", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(bool)

	return out0, nil
}

func (c *USDC) TransferOwnership(ctx context.Context, newOwner common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("transferOwnership", newOwner)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) TransferWithAuthorization(ctx context.Context, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]uint8, signature []uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("transferWithAuthorization", from, to, value, validAfter, validBefore, nonce, signature)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) TransferWithAuthorization0(ctx context.Context, from common.Address, to common.Address, value *big.Int, validAfter *big.Int, validBefore *big.Int, nonce [32]uint8, v uint8, r [32]uint8, s [32]uint8, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("transferWithAuthorization0", from, to, value, validAfter, validBefore, nonce, v, r, s)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) UnBlacklist(ctx context.Context, _account common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("unBlacklist", _account)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Unpause(ctx context.Context, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("unpause")
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) UpdateBlacklister(ctx context.Context, _newBlacklister common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("updateBlacklister", _newBlacklister)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) UpdateMasterMinter(ctx context.Context, _newMasterMinter common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("updateMasterMinter", _newMasterMinter)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) UpdatePauser(ctx context.Context, _newPauser common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("updatePauser", _newPauser)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) UpdateRescuer(ctx context.Context, newRescuer common.Address, blockNumber *big.Int) error {

	input, err := UsdcABI.Pack("updateRescuer", newRescuer)
	if err != nil {
		return fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	_, err = c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return fmt.Errorf("failed call method: %w", err)
	}

	return nil
}

func (c *USDC) Version(ctx context.Context, blockNumber *big.Int) (string, error) {
	var (
		out0 string
	)

	input, err := UsdcABI.Pack("version")
	if err != nil {
		return out0, fmt.Errorf("failed to pack arguments: %w", err)
	}

	msg := ethereum.CallMsg{To: &c.Address, Data: input}
	res, err := c.caller.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return out0, fmt.Errorf("failed call method: %w", err)
	}
	results, err := UsdcABI.Unpack("version", res)
	if err != nil {
		return out0, fmt.Errorf("failed to unpack outputs: %w", err)
	}
	out0 = results[0].(string)

	return out0, nil
}

var UsdcABI = func() abi.ABI {
	a, _ := abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"authorizer","type":"address"},{"indexed":true,"internalType":"bytes32","name":"nonce","type":"bytes32"}],"name":"AuthorizationCanceled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"authorizer","type":"address"},{"indexed":true,"internalType":"bytes32","name":"nonce","type":"bytes32"}],"name":"AuthorizationUsed","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_account","type":"address"}],"name":"Blacklisted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newBlacklister","type":"address"}],"name":"BlacklisterChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"burner","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Burn","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newMasterMinter","type":"address"}],"name":"MasterMinterChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"minter","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Mint","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"minter","type":"address"},{"indexed":false,"internalType":"uint256","name":"minterAllowedAmount","type":"uint256"}],"name":"MinterConfigured","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"oldMinter","type":"address"}],"name":"MinterRemoved","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":false,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[],"name":"Pause","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newAddress","type":"address"}],"name":"PauserChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"newRescuer","type":"address"}],"name":"RescuerChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_account","type":"address"}],"name":"UnBlacklisted","type":"event"},{"anonymous":false,"inputs":[],"name":"Unpause","type":"event"},{"inputs":[],"name":"CANCEL_AUTHORIZATION_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"DOMAIN_SEPARATOR","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"PERMIT_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"RECEIVE_WITH_AUTHORIZATION_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"TRANSFER_WITH_AUTHORIZATION_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"authorizer","type":"address"},{"internalType":"bytes32","name":"nonce","type":"bytes32"}],"name":"authorizationState","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_account","type":"address"}],"name":"blacklist","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"blacklister","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"authorizer","type":"address"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"cancelAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"authorizer","type":"address"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"cancelAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"minter","type":"address"},{"internalType":"uint256","name":"minterAllowedAmount","type":"uint256"}],"name":"configureMinter","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"currency","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"decrement","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"increment","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"tokenName","type":"string"},{"internalType":"string","name":"tokenSymbol","type":"string"},{"internalType":"string","name":"tokenCurrency","type":"string"},{"internalType":"uint8","name":"tokenDecimals","type":"uint8"},{"internalType":"address","name":"newMasterMinter","type":"address"},{"internalType":"address","name":"newPauser","type":"address"},{"internalType":"address","name":"newBlacklister","type":"address"},{"internalType":"address","name":"newOwner","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"newName","type":"string"}],"name":"initializeV2","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"lostAndFound","type":"address"}],"name":"initializeV2_1","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address[]","name":"accountsToBlacklist","type":"address[]"},{"internalType":"string","name":"newSymbol","type":"string"}],"name":"initializeV2_2","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_account","type":"address"}],"name":"isBlacklisted","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"isMinter","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"masterMinter","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"mint","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"minter","type":"address"}],"name":"minterAllowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"nonces","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pauser","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"permit","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"permit","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"validAfter","type":"uint256"},{"internalType":"uint256","name":"validBefore","type":"uint256"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"receiveWithAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"validAfter","type":"uint256"},{"internalType":"uint256","name":"validBefore","type":"uint256"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"receiveWithAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"minter","type":"address"}],"name":"removeMinter","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"contract IERC20","name":"tokenContract","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"rescueERC20","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"rescuer","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"validAfter","type":"uint256"},{"internalType":"uint256","name":"validBefore","type":"uint256"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"bytes","name":"signature","type":"bytes"}],"name":"transferWithAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"validAfter","type":"uint256"},{"internalType":"uint256","name":"validBefore","type":"uint256"},{"internalType":"bytes32","name":"nonce","type":"bytes32"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"transferWithAuthorization","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_account","type":"address"}],"name":"unBlacklist","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newBlacklister","type":"address"}],"name":"updateBlacklister","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newMasterMinter","type":"address"}],"name":"updateMasterMinter","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"_newPauser","type":"address"}],"name":"updatePauser","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newRescuer","type":"address"}],"name":"updateRescuer","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"pure","type":"function"}]`))
	return a
}()
