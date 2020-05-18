// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapterRegistry

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

// AdapterBalance is an auto generated low-level Go binding around an user-defined struct.
type AdapterBalance struct {
	Metadata AdapterMetadata
	Balances []FullTokenBalance
}

// AdapterMetadata is an auto generated low-level Go binding around an user-defined struct.
type AdapterMetadata struct {
	AdapterAddress common.Address
	AdapterType    string
}

// FullTokenBalance is an auto generated low-level Go binding around an user-defined struct.
type FullTokenBalance struct {
	Base       TokenBalance
	Underlying []TokenBalance
}

// ProtocolBalance is an auto generated low-level Go binding around an user-defined struct.
type ProtocolBalance struct {
	Metadata        ProtocolMetadata
	AdapterBalances []AdapterBalance
}

// ProtocolMetadata is an auto generated low-level Go binding around an user-defined struct.
type ProtocolMetadata struct {
	Name        string
	Description string
	WebsiteURL  string
	IconURL     string
	Version     *big.Int
}

// TokenBalance is an auto generated low-level Go binding around an user-defined struct.
type TokenBalance struct {
	Metadata TokenMetadata
	Amount   *big.Int
}

// TokenMetadata is an auto generated low-level Go binding around an user-defined struct.
type TokenMetadata struct {
	Token    common.Address
	Name     string
	Symbol   string
	Decimals uint8
}

// AdapterRegistryABI is the input ABI used to generate the binding from.
const AdapterRegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"},{\"internalType\":\"address[][]\",\"name\":\"tokens\",\"type\":\"address[][]\"}],\"name\":\"addProtocolAdapters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"protocolNames\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"iconURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structProtocolMetadata[]\",\"name\":\"metadata\",\"type\":\"tuple[]\"},{\"internalType\":\"address[][]\",\"name\":\"adapters\",\"type\":\"address[][]\"},{\"internalType\":\"address[][][]\",\"name\":\"tokens\",\"type\":\"address[][][]\"}],\"name\":\"addProtocols\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"tokenAdapterNames\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"}],\"name\":\"addTokenAdapters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getAdapterBalance\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"adapterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"adapterType\",\"type\":\"string\"}],\"internalType\":\"structAdapterMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structAdapterBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"adapters\",\"type\":\"address[]\"}],\"name\":\"getAdapterBalances\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"adapterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"adapterType\",\"type\":\"string\"}],\"internalType\":\"structAdapterMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structAdapterBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBalances\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"iconURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structProtocolMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"adapterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"adapterType\",\"type\":\"string\"}],\"internalType\":\"structAdapterMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structAdapterBalance[]\",\"name\":\"adapterBalances\",\"type\":\"tuple[]\"}],\"internalType\":\"structProtocolBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFinalFullTokenBalance\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenType\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFullTokenBalance\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"}],\"name\":\"getProtocolAdapters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"protocolNames\",\"type\":\"string[]\"}],\"name\":\"getProtocolBalances\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"iconURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structProtocolMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"adapterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"adapterType\",\"type\":\"string\"}],\"internalType\":\"structAdapterMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance\",\"name\":\"base\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structTokenMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTokenBalance[]\",\"name\":\"underlying\",\"type\":\"tuple[]\"}],\"internalType\":\"structFullTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"}],\"internalType\":\"structAdapterBalance[]\",\"name\":\"adapterBalances\",\"type\":\"tuple[]\"}],\"internalType\":\"structProtocolBalance[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"}],\"name\":\"getProtocolMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"iconURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"internalType\":\"structProtocolMetadata\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAdapterName\",\"type\":\"string\"}],\"name\":\"getTokenAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAdapterNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"}],\"name\":\"isValidProtocol\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAdapterName\",\"type\":\"string\"}],\"name\":\"isValidTokenAdapter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"adapterIndices\",\"type\":\"uint256[]\"}],\"name\":\"removeProtocolAdapters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"protocolNames\",\"type\":\"string[]\"}],\"name\":\"removeProtocols\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"tokenAdapterNames\",\"type\":\"string[]\"}],\"name\":\"removeTokenAdapters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newAdapterAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"newSupportedTokens\",\"type\":\"address[]\"}],\"name\":\"updateProtocolAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"protocolName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"iconURL\",\"type\":\"string\"}],\"name\":\"updateProtocolMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenAdapterName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"updateTokenAdapter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdapterRegistry is an auto generated Go binding around an Ethereum contract.
type AdapterRegistry struct {
	AdapterRegistryCaller     // Read-only binding to the contract
	AdapterRegistryTransactor // Write-only binding to the contract
	AdapterRegistryFilterer   // Log filterer for contract events
}

// AdapterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdapterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdapterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdapterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdapterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdapterRegistrySession struct {
	Contract     *AdapterRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdapterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdapterRegistryCallerSession struct {
	Contract *AdapterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AdapterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdapterRegistryTransactorSession struct {
	Contract     *AdapterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AdapterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdapterRegistryRaw struct {
	Contract *AdapterRegistry // Generic contract binding to access the raw methods on
}

// AdapterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdapterRegistryCallerRaw struct {
	Contract *AdapterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AdapterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdapterRegistryTransactorRaw struct {
	Contract *AdapterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdapterRegistry creates a new instance of AdapterRegistry, bound to a specific deployed contract.
func NewAdapterRegistry(address common.Address, backend bind.ContractBackend) (*AdapterRegistry, error) {
	contract, err := bindAdapterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdapterRegistry{AdapterRegistryCaller: AdapterRegistryCaller{contract: contract}, AdapterRegistryTransactor: AdapterRegistryTransactor{contract: contract}, AdapterRegistryFilterer: AdapterRegistryFilterer{contract: contract}}, nil
}

// NewAdapterRegistryCaller creates a new read-only instance of AdapterRegistry, bound to a specific deployed contract.
func NewAdapterRegistryCaller(address common.Address, caller bind.ContractCaller) (*AdapterRegistryCaller, error) {
	contract, err := bindAdapterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterRegistryCaller{contract: contract}, nil
}

// NewAdapterRegistryTransactor creates a new write-only instance of AdapterRegistry, bound to a specific deployed contract.
func NewAdapterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AdapterRegistryTransactor, error) {
	contract, err := bindAdapterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdapterRegistryTransactor{contract: contract}, nil
}

// NewAdapterRegistryFilterer creates a new log filterer instance of AdapterRegistry, bound to a specific deployed contract.
func NewAdapterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AdapterRegistryFilterer, error) {
	contract, err := bindAdapterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdapterRegistryFilterer{contract: contract}, nil
}

// bindAdapterRegistry binds a generic wrapper to an already deployed contract.
func bindAdapterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdapterRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdapterRegistry *AdapterRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdapterRegistry.Contract.AdapterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdapterRegistry *AdapterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AdapterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdapterRegistry *AdapterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AdapterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdapterRegistry *AdapterRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdapterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdapterRegistry *AdapterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdapterRegistry *AdapterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAdapterBalance is a free data retrieval call binding the contract method 0x4bf9649b.
//
// Solidity: function getAdapterBalance(address account, address adapter, address[] tokens) view returns(AdapterBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetAdapterBalance(opts *bind.CallOpts, account common.Address, adapter common.Address, tokens []common.Address) (AdapterBalance, error) {
	var (
		ret0 = new(AdapterBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getAdapterBalance", account, adapter, tokens)
	return *ret0, err
}

// GetAdapterBalance is a free data retrieval call binding the contract method 0x4bf9649b.
//
// Solidity: function getAdapterBalance(address account, address adapter, address[] tokens) view returns(AdapterBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetAdapterBalance(account common.Address, adapter common.Address, tokens []common.Address) (AdapterBalance, error) {
	return _AdapterRegistry.Contract.GetAdapterBalance(&_AdapterRegistry.CallOpts, account, adapter, tokens)
}

// GetAdapterBalance is a free data retrieval call binding the contract method 0x4bf9649b.
//
// Solidity: function getAdapterBalance(address account, address adapter, address[] tokens) view returns(AdapterBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetAdapterBalance(account common.Address, adapter common.Address, tokens []common.Address) (AdapterBalance, error) {
	return _AdapterRegistry.Contract.GetAdapterBalance(&_AdapterRegistry.CallOpts, account, adapter, tokens)
}

// GetAdapterBalances is a free data retrieval call binding the contract method 0xa2aef46d.
//
// Solidity: function getAdapterBalances(address account, address[] adapters) view returns([]AdapterBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetAdapterBalances(opts *bind.CallOpts, account common.Address, adapters []common.Address) ([]AdapterBalance, error) {
	var (
		ret0 = new([]AdapterBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getAdapterBalances", account, adapters)
	return *ret0, err
}

// GetAdapterBalances is a free data retrieval call binding the contract method 0xa2aef46d.
//
// Solidity: function getAdapterBalances(address account, address[] adapters) view returns([]AdapterBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetAdapterBalances(account common.Address, adapters []common.Address) ([]AdapterBalance, error) {
	return _AdapterRegistry.Contract.GetAdapterBalances(&_AdapterRegistry.CallOpts, account, adapters)
}

