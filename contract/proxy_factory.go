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

// ProxyFactoryMetaData contains all meta data concerning the ProxyFactory contract.
var ProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"createProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"computeProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyFactoryMetaData.ABI instead.
var ProxyFactoryABI = ProxyFactoryMetaData.ABI

// ProxyFactory is an auto generated Go binding around an Ethereum contract.
type ProxyFactory struct {
	ProxyFactoryCaller     // Read-only binding to the contract
	ProxyFactoryTransactor // Write-only binding to the contract
	ProxyFactoryFilterer   // Log filterer for contract events
}

// ProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyFactorySession struct {
	Contract     *ProxyFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyFactoryCallerSession struct {
	Contract *ProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyFactoryTransactorSession struct {
	Contract     *ProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyFactoryRaw struct {
	Contract *ProxyFactory // Generic contract binding to access the raw methods on
}

// ProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyFactoryCallerRaw struct {
	Contract *ProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyFactoryTransactorRaw struct {
	Contract *ProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyFactory creates a new instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactory(address common.Address, backend bind.ContractBackend) (*ProxyFactory, error) {
	contract, err := bindProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyFactory{ProxyFactoryCaller: ProxyFactoryCaller{contract: contract}, ProxyFactoryTransactor: ProxyFactoryTransactor{contract: contract}, ProxyFactoryFilterer: ProxyFactoryFilterer{contract: contract}}, nil
}

// NewProxyFactoryCaller creates a new read-only instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*ProxyFactoryCaller, error) {
	contract, err := bindProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryCaller{contract: contract}, nil
}

// NewProxyFactoryTransactor creates a new write-only instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyFactoryTransactor, error) {
	contract, err := bindProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryTransactor{contract: contract}, nil
}

// NewProxyFactoryFilterer creates a new log filterer instance of ProxyFactory, bound to a specific deployed contract.
func NewProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFactoryFilterer, error) {
	contract, err := bindProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFactoryFilterer{contract: contract}, nil
}

// bindProxyFactory binds a generic wrapper to an already deployed contract.
func bindProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyFactory *ProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyFactory.Contract.ProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyFactory *ProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyFactory.Contract.ProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyFactory *ProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyFactory.Contract.ProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyFactory *ProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyFactory *ProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyFactory *ProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0x67c8278e.
//
// Solidity: function computeProxyAddress(address user, bytes32 salt) view returns(address)
func (_ProxyFactory *ProxyFactoryCaller) ComputeProxyAddress(opts *bind.CallOpts, user common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ProxyFactory.contract.Call(opts, &out, "computeProxyAddress", user, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0x67c8278e.
//
// Solidity: function computeProxyAddress(address user, bytes32 salt) view returns(address)
func (_ProxyFactory *ProxyFactorySession) ComputeProxyAddress(user common.Address, salt [32]byte) (common.Address, error) {
	return _ProxyFactory.Contract.ComputeProxyAddress(&_ProxyFactory.CallOpts, user, salt)
}

// ComputeProxyAddress is a free data retrieval call binding the contract method 0x67c8278e.
//
// Solidity: function computeProxyAddress(address user, bytes32 salt) view returns(address)
func (_ProxyFactory *ProxyFactoryCallerSession) ComputeProxyAddress(user common.Address, salt [32]byte) (common.Address, error) {
	return _ProxyFactory.Contract.ComputeProxyAddress(&_ProxyFactory.CallOpts, user, salt)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x81e1f05f.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, bytes32 salt) returns()
func (_ProxyFactory *ProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ProxyFactory.contract.Transact(opts, "createProxy", paymentToken, payment, paymentReceiver, salt)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x81e1f05f.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, bytes32 salt) returns()
func (_ProxyFactory *ProxyFactorySession) CreateProxy(paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.CreateProxy(&_ProxyFactory.TransactOpts, paymentToken, payment, paymentReceiver, salt)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x81e1f05f.
//
// Solidity: function createProxy(address paymentToken, uint256 payment, address paymentReceiver, bytes32 salt) returns()
func (_ProxyFactory *ProxyFactoryTransactorSession) CreateProxy(paymentToken common.Address, payment *big.Int, paymentReceiver common.Address, salt [32]byte) (*types.Transaction, error) {
	return _ProxyFactory.Contract.CreateProxy(&_ProxyFactory.TransactOpts, paymentToken, payment, paymentReceiver, salt)
}
