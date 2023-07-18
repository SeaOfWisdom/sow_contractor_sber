// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package library

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
)

// LibraryMetaData contains all meta data concerning the Library contract.
var LibraryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newParticipant\",\"type\":\"address\"}],\"name\":\"addParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"reviewerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"reviews\",\"type\":\"uint8[]\"}],\"name\":\"addReviewsForWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeAuthor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeReviewer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAuthorRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"}],\"name\":\"claimReviewerRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expires\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractPaperFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"name\":\"getAuthorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"name\":\"getAuthorWorks\",\"outputs\":[{\"internalType\":\"contractSPT[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reviewerAddress\",\"type\":\"address\"}],\"name\":\"getReviewerRewardsForWork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participantAddress\",\"type\":\"address\"}],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"enumParticipantHandler.Role\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorAddress\",\"type\":\"address\"}],\"name\":\"getWorksByAuthor\",\"outputs\":[{\"internalType\":\"contractSPT[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reviewerAddress\",\"type\":\"address\"}],\"name\":\"isAbleToClaimForWork\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"isAuthor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reviewerAddress\",\"type\":\"address\"}],\"name\":\"isReviewer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participantAddress\",\"type\":\"address\"}],\"name\":\"makeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"reviews\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"reviewSignatures\",\"type\":\"bytes[]\"}],\"name\":\"publishReviewsBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"authorAddressese\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_workURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"}],\"name\":\"publishWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workId\",\"type\":\"uint256\"}],\"name\":\"purchaseWork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reviewerDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sowTokenAddress\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"workFactoryAddress\",\"type\":\"address\"}],\"name\":\"setWorkFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sowToken\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LibraryABI is the input ABI used to generate the binding from.
// Deprecated: Use LibraryMetaData.ABI instead.
var LibraryABI = LibraryMetaData.ABI

// Library is an auto generated Go binding around an Ethereum contract.
type Library struct {
	LibraryCaller     // Read-only binding to the contract
	LibraryTransactor // Write-only binding to the contract
	LibraryFilterer   // Log filterer for contract events
}

// LibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibrarySession struct {
	Contract     *Library          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibraryCallerSession struct {
	Contract *LibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibraryTransactorSession struct {
	Contract     *LibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibraryRaw struct {
	Contract *Library // Generic contract binding to access the raw methods on
}

// LibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibraryCallerRaw struct {
	Contract *LibraryCaller // Generic read-only contract binding to access the raw methods on
}

// LibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibraryTransactorRaw struct {
	Contract *LibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibrary creates a new instance of Library, bound to a specific deployed contract.
func NewLibrary(address common.Address, backend bind.ContractBackend) (*Library, error) {
	contract, err := bindLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Library{LibraryCaller: LibraryCaller{contract: contract}, LibraryTransactor: LibraryTransactor{contract: contract}, LibraryFilterer: LibraryFilterer{contract: contract}}, nil
}

// NewLibraryCaller creates a new read-only instance of Library, bound to a specific deployed contract.
func NewLibraryCaller(address common.Address, caller bind.ContractCaller) (*LibraryCaller, error) {
	contract, err := bindLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryCaller{contract: contract}, nil
}

// NewLibraryTransactor creates a new write-only instance of Library, bound to a specific deployed contract.
func NewLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*LibraryTransactor, error) {
	contract, err := bindLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibraryTransactor{contract: contract}, nil
}

// NewLibraryFilterer creates a new log filterer instance of Library, bound to a specific deployed contract.
func NewLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*LibraryFilterer, error) {
	contract, err := bindLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibraryFilterer{contract: contract}, nil
}

// bindLibrary binds a generic wrapper to an already deployed contract.
func bindLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.LibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.LibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Library *LibraryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Library.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Library *LibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Library *LibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Library.Contract.contract.Transact(opts, method, params...)
}