// GetAdapterBalances is a free data retrieval call binding the contract method 0xa2aef46d.
//
// Solidity: function getAdapterBalances(address account, address[] adapters) view returns([]AdapterBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetAdapterBalances(account common.Address, adapters []common.Address) ([]AdapterBalance, error) {
	return _AdapterRegistry.Contract.GetAdapterBalances(&_AdapterRegistry.CallOpts, account, adapters)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address account) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetBalances(opts *bind.CallOpts, account common.Address) ([]ProtocolBalance, error) {
	var (
		ret0 = new([]ProtocolBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getBalances", account)
	return *ret0, err
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address account) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetBalances(account common.Address) ([]ProtocolBalance, error) {
	return _AdapterRegistry.Contract.GetBalances(&_AdapterRegistry.CallOpts, account)
}

// GetBalances is a free data retrieval call binding the contract method 0xc84aae17.
//
// Solidity: function getBalances(address account) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetBalances(account common.Address) ([]ProtocolBalance, error) {
	return _AdapterRegistry.Contract.GetBalances(&_AdapterRegistry.CallOpts, account)
}

// GetFinalFullTokenBalance is a free data retrieval call binding the contract method 0xa81d9a09.
//
// Solidity: function getFinalFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetFinalFullTokenBalance(opts *bind.CallOpts, tokenType string, token common.Address) (FullTokenBalance, error) {
	var (
		ret0 = new(FullTokenBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getFinalFullTokenBalance", tokenType, token)
	return *ret0, err
}

// GetFinalFullTokenBalance is a free data retrieval call binding the contract method 0xa81d9a09.
//
// Solidity: function getFinalFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetFinalFullTokenBalance(tokenType string, token common.Address) (FullTokenBalance, error) {
	return _AdapterRegistry.Contract.GetFinalFullTokenBalance(&_AdapterRegistry.CallOpts, tokenType, token)
}

// GetFinalFullTokenBalance is a free data retrieval call binding the contract method 0xa81d9a09.
//
// Solidity: function getFinalFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetFinalFullTokenBalance(tokenType string, token common.Address) (FullTokenBalance, error) {
	return _AdapterRegistry.Contract.GetFinalFullTokenBalance(&_AdapterRegistry.CallOpts, tokenType, token)
}

// GetFullTokenBalance is a free data retrieval call binding the contract method 0x56098b38.
//
// Solidity: function getFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetFullTokenBalance(opts *bind.CallOpts, tokenType string, token common.Address) (FullTokenBalance, error) {
	var (
		ret0 = new(FullTokenBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getFullTokenBalance", tokenType, token)
	return *ret0, err
}

// GetFullTokenBalance is a free data retrieval call binding the contract method 0x56098b38.
//
// Solidity: function getFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetFullTokenBalance(tokenType string, token common.Address) (FullTokenBalance, error) {
	return _AdapterRegistry.Contract.GetFullTokenBalance(&_AdapterRegistry.CallOpts, tokenType, token)
}

// GetFullTokenBalance is a free data retrieval call binding the contract method 0x56098b38.
//
// Solidity: function getFullTokenBalance(string tokenType, address token) view returns(FullTokenBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetFullTokenBalance(tokenType string, token common.Address) (FullTokenBalance, error) {
	return _AdapterRegistry.Contract.GetFullTokenBalance(&_AdapterRegistry.CallOpts, tokenType, token)
}

// GetProtocolAdapters is a free data retrieval call binding the contract method 0xeff70322.
//
// Solidity: function getProtocolAdapters(string protocolName) view returns(address[])
func (_AdapterRegistry *AdapterRegistryCaller) GetProtocolAdapters(opts *bind.CallOpts, protocolName string) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getProtocolAdapters", protocolName)
	return *ret0, err
}

// GetProtocolAdapters is a free data retrieval call binding the contract method 0xeff70322.
//
// Solidity: function getProtocolAdapters(string protocolName) view returns(address[])
func (_AdapterRegistry *AdapterRegistrySession) GetProtocolAdapters(protocolName string) ([]common.Address, error) {
	return _AdapterRegistry.Contract.GetProtocolAdapters(&_AdapterRegistry.CallOpts, protocolName)
}

