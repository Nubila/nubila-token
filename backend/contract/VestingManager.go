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

// VestingManagerTerm is an auto generated low-level Go binding around an user-defined struct.
type VestingManagerTerm struct {
	Percentage uint32
	Cliff      *big.Int
	Period     uint64
	Num        uint16
	Next       uint16
}

// VestingManagerMetaData contains all meta data concerning the VestingManager contract.
var VestingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tgeTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beneficiaries\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MONTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PPM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEASON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"YEAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"scheduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"scheduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyWithdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getSchedule\",\"inputs\":[{\"name\":\"scheduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"vestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"termIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"terms\",\"type\":\"tuple[]\",\"internalType\":\"structVestingManager.Term[]\",\"components\":[{\"name\":\"percentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"cliff\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"period\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"num\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"next\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numOfSchedules\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tgeTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBeneficiary\",\"inputs\":[{\"name\":\"scheduleId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newBeneficiary\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeneficiaryUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newBeneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ScheduleCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Vested\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"termIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"periodIdx\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// VestingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VestingManagerMetaData.ABI instead.
var VestingManagerABI = VestingManagerMetaData.ABI

// VestingManager is an auto generated Go binding around an Ethereum contract.
type VestingManager struct {
	VestingManagerCaller     // Read-only binding to the contract
	VestingManagerTransactor // Write-only binding to the contract
	VestingManagerFilterer   // Log filterer for contract events
}

// VestingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingManagerSession struct {
	Contract     *VestingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingManagerCallerSession struct {
	Contract *VestingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VestingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingManagerTransactorSession struct {
	Contract     *VestingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VestingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingManagerRaw struct {
	Contract *VestingManager // Generic contract binding to access the raw methods on
}

// VestingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingManagerCallerRaw struct {
	Contract *VestingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VestingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingManagerTransactorRaw struct {
	Contract *VestingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVestingManager creates a new instance of VestingManager, bound to a specific deployed contract.
func NewVestingManager(address common.Address, backend bind.ContractBackend) (*VestingManager, error) {
	contract, err := bindVestingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VestingManager{VestingManagerCaller: VestingManagerCaller{contract: contract}, VestingManagerTransactor: VestingManagerTransactor{contract: contract}, VestingManagerFilterer: VestingManagerFilterer{contract: contract}}, nil
}

// NewVestingManagerCaller creates a new read-only instance of VestingManager, bound to a specific deployed contract.
func NewVestingManagerCaller(address common.Address, caller bind.ContractCaller) (*VestingManagerCaller, error) {
	contract, err := bindVestingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingManagerCaller{contract: contract}, nil
}

// NewVestingManagerTransactor creates a new write-only instance of VestingManager, bound to a specific deployed contract.
func NewVestingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingManagerTransactor, error) {
	contract, err := bindVestingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingManagerTransactor{contract: contract}, nil
}

// NewVestingManagerFilterer creates a new log filterer instance of VestingManager, bound to a specific deployed contract.
func NewVestingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingManagerFilterer, error) {
	contract, err := bindVestingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingManagerFilterer{contract: contract}, nil
}

// bindVestingManager binds a generic wrapper to an already deployed contract.
func bindVestingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VestingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VestingManager *VestingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VestingManager.Contract.VestingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VestingManager *VestingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestingManager.Contract.VestingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VestingManager *VestingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VestingManager.Contract.VestingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VestingManager *VestingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VestingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VestingManager *VestingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VestingManager *VestingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VestingManager.Contract.contract.Transact(opts, method, params...)
}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint64)
func (_VestingManager *VestingManagerCaller) MONTH(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "MONTH")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint64)
func (_VestingManager *VestingManagerSession) MONTH() (uint64, error) {
	return _VestingManager.Contract.MONTH(&_VestingManager.CallOpts)
}

// MONTH is a free data retrieval call binding the contract method 0xd5999a5c.
//
// Solidity: function MONTH() view returns(uint64)
func (_VestingManager *VestingManagerCallerSession) MONTH() (uint64, error) {
	return _VestingManager.Contract.MONTH(&_VestingManager.CallOpts)
}

