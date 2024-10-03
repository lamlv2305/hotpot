// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"unBlacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_currency\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_masterMinter\",\"type\":\"address\"},{\"name\":\"_pauser\",\"type\":\"address\"},{\"name\":\"_blacklister\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_gsnFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"removeMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newGsnFee\",\"type\":\"uint256\"}],\"name\":\"updateGsnFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"},{\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"configureMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHubAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"preRelayedCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"relay\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"encodedFunction\",\"type\":\"bytes\"},{\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"approvalData\",\"type\":\"bytes\"},{\"name\":\"maxPossibleCharge\",\"type\":\"uint256\"}],\"name\":\"acceptRelayedCall\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"minterAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gsnFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newMasterMinter\",\"type\":\"address\"}],\"name\":\"updateMasterMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newBlacklister\",\"type\":\"address\"}],\"name\":\"updateBlacklister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayHubVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blacklister\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"context\",\"type\":\"bytes\"},{\"name\":\"success\",\"type\":\"bool\"},{\"name\":\"actualCharge\",\"type\":\"uint256\"},{\"name\":\"preRetVal\",\"type\":\"bytes32\"}],\"name\":\"postRelayedCall\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"GSNFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GSNFeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"minterAllowedAmount\",\"type\":\"uint256\"}],\"name\":\"MinterConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldMinter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newMasterMinter\",\"type\":\"address\"}],\"name\":\"MasterMinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"UnBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newBlacklister\",\"type\":\"address\"}],\"name\":\"BlacklisterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldRelayHub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newRelayHub\",\"type\":\"address\"}],\"name\":\"RelayHubChanged\",\"type\":\"event\"}],",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address relay, address from, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce, bytes approvalData, uint256 maxPossibleCharge) view returns(uint256, bytes)
func (_ERC20 *ERC20Caller) AcceptRelayedCall(opts *bind.CallOpts, relay common.Address, from common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int, approvalData []byte, maxPossibleCharge *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "acceptRelayedCall", relay, from, encodedFunction, transactionFee, gasPrice, gasLimit, nonce, approvalData, maxPossibleCharge)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address relay, address from, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce, bytes approvalData, uint256 maxPossibleCharge) view returns(uint256, bytes)
func (_ERC20 *ERC20Session) AcceptRelayedCall(relay common.Address, from common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int, approvalData []byte, maxPossibleCharge *big.Int) (*big.Int, []byte, error) {
	return _ERC20.Contract.AcceptRelayedCall(&_ERC20.CallOpts, relay, from, encodedFunction, transactionFee, gasPrice, gasLimit, nonce, approvalData, maxPossibleCharge)
}

// AcceptRelayedCall is a free data retrieval call binding the contract method 0x83947ea0.
//
// Solidity: function acceptRelayedCall(address relay, address from, bytes encodedFunction, uint256 transactionFee, uint256 gasPrice, uint256 gasLimit, uint256 nonce, bytes approvalData, uint256 maxPossibleCharge) view returns(uint256, bytes)
func (_ERC20 *ERC20CallerSession) AcceptRelayedCall(relay common.Address, from common.Address, encodedFunction []byte, transactionFee *big.Int, gasPrice *big.Int, gasLimit *big.Int, nonce *big.Int, approvalData []byte, maxPossibleCharge *big.Int) (*big.Int, []byte, error) {
	return _ERC20.Contract.AcceptRelayedCall(&_ERC20.CallOpts, relay, from, encodedFunction, transactionFee, gasPrice, gasLimit, nonce, approvalData, maxPossibleCharge)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_ERC20 *ERC20Caller) Blacklister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "blacklister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_ERC20 *ERC20Session) Blacklister() (common.Address, error) {
	return _ERC20.Contract.Blacklister(&_ERC20.CallOpts)
}