// GetProtocolAdapters is a free data retrieval call binding the contract method 0xeff70322.
//
// Solidity: function getProtocolAdapters(string protocolName) view returns(address[])
func (_AdapterRegistry *AdapterRegistryCallerSession) GetProtocolAdapters(protocolName string) ([]common.Address, error) {
	return _AdapterRegistry.Contract.GetProtocolAdapters(&_AdapterRegistry.CallOpts, protocolName)
}

// GetProtocolBalances is a free data retrieval call binding the contract method 0x85c6a793.
//
// Solidity: function getProtocolBalances(address account, string[] protocolNames) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistryCaller) GetProtocolBalances(opts *bind.CallOpts, account common.Address, protocolNames []string) ([]ProtocolBalance, error) {
	var (
		ret0 = new([]ProtocolBalance)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getProtocolBalances", account, protocolNames)
	return *ret0, err
}

// GetProtocolBalances is a free data retrieval call binding the contract method 0x85c6a793.
//
// Solidity: function getProtocolBalances(address account, string[] protocolNames) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistrySession) GetProtocolBalances(account common.Address, protocolNames []string) ([]ProtocolBalance, error) {
	return _AdapterRegistry.Contract.GetProtocolBalances(&_AdapterRegistry.CallOpts, account, protocolNames)
}

// GetProtocolBalances is a free data retrieval call binding the contract method 0x85c6a793.
//
// Solidity: function getProtocolBalances(address account, string[] protocolNames) view returns([]ProtocolBalance)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetProtocolBalances(account common.Address, protocolNames []string) ([]ProtocolBalance, error) {
	return _AdapterRegistry.Contract.GetProtocolBalances(&_AdapterRegistry.CallOpts, account, protocolNames)
}

// GetProtocolMetadata is a free data retrieval call binding the contract method 0xd857da49.
//
// Solidity: function getProtocolMetadata(string protocolName) view returns(ProtocolMetadata)
func (_AdapterRegistry *AdapterRegistryCaller) GetProtocolMetadata(opts *bind.CallOpts, protocolName string) (ProtocolMetadata, error) {
	var (
		ret0 = new(ProtocolMetadata)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getProtocolMetadata", protocolName)
	return *ret0, err
}

// GetProtocolMetadata is a free data retrieval call binding the contract method 0xd857da49.
//
// Solidity: function getProtocolMetadata(string protocolName) view returns(ProtocolMetadata)
func (_AdapterRegistry *AdapterRegistrySession) GetProtocolMetadata(protocolName string) (ProtocolMetadata, error) {
	return _AdapterRegistry.Contract.GetProtocolMetadata(&_AdapterRegistry.CallOpts, protocolName)
}

// GetProtocolMetadata is a free data retrieval call binding the contract method 0xd857da49.
//
// Solidity: function getProtocolMetadata(string protocolName) view returns(ProtocolMetadata)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetProtocolMetadata(protocolName string) (ProtocolMetadata, error) {
	return _AdapterRegistry.Contract.GetProtocolMetadata(&_AdapterRegistry.CallOpts, protocolName)
}

// GetProtocolNames is a free data retrieval call binding the contract method 0x3b692f52.
//
// Solidity: function getProtocolNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistryCaller) GetProtocolNames(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getProtocolNames")
	return *ret0, err
}

// GetProtocolNames is a free data retrieval call binding the contract method 0x3b692f52.
//
// Solidity: function getProtocolNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistrySession) GetProtocolNames() ([]string, error) {
	return _AdapterRegistry.Contract.GetProtocolNames(&_AdapterRegistry.CallOpts)
}