// PPM is a free data retrieval call binding the contract method 0x0d720bbc.
//
// Solidity: function PPM() view returns(uint32)
func (_VestingManager *VestingManagerCaller) PPM(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "PPM")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PPM is a free data retrieval call binding the contract method 0x0d720bbc.
//
// Solidity: function PPM() view returns(uint32)
func (_VestingManager *VestingManagerSession) PPM() (uint32, error) {
	return _VestingManager.Contract.PPM(&_VestingManager.CallOpts)
}

// PPM is a free data retrieval call binding the contract method 0x0d720bbc.
//
// Solidity: function PPM() view returns(uint32)
func (_VestingManager *VestingManagerCallerSession) PPM() (uint32, error) {
	return _VestingManager.Contract.PPM(&_VestingManager.CallOpts)
}

// SEASON is a free data retrieval call binding the contract method 0x7bf544d5.
//
// Solidity: function SEASON() view returns(uint64)
func (_VestingManager *VestingManagerCaller) SEASON(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "SEASON")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SEASON is a free data retrieval call binding the contract method 0x7bf544d5.
//
// Solidity: function SEASON() view returns(uint64)
func (_VestingManager *VestingManagerSession) SEASON() (uint64, error) {
	return _VestingManager.Contract.SEASON(&_VestingManager.CallOpts)
}

// SEASON is a free data retrieval call binding the contract method 0x7bf544d5.
//
// Solidity: function SEASON() view returns(uint64)
func (_VestingManager *VestingManagerCallerSession) SEASON() (uint64, error) {
	return _VestingManager.Contract.SEASON(&_VestingManager.CallOpts)
}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint64)
func (_VestingManager *VestingManagerCaller) YEAR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "YEAR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint64)
func (_VestingManager *VestingManagerSession) YEAR() (uint64, error) {
	return _VestingManager.Contract.YEAR(&_VestingManager.CallOpts)
}

// YEAR is a free data retrieval call binding the contract method 0x83914540.
//
// Solidity: function YEAR() view returns(uint64)
func (_VestingManager *VestingManagerCallerSession) YEAR() (uint64, error) {
	return _VestingManager.Contract.YEAR(&_VestingManager.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 scheduleId) view returns(uint256)
func (_VestingManager *VestingManagerCaller) Claimable(opts *bind.CallOpts, scheduleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "claimable", scheduleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 scheduleId) view returns(uint256)
func (_VestingManager *VestingManagerSession) Claimable(scheduleId *big.Int) (*big.Int, error) {
	return _VestingManager.Contract.Claimable(&_VestingManager.CallOpts, scheduleId)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 scheduleId) view returns(uint256)
func (_VestingManager *VestingManagerCallerSession) Claimable(scheduleId *big.Int) (*big.Int, error) {
	return _VestingManager.Contract.Claimable(&_VestingManager.CallOpts, scheduleId)
}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId) view returns(address beneficiary, uint256 totalAmount, uint256 vestedAmount, uint256 termIndex, (uint32,uint256,uint64,uint16,uint16)[] terms)
func (_VestingManager *VestingManagerCaller) GetSchedule(opts *bind.CallOpts, scheduleId *big.Int) (struct {
	Beneficiary  common.Address
	TotalAmount  *big.Int
	VestedAmount *big.Int
	TermIndex    *big.Int
	Terms        []VestingManagerTerm
}, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "getSchedule", scheduleId)

	outstruct := new(struct {
		Beneficiary  common.Address
		TotalAmount  *big.Int
		VestedAmount *big.Int
		TermIndex    *big.Int
		Terms        []VestingManagerTerm
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VestedAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TermIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Terms = *abi.ConvertType(out[4], new([]VestingManagerTerm)).(*[]VestingManagerTerm)

	return *outstruct, err

}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId) view returns(address beneficiary, uint256 totalAmount, uint256 vestedAmount, uint256 termIndex, (uint32,uint256,uint64,uint16,uint16)[] terms)
func (_VestingManager *VestingManagerSession) GetSchedule(scheduleId *big.Int) (struct {
	Beneficiary  common.Address
	TotalAmount  *big.Int
	VestedAmount *big.Int
	TermIndex    *big.Int
	Terms        []VestingManagerTerm
}, error) {
	return _VestingManager.Contract.GetSchedule(&_VestingManager.CallOpts, scheduleId)
}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId) view returns(address beneficiary, uint256 totalAmount, uint256 vestedAmount, uint256 termIndex, (uint32,uint256,uint64,uint16,uint16)[] terms)
func (_VestingManager *VestingManagerCallerSession) GetSchedule(scheduleId *big.Int) (struct {
	Beneficiary  common.Address
	TotalAmount  *big.Int
	VestedAmount *big.Int
	TermIndex    *big.Int
	Terms        []VestingManagerTerm
}, error) {
	return _VestingManager.Contract.GetSchedule(&_VestingManager.CallOpts, scheduleId)
}