// Blacklister is a free data retrieval call binding the contract method 0xbd102430.
//
// Solidity: function blacklister() view returns(address)
func (_ERC20 *ERC20CallerSession) Blacklister() (common.Address, error) {
	return _ERC20.Contract.Blacklister(&_ERC20.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_ERC20 *ERC20Caller) Currency(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "currency")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_ERC20 *ERC20Session) Currency() (string, error) {
	return _ERC20.Contract.Currency(&_ERC20.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() view returns(string)
func (_ERC20 *ERC20CallerSession) Currency() (string, error) {
	return _ERC20.Contract.Currency(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ERC20 *ERC20Caller) GetHubAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getHubAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ERC20 *ERC20Session) GetHubAddr() (common.Address, error) {
	return _ERC20.Contract.GetHubAddr(&_ERC20.CallOpts)
}

// GetHubAddr is a free data retrieval call binding the contract method 0x74e861d6.
//
// Solidity: function getHubAddr() view returns(address)
func (_ERC20 *ERC20CallerSession) GetHubAddr() (common.Address, error) {
	return _ERC20.Contract.GetHubAddr(&_ERC20.CallOpts)
}

// GsnFee is a free data retrieval call binding the contract method 0x9d440f68.
//
// Solidity: function gsnFee() view returns(uint256)
func (_ERC20 *ERC20Caller) GsnFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "gsnFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GsnFee is a free data retrieval call binding the contract method 0x9d440f68.
//
// Solidity: function gsnFee() view returns(uint256)
func (_ERC20 *ERC20Session) GsnFee() (*big.Int, error) {
	return _ERC20.Contract.GsnFee(&_ERC20.CallOpts)
}

// GsnFee is a free data retrieval call binding the contract method 0x9d440f68.
//
// Solidity: function gsnFee() view returns(uint256)
func (_ERC20 *ERC20CallerSession) GsnFee() (*big.Int, error) {
	return _ERC20.Contract.GsnFee(&_ERC20.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_ERC20 *ERC20Caller) IsBlacklisted(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "isBlacklisted", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_ERC20 *ERC20Session) IsBlacklisted(_account common.Address) (bool, error) {
	return _ERC20.Contract.IsBlacklisted(&_ERC20.CallOpts, _account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address _account) view returns(bool)
func (_ERC20 *ERC20CallerSession) IsBlacklisted(_account common.Address) (bool, error) {
	return _ERC20.Contract.IsBlacklisted(&_ERC20.CallOpts, _account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC20 *ERC20Caller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC20 *ERC20Session) IsMinter(account common.Address) (bool, error) {
	return _ERC20.Contract.IsMinter(&_ERC20.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_ERC20 *ERC20CallerSession) IsMinter(account common.Address) (bool, error) {
	return _ERC20.Contract.IsMinter(&_ERC20.CallOpts, account)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_ERC20 *ERC20Caller) MasterMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "masterMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_ERC20 *ERC20Session) MasterMinter() (common.Address, error) {
	return _ERC20.Contract.MasterMinter(&_ERC20.CallOpts)
}

// MasterMinter is a free data retrieval call binding the contract method 0x35d99f35.
//
// Solidity: function masterMinter() view returns(address)
func (_ERC20 *ERC20CallerSession) MasterMinter() (common.Address, error) {
	return _ERC20.Contract.MasterMinter(&_ERC20.CallOpts)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_ERC20 *ERC20Caller) MinterAllowance(opts *bind.CallOpts, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "minterAllowance", minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_ERC20 *ERC20Session) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _ERC20.Contract.MinterAllowance(&_ERC20.CallOpts, minter)
}

// MinterAllowance is a free data retrieval call binding the contract method 0x8a6db9c3.
//
// Solidity: function minterAllowance(address minter) view returns(uint256)
func (_ERC20 *ERC20CallerSession) MinterAllowance(minter common.Address) (*big.Int, error) {
	return _ERC20.Contract.MinterAllowance(&_ERC20.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20Session) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20 *ERC20CallerSession) Owner() (common.Address, error) {
	return _ERC20.Contract.Owner(&_ERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20 *ERC20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20 *ERC20Session) Paused() (bool, error) {
	return _ERC20.Contract.Paused(&_ERC20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20 *ERC20CallerSession) Paused() (bool, error) {
	return _ERC20.Contract.Paused(&_ERC20.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_ERC20 *ERC20Caller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_ERC20 *ERC20Session) Pauser() (common.Address, error) {
	return _ERC20.Contract.Pauser(&_ERC20.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_ERC20 *ERC20CallerSession) Pauser() (common.Address, error) {
	return _ERC20.Contract.Pauser(&_ERC20.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ERC20 *ERC20Caller) RelayHubVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "relayHubVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ERC20 *ERC20Session) RelayHubVersion() (string, error) {
	return _ERC20.Contract.RelayHubVersion(&_ERC20.CallOpts)
}

// RelayHubVersion is a free data retrieval call binding the contract method 0xad61ccd5.
//
// Solidity: function relayHubVersion() view returns(string)
func (_ERC20 *ERC20CallerSession) RelayHubVersion() (string, error) {
	return _ERC20.Contract.RelayHubVersion(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_ERC20 *ERC20Transactor) Blacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "blacklist", _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_ERC20 *ERC20Session) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.Blacklist(&_ERC20.TransactOpts, _account)
}

// Blacklist is a paid mutator transaction binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address _account) returns()
func (_ERC20 *ERC20TransactorSession) Blacklist(_account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.Blacklist(&_ERC20.TransactOpts, _account)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_ERC20 *ERC20Transactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_ERC20 *ERC20Session) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Burn(&_ERC20.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_ERC20 *ERC20TransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Burn(&_ERC20.TransactOpts, _amount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_ERC20 *ERC20Transactor) ConfigureMinter(opts *bind.TransactOpts, minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "configureMinter", minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_ERC20 *ERC20Session) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.ConfigureMinter(&_ERC20.TransactOpts, minter, minterAllowedAmount)
}

// ConfigureMinter is a paid mutator transaction binding the contract method 0x4e44d956.
//
// Solidity: function configureMinter(address minter, uint256 minterAllowedAmount) returns(bool)
func (_ERC20 *ERC20TransactorSession) ConfigureMinter(minter common.Address, minterAllowedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.ConfigureMinter(&_ERC20.TransactOpts, minter, minterAllowedAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, _spender, _subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x2766ba58.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner, uint256 _gsnFee) returns()
func (_ERC20 *ERC20Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address, _gsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "initialize", _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner, _gsnFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x2766ba58.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner, uint256 _gsnFee) returns()
func (_ERC20 *ERC20Session) Initialize(_name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address, _gsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Initialize(&_ERC20.TransactOpts, _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner, _gsnFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x2766ba58.
//
// Solidity: function initialize(string _name, string _symbol, string _currency, uint8 _decimals, address _masterMinter, address _pauser, address _blacklister, address _owner, uint256 _gsnFee) returns()
func (_ERC20 *ERC20TransactorSession) Initialize(_name string, _symbol string, _currency string, _decimals uint8, _masterMinter common.Address, _pauser common.Address, _blacklister common.Address, _owner common.Address, _gsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Initialize(&_ERC20.TransactOpts, _name, _symbol, _currency, _decimals, _masterMinter, _pauser, _blacklister, _owner, _gsnFee)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20 *ERC20Transactor) Initialize0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "initialize0")
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20 *ERC20Session) Initialize0() (*types.Transaction, error) {
	return _ERC20.Contract.Initialize0(&_ERC20.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC20 *ERC20TransactorSession) Initialize0() (*types.Transaction, error) {
	return _ERC20.Contract.Initialize0(&_ERC20.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ERC20 *ERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ERC20 *ERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Mint(&_ERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Mint(&_ERC20.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20 *ERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20 *ERC20Session) Pause() (*types.Transaction, error) {
	return _ERC20.Contract.Pause(&_ERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20 *ERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20.Contract.Pause(&_ERC20.TransactOpts)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ERC20 *ERC20Transactor) PostRelayedCall(opts *bind.TransactOpts, context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "postRelayedCall", context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ERC20 *ERC20Session) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ERC20.Contract.PostRelayedCall(&_ERC20.TransactOpts, context, success, actualCharge, preRetVal)
}

// PostRelayedCall is a paid mutator transaction binding the contract method 0xe06e0e22.
//
// Solidity: function postRelayedCall(bytes context, bool success, uint256 actualCharge, bytes32 preRetVal) returns()
func (_ERC20 *ERC20TransactorSession) PostRelayedCall(context []byte, success bool, actualCharge *big.Int, preRetVal [32]byte) (*types.Transaction, error) {
	return _ERC20.Contract.PostRelayedCall(&_ERC20.TransactOpts, context, success, actualCharge, preRetVal)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ERC20 *ERC20Transactor) PreRelayedCall(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "preRelayedCall", context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ERC20 *ERC20Session) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _ERC20.Contract.PreRelayedCall(&_ERC20.TransactOpts, context)
}

// PreRelayedCall is a paid mutator transaction binding the contract method 0x80274db7.
//
// Solidity: function preRelayedCall(bytes context) returns(bytes32)
func (_ERC20 *ERC20TransactorSession) PreRelayedCall(context []byte) (*types.Transaction, error) {
	return _ERC20.Contract.PreRelayedCall(&_ERC20.TransactOpts, context)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_ERC20 *ERC20Transactor) RemoveMinter(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "removeMinter", minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_ERC20 *ERC20Session) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.RemoveMinter(&_ERC20.TransactOpts, minter)
}

// RemoveMinter is a paid mutator transaction binding the contract method 0x3092afd5.
//
// Solidity: function removeMinter(address minter) returns(bool)
func (_ERC20 *ERC20TransactorSession) RemoveMinter(minter common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.RemoveMinter(&_ERC20.TransactOpts, minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20 *ERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.TransferOwnership(&_ERC20.TransactOpts, newOwner)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_ERC20 *ERC20Transactor) UnBlacklist(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "unBlacklist", _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_ERC20 *ERC20Session) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UnBlacklist(&_ERC20.TransactOpts, _account)
}

// UnBlacklist is a paid mutator transaction binding the contract method 0x1a895266.
//
// Solidity: function unBlacklist(address _account) returns()
func (_ERC20 *ERC20TransactorSession) UnBlacklist(_account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UnBlacklist(&_ERC20.TransactOpts, _account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20 *ERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20 *ERC20Session) Unpause() (*types.Transaction, error) {
	return _ERC20.Contract.Unpause(&_ERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20 *ERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20.Contract.Unpause(&_ERC20.TransactOpts)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_ERC20 *ERC20Transactor) UpdateBlacklister(opts *bind.TransactOpts, _newBlacklister common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "updateBlacklister", _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_ERC20 *ERC20Session) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateBlacklister(&_ERC20.TransactOpts, _newBlacklister)
}

// UpdateBlacklister is a paid mutator transaction binding the contract method 0xad38bf22.
//
// Solidity: function updateBlacklister(address _newBlacklister) returns()
func (_ERC20 *ERC20TransactorSession) UpdateBlacklister(_newBlacklister common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateBlacklister(&_ERC20.TransactOpts, _newBlacklister)
}

// UpdateGsnFee is a paid mutator transaction binding the contract method 0x3d88571d.
//
// Solidity: function updateGsnFee(uint256 _newGsnFee) returns()
func (_ERC20 *ERC20Transactor) UpdateGsnFee(opts *bind.TransactOpts, _newGsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "updateGsnFee", _newGsnFee)
}

// UpdateGsnFee is a paid mutator transaction binding the contract method 0x3d88571d.
//
// Solidity: function updateGsnFee(uint256 _newGsnFee) returns()
func (_ERC20 *ERC20Session) UpdateGsnFee(_newGsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateGsnFee(&_ERC20.TransactOpts, _newGsnFee)
}

// UpdateGsnFee is a paid mutator transaction binding the contract method 0x3d88571d.
//
// Solidity: function updateGsnFee(uint256 _newGsnFee) returns()
func (_ERC20 *ERC20TransactorSession) UpdateGsnFee(_newGsnFee *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateGsnFee(&_ERC20.TransactOpts, _newGsnFee)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_ERC20 *ERC20Transactor) UpdateMasterMinter(opts *bind.TransactOpts, _newMasterMinter common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "updateMasterMinter", _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_ERC20 *ERC20Session) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateMasterMinter(&_ERC20.TransactOpts, _newMasterMinter)
}

// UpdateMasterMinter is a paid mutator transaction binding the contract method 0xaa20e1e4.
//
// Solidity: function updateMasterMinter(address _newMasterMinter) returns()
func (_ERC20 *ERC20TransactorSession) UpdateMasterMinter(_newMasterMinter common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdateMasterMinter(&_ERC20.TransactOpts, _newMasterMinter)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_ERC20 *ERC20Transactor) UpdatePauser(opts *bind.TransactOpts, _newPauser common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "updatePauser", _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_ERC20 *ERC20Session) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdatePauser(&_ERC20.TransactOpts, _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_ERC20 *ERC20TransactorSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.UpdatePauser(&_ERC20.TransactOpts, _newPauser)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the ERC20 contract.
type ERC20BlacklistedIterator struct {
	Event *ERC20Blacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Blacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Blacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Blacklisted represents a Blacklisted event raised by the ERC20 contract.
type ERC20Blacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) FilterBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*ERC20BlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BlacklistedIterator{contract: _ERC20.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *ERC20Blacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Blacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Blacklisted)
				if err := _ERC20.contract.UnpackLog(event, "Blacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisted is a log parse operation binding the contract event 0xffa4e6181777692565cf28528fc88fd1516ea86b56da075235fa575af6a4b855.
//
// Solidity: event Blacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) ParseBlacklisted(log types.Log) (*ERC20Blacklisted, error) {
	event := new(ERC20Blacklisted)
	if err := _ERC20.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BlacklisterChangedIterator is returned from FilterBlacklisterChanged and is used to iterate over the raw logs and unpacked data for BlacklisterChanged events raised by the ERC20 contract.
type ERC20BlacklisterChangedIterator struct {
	Event *ERC20BlacklisterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BlacklisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BlacklisterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BlacklisterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BlacklisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BlacklisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BlacklisterChanged represents a BlacklisterChanged event raised by the ERC20 contract.
type ERC20BlacklisterChanged struct {
	NewBlacklister common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlacklisterChanged is a free log retrieval operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_ERC20 *ERC20Filterer) FilterBlacklisterChanged(opts *bind.FilterOpts, newBlacklister []common.Address) (*ERC20BlacklisterChangedIterator, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BlacklisterChangedIterator{contract: _ERC20.contract, event: "BlacklisterChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklisterChanged is a free log subscription operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_ERC20 *ERC20Filterer) WatchBlacklisterChanged(opts *bind.WatchOpts, sink chan<- *ERC20BlacklisterChanged, newBlacklister []common.Address) (event.Subscription, error) {

	var newBlacklisterRule []interface{}
	for _, newBlacklisterItem := range newBlacklister {
		newBlacklisterRule = append(newBlacklisterRule, newBlacklisterItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "BlacklisterChanged", newBlacklisterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BlacklisterChanged)
				if err := _ERC20.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklisterChanged is a log parse operation binding the contract event 0xc67398012c111ce95ecb7429b933096c977380ee6c421175a71a4a4c6c88c06e.
//
// Solidity: event BlacklisterChanged(address indexed newBlacklister)
func (_ERC20 *ERC20Filterer) ParseBlacklisterChanged(log types.Log) (*ERC20BlacklisterChanged, error) {
	event := new(ERC20BlacklisterChanged)
	if err := _ERC20.contract.UnpackLog(event, "BlacklisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ERC20 contract.
type ERC20BurnIterator struct {
	Event *ERC20Burn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Burn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Burn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Burn represents a Burn event raised by the ERC20 contract.
type ERC20Burn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_ERC20 *ERC20Filterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*ERC20BurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnIterator{contract: _ERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_ERC20 *ERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ERC20Burn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Burn)
				if err := _ERC20.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_ERC20 *ERC20Filterer) ParseBurn(log types.Log) (*ERC20Burn, error) {
	event := new(ERC20Burn)
	if err := _ERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20GSNFeeChargedIterator is returned from FilterGSNFeeCharged and is used to iterate over the raw logs and unpacked data for GSNFeeCharged events raised by the ERC20 contract.
type ERC20GSNFeeChargedIterator struct {
	Event *ERC20GSNFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20GSNFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20GSNFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20GSNFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20GSNFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20GSNFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20GSNFeeCharged represents a GSNFeeCharged event raised by the ERC20 contract.
type ERC20GSNFeeCharged struct {
	Fee  *big.Int
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGSNFeeCharged is a free log retrieval operation binding the contract event 0x690a348690cbfffd8ee5c69b32f9f4329d86b58c4905fa88ef6800ed93c1a556.
//
// Solidity: event GSNFeeCharged(uint256 fee, address user)
func (_ERC20 *ERC20Filterer) FilterGSNFeeCharged(opts *bind.FilterOpts) (*ERC20GSNFeeChargedIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "GSNFeeCharged")
	if err != nil {
		return nil, err
	}
	return &ERC20GSNFeeChargedIterator{contract: _ERC20.contract, event: "GSNFeeCharged", logs: logs, sub: sub}, nil
}

// WatchGSNFeeCharged is a free log subscription operation binding the contract event 0x690a348690cbfffd8ee5c69b32f9f4329d86b58c4905fa88ef6800ed93c1a556.
//
// Solidity: event GSNFeeCharged(uint256 fee, address user)
func (_ERC20 *ERC20Filterer) WatchGSNFeeCharged(opts *bind.WatchOpts, sink chan<- *ERC20GSNFeeCharged) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "GSNFeeCharged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20GSNFeeCharged)
				if err := _ERC20.contract.UnpackLog(event, "GSNFeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGSNFeeCharged is a log parse operation binding the contract event 0x690a348690cbfffd8ee5c69b32f9f4329d86b58c4905fa88ef6800ed93c1a556.
//
// Solidity: event GSNFeeCharged(uint256 fee, address user)
func (_ERC20 *ERC20Filterer) ParseGSNFeeCharged(log types.Log) (*ERC20GSNFeeCharged, error) {
	event := new(ERC20GSNFeeCharged)
	if err := _ERC20.contract.UnpackLog(event, "GSNFeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20GSNFeeUpdatedIterator is returned from FilterGSNFeeUpdated and is used to iterate over the raw logs and unpacked data for GSNFeeUpdated events raised by the ERC20 contract.
type ERC20GSNFeeUpdatedIterator struct {
	Event *ERC20GSNFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20GSNFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20GSNFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20GSNFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20GSNFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20GSNFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20GSNFeeUpdated represents a GSNFeeUpdated event raised by the ERC20 contract.
type ERC20GSNFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGSNFeeUpdated is a free log retrieval operation binding the contract event 0x1ad60cf17fda9805c1b11d0d7883d02ae846fdbb9ed9c6b3930f285ca127483c.
//
// Solidity: event GSNFeeUpdated(uint256 oldFee, uint256 newFee)
func (_ERC20 *ERC20Filterer) FilterGSNFeeUpdated(opts *bind.FilterOpts) (*ERC20GSNFeeUpdatedIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "GSNFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ERC20GSNFeeUpdatedIterator{contract: _ERC20.contract, event: "GSNFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchGSNFeeUpdated is a free log subscription operation binding the contract event 0x1ad60cf17fda9805c1b11d0d7883d02ae846fdbb9ed9c6b3930f285ca127483c.
//
// Solidity: event GSNFeeUpdated(uint256 oldFee, uint256 newFee)
func (_ERC20 *ERC20Filterer) WatchGSNFeeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20GSNFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "GSNFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20GSNFeeUpdated)
				if err := _ERC20.contract.UnpackLog(event, "GSNFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGSNFeeUpdated is a log parse operation binding the contract event 0x1ad60cf17fda9805c1b11d0d7883d02ae846fdbb9ed9c6b3930f285ca127483c.
//
// Solidity: event GSNFeeUpdated(uint256 oldFee, uint256 newFee)
func (_ERC20 *ERC20Filterer) ParseGSNFeeUpdated(log types.Log) (*ERC20GSNFeeUpdated, error) {
	event := new(ERC20GSNFeeUpdated)
	if err := _ERC20.contract.UnpackLog(event, "GSNFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MasterMinterChangedIterator is returned from FilterMasterMinterChanged and is used to iterate over the raw logs and unpacked data for MasterMinterChanged events raised by the ERC20 contract.
type ERC20MasterMinterChangedIterator struct {
	Event *ERC20MasterMinterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20MasterMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MasterMinterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20MasterMinterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20MasterMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MasterMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MasterMinterChanged represents a MasterMinterChanged event raised by the ERC20 contract.
type ERC20MasterMinterChanged struct {
	NewMasterMinter common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMasterMinterChanged is a free log retrieval operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_ERC20 *ERC20Filterer) FilterMasterMinterChanged(opts *bind.FilterOpts, newMasterMinter []common.Address) (*ERC20MasterMinterChangedIterator, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MasterMinterChangedIterator{contract: _ERC20.contract, event: "MasterMinterChanged", logs: logs, sub: sub}, nil
}

// WatchMasterMinterChanged is a free log subscription operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_ERC20 *ERC20Filterer) WatchMasterMinterChanged(opts *bind.WatchOpts, sink chan<- *ERC20MasterMinterChanged, newMasterMinter []common.Address) (event.Subscription, error) {

	var newMasterMinterRule []interface{}
	for _, newMasterMinterItem := range newMasterMinter {
		newMasterMinterRule = append(newMasterMinterRule, newMasterMinterItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "MasterMinterChanged", newMasterMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MasterMinterChanged)
				if err := _ERC20.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMasterMinterChanged is a log parse operation binding the contract event 0xdb66dfa9c6b8f5226fe9aac7e51897ae8ee94ac31dc70bb6c9900b2574b707e6.
//
// Solidity: event MasterMinterChanged(address indexed newMasterMinter)
func (_ERC20 *ERC20Filterer) ParseMasterMinterChanged(log types.Log) (*ERC20MasterMinterChanged, error) {
	event := new(ERC20MasterMinterChanged)
	if err := _ERC20.contract.UnpackLog(event, "MasterMinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ERC20 contract.
type ERC20MintIterator struct {
	Event *ERC20Mint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Mint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Mint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Mint represents a Mint event raised by the ERC20 contract.
type ERC20Mint struct {
	Minter common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_ERC20 *ERC20Filterer) FilterMint(opts *bind.FilterOpts, minter []common.Address, to []common.Address) (*ERC20MintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintIterator{contract: _ERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_ERC20 *ERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ERC20Mint, minter []common.Address, to []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Mint", minterRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Mint)
				if err := _ERC20.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0xab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f8.
//
// Solidity: event Mint(address indexed minter, address indexed to, uint256 amount)
func (_ERC20 *ERC20Filterer) ParseMint(log types.Log) (*ERC20Mint, error) {
	event := new(ERC20Mint)
	if err := _ERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MinterConfiguredIterator is returned from FilterMinterConfigured and is used to iterate over the raw logs and unpacked data for MinterConfigured events raised by the ERC20 contract.
type ERC20MinterConfiguredIterator struct {
	Event *ERC20MinterConfigured // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20MinterConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MinterConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20MinterConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20MinterConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MinterConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MinterConfigured represents a MinterConfigured event raised by the ERC20 contract.
type ERC20MinterConfigured struct {
	Minter              common.Address
	MinterAllowedAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinterConfigured is a free log retrieval operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_ERC20 *ERC20Filterer) FilterMinterConfigured(opts *bind.FilterOpts, minter []common.Address) (*ERC20MinterConfiguredIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MinterConfiguredIterator{contract: _ERC20.contract, event: "MinterConfigured", logs: logs, sub: sub}, nil
}

// WatchMinterConfigured is a free log subscription operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_ERC20 *ERC20Filterer) WatchMinterConfigured(opts *bind.WatchOpts, sink chan<- *ERC20MinterConfigured, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "MinterConfigured", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MinterConfigured)
				if err := _ERC20.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterConfigured is a log parse operation binding the contract event 0x46980fca912ef9bcdbd36877427b6b90e860769f604e89c0e67720cece530d20.
//
// Solidity: event MinterConfigured(address indexed minter, uint256 minterAllowedAmount)
func (_ERC20 *ERC20Filterer) ParseMinterConfigured(log types.Log) (*ERC20MinterConfigured, error) {
	event := new(ERC20MinterConfigured)
	if err := _ERC20.contract.UnpackLog(event, "MinterConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ERC20 contract.
type ERC20MinterRemovedIterator struct {
	Event *ERC20MinterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20MinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MinterRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20MinterRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20MinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MinterRemoved represents a MinterRemoved event raised by the ERC20 contract.
type ERC20MinterRemoved struct {
	OldMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_ERC20 *ERC20Filterer) FilterMinterRemoved(opts *bind.FilterOpts, oldMinter []common.Address) (*ERC20MinterRemovedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MinterRemovedIterator{contract: _ERC20.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_ERC20 *ERC20Filterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ERC20MinterRemoved, oldMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "MinterRemoved", oldMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MinterRemoved)
				if err := _ERC20.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed oldMinter)
func (_ERC20 *ERC20Filterer) ParseMinterRemoved(log types.Log) (*ERC20MinterRemoved, error) {
	event := new(ERC20MinterRemoved)
	if err := _ERC20.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20 contract.
type ERC20OwnershipTransferredIterator struct {
	Event *ERC20OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20OwnershipTransferred represents a OwnershipTransferred event raised by the ERC20 contract.
type ERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_ERC20 *ERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*ERC20OwnershipTransferredIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &ERC20OwnershipTransferredIterator{contract: _ERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_ERC20 *ERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20OwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20OwnershipTransferred)
				if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (_ERC20 *ERC20Filterer) ParseOwnershipTransferred(log types.Log) (*ERC20OwnershipTransferred, error) {
	event := new(ERC20OwnershipTransferred)
	if err := _ERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ERC20 contract.
type ERC20PauseIterator struct {
	Event *ERC20Pause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Pause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Pause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Pause represents a Pause event raised by the ERC20 contract.
type ERC20Pause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ERC20 *ERC20Filterer) FilterPause(opts *bind.FilterOpts) (*ERC20PauseIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ERC20PauseIterator{contract: _ERC20.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ERC20 *ERC20Filterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ERC20Pause) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Pause)
				if err := _ERC20.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ERC20 *ERC20Filterer) ParsePause(log types.Log) (*ERC20Pause, error) {
	event := new(ERC20Pause)
	if err := _ERC20.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20PauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the ERC20 contract.
type ERC20PauserChangedIterator struct {
	Event *ERC20PauserChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20PauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20PauserChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20PauserChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20PauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20PauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20PauserChanged represents a PauserChanged event raised by the ERC20 contract.
type ERC20PauserChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_ERC20 *ERC20Filterer) FilterPauserChanged(opts *bind.FilterOpts, newAddress []common.Address) (*ERC20PauserChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20PauserChangedIterator{contract: _ERC20.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_ERC20 *ERC20Filterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *ERC20PauserChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20PauserChanged)
				if err := _ERC20.contract.UnpackLog(event, "PauserChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserChanged is a log parse operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_ERC20 *ERC20Filterer) ParsePauserChanged(log types.Log) (*ERC20PauserChanged, error) {
	event := new(ERC20PauserChanged)
	if err := _ERC20.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20RelayHubChangedIterator is returned from FilterRelayHubChanged and is used to iterate over the raw logs and unpacked data for RelayHubChanged events raised by the ERC20 contract.
type ERC20RelayHubChangedIterator struct {
	Event *ERC20RelayHubChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20RelayHubChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20RelayHubChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20RelayHubChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20RelayHubChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20RelayHubChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20RelayHubChanged represents a RelayHubChanged event raised by the ERC20 contract.
type ERC20RelayHubChanged struct {
	OldRelayHub common.Address
	NewRelayHub common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayHubChanged is a free log retrieval operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ERC20 *ERC20Filterer) FilterRelayHubChanged(opts *bind.FilterOpts, oldRelayHub []common.Address, newRelayHub []common.Address) (*ERC20RelayHubChangedIterator, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return &ERC20RelayHubChangedIterator{contract: _ERC20.contract, event: "RelayHubChanged", logs: logs, sub: sub}, nil
}

// WatchRelayHubChanged is a free log subscription operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ERC20 *ERC20Filterer) WatchRelayHubChanged(opts *bind.WatchOpts, sink chan<- *ERC20RelayHubChanged, oldRelayHub []common.Address, newRelayHub []common.Address) (event.Subscription, error) {

	var oldRelayHubRule []interface{}
	for _, oldRelayHubItem := range oldRelayHub {
		oldRelayHubRule = append(oldRelayHubRule, oldRelayHubItem)
	}
	var newRelayHubRule []interface{}
	for _, newRelayHubItem := range newRelayHub {
		newRelayHubRule = append(newRelayHubRule, newRelayHubItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "RelayHubChanged", oldRelayHubRule, newRelayHubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20RelayHubChanged)
				if err := _ERC20.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayHubChanged is a log parse operation binding the contract event 0xb9f84b8e65164b14439ae3620df0a4d8786d896996c0282b683f9d8c08f046e8.
//
// Solidity: event RelayHubChanged(address indexed oldRelayHub, address indexed newRelayHub)
func (_ERC20 *ERC20Filterer) ParseRelayHubChanged(log types.Log) (*ERC20RelayHubChanged, error) {
	event := new(ERC20RelayHubChanged)
	if err := _ERC20.contract.UnpackLog(event, "RelayHubChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UnBlacklistedIterator is returned from FilterUnBlacklisted and is used to iterate over the raw logs and unpacked data for UnBlacklisted events raised by the ERC20 contract.
type ERC20UnBlacklistedIterator struct {
	Event *ERC20UnBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20UnBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UnBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20UnBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20UnBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UnBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UnBlacklisted represents a UnBlacklisted event raised by the ERC20 contract.
type ERC20UnBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnBlacklisted is a free log retrieval operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) FilterUnBlacklisted(opts *bind.FilterOpts, _account []common.Address) (*ERC20UnBlacklistedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UnBlacklistedIterator{contract: _ERC20.contract, event: "UnBlacklisted", logs: logs, sub: sub}, nil
}

// WatchUnBlacklisted is a free log subscription operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) WatchUnBlacklisted(opts *bind.WatchOpts, sink chan<- *ERC20UnBlacklisted, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "UnBlacklisted", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UnBlacklisted)
				if err := _ERC20.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBlacklisted is a log parse operation binding the contract event 0x117e3210bb9aa7d9baff172026820255c6f6c30ba8999d1c2fd88e2848137c4e.
//
// Solidity: event UnBlacklisted(address indexed _account)
func (_ERC20 *ERC20Filterer) ParseUnBlacklisted(log types.Log) (*ERC20UnBlacklisted, error) {
	event := new(ERC20UnBlacklisted)
	if err := _ERC20.contract.UnpackLog(event, "UnBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ERC20 contract.
type ERC20UnpauseIterator struct {
	Event *ERC20Unpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20UnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Unpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Unpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20UnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Unpause represents a Unpause event raised by the ERC20 contract.
type ERC20Unpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ERC20 *ERC20Filterer) FilterUnpause(opts *bind.FilterOpts) (*ERC20UnpauseIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ERC20UnpauseIterator{contract: _ERC20.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ERC20 *ERC20Filterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ERC20Unpause) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Unpause)
				if err := _ERC20.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ERC20 *ERC20Filterer) ParseUnpause(log types.Log) (*ERC20Unpause, error) {
	event := new(ERC20Unpause)
	if err := _ERC20.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