// GetProtocolNames is a free data retrieval call binding the contract method 0x3b692f52.
//
// Solidity: function getProtocolNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistryCallerSession) GetProtocolNames() ([]string, error) {
	return _AdapterRegistry.Contract.GetProtocolNames(&_AdapterRegistry.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0x07526acf.
//
// Solidity: function getSupportedTokens(address adapter) view returns(address[])
func (_AdapterRegistry *AdapterRegistryCaller) GetSupportedTokens(opts *bind.CallOpts, adapter common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getSupportedTokens", adapter)
	return *ret0, err
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0x07526acf.
//
// Solidity: function getSupportedTokens(address adapter) view returns(address[])
func (_AdapterRegistry *AdapterRegistrySession) GetSupportedTokens(adapter common.Address) ([]common.Address, error) {
	return _AdapterRegistry.Contract.GetSupportedTokens(&_AdapterRegistry.CallOpts, adapter)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0x07526acf.
//
// Solidity: function getSupportedTokens(address adapter) view returns(address[])
func (_AdapterRegistry *AdapterRegistryCallerSession) GetSupportedTokens(adapter common.Address) ([]common.Address, error) {
	return _AdapterRegistry.Contract.GetSupportedTokens(&_AdapterRegistry.CallOpts, adapter)
}

// GetTokenAdapter is a free data retrieval call binding the contract method 0x54fa3382.
//
// Solidity: function getTokenAdapter(string tokenAdapterName) view returns(address)
func (_AdapterRegistry *AdapterRegistryCaller) GetTokenAdapter(opts *bind.CallOpts, tokenAdapterName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getTokenAdapter", tokenAdapterName)
	return *ret0, err
}

// GetTokenAdapter is a free data retrieval call binding the contract method 0x54fa3382.
//
// Solidity: function getTokenAdapter(string tokenAdapterName) view returns(address)
func (_AdapterRegistry *AdapterRegistrySession) GetTokenAdapter(tokenAdapterName string) (common.Address, error) {
	return _AdapterRegistry.Contract.GetTokenAdapter(&_AdapterRegistry.CallOpts, tokenAdapterName)
}

// GetTokenAdapter is a free data retrieval call binding the contract method 0x54fa3382.
//
// Solidity: function getTokenAdapter(string tokenAdapterName) view returns(address)
func (_AdapterRegistry *AdapterRegistryCallerSession) GetTokenAdapter(tokenAdapterName string) (common.Address, error) {
	return _AdapterRegistry.Contract.GetTokenAdapter(&_AdapterRegistry.CallOpts, tokenAdapterName)
}

// GetTokenAdapterNames is a free data retrieval call binding the contract method 0xc0680cd3.
//
// Solidity: function getTokenAdapterNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistryCaller) GetTokenAdapterNames(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "getTokenAdapterNames")
	return *ret0, err
}

// GetTokenAdapterNames is a free data retrieval call binding the contract method 0xc0680cd3.
//
// Solidity: function getTokenAdapterNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistrySession) GetTokenAdapterNames() ([]string, error) {
	return _AdapterRegistry.Contract.GetTokenAdapterNames(&_AdapterRegistry.CallOpts)
}

// GetTokenAdapterNames is a free data retrieval call binding the contract method 0xc0680cd3.
//
// Solidity: function getTokenAdapterNames() view returns(string[])
func (_AdapterRegistry *AdapterRegistryCallerSession) GetTokenAdapterNames() ([]string, error) {
	return _AdapterRegistry.Contract.GetTokenAdapterNames(&_AdapterRegistry.CallOpts)
}

// IsValidProtocol is a free data retrieval call binding the contract method 0xcb874d33.
//
// Solidity: function isValidProtocol(string protocolName) view returns(bool)
func (_AdapterRegistry *AdapterRegistryCaller) IsValidProtocol(opts *bind.CallOpts, protocolName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "isValidProtocol", protocolName)
	return *ret0, err
}

// IsValidProtocol is a free data retrieval call binding the contract method 0xcb874d33.
//
// Solidity: function isValidProtocol(string protocolName) view returns(bool)
func (_AdapterRegistry *AdapterRegistrySession) IsValidProtocol(protocolName string) (bool, error) {
	return _AdapterRegistry.Contract.IsValidProtocol(&_AdapterRegistry.CallOpts, protocolName)
}

// IsValidProtocol is a free data retrieval call binding the contract method 0xcb874d33.
//
// Solidity: function isValidProtocol(string protocolName) view returns(bool)
func (_AdapterRegistry *AdapterRegistryCallerSession) IsValidProtocol(protocolName string) (bool, error) {
	return _AdapterRegistry.Contract.IsValidProtocol(&_AdapterRegistry.CallOpts, protocolName)
}

// IsValidTokenAdapter is a free data retrieval call binding the contract method 0xe9dd5f25.
//
// Solidity: function isValidTokenAdapter(string tokenAdapterName) view returns(bool)
func (_AdapterRegistry *AdapterRegistryCaller) IsValidTokenAdapter(opts *bind.CallOpts, tokenAdapterName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "isValidTokenAdapter", tokenAdapterName)
	return *ret0, err
}

// IsValidTokenAdapter is a free data retrieval call binding the contract method 0xe9dd5f25.
//
// Solidity: function isValidTokenAdapter(string tokenAdapterName) view returns(bool)
func (_AdapterRegistry *AdapterRegistrySession) IsValidTokenAdapter(tokenAdapterName string) (bool, error) {
	return _AdapterRegistry.Contract.IsValidTokenAdapter(&_AdapterRegistry.CallOpts, tokenAdapterName)
}

