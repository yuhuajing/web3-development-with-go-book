// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stonedrop

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

// StonesDropItem is an auto generated low-level Go binding around an user-defined struct.
type StonesDropItem struct {
	StartTokenId *big.Int
	EndTokenId   *big.Int
	Collection   common.Address
}

// StonesDropTier is an auto generated low-level Go binding around an user-defined struct.
type StonesDropTier struct {
	Available uint16
	Amount    *big.Int
	Rewards   uint16
}

// StonedropMetaData contains all meta data concerning the Stonedrop contract.
var StonedropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tier\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"}],\"name\":\"MysteryBoxDropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractICollectionV3\",\"name\":\"collection\",\"type\":\"address\"}],\"internalType\":\"structStonesDrop.Item[]\",\"name\":\"_items\",\"type\":\"tuple[]\"}],\"name\":\"addItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bought\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"quantity\",\"type\":\"uint8\"}],\"name\":\"buyMysteryBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"}],\"name\":\"getAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_tier\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"giveAway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractICollectionV3\",\"name\":\"collection\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preCompute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stone\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"setGeneralDropParamaters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nftHolderAddresses\",\"type\":\"address[]\"}],\"name\":\"setNftHolderAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"available\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rewards\",\"type\":\"uint16\"}],\"internalType\":\"structStonesDrop.Tier[]\",\"name\":\"_tiers\",\"type\":\"tuple[]\"}],\"name\":\"setTiersParamaters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stone\",\"outputs\":[{\"internalType\":\"contractIFarm\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tiers\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"available\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rewards\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StonedropABI is the input ABI used to generate the binding from.
// Deprecated: Use StonedropMetaData.ABI instead.
var StonedropABI = StonedropMetaData.ABI

// Stonedrop is an auto generated Go binding around an Ethereum contract.
type Stonedrop struct {
	StonedropCaller     // Read-only binding to the contract
	StonedropTransactor // Write-only binding to the contract
	StonedropFilterer   // Log filterer for contract events
}

