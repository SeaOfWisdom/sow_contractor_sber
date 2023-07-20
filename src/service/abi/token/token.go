// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bond

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BondABI is the input ABI used to generate the binding from.
const BondABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnAndSetPending\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"binancePool\",\"type\":\"address\"}],\"name\":\"changeBinancePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crossChainBridge\",\"type\":\"address\"}],\"name\":\"changeCrossChainBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectableFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedSupply\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintBonds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"repairCollectableFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"repairRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSharesSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUnbondedBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updatePendingBurning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"updateRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"updateRatioAndFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Bond is an auto generated Go binding around an Ethereum contract.
type Bond struct {
	BondCaller     // Read-only binding to the contract
	BondTransactor // Write-only binding to the contract
	BondFilterer   // Log filterer for contract events
}

// BondCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondSession struct {
	Contract     *Bond             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondCallerSession struct {
	Contract *BondCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondTransactorSession struct {
	Contract     *BondTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondRaw struct {
	Contract *Bond // Generic contract binding to access the raw methods on
}

// BondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondCallerRaw struct {
	Contract *BondCaller // Generic read-only contract binding to access the raw methods on
}

// BondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondTransactorRaw struct {
	Contract *BondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBond creates a new instance of Bond, bound to a specific deployed contract.
func NewBond(address common.Address, backend bind.ContractBackend) (*Bond, error) {
	contract, err := bindBond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bond{BondCaller: BondCaller{contract: contract}, BondTransactor: BondTransactor{contract: contract}, BondFilterer: BondFilterer{contract: contract}}, nil
}

// NewBondCaller creates a new read-only instance of Bond, bound to a specific deployed contract.
func NewBondCaller(address common.Address, caller bind.ContractCaller) (*BondCaller, error) {
	contract, err := bindBond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondCaller{contract: contract}, nil
}

// NewBondTransactor creates a new write-only instance of Bond, bound to a specific deployed contract.
func NewBondTransactor(address common.Address, transactor bind.ContractTransactor) (*BondTransactor, error) {
	contract, err := bindBond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondTransactor{contract: contract}, nil
}

// NewBondFilterer creates a new log filterer instance of Bond, bound to a specific deployed contract.
func NewBondFilterer(address common.Address, filterer bind.ContractFilterer) (*BondFilterer, error) {
	contract, err := bindBond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondFilterer{contract: contract}, nil
}

// bindBond binds a generic wrapper to an already deployed contract.
func bindBond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bond *BondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bond.Contract.BondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bond *BondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bond.Contract.BondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bond *BondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bond.Contract.BondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bond *BondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bond *BondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bond *BondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bond.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bond *BondCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bond *BondSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bond.Contract.Allowance(&_Bond.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bond *BondCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bond.Contract.Allowance(&_Bond.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bond *BondCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bond *BondSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bond.Contract.BalanceOf(&_Bond.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bond *BondCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bond.Contract.BalanceOf(&_Bond.CallOpts, account)
}

// CollectableFee is a free data retrieval call binding the contract method 0x6406dd5c.
//
// Solidity: function collectableFee() view returns(uint256)
func (_Bond *BondCaller) CollectableFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "collectableFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectableFee is a free data retrieval call binding the contract method 0x6406dd5c.
//
// Solidity: function collectableFee() view returns(uint256)
func (_Bond *BondSession) CollectableFee() (*big.Int, error) {
	return _Bond.Contract.CollectableFee(&_Bond.CallOpts)
}

// CollectableFee is a free data retrieval call binding the contract method 0x6406dd5c.
//
// Solidity: function collectableFee() view returns(uint256)
func (_Bond *BondCallerSession) CollectableFee() (*big.Int, error) {
	return _Bond.Contract.CollectableFee(&_Bond.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bond *BondCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bond *BondSession) Decimals() (uint8, error) {
	return _Bond.Contract.Decimals(&_Bond.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bond *BondCallerSession) Decimals() (uint8, error) {
	return _Bond.Contract.Decimals(&_Bond.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(int256)
func (_Bond *BondCaller) LockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "lockedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(int256)
func (_Bond *BondSession) LockedSupply() (*big.Int, error) {
	return _Bond.Contract.LockedSupply(&_Bond.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(int256)
func (_Bond *BondCallerSession) LockedSupply() (*big.Int, error) {
	return _Bond.Contract.LockedSupply(&_Bond.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bond *BondCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bond *BondSession) Name() (string, error) {
	return _Bond.Contract.Name(&_Bond.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bond *BondCallerSession) Name() (string, error) {
	return _Bond.Contract.Name(&_Bond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bond *BondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bond *BondSession) Owner() (common.Address, error) {
	return _Bond.Contract.Owner(&_Bond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bond *BondCallerSession) Owner() (common.Address, error) {
	return _Bond.Contract.Owner(&_Bond.CallOpts)
}

// PendingBurn is a free data retrieval call binding the contract method 0x583585f4.
//
// Solidity: function pendingBurn(address account) view returns(uint256)
func (_Bond *BondCaller) PendingBurn(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "pendingBurn", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingBurn is a free data retrieval call binding the contract method 0x583585f4.
//
// Solidity: function pendingBurn(address account) view returns(uint256)
func (_Bond *BondSession) PendingBurn(account common.Address) (*big.Int, error) {
	return _Bond.Contract.PendingBurn(&_Bond.CallOpts, account)
}

// PendingBurn is a free data retrieval call binding the contract method 0x583585f4.
//
// Solidity: function pendingBurn(address account) view returns(uint256)
func (_Bond *BondCallerSession) PendingBurn(account common.Address) (*big.Int, error) {
	return _Bond.Contract.PendingBurn(&_Bond.CallOpts, account)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() view returns(uint256)
func (_Bond *BondCaller) Ratio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "ratio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() view returns(uint256)
func (_Bond *BondSession) Ratio() (*big.Int, error) {
	return _Bond.Contract.Ratio(&_Bond.CallOpts)
}

// Ratio is a free data retrieval call binding the contract method 0x71ca337d.
//
// Solidity: function ratio() view returns(uint256)
func (_Bond *BondCallerSession) Ratio() (*big.Int, error) {
	return _Bond.Contract.Ratio(&_Bond.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bond *BondCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bond *BondSession) Symbol() (string, error) {
	return _Bond.Contract.Symbol(&_Bond.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bond *BondCallerSession) Symbol() (string, error) {
	return _Bond.Contract.Symbol(&_Bond.CallOpts)
}

// TotalSharesSupply is a free data retrieval call binding the contract method 0xd50619cc.
//
// Solidity: function totalSharesSupply() view returns(uint256)
func (_Bond *BondCaller) TotalSharesSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "totalSharesSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSharesSupply is a free data retrieval call binding the contract method 0xd50619cc.
//
// Solidity: function totalSharesSupply() view returns(uint256)
func (_Bond *BondSession) TotalSharesSupply() (*big.Int, error) {
	return _Bond.Contract.TotalSharesSupply(&_Bond.CallOpts)
}

// TotalSharesSupply is a free data retrieval call binding the contract method 0xd50619cc.
//
// Solidity: function totalSharesSupply() view returns(uint256)
func (_Bond *BondCallerSession) TotalSharesSupply() (*big.Int, error) {
	return _Bond.Contract.TotalSharesSupply(&_Bond.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Bond *BondCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Bond *BondSession) TotalStaked() (*big.Int, error) {
	return _Bond.Contract.TotalStaked(&_Bond.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Bond *BondCallerSession) TotalStaked() (*big.Int, error) {
	return _Bond.Contract.TotalStaked(&_Bond.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bond *BondCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bond *BondSession) TotalSupply() (*big.Int, error) {
	return _Bond.Contract.TotalSupply(&_Bond.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bond *BondCallerSession) TotalSupply() (*big.Int, error) {
	return _Bond.Contract.TotalSupply(&_Bond.CallOpts)
}

// TotalUnbondedBonds is a free data retrieval call binding the contract method 0x87130f03.
//
// Solidity: function totalUnbondedBonds() view returns(uint256)
func (_Bond *BondCaller) TotalUnbondedBonds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bond.contract.Call(opts, &out, "totalUnbondedBonds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUnbondedBonds is a free data retrieval call binding the contract method 0x87130f03.
//
// Solidity: function totalUnbondedBonds() view returns(uint256)
func (_Bond *BondSession) TotalUnbondedBonds() (*big.Int, error) {
	return _Bond.Contract.TotalUnbondedBonds(&_Bond.CallOpts)
}

// TotalUnbondedBonds is a free data retrieval call binding the contract method 0x87130f03.
//
// Solidity: function totalUnbondedBonds() view returns(uint256)
func (_Bond *BondCallerSession) TotalUnbondedBonds() (*big.Int, error) {
	return _Bond.Contract.TotalUnbondedBonds(&_Bond.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bond *BondTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bond *BondSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Approve(&_Bond.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bond *BondTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Approve(&_Bond.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Bond *BondTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Bond *BondSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Burn(&_Bond.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Bond *BondTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Burn(&_Bond.TransactOpts, account, amount)
}

// BurnAndSetPending is a paid mutator transaction binding the contract method 0x1c52cfae.
//
// Solidity: function burnAndSetPending(address account, uint256 amount) returns()
func (_Bond *BondTransactor) BurnAndSetPending(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "burnAndSetPending", account, amount)
}

// BurnAndSetPending is a paid mutator transaction binding the contract method 0x1c52cfae.
//
// Solidity: function burnAndSetPending(address account, uint256 amount) returns()
func (_Bond *BondSession) BurnAndSetPending(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.BurnAndSetPending(&_Bond.TransactOpts, account, amount)
}

// BurnAndSetPending is a paid mutator transaction binding the contract method 0x1c52cfae.
//
// Solidity: function burnAndSetPending(address account, uint256 amount) returns()
func (_Bond *BondTransactorSession) BurnAndSetPending(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.BurnAndSetPending(&_Bond.TransactOpts, account, amount)
}

// ChangeBinancePool is a paid mutator transaction binding the contract method 0x83a8687b.
//
// Solidity: function changeBinancePool(address binancePool) returns()
func (_Bond *BondTransactor) ChangeBinancePool(opts *bind.TransactOpts, binancePool common.Address) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "changeBinancePool", binancePool)
}

// ChangeBinancePool is a paid mutator transaction binding the contract method 0x83a8687b.
//
// Solidity: function changeBinancePool(address binancePool) returns()
func (_Bond *BondSession) ChangeBinancePool(binancePool common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeBinancePool(&_Bond.TransactOpts, binancePool)
}

// ChangeBinancePool is a paid mutator transaction binding the contract method 0x83a8687b.
//
// Solidity: function changeBinancePool(address binancePool) returns()
func (_Bond *BondTransactorSession) ChangeBinancePool(binancePool common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeBinancePool(&_Bond.TransactOpts, binancePool)
}

// ChangeCrossChainBridge is a paid mutator transaction binding the contract method 0xe3ec72be.
//
// Solidity: function changeCrossChainBridge(address crossChainBridge) returns()
func (_Bond *BondTransactor) ChangeCrossChainBridge(opts *bind.TransactOpts, crossChainBridge common.Address) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "changeCrossChainBridge", crossChainBridge)
}

// ChangeCrossChainBridge is a paid mutator transaction binding the contract method 0xe3ec72be.
//
// Solidity: function changeCrossChainBridge(address crossChainBridge) returns()
func (_Bond *BondSession) ChangeCrossChainBridge(crossChainBridge common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeCrossChainBridge(&_Bond.TransactOpts, crossChainBridge)
}

// ChangeCrossChainBridge is a paid mutator transaction binding the contract method 0xe3ec72be.
//
// Solidity: function changeCrossChainBridge(address crossChainBridge) returns()
func (_Bond *BondTransactorSession) ChangeCrossChainBridge(crossChainBridge common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeCrossChainBridge(&_Bond.TransactOpts, crossChainBridge)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator) returns()
func (_Bond *BondTransactor) ChangeOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "changeOperator", operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator) returns()
func (_Bond *BondSession) ChangeOperator(operator common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeOperator(&_Bond.TransactOpts, operator)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(address operator) returns()
func (_Bond *BondTransactorSession) ChangeOperator(operator common.Address) (*types.Transaction, error) {
	return _Bond.Contract.ChangeOperator(&_Bond.TransactOpts, operator)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bond *BondTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bond *BondSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.DecreaseAllowance(&_Bond.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bond *BondTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.DecreaseAllowance(&_Bond.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bond *BondTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bond *BondSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.IncreaseAllowance(&_Bond.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bond *BondTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.IncreaseAllowance(&_Bond.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address operator) returns()
func (_Bond *BondTransactor) Initialize(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "initialize", operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address operator) returns()
func (_Bond *BondSession) Initialize(operator common.Address) (*types.Transaction, error) {
	return _Bond.Contract.Initialize(&_Bond.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address operator) returns()
func (_Bond *BondTransactorSession) Initialize(operator common.Address) (*types.Transaction, error) {
	return _Bond.Contract.Initialize(&_Bond.TransactOpts, operator)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 shares) returns()
func (_Bond *BondTransactor) Mint(opts *bind.TransactOpts, account common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "mint", account, shares)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 shares) returns()
func (_Bond *BondSession) Mint(account common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Mint(&_Bond.TransactOpts, account, shares)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 shares) returns()
func (_Bond *BondTransactorSession) Mint(account common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Mint(&_Bond.TransactOpts, account, shares)
}

// MintBonds is a paid mutator transaction binding the contract method 0xa85374e1.
//
// Solidity: function mintBonds(address account, uint256 amount) returns()
func (_Bond *BondTransactor) MintBonds(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "mintBonds", account, amount)
}

// MintBonds is a paid mutator transaction binding the contract method 0xa85374e1.
//
// Solidity: function mintBonds(address account, uint256 amount) returns()
func (_Bond *BondSession) MintBonds(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.MintBonds(&_Bond.TransactOpts, account, amount)
}

// MintBonds is a paid mutator transaction binding the contract method 0xa85374e1.
//
// Solidity: function mintBonds(address account, uint256 amount) returns()
func (_Bond *BondTransactorSession) MintBonds(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.MintBonds(&_Bond.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bond *BondTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bond *BondSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bond.Contract.RenounceOwnership(&_Bond.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bond *BondTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bond.Contract.RenounceOwnership(&_Bond.TransactOpts)
}

// RepairCollectableFee is a paid mutator transaction binding the contract method 0x1d62f87c.
//
// Solidity: function repairCollectableFee(uint256 newFee) returns()
func (_Bond *BondTransactor) RepairCollectableFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "repairCollectableFee", newFee)
}

// RepairCollectableFee is a paid mutator transaction binding the contract method 0x1d62f87c.
//
// Solidity: function repairCollectableFee(uint256 newFee) returns()
func (_Bond *BondSession) RepairCollectableFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.RepairCollectableFee(&_Bond.TransactOpts, newFee)
}

// RepairCollectableFee is a paid mutator transaction binding the contract method 0x1d62f87c.
//
// Solidity: function repairCollectableFee(uint256 newFee) returns()
func (_Bond *BondTransactorSession) RepairCollectableFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.RepairCollectableFee(&_Bond.TransactOpts, newFee)
}

// RepairRatio is a paid mutator transaction binding the contract method 0xe5683ad9.
//
// Solidity: function repairRatio(uint256 newRatio) returns()
func (_Bond *BondTransactor) RepairRatio(opts *bind.TransactOpts, newRatio *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "repairRatio", newRatio)
}

// RepairRatio is a paid mutator transaction binding the contract method 0xe5683ad9.
//
// Solidity: function repairRatio(uint256 newRatio) returns()
func (_Bond *BondSession) RepairRatio(newRatio *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.RepairRatio(&_Bond.TransactOpts, newRatio)
}

// RepairRatio is a paid mutator transaction binding the contract method 0xe5683ad9.
//
// Solidity: function repairRatio(uint256 newRatio) returns()
func (_Bond *BondTransactorSession) RepairRatio(newRatio *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.RepairRatio(&_Bond.TransactOpts, newRatio)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bond *BondTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bond *BondSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Transfer(&_Bond.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Bond *BondTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.Transfer(&_Bond.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bond *BondTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bond *BondSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.TransferFrom(&_Bond.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Bond *BondTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.TransferFrom(&_Bond.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bond *BondTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bond *BondSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bond.Contract.TransferOwnership(&_Bond.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bond *BondTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bond.Contract.TransferOwnership(&_Bond.TransactOpts, newOwner)
}

// UpdatePendingBurning is a paid mutator transaction binding the contract method 0x3b0f4755.
//
// Solidity: function updatePendingBurning(address account, uint256 amount) returns()
func (_Bond *BondTransactor) UpdatePendingBurning(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "updatePendingBurning", account, amount)
}

// UpdatePendingBurning is a paid mutator transaction binding the contract method 0x3b0f4755.
//
// Solidity: function updatePendingBurning(address account, uint256 amount) returns()
func (_Bond *BondSession) UpdatePendingBurning(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdatePendingBurning(&_Bond.TransactOpts, account, amount)
}

// UpdatePendingBurning is a paid mutator transaction binding the contract method 0x3b0f4755.
//
// Solidity: function updatePendingBurning(address account, uint256 amount) returns()
func (_Bond *BondTransactorSession) UpdatePendingBurning(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdatePendingBurning(&_Bond.TransactOpts, account, amount)
}

// UpdateRatio is a paid mutator transaction binding the contract method 0xdc647e29.
//
// Solidity: function updateRatio(uint256 totalRewards) returns()
func (_Bond *BondTransactor) UpdateRatio(opts *bind.TransactOpts, totalRewards *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "updateRatio", totalRewards)
}

// UpdateRatio is a paid mutator transaction binding the contract method 0xdc647e29.
//
// Solidity: function updateRatio(uint256 totalRewards) returns()
func (_Bond *BondSession) UpdateRatio(totalRewards *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdateRatio(&_Bond.TransactOpts, totalRewards)
}

// UpdateRatio is a paid mutator transaction binding the contract method 0xdc647e29.
//
// Solidity: function updateRatio(uint256 totalRewards) returns()
func (_Bond *BondTransactorSession) UpdateRatio(totalRewards *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdateRatio(&_Bond.TransactOpts, totalRewards)
}

// UpdateRatioAndFee is a paid mutator transaction binding the contract method 0xc814cf89.
//
// Solidity: function updateRatioAndFee(uint256 newRatio, uint256 newFee) returns()
func (_Bond *BondTransactor) UpdateRatioAndFee(opts *bind.TransactOpts, newRatio *big.Int, newFee *big.Int) (*types.Transaction, error) {
	return _Bond.contract.Transact(opts, "updateRatioAndFee", newRatio, newFee)
}

// UpdateRatioAndFee is a paid mutator transaction binding the contract method 0xc814cf89.
//
// Solidity: function updateRatioAndFee(uint256 newRatio, uint256 newFee) returns()
func (_Bond *BondSession) UpdateRatioAndFee(newRatio *big.Int, newFee *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdateRatioAndFee(&_Bond.TransactOpts, newRatio, newFee)
}

// UpdateRatioAndFee is a paid mutator transaction binding the contract method 0xc814cf89.
//
// Solidity: function updateRatioAndFee(uint256 newRatio, uint256 newFee) returns()
func (_Bond *BondTransactorSession) UpdateRatioAndFee(newRatio *big.Int, newFee *big.Int) (*types.Transaction, error) {
	return _Bond.Contract.UpdateRatioAndFee(&_Bond.TransactOpts, newRatio, newFee)
}

// BondApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bond contract.
type BondApprovalIterator struct {
	Event *BondApproval // Event containing the contract specifics and raw log

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
func (it *BondApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondApproval)
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
		it.Event = new(BondApproval)
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
func (it *BondApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondApproval represents a Approval event raised by the Bond contract.
type BondApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bond *BondFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BondApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bond.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BondApprovalIterator{contract: _Bond.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bond *BondFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BondApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bond.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondApproval)
				if err := _Bond.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Bond *BondFilterer) ParseApproval(log types.Log) (*BondApproval, error) {
	event := new(BondApproval)
	if err := _Bond.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bond contract.
type BondOwnershipTransferredIterator struct {
	Event *BondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondOwnershipTransferred)
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
		it.Event = new(BondOwnershipTransferred)
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
func (it *BondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondOwnershipTransferred represents a OwnershipTransferred event raised by the Bond contract.
type BondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bond *BondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BondOwnershipTransferredIterator{contract: _Bond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bond *BondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondOwnershipTransferred)
				if err := _Bond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bond *BondFilterer) ParseOwnershipTransferred(log types.Log) (*BondOwnershipTransferred, error) {
	event := new(BondOwnershipTransferred)
	if err := _Bond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondRatioUpdatedIterator is returned from FilterRatioUpdated and is used to iterate over the raw logs and unpacked data for RatioUpdated events raised by the Bond contract.
type BondRatioUpdatedIterator struct {
	Event *BondRatioUpdated // Event containing the contract specifics and raw log

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
func (it *BondRatioUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondRatioUpdated)
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
		it.Event = new(BondRatioUpdated)
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
func (it *BondRatioUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondRatioUpdated represents a RatioUpdated event raised by the Bond contract.
type BondRatioUpdated struct {
	NewRatio *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRatioUpdated is a free log retrieval operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 newRatio)
func (_Bond *BondFilterer) FilterRatioUpdated(opts *bind.FilterOpts) (*BondRatioUpdatedIterator, error) {

	logs, sub, err := _Bond.contract.FilterLogs(opts, "RatioUpdated")
	if err != nil {
		return nil, err
	}
	return &BondRatioUpdatedIterator{contract: _Bond.contract, event: "RatioUpdated", logs: logs, sub: sub}, nil
}

// WatchRatioUpdated is a free log subscription operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 newRatio)
func (_Bond *BondFilterer) WatchRatioUpdated(opts *bind.WatchOpts, sink chan<- *BondRatioUpdated) (event.Subscription, error) {

	logs, sub, err := _Bond.contract.WatchLogs(opts, "RatioUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondRatioUpdated)
				if err := _Bond.contract.UnpackLog(event, "RatioUpdated", log); err != nil {
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

// ParseRatioUpdated is a log parse operation binding the contract event 0x8bfafad5cbd7165952ec3370b6dcf547a709f2dad989ecb117d6d361038dd8c5.
//
// Solidity: event RatioUpdated(uint256 newRatio)
func (_Bond *BondFilterer) ParseRatioUpdated(log types.Log) (*BondRatioUpdated, error) {
	event := new(BondRatioUpdated)
	if err := _Bond.contract.UnpackLog(event, "RatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bond contract.
type BondTransferIterator struct {
	Event *BondTransfer // Event containing the contract specifics and raw log

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
func (it *BondTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondTransfer)
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
		it.Event = new(BondTransfer)
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
func (it *BondTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondTransfer represents a Transfer event raised by the Bond contract.
type BondTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bond *BondFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BondTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bond.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BondTransferIterator{contract: _Bond.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bond *BondFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BondTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bond.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondTransfer)
				if err := _Bond.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Bond *BondFilterer) ParseTransfer(log types.Log) (*BondTransfer, error) {
	event := new(BondTransfer)
	if err := _Bond.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