// IsValidTokenAdapter is a free data retrieval call binding the contract method 0xe9dd5f25.
//
// Solidity: function isValidTokenAdapter(string tokenAdapterName) view returns(bool)
func (_AdapterRegistry *AdapterRegistryCallerSession) IsValidTokenAdapter(tokenAdapterName string) (bool, error) {
	return _AdapterRegistry.Contract.IsValidTokenAdapter(&_AdapterRegistry.CallOpts, tokenAdapterName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdapterRegistry *AdapterRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdapterRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdapterRegistry *AdapterRegistrySession) Owner() (common.Address, error) {
	return _AdapterRegistry.Contract.Owner(&_AdapterRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdapterRegistry *AdapterRegistryCallerSession) Owner() (common.Address, error) {
	return _AdapterRegistry.Contract.Owner(&_AdapterRegistry.CallOpts)
}

// AddProtocolAdapters is a paid mutator transaction binding the contract method 0xbefe43a3.
//
// Solidity: function addProtocolAdapters(string protocolName, address[] adapters, address[][] tokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) AddProtocolAdapters(opts *bind.TransactOpts, protocolName string, adapters []common.Address, tokens [][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "addProtocolAdapters", protocolName, adapters, tokens)
}

// AddProtocolAdapters is a paid mutator transaction binding the contract method 0xbefe43a3.
//
// Solidity: function addProtocolAdapters(string protocolName, address[] adapters, address[][] tokens) returns()
func (_AdapterRegistry *AdapterRegistrySession) AddProtocolAdapters(protocolName string, adapters []common.Address, tokens [][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddProtocolAdapters(&_AdapterRegistry.TransactOpts, protocolName, adapters, tokens)
}

// AddProtocolAdapters is a paid mutator transaction binding the contract method 0xbefe43a3.
//
// Solidity: function addProtocolAdapters(string protocolName, address[] adapters, address[][] tokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) AddProtocolAdapters(protocolName string, adapters []common.Address, tokens [][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddProtocolAdapters(&_AdapterRegistry.TransactOpts, protocolName, adapters, tokens)
}

// AddProtocols is a paid mutator transaction binding the contract method 0x83c0a76e.
//
// Solidity: function addProtocols(string[] protocolNames, []ProtocolMetadata metadata, address[][] adapters, address[][][] tokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) AddProtocols(opts *bind.TransactOpts, protocolNames []string, metadata []ProtocolMetadata, adapters [][]common.Address, tokens [][][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "addProtocols", protocolNames, metadata, adapters, tokens)
}

// AddProtocols is a paid mutator transaction binding the contract method 0x83c0a76e.
//
// Solidity: function addProtocols(string[] protocolNames, []ProtocolMetadata metadata, address[][] adapters, address[][][] tokens) returns()
func (_AdapterRegistry *AdapterRegistrySession) AddProtocols(protocolNames []string, metadata []ProtocolMetadata, adapters [][]common.Address, tokens [][][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddProtocols(&_AdapterRegistry.TransactOpts, protocolNames, metadata, adapters, tokens)
}

// AddProtocols is a paid mutator transaction binding the contract method 0x83c0a76e.
//
// Solidity: function addProtocols(string[] protocolNames, []ProtocolMetadata metadata, address[][] adapters, address[][][] tokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) AddProtocols(protocolNames []string, metadata []ProtocolMetadata, adapters [][]common.Address, tokens [][][]common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddProtocols(&_AdapterRegistry.TransactOpts, protocolNames, metadata, adapters, tokens)
}

// AddTokenAdapters is a paid mutator transaction binding the contract method 0x70be4eda.
//
// Solidity: function addTokenAdapters(string[] tokenAdapterNames, address[] adapters) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) AddTokenAdapters(opts *bind.TransactOpts, tokenAdapterNames []string, adapters []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "addTokenAdapters", tokenAdapterNames, adapters)
}

// AddTokenAdapters is a paid mutator transaction binding the contract method 0x70be4eda.
//
// Solidity: function addTokenAdapters(string[] tokenAdapterNames, address[] adapters) returns()
func (_AdapterRegistry *AdapterRegistrySession) AddTokenAdapters(tokenAdapterNames []string, adapters []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddTokenAdapters(&_AdapterRegistry.TransactOpts, tokenAdapterNames, adapters)
}

