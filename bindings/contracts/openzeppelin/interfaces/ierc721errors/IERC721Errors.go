// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc721errors

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

// IERC721ErrorsMetaData contains all meta data concerning the IERC721Errors contract.
var IERC721ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"}]",
}

// IERC721ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721ErrorsMetaData.ABI instead.
var IERC721ErrorsABI = IERC721ErrorsMetaData.ABI

// IERC721Errors is an auto generated Go binding around an Ethereum contract.
type IERC721Errors struct {
	IERC721ErrorsCaller     // Read-only binding to the contract
	IERC721ErrorsTransactor // Write-only binding to the contract
	IERC721ErrorsFilterer   // Log filterer for contract events
}

// IERC721ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721ErrorsSession struct {
	Contract     *IERC721Errors    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721ErrorsCallerSession struct {
	Contract *IERC721ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC721ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721ErrorsTransactorSession struct {
	Contract     *IERC721ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC721ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721ErrorsRaw struct {
	Contract *IERC721Errors // Generic contract binding to access the raw methods on
}

// IERC721ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721ErrorsCallerRaw struct {
	Contract *IERC721ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721ErrorsTransactorRaw struct {
	Contract *IERC721ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Errors creates a new instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721Errors(address common.Address, backend bind.ContractBackend) (*IERC721Errors, error) {
	contract, err := bindIERC721Errors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Errors{IERC721ErrorsCaller: IERC721ErrorsCaller{contract: contract}, IERC721ErrorsTransactor: IERC721ErrorsTransactor{contract: contract}, IERC721ErrorsFilterer: IERC721ErrorsFilterer{contract: contract}}, nil
}

// NewIERC721ErrorsCaller creates a new read-only instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC721ErrorsCaller, error) {
	contract, err := bindIERC721Errors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsCaller{contract: contract}, nil
}

// NewIERC721ErrorsTransactor creates a new write-only instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ErrorsTransactor, error) {
	contract, err := bindIERC721Errors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsTransactor{contract: contract}, nil
}

// NewIERC721ErrorsFilterer creates a new log filterer instance of IERC721Errors, bound to a specific deployed contract.
func NewIERC721ErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ErrorsFilterer, error) {
	contract, err := bindIERC721Errors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ErrorsFilterer{contract: contract}, nil
}

// bindIERC721Errors binds a generic wrapper to an already deployed contract.
func bindIERC721Errors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC721ErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Errors *IERC721ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Errors.Contract.IERC721ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Errors *IERC721ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Errors.Contract.IERC721ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Errors *IERC721ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Errors.Contract.IERC721ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Errors *IERC721ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Errors *IERC721ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Errors *IERC721ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721Errors.Contract.contract.Transact(opts, method, params...)
}
