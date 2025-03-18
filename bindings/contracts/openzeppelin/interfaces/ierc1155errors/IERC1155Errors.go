// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc1155errors

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

// IERC1155ErrorsMetaData contains all meta data concerning the IERC1155Errors contract.
var IERC1155ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"}]",
}

// IERC1155ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155ErrorsMetaData.ABI instead.
var IERC1155ErrorsABI = IERC1155ErrorsMetaData.ABI

// IERC1155Errors is an auto generated Go binding around an Ethereum contract.
type IERC1155Errors struct {
	IERC1155ErrorsCaller     // Read-only binding to the contract
	IERC1155ErrorsTransactor // Write-only binding to the contract
	IERC1155ErrorsFilterer   // Log filterer for contract events
}

// IERC1155ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155ErrorsSession struct {
	Contract     *IERC1155Errors   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155ErrorsCallerSession struct {
	Contract *IERC1155ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC1155ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155ErrorsTransactorSession struct {
	Contract     *IERC1155ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC1155ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155ErrorsRaw struct {
	Contract *IERC1155Errors // Generic contract binding to access the raw methods on
}

// IERC1155ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155ErrorsCallerRaw struct {
	Contract *IERC1155ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1155ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155ErrorsTransactorRaw struct {
	Contract *IERC1155ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155Errors creates a new instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155Errors(address common.Address, backend bind.ContractBackend) (*IERC1155Errors, error) {
	contract, err := bindIERC1155Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155Errors{IERC1155ErrorsCaller: IERC1155ErrorsCaller{contract: contract}, IERC1155ErrorsTransactor: IERC1155ErrorsTransactor{contract: contract}, IERC1155ErrorsFilterer: IERC1155ErrorsFilterer{contract: contract}}, nil
}

// NewIERC1155ErrorsCaller creates a new read-only instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC1155ErrorsCaller, error) {
	contract, err := bindIERC1155Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsCaller{contract: contract}, nil
}

// NewIERC1155ErrorsTransactor creates a new write-only instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155ErrorsTransactor, error) {
	contract, err := bindIERC1155Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsTransactor{contract: contract}, nil
}

// NewIERC1155ErrorsFilterer creates a new log filterer instance of IERC1155Errors, bound to a specific deployed contract.
func NewIERC1155ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155ErrorsFilterer, error) {
	contract, err := bindIERC1155Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155ErrorsFilterer{contract: contract}, nil
}

// bindIERC1155Errors binds a generic wrapper to an already deployed contract.
func bindIERC1155Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1155ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Errors *IERC1155ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Errors.Contract.IERC1155ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Errors *IERC1155ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.IERC1155ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Errors *IERC1155ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.IERC1155ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155Errors *IERC1155ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155Errors *IERC1155ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155Errors *IERC1155ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155Errors.Contract.contract.Transact(opts, method, params...)
}