// AddTokenAdapters is a paid mutator transaction binding the contract method 0x70be4eda.
//
// Solidity: function addTokenAdapters(string[] tokenAdapterNames, address[] adapters) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) AddTokenAdapters(tokenAdapterNames []string, adapters []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.AddTokenAdapters(&_AdapterRegistry.TransactOpts, tokenAdapterNames, adapters)
}

// RemoveProtocolAdapters is a paid mutator transaction binding the contract method 0x00778353.
//
// Solidity: function removeProtocolAdapters(string protocolName, uint256[] adapterIndices) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) RemoveProtocolAdapters(opts *bind.TransactOpts, protocolName string, adapterIndices []*big.Int) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "removeProtocolAdapters", protocolName, adapterIndices)
}

// RemoveProtocolAdapters is a paid mutator transaction binding the contract method 0x00778353.
//
// Solidity: function removeProtocolAdapters(string protocolName, uint256[] adapterIndices) returns()
func (_AdapterRegistry *AdapterRegistrySession) RemoveProtocolAdapters(protocolName string, adapterIndices []*big.Int) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveProtocolAdapters(&_AdapterRegistry.TransactOpts, protocolName, adapterIndices)
}

// RemoveProtocolAdapters is a paid mutator transaction binding the contract method 0x00778353.
//
// Solidity: function removeProtocolAdapters(string protocolName, uint256[] adapterIndices) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) RemoveProtocolAdapters(protocolName string, adapterIndices []*big.Int) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveProtocolAdapters(&_AdapterRegistry.TransactOpts, protocolName, adapterIndices)
}

// RemoveProtocols is a paid mutator transaction binding the contract method 0x657ea762.
//
// Solidity: function removeProtocols(string[] protocolNames) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) RemoveProtocols(opts *bind.TransactOpts, protocolNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "removeProtocols", protocolNames)
}

// RemoveProtocols is a paid mutator transaction binding the contract method 0x657ea762.
//
// Solidity: function removeProtocols(string[] protocolNames) returns()
func (_AdapterRegistry *AdapterRegistrySession) RemoveProtocols(protocolNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveProtocols(&_AdapterRegistry.TransactOpts, protocolNames)
}

// RemoveProtocols is a paid mutator transaction binding the contract method 0x657ea762.
//
// Solidity: function removeProtocols(string[] protocolNames) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) RemoveProtocols(protocolNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveProtocols(&_AdapterRegistry.TransactOpts, protocolNames)
}

// RemoveTokenAdapters is a paid mutator transaction binding the contract method 0x9d392fa0.
//
// Solidity: function removeTokenAdapters(string[] tokenAdapterNames) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) RemoveTokenAdapters(opts *bind.TransactOpts, tokenAdapterNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "removeTokenAdapters", tokenAdapterNames)
}

// RemoveTokenAdapters is a paid mutator transaction binding the contract method 0x9d392fa0.
//
// Solidity: function removeTokenAdapters(string[] tokenAdapterNames) returns()
func (_AdapterRegistry *AdapterRegistrySession) RemoveTokenAdapters(tokenAdapterNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveTokenAdapters(&_AdapterRegistry.TransactOpts, tokenAdapterNames)
}

// RemoveTokenAdapters is a paid mutator transaction binding the contract method 0x9d392fa0.
//
// Solidity: function removeTokenAdapters(string[] tokenAdapterNames) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) RemoveTokenAdapters(tokenAdapterNames []string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.RemoveTokenAdapters(&_AdapterRegistry.TransactOpts, tokenAdapterNames)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "transferOwnership", _owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_AdapterRegistry *AdapterRegistrySession) TransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.TransferOwnership(&_AdapterRegistry.TransactOpts, _owner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _owner) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) TransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.TransferOwnership(&_AdapterRegistry.TransactOpts, _owner)
}

// UpdateProtocolAdapter is a paid mutator transaction binding the contract method 0xdeebabcc.
//
// Solidity: function updateProtocolAdapter(string protocolName, uint256 index, address newAdapterAddress, address[] newSupportedTokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) UpdateProtocolAdapter(opts *bind.TransactOpts, protocolName string, index *big.Int, newAdapterAddress common.Address, newSupportedTokens []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "updateProtocolAdapter", protocolName, index, newAdapterAddress, newSupportedTokens)
}