// AuthorPercent is a free data retrieval call binding the contract method 0x68e98aa6.
//
// Solidity: function authorPercent() view returns(uint256)
func (_Library *LibraryCaller) AuthorPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "authorPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuthorPercent is a free data retrieval call binding the contract method 0x68e98aa6.
//
// Solidity: function authorPercent() view returns(uint256)
func (_Library *LibrarySession) AuthorPercent() (*big.Int, error) {
	return _Library.Contract.AuthorPercent(&_Library.CallOpts)
}

// AuthorPercent is a free data retrieval call binding the contract method 0x68e98aa6.
//
// Solidity: function authorPercent() view returns(uint256)
func (_Library *LibraryCallerSession) AuthorPercent() (*big.Int, error) {
	return _Library.Contract.AuthorPercent(&_Library.CallOpts)
}

// Expires is a free data retrieval call binding the contract method 0xb1cb0db3.
//
// Solidity: function expires() view returns(uint64)
func (_Library *LibraryCaller) Expires(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "expires")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Expires is a free data retrieval call binding the contract method 0xb1cb0db3.
//
// Solidity: function expires() view returns(uint64)
func (_Library *LibrarySession) Expires() (uint64, error) {
	return _Library.Contract.Expires(&_Library.CallOpts)
}

// Expires is a free data retrieval call binding the contract method 0xb1cb0db3.
//
// Solidity: function expires() view returns(uint64)
func (_Library *LibraryCallerSession) Expires() (uint64, error) {
	return _Library.Contract.Expires(&_Library.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Library *LibraryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Library *LibrarySession) Factory() (common.Address, error) {
	return _Library.Contract.Factory(&_Library.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Library *LibraryCallerSession) Factory() (common.Address, error) {
	return _Library.Contract.Factory(&_Library.CallOpts)
}

// GetAuthorRewards is a free data retrieval call binding the contract method 0x49e72fc6.
//
// Solidity: function getAuthorRewards(address author) view returns(uint256)
func (_Library *LibraryCaller) GetAuthorRewards(opts *bind.CallOpts, author common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getAuthorRewards", author)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuthorRewards is a free data retrieval call binding the contract method 0x49e72fc6.
//
// Solidity: function getAuthorRewards(address author) view returns(uint256)
func (_Library *LibrarySession) GetAuthorRewards(author common.Address) (*big.Int, error) {
	return _Library.Contract.GetAuthorRewards(&_Library.CallOpts, author)
}

// GetAuthorRewards is a free data retrieval call binding the contract method 0x49e72fc6.
//
// Solidity: function getAuthorRewards(address author) view returns(uint256)
func (_Library *LibraryCallerSession) GetAuthorRewards(author common.Address) (*big.Int, error) {
	return _Library.Contract.GetAuthorRewards(&_Library.CallOpts, author)
}

// GetAuthorWorks is a free data retrieval call binding the contract method 0x68d4d6c8.
//
// Solidity: function getAuthorWorks(address author) view returns(address[])
func (_Library *LibraryCaller) GetAuthorWorks(opts *bind.CallOpts, author common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getAuthorWorks", author)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorWorks is a free data retrieval call binding the contract method 0x68d4d6c8.
//
// Solidity: function getAuthorWorks(address author) view returns(address[])
func (_Library *LibrarySession) GetAuthorWorks(author common.Address) ([]common.Address, error) {
	return _Library.Contract.GetAuthorWorks(&_Library.CallOpts, author)
}

// GetAuthorWorks is a free data retrieval call binding the contract method 0x68d4d6c8.
//
// Solidity: function getAuthorWorks(address author) view returns(address[])
func (_Library *LibraryCallerSession) GetAuthorWorks(author common.Address) ([]common.Address, error) {
	return _Library.Contract.GetAuthorWorks(&_Library.CallOpts, author)
}

// GetReviewerRewardsForWork is a free data retrieval call binding the contract method 0x5e1dc696.
//
// Solidity: function getReviewerRewardsForWork(uint256 workId, address reviewerAddress) view returns(uint256)
func (_Library *LibraryCaller) GetReviewerRewardsForWork(opts *bind.CallOpts, workId *big.Int, reviewerAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getReviewerRewardsForWork", workId, reviewerAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReviewerRewardsForWork is a free data retrieval call binding the contract method 0x5e1dc696.
//
// Solidity: function getReviewerRewardsForWork(uint256 workId, address reviewerAddress) view returns(uint256)
func (_Library *LibrarySession) GetReviewerRewardsForWork(workId *big.Int, reviewerAddress common.Address) (*big.Int, error) {
	return _Library.Contract.GetReviewerRewardsForWork(&_Library.CallOpts, workId, reviewerAddress)
}

// GetReviewerRewardsForWork is a free data retrieval call binding the contract method 0x5e1dc696.
//
// Solidity: function getReviewerRewardsForWork(uint256 workId, address reviewerAddress) view returns(uint256)
func (_Library *LibraryCallerSession) GetReviewerRewardsForWork(workId *big.Int, reviewerAddress common.Address) (*big.Int, error) {
	return _Library.Contract.GetReviewerRewardsForWork(&_Library.CallOpts, workId, reviewerAddress)
}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address participantAddress) view returns(uint8)
func (_Library *LibraryCaller) GetRole(opts *bind.CallOpts, participantAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getRole", participantAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address participantAddress) view returns(uint8)
func (_Library *LibrarySession) GetRole(participantAddress common.Address) (uint8, error) {
	return _Library.Contract.GetRole(&_Library.CallOpts, participantAddress)
}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address participantAddress) view returns(uint8)
func (_Library *LibraryCallerSession) GetRole(participantAddress common.Address) (uint8, error) {
	return _Library.Contract.GetRole(&_Library.CallOpts, participantAddress)
}

// GetWorksByAuthor is a free data retrieval call binding the contract method 0x79ed8687.
//
// Solidity: function getWorksByAuthor(address authorAddress) view returns(address[])
func (_Library *LibraryCaller) GetWorksByAuthor(opts *bind.CallOpts, authorAddress common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "getWorksByAuthor", authorAddress)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWorksByAuthor is a free data retrieval call binding the contract method 0x79ed8687.
//
// Solidity: function getWorksByAuthor(address authorAddress) view returns(address[])
func (_Library *LibrarySession) GetWorksByAuthor(authorAddress common.Address) ([]common.Address, error) {
	return _Library.Contract.GetWorksByAuthor(&_Library.CallOpts, authorAddress)
}

// GetWorksByAuthor is a free data retrieval call binding the contract method 0x79ed8687.
//
// Solidity: function getWorksByAuthor(address authorAddress) view returns(address[])
func (_Library *LibraryCallerSession) GetWorksByAuthor(authorAddress common.Address) ([]common.Address, error) {
	return _Library.Contract.GetWorksByAuthor(&_Library.CallOpts, authorAddress)
}

// IsAbleToClaimForWork is a free data retrieval call binding the contract method 0x44fe7a12.
//
// Solidity: function isAbleToClaimForWork(uint256 workId, address reviewerAddress) view returns(bool)
func (_Library *LibraryCaller) IsAbleToClaimForWork(opts *bind.CallOpts, workId *big.Int, reviewerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "isAbleToClaimForWork", workId, reviewerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAbleToClaimForWork is a free data retrieval call binding the contract method 0x44fe7a12.
//
// Solidity: function isAbleToClaimForWork(uint256 workId, address reviewerAddress) view returns(bool)
func (_Library *LibrarySession) IsAbleToClaimForWork(workId *big.Int, reviewerAddress common.Address) (bool, error) {
	return _Library.Contract.IsAbleToClaimForWork(&_Library.CallOpts, workId, reviewerAddress)
}

// IsAbleToClaimForWork is a free data retrieval call binding the contract method 0x44fe7a12.
//
// Solidity: function isAbleToClaimForWork(uint256 workId, address reviewerAddress) view returns(bool)
func (_Library *LibraryCallerSession) IsAbleToClaimForWork(workId *big.Int, reviewerAddress common.Address) (bool, error) {
	return _Library.Contract.IsAbleToClaimForWork(&_Library.CallOpts, workId, reviewerAddress)
}

// IsAuthor is a free data retrieval call binding the contract method 0xd4b12530.
//
// Solidity: function isAuthor(address participant) view returns(bool)
func (_Library *LibraryCaller) IsAuthor(opts *bind.CallOpts, participant common.Address) (bool, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "isAuthor", participant)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthor is a free data retrieval call binding the contract method 0xd4b12530.
//
// Solidity: function isAuthor(address participant) view returns(bool)
func (_Library *LibrarySession) IsAuthor(participant common.Address) (bool, error) {
	return _Library.Contract.IsAuthor(&_Library.CallOpts, participant)
}

// IsAuthor is a free data retrieval call binding the contract method 0xd4b12530.
//
// Solidity: function isAuthor(address participant) view returns(bool)
func (_Library *LibraryCallerSession) IsAuthor(participant common.Address) (bool, error) {
	return _Library.Contract.IsAuthor(&_Library.CallOpts, participant)
}

// IsReviewer is a free data retrieval call binding the contract method 0xfdc6258a.
//
// Solidity: function isReviewer(address reviewerAddress) view returns(bool)
func (_Library *LibraryCaller) IsReviewer(opts *bind.CallOpts, reviewerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "isReviewer", reviewerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReviewer is a free data retrieval call binding the contract method 0xfdc6258a.
//
// Solidity: function isReviewer(address reviewerAddress) view returns(bool)
func (_Library *LibrarySession) IsReviewer(reviewerAddress common.Address) (bool, error) {
	return _Library.Contract.IsReviewer(&_Library.CallOpts, reviewerAddress)
}

// IsReviewer is a free data retrieval call binding the contract method 0xfdc6258a.
//
// Solidity: function isReviewer(address reviewerAddress) view returns(bool)
func (_Library *LibraryCallerSession) IsReviewer(reviewerAddress common.Address) (bool, error) {
	return _Library.Contract.IsReviewer(&_Library.CallOpts, reviewerAddress)
}

// MaxPercent is a free data retrieval call binding the contract method 0x59a153ba.
//
// Solidity: function maxPercent() view returns(uint256)
func (_Library *LibraryCaller) MaxPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "maxPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPercent is a free data retrieval call binding the contract method 0x59a153ba.
//
// Solidity: function maxPercent() view returns(uint256)
func (_Library *LibrarySession) MaxPercent() (*big.Int, error) {
	return _Library.Contract.MaxPercent(&_Library.CallOpts)
}

// MaxPercent is a free data retrieval call binding the contract method 0x59a153ba.
//
// Solidity: function maxPercent() view returns(uint256)
func (_Library *LibraryCallerSession) MaxPercent() (*big.Int, error) {
	return _Library.Contract.MaxPercent(&_Library.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Library *LibraryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Library *LibrarySession) Owner() (common.Address, error) {
	return _Library.Contract.Owner(&_Library.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Library *LibraryCallerSession) Owner() (common.Address, error) {
	return _Library.Contract.Owner(&_Library.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Library *LibraryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Library *LibrarySession) Paused() (bool, error) {
	return _Library.Contract.Paused(&_Library.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Library *LibraryCallerSession) Paused() (bool, error) {
	return _Library.Contract.Paused(&_Library.CallOpts)
}

// ReviewerDepositAmount is a free data retrieval call binding the contract method 0xbae5a99e.
//
// Solidity: function reviewerDepositAmount() view returns(uint256)
func (_Library *LibraryCaller) ReviewerDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "reviewerDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReviewerDepositAmount is a free data retrieval call binding the contract method 0xbae5a99e.
//
// Solidity: function reviewerDepositAmount() view returns(uint256)
func (_Library *LibrarySession) ReviewerDepositAmount() (*big.Int, error) {
	return _Library.Contract.ReviewerDepositAmount(&_Library.CallOpts)
}

// ReviewerDepositAmount is a free data retrieval call binding the contract method 0xbae5a99e.
//
// Solidity: function reviewerDepositAmount() view returns(uint256)
func (_Library *LibraryCallerSession) ReviewerDepositAmount() (*big.Int, error) {
	return _Library.Contract.ReviewerDepositAmount(&_Library.CallOpts)
}

// SowToken is a free data retrieval call binding the contract method 0x8f9eee84.
//
// Solidity: function sowToken() view returns(address)
func (_Library *LibraryCaller) SowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Library.contract.Call(opts, &out, "sowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SowToken is a free data retrieval call binding the contract method 0x8f9eee84.
//
// Solidity: function sowToken() view returns(address)
func (_Library *LibrarySession) SowToken() (common.Address, error) {
	return _Library.Contract.SowToken(&_Library.CallOpts)
}

// SowToken is a free data retrieval call binding the contract method 0x8f9eee84.
//
// Solidity: function sowToken() view returns(address)
func (_Library *LibraryCallerSession) SowToken() (common.Address, error) {
	return _Library.Contract.SowToken(&_Library.CallOpts)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_Library *LibraryTransactor) AddParticipant(opts *bind.TransactOpts, newParticipant common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "addParticipant", newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_Library *LibrarySession) AddParticipant(newParticipant common.Address) (*types.Transaction, error) {
	return _Library.Contract.AddParticipant(&_Library.TransactOpts, newParticipant)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address newParticipant) returns()
func (_Library *LibraryTransactorSession) AddParticipant(newParticipant common.Address) (*types.Transaction, error) {
	return _Library.Contract.AddParticipant(&_Library.TransactOpts, newParticipant)
}

// AddReviewsForWork is a paid mutator transaction binding the contract method 0x5f188228.
//
// Solidity: function addReviewsForWork(uint256 workId, address[] reviewerAddresses, uint8[] reviews) returns()
func (_Library *LibraryTransactor) AddReviewsForWork(opts *bind.TransactOpts, workId *big.Int, reviewerAddresses []common.Address, reviews []uint8) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "addReviewsForWork", workId, reviewerAddresses, reviews)
}

// AddReviewsForWork is a paid mutator transaction binding the contract method 0x5f188228.
//
// Solidity: function addReviewsForWork(uint256 workId, address[] reviewerAddresses, uint8[] reviews) returns()
func (_Library *LibrarySession) AddReviewsForWork(workId *big.Int, reviewerAddresses []common.Address, reviews []uint8) (*types.Transaction, error) {
	return _Library.Contract.AddReviewsForWork(&_Library.TransactOpts, workId, reviewerAddresses, reviews)
}

// AddReviewsForWork is a paid mutator transaction binding the contract method 0x5f188228.
//
// Solidity: function addReviewsForWork(uint256 workId, address[] reviewerAddresses, uint8[] reviews) returns()
func (_Library *LibraryTransactorSession) AddReviewsForWork(workId *big.Int, reviewerAddresses []common.Address, reviews []uint8) (*types.Transaction, error) {
	return _Library.Contract.AddReviewsForWork(&_Library.TransactOpts, workId, reviewerAddresses, reviews)
}

// BecomeAuthor is a paid mutator transaction binding the contract method 0x85f8e4c5.
//
// Solidity: function becomeAuthor() returns()
func (_Library *LibraryTransactor) BecomeAuthor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "becomeAuthor")
}

// BecomeAuthor is a paid mutator transaction binding the contract method 0x85f8e4c5.
//
// Solidity: function becomeAuthor() returns()
func (_Library *LibrarySession) BecomeAuthor() (*types.Transaction, error) {
	return _Library.Contract.BecomeAuthor(&_Library.TransactOpts)
}

// BecomeAuthor is a paid mutator transaction binding the contract method 0x85f8e4c5.
//
// Solidity: function becomeAuthor() returns()
func (_Library *LibraryTransactorSession) BecomeAuthor() (*types.Transaction, error) {
	return _Library.Contract.BecomeAuthor(&_Library.TransactOpts)
}

// BecomeReviewer is a paid mutator transaction binding the contract method 0xe7d22fec.
//
// Solidity: function becomeReviewer() returns()
func (_Library *LibraryTransactor) BecomeReviewer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "becomeReviewer")
}

// BecomeReviewer is a paid mutator transaction binding the contract method 0xe7d22fec.
//
// Solidity: function becomeReviewer() returns()
func (_Library *LibrarySession) BecomeReviewer() (*types.Transaction, error) {
	return _Library.Contract.BecomeReviewer(&_Library.TransactOpts)
}

// BecomeReviewer is a paid mutator transaction binding the contract method 0xe7d22fec.
//
// Solidity: function becomeReviewer() returns()
func (_Library *LibraryTransactorSession) BecomeReviewer() (*types.Transaction, error) {
	return _Library.Contract.BecomeReviewer(&_Library.TransactOpts)
}

// ClaimAuthorRewards is a paid mutator transaction binding the contract method 0x93416375.
//
// Solidity: function claimAuthorRewards() returns()
func (_Library *LibraryTransactor) ClaimAuthorRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "claimAuthorRewards")
}

// ClaimAuthorRewards is a paid mutator transaction binding the contract method 0x93416375.
//
// Solidity: function claimAuthorRewards() returns()
func (_Library *LibrarySession) ClaimAuthorRewards() (*types.Transaction, error) {
	return _Library.Contract.ClaimAuthorRewards(&_Library.TransactOpts)
}

// ClaimAuthorRewards is a paid mutator transaction binding the contract method 0x93416375.
//
// Solidity: function claimAuthorRewards() returns()
func (_Library *LibraryTransactorSession) ClaimAuthorRewards() (*types.Transaction, error) {
	return _Library.Contract.ClaimAuthorRewards(&_Library.TransactOpts)
}

// ClaimReviewerRewards is a paid mutator transaction binding the contract method 0xd5a2108a.
//
// Solidity: function claimReviewerRewards(uint256 workId) returns()
func (_Library *LibraryTransactor) ClaimReviewerRewards(opts *bind.TransactOpts, workId *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "claimReviewerRewards", workId)
}

// ClaimReviewerRewards is a paid mutator transaction binding the contract method 0xd5a2108a.
//
// Solidity: function claimReviewerRewards(uint256 workId) returns()
func (_Library *LibrarySession) ClaimReviewerRewards(workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.ClaimReviewerRewards(&_Library.TransactOpts, workId)
}

// ClaimReviewerRewards is a paid mutator transaction binding the contract method 0xd5a2108a.
//
// Solidity: function claimReviewerRewards(uint256 workId) returns()
func (_Library *LibraryTransactorSession) ClaimReviewerRewards(workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.ClaimReviewerRewards(&_Library.TransactOpts, workId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Library *LibraryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Library *LibrarySession) Initialize() (*types.Transaction, error) {
	return _Library.Contract.Initialize(&_Library.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Library *LibraryTransactorSession) Initialize() (*types.Transaction, error) {
	return _Library.Contract.Initialize(&_Library.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Library *LibraryTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Library *LibrarySession) Join() (*types.Transaction, error) {
	return _Library.Contract.Join(&_Library.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Library *LibraryTransactorSession) Join() (*types.Transaction, error) {
	return _Library.Contract.Join(&_Library.TransactOpts)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(address participantAddress) returns()
func (_Library *LibraryTransactor) MakeAdmin(opts *bind.TransactOpts, participantAddress common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "makeAdmin", participantAddress)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(address participantAddress) returns()
func (_Library *LibrarySession) MakeAdmin(participantAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.MakeAdmin(&_Library.TransactOpts, participantAddress)
}

// MakeAdmin is a paid mutator transaction binding the contract method 0x472905ca.
//
// Solidity: function makeAdmin(address participantAddress) returns()
func (_Library *LibraryTransactorSession) MakeAdmin(participantAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.MakeAdmin(&_Library.TransactOpts, participantAddress)
}

// PublishReviewsBatch is a paid mutator transaction binding the contract method 0x3cf921c7.
//
// Solidity: function publishReviewsBatch(uint256 workId, string[] reviews, bytes[] reviewSignatures) returns()
func (_Library *LibraryTransactor) PublishReviewsBatch(opts *bind.TransactOpts, workId *big.Int, reviews []string, reviewSignatures [][]byte) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "publishReviewsBatch", workId, reviews, reviewSignatures)
}

// PublishReviewsBatch is a paid mutator transaction binding the contract method 0x3cf921c7.
//
// Solidity: function publishReviewsBatch(uint256 workId, string[] reviews, bytes[] reviewSignatures) returns()
func (_Library *LibrarySession) PublishReviewsBatch(workId *big.Int, reviews []string, reviewSignatures [][]byte) (*types.Transaction, error) {
	return _Library.Contract.PublishReviewsBatch(&_Library.TransactOpts, workId, reviews, reviewSignatures)
}

// PublishReviewsBatch is a paid mutator transaction binding the contract method 0x3cf921c7.
//
// Solidity: function publishReviewsBatch(uint256 workId, string[] reviews, bytes[] reviewSignatures) returns()
func (_Library *LibraryTransactorSession) PublishReviewsBatch(workId *big.Int, reviews []string, reviewSignatures [][]byte) (*types.Transaction, error) {
	return _Library.Contract.PublishReviewsBatch(&_Library.TransactOpts, workId, reviews, reviewSignatures)
}

// PublishWork is a paid mutator transaction binding the contract method 0x468df9fc.
//
// Solidity: function publishWork(address[] authorAddressese, string name, string _workURI, uint256 workId) returns()
func (_Library *LibraryTransactor) PublishWork(opts *bind.TransactOpts, authorAddressese []common.Address, name string, _workURI string, workId *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "publishWork", authorAddressese, name, _workURI, workId)
}

// PublishWork is a paid mutator transaction binding the contract method 0x468df9fc.
//
// Solidity: function publishWork(address[] authorAddressese, string name, string _workURI, uint256 workId) returns()
func (_Library *LibrarySession) PublishWork(authorAddressese []common.Address, name string, _workURI string, workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.PublishWork(&_Library.TransactOpts, authorAddressese, name, _workURI, workId)
}

// PublishWork is a paid mutator transaction binding the contract method 0x468df9fc.
//
// Solidity: function publishWork(address[] authorAddressese, string name, string _workURI, uint256 workId) returns()
func (_Library *LibraryTransactorSession) PublishWork(authorAddressese []common.Address, name string, _workURI string, workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.PublishWork(&_Library.TransactOpts, authorAddressese, name, _workURI, workId)
}

// PurchaseWork is a paid mutator transaction binding the contract method 0x07bc1373.
//
// Solidity: function purchaseWork(uint256 workId) returns()
func (_Library *LibraryTransactor) PurchaseWork(opts *bind.TransactOpts, workId *big.Int) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "purchaseWork", workId)
}

// PurchaseWork is a paid mutator transaction binding the contract method 0x07bc1373.
//
// Solidity: function purchaseWork(uint256 workId) returns()
func (_Library *LibrarySession) PurchaseWork(workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.PurchaseWork(&_Library.TransactOpts, workId)
}

// PurchaseWork is a paid mutator transaction binding the contract method 0x07bc1373.
//
// Solidity: function purchaseWork(uint256 workId) returns()
func (_Library *LibraryTransactorSession) PurchaseWork(workId *big.Int) (*types.Transaction, error) {
	return _Library.Contract.PurchaseWork(&_Library.TransactOpts, workId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Library *LibraryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Library *LibrarySession) RenounceOwnership() (*types.Transaction, error) {
	return _Library.Contract.RenounceOwnership(&_Library.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Library *LibraryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Library.Contract.RenounceOwnership(&_Library.TransactOpts)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address sowTokenAddress) returns()
func (_Library *LibraryTransactor) SetToken(opts *bind.TransactOpts, sowTokenAddress common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "setToken", sowTokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address sowTokenAddress) returns()
func (_Library *LibrarySession) SetToken(sowTokenAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetToken(&_Library.TransactOpts, sowTokenAddress)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address sowTokenAddress) returns()
func (_Library *LibraryTransactorSession) SetToken(sowTokenAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetToken(&_Library.TransactOpts, sowTokenAddress)
}

// SetWorkFactory is a paid mutator transaction binding the contract method 0x83349fe9.
//
// Solidity: function setWorkFactory(address workFactoryAddress) returns()
func (_Library *LibraryTransactor) SetWorkFactory(opts *bind.TransactOpts, workFactoryAddress common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "setWorkFactory", workFactoryAddress)
}

// SetWorkFactory is a paid mutator transaction binding the contract method 0x83349fe9.
//
// Solidity: function setWorkFactory(address workFactoryAddress) returns()
func (_Library *LibrarySession) SetWorkFactory(workFactoryAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetWorkFactory(&_Library.TransactOpts, workFactoryAddress)
}

// SetWorkFactory is a paid mutator transaction binding the contract method 0x83349fe9.
//
// Solidity: function setWorkFactory(address workFactoryAddress) returns()
func (_Library *LibraryTransactorSession) SetWorkFactory(workFactoryAddress common.Address) (*types.Transaction, error) {
	return _Library.Contract.SetWorkFactory(&_Library.TransactOpts, workFactoryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Library *LibraryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Library.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Library *LibrarySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Library.Contract.TransferOwnership(&_Library.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Library *LibraryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Library.Contract.TransferOwnership(&_Library.TransactOpts, newOwner)
}

// LibraryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Library contract.
type LibraryInitializedIterator struct {
	Event *LibraryInitialized // Event containing the contract specifics and raw log

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
func (it *LibraryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryInitialized)
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
		it.Event = new(LibraryInitialized)
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
func (it *LibraryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryInitialized represents a Initialized event raised by the Library contract.
type LibraryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Library *LibraryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LibraryInitializedIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LibraryInitializedIterator{contract: _Library.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Library *LibraryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LibraryInitialized) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryInitialized)
				if err := _Library.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Library *LibraryFilterer) ParseInitialized(log types.Log) (*LibraryInitialized, error) {
	event := new(LibraryInitialized)
	if err := _Library.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Library contract.
type LibraryOwnershipTransferredIterator struct {
	Event *LibraryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryOwnershipTransferred)
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
		it.Event = new(LibraryOwnershipTransferred)
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
func (it *LibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryOwnershipTransferred represents a OwnershipTransferred event raised by the Library contract.
type LibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Library *LibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LibraryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Library.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LibraryOwnershipTransferredIterator{contract: _Library.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Library *LibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LibraryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Library.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryOwnershipTransferred)
				if err := _Library.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Library *LibraryFilterer) ParseOwnershipTransferred(log types.Log) (*LibraryOwnershipTransferred, error) {
	event := new(LibraryOwnershipTransferred)
	if err := _Library.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Library contract.
type LibraryPausedIterator struct {
	Event *LibraryPaused // Event containing the contract specifics and raw log

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
func (it *LibraryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryPaused)
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
		it.Event = new(LibraryPaused)
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
func (it *LibraryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryPaused represents a Paused event raised by the Library contract.
type LibraryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Library *LibraryFilterer) FilterPaused(opts *bind.FilterOpts) (*LibraryPausedIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LibraryPausedIterator{contract: _Library.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Library *LibraryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LibraryPaused) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryPaused)
				if err := _Library.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Library *LibraryFilterer) ParsePaused(log types.Log) (*LibraryPaused, error) {
	event := new(LibraryPaused)
	if err := _Library.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibraryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Library contract.
type LibraryUnpausedIterator struct {
	Event *LibraryUnpaused // Event containing the contract specifics and raw log

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
func (it *LibraryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibraryUnpaused)
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
		it.Event = new(LibraryUnpaused)
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
func (it *LibraryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibraryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibraryUnpaused represents a Unpaused event raised by the Library contract.
type LibraryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Library *LibraryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LibraryUnpausedIterator, error) {

	logs, sub, err := _Library.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LibraryUnpausedIterator{contract: _Library.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Library *LibraryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LibraryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Library.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibraryUnpaused)
				if err := _Library.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Library *LibraryFilterer) ParseUnpaused(log types.Log) (*LibraryUnpaused, error) {
	event := new(LibraryUnpaused)
	if err := _Library.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