// NumOfSchedules is a free data retrieval call binding the contract method 0x5397c099.
//
// Solidity: function numOfSchedules() view returns(uint256)
func (_VestingManager *VestingManagerCaller) NumOfSchedules(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "numOfSchedules")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfSchedules is a free data retrieval call binding the contract method 0x5397c099.
//
// Solidity: function numOfSchedules() view returns(uint256)
func (_VestingManager *VestingManagerSession) NumOfSchedules() (*big.Int, error) {
	return _VestingManager.Contract.NumOfSchedules(&_VestingManager.CallOpts)
}

// NumOfSchedules is a free data retrieval call binding the contract method 0x5397c099.
//
// Solidity: function numOfSchedules() view returns(uint256)
func (_VestingManager *VestingManagerCallerSession) NumOfSchedules() (*big.Int, error) {
	return _VestingManager.Contract.NumOfSchedules(&_VestingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VestingManager *VestingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VestingManager *VestingManagerSession) Owner() (common.Address, error) {
	return _VestingManager.Contract.Owner(&_VestingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VestingManager *VestingManagerCallerSession) Owner() (common.Address, error) {
	return _VestingManager.Contract.Owner(&_VestingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VestingManager *VestingManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VestingManager *VestingManagerSession) Paused() (bool, error) {
	return _VestingManager.Contract.Paused(&_VestingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VestingManager *VestingManagerCallerSession) Paused() (bool, error) {
	return _VestingManager.Contract.Paused(&_VestingManager.CallOpts)
}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_VestingManager *VestingManagerCaller) TgeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "tgeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_VestingManager *VestingManagerSession) TgeTimestamp() (*big.Int, error) {
	return _VestingManager.Contract.TgeTimestamp(&_VestingManager.CallOpts)
}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_VestingManager *VestingManagerCallerSession) TgeTimestamp() (*big.Int, error) {
	return _VestingManager.Contract.TgeTimestamp(&_VestingManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VestingManager *VestingManagerCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VestingManager.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VestingManager *VestingManagerSession) Token() (common.Address, error) {
	return _VestingManager.Contract.Token(&_VestingManager.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_VestingManager *VestingManagerCallerSession) Token() (common.Address, error) {
	return _VestingManager.Contract.Token(&_VestingManager.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 scheduleId) returns()
func (_VestingManager *VestingManagerTransactor) Claim(opts *bind.TransactOpts, scheduleId *big.Int) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "claim", scheduleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 scheduleId) returns()
func (_VestingManager *VestingManagerSession) Claim(scheduleId *big.Int) (*types.Transaction, error) {
	return _VestingManager.Contract.Claim(&_VestingManager.TransactOpts, scheduleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 scheduleId) returns()
func (_VestingManager *VestingManagerTransactorSession) Claim(scheduleId *big.Int) (*types.Transaction, error) {
	return _VestingManager.Contract.Claim(&_VestingManager.TransactOpts, scheduleId)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 amount) returns()
func (_VestingManager *VestingManagerTransactor) EmergencyWithdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "emergencyWithdraw", amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 amount) returns()
func (_VestingManager *VestingManagerSession) EmergencyWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _VestingManager.Contract.EmergencyWithdraw(&_VestingManager.TransactOpts, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 amount) returns()
func (_VestingManager *VestingManagerTransactorSession) EmergencyWithdraw(amount *big.Int) (*types.Transaction, error) {
	return _VestingManager.Contract.EmergencyWithdraw(&_VestingManager.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VestingManager *VestingManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VestingManager *VestingManagerSession) Pause() (*types.Transaction, error) {
	return _VestingManager.Contract.Pause(&_VestingManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VestingManager *VestingManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _VestingManager.Contract.Pause(&_VestingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VestingManager *VestingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VestingManager *VestingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VestingManager.Contract.RenounceOwnership(&_VestingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VestingManager *VestingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VestingManager.Contract.RenounceOwnership(&_VestingManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VestingManager *VestingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VestingManager *VestingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VestingManager.Contract.TransferOwnership(&_VestingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VestingManager *VestingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VestingManager.Contract.TransferOwnership(&_VestingManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VestingManager *VestingManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VestingManager *VestingManagerSession) Unpause() (*types.Transaction, error) {
	return _VestingManager.Contract.Unpause(&_VestingManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VestingManager *VestingManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _VestingManager.Contract.Unpause(&_VestingManager.TransactOpts)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0x3e8eb5a4.
//
// Solidity: function updateBeneficiary(uint256 scheduleId, address newBeneficiary) returns()
func (_VestingManager *VestingManagerTransactor) UpdateBeneficiary(opts *bind.TransactOpts, scheduleId *big.Int, newBeneficiary common.Address) (*types.Transaction, error) {
	return _VestingManager.contract.Transact(opts, "updateBeneficiary", scheduleId, newBeneficiary)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0x3e8eb5a4.
//
// Solidity: function updateBeneficiary(uint256 scheduleId, address newBeneficiary) returns()
func (_VestingManager *VestingManagerSession) UpdateBeneficiary(scheduleId *big.Int, newBeneficiary common.Address) (*types.Transaction, error) {
	return _VestingManager.Contract.UpdateBeneficiary(&_VestingManager.TransactOpts, scheduleId, newBeneficiary)
}

// UpdateBeneficiary is a paid mutator transaction binding the contract method 0x3e8eb5a4.
//
// Solidity: function updateBeneficiary(uint256 scheduleId, address newBeneficiary) returns()
func (_VestingManager *VestingManagerTransactorSession) UpdateBeneficiary(scheduleId *big.Int, newBeneficiary common.Address) (*types.Transaction, error) {
	return _VestingManager.Contract.UpdateBeneficiary(&_VestingManager.TransactOpts, scheduleId, newBeneficiary)
}

// VestingManagerBeneficiaryUpdatedIterator is returned from FilterBeneficiaryUpdated and is used to iterate over the raw logs and unpacked data for BeneficiaryUpdated events raised by the VestingManager contract.
type VestingManagerBeneficiaryUpdatedIterator struct {
	Event *VestingManagerBeneficiaryUpdated // Event containing the contract specifics and raw log

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
func (it *VestingManagerBeneficiaryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerBeneficiaryUpdated)
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
		it.Event = new(VestingManagerBeneficiaryUpdated)
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
func (it *VestingManagerBeneficiaryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerBeneficiaryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerBeneficiaryUpdated represents a BeneficiaryUpdated event raised by the VestingManager contract.
type VestingManagerBeneficiaryUpdated struct {
	Id             *big.Int
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryUpdated is a free log retrieval operation binding the contract event 0xa6651cd315e1b1d3d439de42416a9f109ebffaa21aed0412debe15da701f4412.
//
// Solidity: event BeneficiaryUpdated(uint256 indexed id, address indexed newBeneficiary)
func (_VestingManager *VestingManagerFilterer) FilterBeneficiaryUpdated(opts *bind.FilterOpts, id []*big.Int, newBeneficiary []common.Address) (*VestingManagerBeneficiaryUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newBeneficiaryRule []interface{}
	for _, newBeneficiaryItem := range newBeneficiary {
		newBeneficiaryRule = append(newBeneficiaryRule, newBeneficiaryItem)
	}

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "BeneficiaryUpdated", idRule, newBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &VestingManagerBeneficiaryUpdatedIterator{contract: _VestingManager.contract, event: "BeneficiaryUpdated", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryUpdated is a free log subscription operation binding the contract event 0xa6651cd315e1b1d3d439de42416a9f109ebffaa21aed0412debe15da701f4412.
//
// Solidity: event BeneficiaryUpdated(uint256 indexed id, address indexed newBeneficiary)
func (_VestingManager *VestingManagerFilterer) WatchBeneficiaryUpdated(opts *bind.WatchOpts, sink chan<- *VestingManagerBeneficiaryUpdated, id []*big.Int, newBeneficiary []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var newBeneficiaryRule []interface{}
	for _, newBeneficiaryItem := range newBeneficiary {
		newBeneficiaryRule = append(newBeneficiaryRule, newBeneficiaryItem)
	}

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "BeneficiaryUpdated", idRule, newBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerBeneficiaryUpdated)
				if err := _VestingManager.contract.UnpackLog(event, "BeneficiaryUpdated", log); err != nil {
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

// ParseBeneficiaryUpdated is a log parse operation binding the contract event 0xa6651cd315e1b1d3d439de42416a9f109ebffaa21aed0412debe15da701f4412.
//
// Solidity: event BeneficiaryUpdated(uint256 indexed id, address indexed newBeneficiary)
func (_VestingManager *VestingManagerFilterer) ParseBeneficiaryUpdated(log types.Log) (*VestingManagerBeneficiaryUpdated, error) {
	event := new(VestingManagerBeneficiaryUpdated)
	if err := _VestingManager.contract.UnpackLog(event, "BeneficiaryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VestingManager contract.
type VestingManagerOwnershipTransferredIterator struct {
	Event *VestingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VestingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerOwnershipTransferred)
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
		it.Event = new(VestingManagerOwnershipTransferred)
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
func (it *VestingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the VestingManager contract.
type VestingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VestingManager *VestingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VestingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VestingManagerOwnershipTransferredIterator{contract: _VestingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VestingManager *VestingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VestingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerOwnershipTransferred)
				if err := _VestingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VestingManager *VestingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*VestingManagerOwnershipTransferred, error) {
	event := new(VestingManagerOwnershipTransferred)
	if err := _VestingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VestingManager contract.
type VestingManagerPausedIterator struct {
	Event *VestingManagerPaused // Event containing the contract specifics and raw log

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
func (it *VestingManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerPaused)
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
		it.Event = new(VestingManagerPaused)
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
func (it *VestingManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerPaused represents a Paused event raised by the VestingManager contract.
type VestingManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VestingManager *VestingManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*VestingManagerPausedIterator, error) {

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VestingManagerPausedIterator{contract: _VestingManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VestingManager *VestingManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VestingManagerPaused) (event.Subscription, error) {

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerPaused)
				if err := _VestingManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VestingManager *VestingManagerFilterer) ParsePaused(log types.Log) (*VestingManagerPaused, error) {
	event := new(VestingManagerPaused)
	if err := _VestingManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingManagerScheduleCreatedIterator is returned from FilterScheduleCreated and is used to iterate over the raw logs and unpacked data for ScheduleCreated events raised by the VestingManager contract.
type VestingManagerScheduleCreatedIterator struct {
	Event *VestingManagerScheduleCreated // Event containing the contract specifics and raw log

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
func (it *VestingManagerScheduleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerScheduleCreated)
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
		it.Event = new(VestingManagerScheduleCreated)
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
func (it *VestingManagerScheduleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerScheduleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerScheduleCreated represents a ScheduleCreated event raised by the VestingManager contract.
type VestingManagerScheduleCreated struct {
	Id          *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterScheduleCreated is a free log retrieval operation binding the contract event 0x721ddaabe5cc50a5a5fdc5ce1e5636612da513f19115ed78a77b4e60665f07b4.
//
// Solidity: event ScheduleCreated(uint256 indexed id, uint256 totalAmount)
func (_VestingManager *VestingManagerFilterer) FilterScheduleCreated(opts *bind.FilterOpts, id []*big.Int) (*VestingManagerScheduleCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "ScheduleCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &VestingManagerScheduleCreatedIterator{contract: _VestingManager.contract, event: "ScheduleCreated", logs: logs, sub: sub}, nil
}

// WatchScheduleCreated is a free log subscription operation binding the contract event 0x721ddaabe5cc50a5a5fdc5ce1e5636612da513f19115ed78a77b4e60665f07b4.
//
// Solidity: event ScheduleCreated(uint256 indexed id, uint256 totalAmount)
func (_VestingManager *VestingManagerFilterer) WatchScheduleCreated(opts *bind.WatchOpts, sink chan<- *VestingManagerScheduleCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "ScheduleCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerScheduleCreated)
				if err := _VestingManager.contract.UnpackLog(event, "ScheduleCreated", log); err != nil {
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

// ParseScheduleCreated is a log parse operation binding the contract event 0x721ddaabe5cc50a5a5fdc5ce1e5636612da513f19115ed78a77b4e60665f07b4.
//
// Solidity: event ScheduleCreated(uint256 indexed id, uint256 totalAmount)
func (_VestingManager *VestingManagerFilterer) ParseScheduleCreated(log types.Log) (*VestingManagerScheduleCreated, error) {
	event := new(VestingManagerScheduleCreated)
	if err := _VestingManager.contract.UnpackLog(event, "ScheduleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VestingManager contract.
type VestingManagerUnpausedIterator struct {
	Event *VestingManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *VestingManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerUnpaused)
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
		it.Event = new(VestingManagerUnpaused)
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
func (it *VestingManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerUnpaused represents a Unpaused event raised by the VestingManager contract.
type VestingManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VestingManager *VestingManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VestingManagerUnpausedIterator, error) {

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VestingManagerUnpausedIterator{contract: _VestingManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VestingManager *VestingManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VestingManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerUnpaused)
				if err := _VestingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VestingManager *VestingManagerFilterer) ParseUnpaused(log types.Log) (*VestingManagerUnpaused, error) {
	event := new(VestingManagerUnpaused)
	if err := _VestingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VestingManagerVestedIterator is returned from FilterVested and is used to iterate over the raw logs and unpacked data for Vested events raised by the VestingManager contract.
type VestingManagerVestedIterator struct {
	Event *VestingManagerVested // Event containing the contract specifics and raw log

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
func (it *VestingManagerVestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingManagerVested)
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
		it.Event = new(VestingManagerVested)
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
func (it *VestingManagerVestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingManagerVestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingManagerVested represents a Vested event raised by the VestingManager contract.
type VestingManagerVested struct {
	Id          *big.Int
	TermIndex   *big.Int
	PeriodIdx   *big.Int
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVested is a free log retrieval operation binding the contract event 0x70636b29f83d92d8138de80f0a8a8930db90115322c692b3e570a75d44d4b854.
//
// Solidity: event Vested(uint256 indexed id, uint256 indexed termIndex, uint256 indexed periodIdx, address beneficiary, uint256 amount)
func (_VestingManager *VestingManagerFilterer) FilterVested(opts *bind.FilterOpts, id []*big.Int, termIndex []*big.Int, periodIdx []*big.Int) (*VestingManagerVestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var termIndexRule []interface{}
	for _, termIndexItem := range termIndex {
		termIndexRule = append(termIndexRule, termIndexItem)
	}
	var periodIdxRule []interface{}
	for _, periodIdxItem := range periodIdx {
		periodIdxRule = append(periodIdxRule, periodIdxItem)
	}

	logs, sub, err := _VestingManager.contract.FilterLogs(opts, "Vested", idRule, termIndexRule, periodIdxRule)
	if err != nil {
		return nil, err
	}
	return &VestingManagerVestedIterator{contract: _VestingManager.contract, event: "Vested", logs: logs, sub: sub}, nil
}

// WatchVested is a free log subscription operation binding the contract event 0x70636b29f83d92d8138de80f0a8a8930db90115322c692b3e570a75d44d4b854.
//
// Solidity: event Vested(uint256 indexed id, uint256 indexed termIndex, uint256 indexed periodIdx, address beneficiary, uint256 amount)
func (_VestingManager *VestingManagerFilterer) WatchVested(opts *bind.WatchOpts, sink chan<- *VestingManagerVested, id []*big.Int, termIndex []*big.Int, periodIdx []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var termIndexRule []interface{}
	for _, termIndexItem := range termIndex {
		termIndexRule = append(termIndexRule, termIndexItem)
	}
	var periodIdxRule []interface{}
	for _, periodIdxItem := range periodIdx {
		periodIdxRule = append(periodIdxRule, periodIdxItem)
	}

	logs, sub, err := _VestingManager.contract.WatchLogs(opts, "Vested", idRule, termIndexRule, periodIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingManagerVested)
				if err := _VestingManager.contract.UnpackLog(event, "Vested", log); err != nil {
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

// ParseVested is a log parse operation binding the contract event 0x70636b29f83d92d8138de80f0a8a8930db90115322c692b3e570a75d44d4b854.
//
// Solidity: event Vested(uint256 indexed id, uint256 indexed termIndex, uint256 indexed periodIdx, address beneficiary, uint256 amount)
func (_VestingManager *VestingManagerFilterer) ParseVested(log types.Log) (*VestingManagerVested, error) {
	event := new(VestingManagerVested)
	if err := _VestingManager.contract.UnpackLog(event, "Vested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