// UpdateProtocolAdapter is a paid mutator transaction binding the contract method 0xdeebabcc.
//
// Solidity: function updateProtocolAdapter(string protocolName, uint256 index, address newAdapterAddress, address[] newSupportedTokens) returns()
func (_AdapterRegistry *AdapterRegistrySession) UpdateProtocolAdapter(protocolName string, index *big.Int, newAdapterAddress common.Address, newSupportedTokens []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateProtocolAdapter(&_AdapterRegistry.TransactOpts, protocolName, index, newAdapterAddress, newSupportedTokens)
}

// UpdateProtocolAdapter is a paid mutator transaction binding the contract method 0xdeebabcc.
//
// Solidity: function updateProtocolAdapter(string protocolName, uint256 index, address newAdapterAddress, address[] newSupportedTokens) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) UpdateProtocolAdapter(protocolName string, index *big.Int, newAdapterAddress common.Address, newSupportedTokens []common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateProtocolAdapter(&_AdapterRegistry.TransactOpts, protocolName, index, newAdapterAddress, newSupportedTokens)
}

// UpdateProtocolMetadata is a paid mutator transaction binding the contract method 0xb2c687d0.
//
// Solidity: function updateProtocolMetadata(string protocolName, string name, string description, string websiteURL, string iconURL) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) UpdateProtocolMetadata(opts *bind.TransactOpts, protocolName string, name string, description string, websiteURL string, iconURL string) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "updateProtocolMetadata", protocolName, name, description, websiteURL, iconURL)
}

// UpdateProtocolMetadata is a paid mutator transaction binding the contract method 0xb2c687d0.
//
// Solidity: function updateProtocolMetadata(string protocolName, string name, string description, string websiteURL, string iconURL) returns()
func (_AdapterRegistry *AdapterRegistrySession) UpdateProtocolMetadata(protocolName string, name string, description string, websiteURL string, iconURL string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateProtocolMetadata(&_AdapterRegistry.TransactOpts, protocolName, name, description, websiteURL, iconURL)
}

// UpdateProtocolMetadata is a paid mutator transaction binding the contract method 0xb2c687d0.
//
// Solidity: function updateProtocolMetadata(string protocolName, string name, string description, string websiteURL, string iconURL) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) UpdateProtocolMetadata(protocolName string, name string, description string, websiteURL string, iconURL string) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateProtocolMetadata(&_AdapterRegistry.TransactOpts, protocolName, name, description, websiteURL, iconURL)
}

// UpdateTokenAdapter is a paid mutator transaction binding the contract method 0x2dc909f6.
//
// Solidity: function updateTokenAdapter(string tokenAdapterName, address adapter) returns()
func (_AdapterRegistry *AdapterRegistryTransactor) UpdateTokenAdapter(opts *bind.TransactOpts, tokenAdapterName string, adapter common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.contract.Transact(opts, "updateTokenAdapter", tokenAdapterName, adapter)
}

// UpdateTokenAdapter is a paid mutator transaction binding the contract method 0x2dc909f6.
//
// Solidity: function updateTokenAdapter(string tokenAdapterName, address adapter) returns()
func (_AdapterRegistry *AdapterRegistrySession) UpdateTokenAdapter(tokenAdapterName string, adapter common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateTokenAdapter(&_AdapterRegistry.TransactOpts, tokenAdapterName, adapter)
}

// UpdateTokenAdapter is a paid mutator transaction binding the contract method 0x2dc909f6.
//
// Solidity: function updateTokenAdapter(string tokenAdapterName, address adapter) returns()
func (_AdapterRegistry *AdapterRegistryTransactorSession) UpdateTokenAdapter(tokenAdapterName string, adapter common.Address) (*types.Transaction, error) {
	return _AdapterRegistry.Contract.UpdateTokenAdapter(&_AdapterRegistry.TransactOpts, tokenAdapterName, adapter)
}

// AdapterRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AdapterRegistry contract.
type AdapterRegistryOwnershipTransferredIterator struct {
	Event *AdapterRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AdapterRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdapterRegistryOwnershipTransferred)
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
		it.Event = new(AdapterRegistryOwnershipTransferred)
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
func (it *AdapterRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdapterRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdapterRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the AdapterRegistry contract.
type AdapterRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdapterRegistry *AdapterRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AdapterRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdapterRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdapterRegistryOwnershipTransferredIterator{contract: _AdapterRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdapterRegistry *AdapterRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdapterRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdapterRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdapterRegistryOwnershipTransferred)
				if err := _AdapterRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AdapterRegistry *AdapterRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*AdapterRegistryOwnershipTransferred, error) {
	event := new(AdapterRegistryOwnershipTransferred)
	if err := _AdapterRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