// StonedropCaller is an auto generated read-only Go binding around an Ethereum contract.
type StonedropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StonedropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StonedropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StonedropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StonedropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StonedropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StonedropSession struct {
	Contract     *Stonedrop        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StonedropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StonedropCallerSession struct {
	Contract *StonedropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StonedropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StonedropTransactorSession struct {
	Contract     *StonedropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StonedropRaw is an auto generated low-level Go binding around an Ethereum contract.
type StonedropRaw struct {
	Contract *Stonedrop // Generic contract binding to access the raw methods on
}

// StonedropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StonedropCallerRaw struct {
	Contract *StonedropCaller // Generic read-only contract binding to access the raw methods on
}

// StonedropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StonedropTransactorRaw struct {
	Contract *StonedropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStonedrop creates a new instance of Stonedrop, bound to a specific deployed contract.
func NewStonedrop(address common.Address, backend bind.ContractBackend) (*Stonedrop, error) {
	contract, err := bindStonedrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stonedrop{StonedropCaller: StonedropCaller{contract: contract}, StonedropTransactor: StonedropTransactor{contract: contract}, StonedropFilterer: StonedropFilterer{contract: contract}}, nil
}

// NewStonedropCaller creates a new read-only instance of Stonedrop, bound to a specific deployed contract.
func NewStonedropCaller(address common.Address, caller bind.ContractCaller) (*StonedropCaller, error) {
	contract, err := bindStonedrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StonedropCaller{contract: contract}, nil
}

// NewStonedropTransactor creates a new write-only instance of Stonedrop, bound to a specific deployed contract.
func NewStonedropTransactor(address common.Address, transactor bind.ContractTransactor) (*StonedropTransactor, error) {
	contract, err := bindStonedrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StonedropTransactor{contract: contract}, nil
}

// NewStonedropFilterer creates a new log filterer instance of Stonedrop, bound to a specific deployed contract.
func NewStonedropFilterer(address common.Address, filterer bind.ContractFilterer) (*StonedropFilterer, error) {
	contract, err := bindStonedrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StonedropFilterer{contract: contract}, nil
}

// bindStonedrop binds a generic wrapper to an already deployed contract.
func bindStonedrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StonedropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stonedrop *StonedropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stonedrop.Contract.StonedropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stonedrop *StonedropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stonedrop.Contract.StonedropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stonedrop *StonedropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stonedrop.Contract.StonedropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stonedrop *StonedropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stonedrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stonedrop *StonedropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stonedrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stonedrop *StonedropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stonedrop.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stonedrop *StonedropCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stonedrop *StonedropSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stonedrop.Contract.DEFAULTADMINROLE(&_Stonedrop.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stonedrop *StonedropCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stonedrop.Contract.DEFAULTADMINROLE(&_Stonedrop.CallOpts)
}

// Bought is a free data retrieval call binding the contract method 0x667022fd.
//
// Solidity: function bought(address ) view returns(uint256 amount, uint32 version)
func (_Stonedrop *StonedropCaller) Bought(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount  *big.Int
	Version uint32
}, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "bought", arg0)

	outstruct := new(struct {
		Amount  *big.Int
		Version uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Bought is a free data retrieval call binding the contract method 0x667022fd.
//
// Solidity: function bought(address ) view returns(uint256 amount, uint32 version)
func (_Stonedrop *StonedropSession) Bought(arg0 common.Address) (struct {
	Amount  *big.Int
	Version uint32
}, error) {
	return _Stonedrop.Contract.Bought(&_Stonedrop.CallOpts, arg0)
}

// Bought is a free data retrieval call binding the contract method 0x667022fd.
//
// Solidity: function bought(address ) view returns(uint256 amount, uint32 version)
func (_Stonedrop *StonedropCallerSession) Bought(arg0 common.Address) (struct {
	Amount  *big.Int
	Version uint32
}, error) {
	return _Stonedrop.Contract.Bought(&_Stonedrop.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Stonedrop *StonedropCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Stonedrop *StonedropSession) EndTime() (*big.Int, error) {
	return _Stonedrop.Contract.EndTime(&_Stonedrop.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Stonedrop *StonedropCallerSession) EndTime() (*big.Int, error) {
	return _Stonedrop.Contract.EndTime(&_Stonedrop.CallOpts)
}

// GetAvailable is a free data retrieval call binding the contract method 0x1e1a0c0c.
//
// Solidity: function getAvailable(uint8 _tier) view returns(uint256)
func (_Stonedrop *StonedropCaller) GetAvailable(opts *bind.CallOpts, _tier uint8) (*big.Int, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "getAvailable", _tier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailable is a free data retrieval call binding the contract method 0x1e1a0c0c.
//
// Solidity: function getAvailable(uint8 _tier) view returns(uint256)
func (_Stonedrop *StonedropSession) GetAvailable(_tier uint8) (*big.Int, error) {
	return _Stonedrop.Contract.GetAvailable(&_Stonedrop.CallOpts, _tier)
}

// GetAvailable is a free data retrieval call binding the contract method 0x1e1a0c0c.
//
// Solidity: function getAvailable(uint8 _tier) view returns(uint256)
func (_Stonedrop *StonedropCallerSession) GetAvailable(_tier uint8) (*big.Int, error) {
	return _Stonedrop.Contract.GetAvailable(&_Stonedrop.CallOpts, _tier)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stonedrop *StonedropCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stonedrop *StonedropSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stonedrop.Contract.GetRoleAdmin(&_Stonedrop.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stonedrop *StonedropCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stonedrop.Contract.GetRoleAdmin(&_Stonedrop.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stonedrop *StonedropCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stonedrop *StonedropSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stonedrop.Contract.HasRole(&_Stonedrop.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stonedrop *StonedropCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stonedrop.Contract.HasRole(&_Stonedrop.CallOpts, role, account)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 startTokenId, uint256 endTokenId, address collection)
func (_Stonedrop *StonedropCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTokenId *big.Int
	EndTokenId   *big.Int
	Collection   common.Address
}, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "items", arg0)

	outstruct := new(struct {
		StartTokenId *big.Int
		EndTokenId   *big.Int
		Collection   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Collection = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 startTokenId, uint256 endTokenId, address collection)
func (_Stonedrop *StonedropSession) Items(arg0 *big.Int) (struct {
	StartTokenId *big.Int
	EndTokenId   *big.Int
	Collection   common.Address
}, error) {
	return _Stonedrop.Contract.Items(&_Stonedrop.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 startTokenId, uint256 endTokenId, address collection)
func (_Stonedrop *StonedropCallerSession) Items(arg0 *big.Int) (struct {
	StartTokenId *big.Int
	EndTokenId   *big.Int
	Collection   common.Address
}, error) {
	return _Stonedrop.Contract.Items(&_Stonedrop.CallOpts, arg0)
}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint256)
func (_Stonedrop *StonedropCaller) MaxPerWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "maxPerWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint256)
func (_Stonedrop *StonedropSession) MaxPerWallet() (*big.Int, error) {
	return _Stonedrop.Contract.MaxPerWallet(&_Stonedrop.CallOpts)
}

// MaxPerWallet is a free data retrieval call binding the contract method 0x453c2310.
//
// Solidity: function maxPerWallet() view returns(uint256)
func (_Stonedrop *StonedropCallerSession) MaxPerWallet() (*big.Int, error) {
	return _Stonedrop.Contract.MaxPerWallet(&_Stonedrop.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Stonedrop *StonedropCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Stonedrop *StonedropSession) StartTime() (*big.Int, error) {
	return _Stonedrop.Contract.StartTime(&_Stonedrop.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Stonedrop *StonedropCallerSession) StartTime() (*big.Int, error) {
	return _Stonedrop.Contract.StartTime(&_Stonedrop.CallOpts)
}

// Stone is a free data retrieval call binding the contract method 0x0167eb85.
//
// Solidity: function stone() view returns(address)
func (_Stonedrop *StonedropCaller) Stone(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "stone")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stone is a free data retrieval call binding the contract method 0x0167eb85.
//
// Solidity: function stone() view returns(address)
func (_Stonedrop *StonedropSession) Stone() (common.Address, error) {
	return _Stonedrop.Contract.Stone(&_Stonedrop.CallOpts)
}

// Stone is a free data retrieval call binding the contract method 0x0167eb85.
//
// Solidity: function stone() view returns(address)
func (_Stonedrop *StonedropCallerSession) Stone() (common.Address, error) {
	return _Stonedrop.Contract.Stone(&_Stonedrop.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stonedrop *StonedropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stonedrop *StonedropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stonedrop.Contract.SupportsInterface(&_Stonedrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stonedrop *StonedropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stonedrop.Contract.SupportsInterface(&_Stonedrop.CallOpts, interfaceId)
}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint16 available, uint256 amount, uint16 rewards)
func (_Stonedrop *StonedropCaller) Tiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Available uint16
	Amount    *big.Int
	Rewards   uint16
}, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "tiers", arg0)

	outstruct := new(struct {
		Available uint16
		Amount    *big.Int
		Rewards   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Available = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint16 available, uint256 amount, uint16 rewards)
func (_Stonedrop *StonedropSession) Tiers(arg0 *big.Int) (struct {
	Available uint16
	Amount    *big.Int
	Rewards   uint16
}, error) {
	return _Stonedrop.Contract.Tiers(&_Stonedrop.CallOpts, arg0)
}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint16 available, uint256 amount, uint16 rewards)
func (_Stonedrop *StonedropCallerSession) Tiers(arg0 *big.Int) (struct {
	Available uint16
	Amount    *big.Int
	Rewards   uint16
}, error) {
	return _Stonedrop.Contract.Tiers(&_Stonedrop.CallOpts, arg0)
}

// Tree is a free data retrieval call binding the contract method 0x31352936.
//
// Solidity: function tree(uint256 ) view returns(uint16)
func (_Stonedrop *StonedropCaller) Tree(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Stonedrop.contract.Call(opts, &out, "tree", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0x31352936.
//
// Solidity: function tree(uint256 ) view returns(uint16)
func (_Stonedrop *StonedropSession) Tree(arg0 *big.Int) (uint16, error) {
	return _Stonedrop.Contract.Tree(&_Stonedrop.CallOpts, arg0)
}

// Tree is a free data retrieval call binding the contract method 0x31352936.
//
// Solidity: function tree(uint256 ) view returns(uint16)
func (_Stonedrop *StonedropCallerSession) Tree(arg0 *big.Int) (uint16, error) {
	return _Stonedrop.Contract.Tree(&_Stonedrop.CallOpts, arg0)
}

// AddItems is a paid mutator transaction binding the contract method 0x457110e8.
//
// Solidity: function addItems((uint256,uint256,address)[] _items) returns()
func (_Stonedrop *StonedropTransactor) AddItems(opts *bind.TransactOpts, _items []StonesDropItem) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "addItems", _items)
}

// AddItems is a paid mutator transaction binding the contract method 0x457110e8.
//
// Solidity: function addItems((uint256,uint256,address)[] _items) returns()
func (_Stonedrop *StonedropSession) AddItems(_items []StonesDropItem) (*types.Transaction, error) {
	return _Stonedrop.Contract.AddItems(&_Stonedrop.TransactOpts, _items)
}

// AddItems is a paid mutator transaction binding the contract method 0x457110e8.
//
// Solidity: function addItems((uint256,uint256,address)[] _items) returns()
func (_Stonedrop *StonedropTransactorSession) AddItems(_items []StonesDropItem) (*types.Transaction, error) {
	return _Stonedrop.Contract.AddItems(&_Stonedrop.TransactOpts, _items)
}

// BuyMysteryBox is a paid mutator transaction binding the contract method 0x142cbfe4.
//
// Solidity: function buyMysteryBox(uint8 _tier, uint8 quantity) returns()
func (_Stonedrop *StonedropTransactor) BuyMysteryBox(opts *bind.TransactOpts, _tier uint8, quantity uint8) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "buyMysteryBox", _tier, quantity)
}

// BuyMysteryBox is a paid mutator transaction binding the contract method 0x142cbfe4.
//
// Solidity: function buyMysteryBox(uint8 _tier, uint8 quantity) returns()
func (_Stonedrop *StonedropSession) BuyMysteryBox(_tier uint8, quantity uint8) (*types.Transaction, error) {
	return _Stonedrop.Contract.BuyMysteryBox(&_Stonedrop.TransactOpts, _tier, quantity)
}

// BuyMysteryBox is a paid mutator transaction binding the contract method 0x142cbfe4.
//
// Solidity: function buyMysteryBox(uint8 _tier, uint8 quantity) returns()
func (_Stonedrop *StonedropTransactorSession) BuyMysteryBox(_tier uint8, quantity uint8) (*types.Transaction, error) {
	return _Stonedrop.Contract.BuyMysteryBox(&_Stonedrop.TransactOpts, _tier, quantity)
}

// GiveAway is a paid mutator transaction binding the contract method 0xe379be89.
//
// Solidity: function giveAway(uint8 _tier, address _user) returns()
func (_Stonedrop *StonedropTransactor) GiveAway(opts *bind.TransactOpts, _tier uint8, _user common.Address) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "giveAway", _tier, _user)
}

// GiveAway is a paid mutator transaction binding the contract method 0xe379be89.
//
// Solidity: function giveAway(uint8 _tier, address _user) returns()
func (_Stonedrop *StonedropSession) GiveAway(_tier uint8, _user common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.GiveAway(&_Stonedrop.TransactOpts, _tier, _user)
}

// GiveAway is a paid mutator transaction binding the contract method 0xe379be89.
//
// Solidity: function giveAway(uint8 _tier, address _user) returns()
func (_Stonedrop *StonedropTransactorSession) GiveAway(_tier uint8, _user common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.GiveAway(&_Stonedrop.TransactOpts, _tier, _user)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.GrantRole(&_Stonedrop.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.GrantRole(&_Stonedrop.TransactOpts, role, account)
}

// PreCompute is a paid mutator transaction binding the contract method 0xdc317550.
//
// Solidity: function preCompute() returns()
func (_Stonedrop *StonedropTransactor) PreCompute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "preCompute")
}

// PreCompute is a paid mutator transaction binding the contract method 0xdc317550.
//
// Solidity: function preCompute() returns()
func (_Stonedrop *StonedropSession) PreCompute() (*types.Transaction, error) {
	return _Stonedrop.Contract.PreCompute(&_Stonedrop.TransactOpts)
}

// PreCompute is a paid mutator transaction binding the contract method 0xdc317550.
//
// Solidity: function preCompute() returns()
func (_Stonedrop *StonedropTransactorSession) PreCompute() (*types.Transaction, error) {
	return _Stonedrop.Contract.PreCompute(&_Stonedrop.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.RenounceRole(&_Stonedrop.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.RenounceRole(&_Stonedrop.TransactOpts, role, account)
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_Stonedrop *StonedropTransactor) ResetItems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "resetItems")
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_Stonedrop *StonedropSession) ResetItems() (*types.Transaction, error) {
	return _Stonedrop.Contract.ResetItems(&_Stonedrop.TransactOpts)
}

// ResetItems is a paid mutator transaction binding the contract method 0xacec599e.
//
// Solidity: function resetItems() returns()
func (_Stonedrop *StonedropTransactorSession) ResetItems() (*types.Transaction, error) {
	return _Stonedrop.Contract.ResetItems(&_Stonedrop.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.RevokeRole(&_Stonedrop.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stonedrop *StonedropTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.RevokeRole(&_Stonedrop.TransactOpts, role, account)
}

// SetGeneralDropParamaters is a paid mutator transaction binding the contract method 0xd5f1c5aa.
//
// Solidity: function setGeneralDropParamaters(address _stone, uint256 _maxPerWallet, uint256 _startTime, uint256 _endTime) returns()
func (_Stonedrop *StonedropTransactor) SetGeneralDropParamaters(opts *bind.TransactOpts, _stone common.Address, _maxPerWallet *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "setGeneralDropParamaters", _stone, _maxPerWallet, _startTime, _endTime)
}

// SetGeneralDropParamaters is a paid mutator transaction binding the contract method 0xd5f1c5aa.
//
// Solidity: function setGeneralDropParamaters(address _stone, uint256 _maxPerWallet, uint256 _startTime, uint256 _endTime) returns()
func (_Stonedrop *StonedropSession) SetGeneralDropParamaters(_stone common.Address, _maxPerWallet *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetGeneralDropParamaters(&_Stonedrop.TransactOpts, _stone, _maxPerWallet, _startTime, _endTime)
}

// SetGeneralDropParamaters is a paid mutator transaction binding the contract method 0xd5f1c5aa.
//
// Solidity: function setGeneralDropParamaters(address _stone, uint256 _maxPerWallet, uint256 _startTime, uint256 _endTime) returns()
func (_Stonedrop *StonedropTransactorSession) SetGeneralDropParamaters(_stone common.Address, _maxPerWallet *big.Int, _startTime *big.Int, _endTime *big.Int) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetGeneralDropParamaters(&_Stonedrop.TransactOpts, _stone, _maxPerWallet, _startTime, _endTime)
}

// SetNftHolderAddresses is a paid mutator transaction binding the contract method 0x3c6f3829.
//
// Solidity: function setNftHolderAddresses(address[] _nftHolderAddresses) returns()
func (_Stonedrop *StonedropTransactor) SetNftHolderAddresses(opts *bind.TransactOpts, _nftHolderAddresses []common.Address) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "setNftHolderAddresses", _nftHolderAddresses)
}

// SetNftHolderAddresses is a paid mutator transaction binding the contract method 0x3c6f3829.
//
// Solidity: function setNftHolderAddresses(address[] _nftHolderAddresses) returns()
func (_Stonedrop *StonedropSession) SetNftHolderAddresses(_nftHolderAddresses []common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetNftHolderAddresses(&_Stonedrop.TransactOpts, _nftHolderAddresses)
}

// SetNftHolderAddresses is a paid mutator transaction binding the contract method 0x3c6f3829.
//
// Solidity: function setNftHolderAddresses(address[] _nftHolderAddresses) returns()
func (_Stonedrop *StonedropTransactorSession) SetNftHolderAddresses(_nftHolderAddresses []common.Address) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetNftHolderAddresses(&_Stonedrop.TransactOpts, _nftHolderAddresses)
}

// SetTiersParamaters is a paid mutator transaction binding the contract method 0x4196d0b8.
//
// Solidity: function setTiersParamaters((uint16,uint256,uint16)[] _tiers) returns()
func (_Stonedrop *StonedropTransactor) SetTiersParamaters(opts *bind.TransactOpts, _tiers []StonesDropTier) (*types.Transaction, error) {
	return _Stonedrop.contract.Transact(opts, "setTiersParamaters", _tiers)
}

// SetTiersParamaters is a paid mutator transaction binding the contract method 0x4196d0b8.
//
// Solidity: function setTiersParamaters((uint16,uint256,uint16)[] _tiers) returns()
func (_Stonedrop *StonedropSession) SetTiersParamaters(_tiers []StonesDropTier) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetTiersParamaters(&_Stonedrop.TransactOpts, _tiers)
}

// SetTiersParamaters is a paid mutator transaction binding the contract method 0x4196d0b8.
//
// Solidity: function setTiersParamaters((uint16,uint256,uint16)[] _tiers) returns()
func (_Stonedrop *StonedropTransactorSession) SetTiersParamaters(_tiers []StonesDropTier) (*types.Transaction, error) {
	return _Stonedrop.Contract.SetTiersParamaters(&_Stonedrop.TransactOpts, _tiers)
}

// StonedropMysteryBoxDroppedIterator is returned from FilterMysteryBoxDropped and is used to iterate over the raw logs and unpacked data for MysteryBoxDropped events raised by the Stonedrop contract.
type StonedropMysteryBoxDroppedIterator struct {
	Event *StonedropMysteryBoxDropped // Event containing the contract specifics and raw log

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
func (it *StonedropMysteryBoxDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StonedropMysteryBoxDropped)
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
		it.Event = new(StonedropMysteryBoxDropped)
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
func (it *StonedropMysteryBoxDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StonedropMysteryBoxDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StonedropMysteryBoxDropped represents a MysteryBoxDropped event raised by the Stonedrop contract.
type StonedropMysteryBoxDropped struct {
	Tier       uint8
	Collection common.Address
	Id         *big.Int
	User       common.Address
	Version    uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMysteryBoxDropped is a free log retrieval operation binding the contract event 0xb48a3f4dab3b7d9c6a59f8fd2b97e1336ce85b34bc2c9bc96b59f964cfa7b0c3.
//
// Solidity: event MysteryBoxDropped(uint8 tier, address collection, uint256 id, address user, uint32 version)
func (_Stonedrop *StonedropFilterer) FilterMysteryBoxDropped(opts *bind.FilterOpts) (*StonedropMysteryBoxDroppedIterator, error) {

	logs, sub, err := _Stonedrop.contract.FilterLogs(opts, "MysteryBoxDropped")
	if err != nil {
		return nil, err
	}
	return &StonedropMysteryBoxDroppedIterator{contract: _Stonedrop.contract, event: "MysteryBoxDropped", logs: logs, sub: sub}, nil
}

// WatchMysteryBoxDropped is a free log subscription operation binding the contract event 0xb48a3f4dab3b7d9c6a59f8fd2b97e1336ce85b34bc2c9bc96b59f964cfa7b0c3.
//
// Solidity: event MysteryBoxDropped(uint8 tier, address collection, uint256 id, address user, uint32 version)
func (_Stonedrop *StonedropFilterer) WatchMysteryBoxDropped(opts *bind.WatchOpts, sink chan<- *StonedropMysteryBoxDropped) (event.Subscription, error) {

	logs, sub, err := _Stonedrop.contract.WatchLogs(opts, "MysteryBoxDropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StonedropMysteryBoxDropped)
				if err := _Stonedrop.contract.UnpackLog(event, "MysteryBoxDropped", log); err != nil {
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

// ParseMysteryBoxDropped is a log parse operation binding the contract event 0xb48a3f4dab3b7d9c6a59f8fd2b97e1336ce85b34bc2c9bc96b59f964cfa7b0c3.
//
// Solidity: event MysteryBoxDropped(uint8 tier, address collection, uint256 id, address user, uint32 version)
func (_Stonedrop *StonedropFilterer) ParseMysteryBoxDropped(log types.Log) (*StonedropMysteryBoxDropped, error) {
	event := new(StonedropMysteryBoxDropped)
	if err := _Stonedrop.contract.UnpackLog(event, "MysteryBoxDropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StonedropRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stonedrop contract.
type StonedropRoleAdminChangedIterator struct {
	Event *StonedropRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StonedropRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StonedropRoleAdminChanged)
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
		it.Event = new(StonedropRoleAdminChanged)
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
func (it *StonedropRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StonedropRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StonedropRoleAdminChanged represents a RoleAdminChanged event raised by the Stonedrop contract.
type StonedropRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stonedrop *StonedropFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StonedropRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Stonedrop.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StonedropRoleAdminChangedIterator{contract: _Stonedrop.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stonedrop *StonedropFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StonedropRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Stonedrop.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StonedropRoleAdminChanged)
				if err := _Stonedrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stonedrop *StonedropFilterer) ParseRoleAdminChanged(log types.Log) (*StonedropRoleAdminChanged, error) {
	event := new(StonedropRoleAdminChanged)
	if err := _Stonedrop.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StonedropRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stonedrop contract.
type StonedropRoleGrantedIterator struct {
	Event *StonedropRoleGranted // Event containing the contract specifics and raw log

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
func (it *StonedropRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StonedropRoleGranted)
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
		it.Event = new(StonedropRoleGranted)
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
func (it *StonedropRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StonedropRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StonedropRoleGranted represents a RoleGranted event raised by the Stonedrop contract.
type StonedropRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StonedropRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stonedrop.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StonedropRoleGrantedIterator{contract: _Stonedrop.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StonedropRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stonedrop.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StonedropRoleGranted)
				if err := _Stonedrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) ParseRoleGranted(log types.Log) (*StonedropRoleGranted, error) {
	event := new(StonedropRoleGranted)
	if err := _Stonedrop.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StonedropRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stonedrop contract.
type StonedropRoleRevokedIterator struct {
	Event *StonedropRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StonedropRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StonedropRoleRevoked)
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
		it.Event = new(StonedropRoleRevoked)
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
func (it *StonedropRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StonedropRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StonedropRoleRevoked represents a RoleRevoked event raised by the Stonedrop contract.
type StonedropRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StonedropRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stonedrop.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StonedropRoleRevokedIterator{contract: _Stonedrop.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StonedropRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Stonedrop.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StonedropRoleRevoked)
				if err := _Stonedrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stonedrop *StonedropFilterer) ParseRoleRevoked(log types.Log) (*StonedropRoleRevoked, error) {
	event := new(StonedropRoleRevoked)
	if err := _Stonedrop.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
