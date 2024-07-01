// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registery

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

// RegisteryMetaData contains all meta data concerning the Registery contract.
var RegisteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"registerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506109858061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80630b7914301461003857806370bb43e41461006b575b5f80fd5b610052600480360381019061004d9190610400565b610087565b60405161006294939291906104da565b60405180910390f35b610085600480360381019061008091906105b6565b61020b565b005b5f8181548110610095575f80fd5b905f5260205f2090600402015f91509050805f0180546100b490610686565b80601f01602080910402602001604051908101604052809291908181526020018280546100e090610686565b801561012b5780601f106101025761010080835404028352916020019161012b565b820191905f5260205f20905b81548152906001019060200180831161010e57829003601f168201915b50505050509080600101805461014090610686565b80601f016020809104026020016040519081016040528092919081815260200182805461016c90610686565b80156101b75780601f1061018e576101008083540402835291602001916101b7565b820191905f5260205f20905b81548152906001019060200180831161019a57829003601f168201915b505050505090806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806003015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905084565b5f604051806080016040528088888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f0190816103189190610880565b50602082015181600101908161032e9190610880565b506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050565b5f80fd5b5f80fd5b5f819050919050565b6103df816103cd565b81146103e9575f80fd5b50565b5f813590506103fa816103d6565b92915050565b5f60208284031215610415576104146103c5565b5b5f610422848285016103ec565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61046d8261042b565b6104778185610435565b9350610487818560208601610445565b61049081610453565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104c48261049b565b9050919050565b6104d4816104ba565b82525050565b5f6080820190508181035f8301526104f28187610463565b905081810360208301526105068186610463565b905061051560408301856104cb565b61052260608301846104cb565b95945050505050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261054c5761054b61052b565b5b8235905067ffffffffffffffff8111156105695761056861052f565b5b60208301915083600182028301111561058557610584610533565b5b9250929050565b610595816104ba565b811461059f575f80fd5b50565b5f813590506105b08161058c565b92915050565b5f805f805f80608087890312156105d0576105cf6103c5565b5b5f87013567ffffffffffffffff8111156105ed576105ec6103c9565b5b6105f989828a01610537565b9650965050602087013567ffffffffffffffff81111561061c5761061b6103c9565b5b61062889828a01610537565b9450945050604061063b89828a016105a2565b925050606061064c89828a016105a2565b9150509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061069d57607f821691505b6020821081036106b0576106af610659565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261073f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610704565b6107498683610704565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61078461077f61077a846103cd565b610761565b6103cd565b9050919050565b5f819050919050565b61079d8361076a565b6107b16107a98261078b565b848454610710565b825550505050565b5f90565b6107c56107b9565b6107d0818484610794565b505050565b5b818110156107f3576107e85f826107bd565b6001810190506107d6565b5050565b601f82111561083857610809816106e3565b610812846106f5565b81016020851015610821578190505b61083561082d856106f5565b8301826107d5565b50505b505050565b5f82821c905092915050565b5f6108585f198460080261083d565b1980831691505092915050565b5f6108708383610849565b9150826002028217905092915050565b6108898261042b565b67ffffffffffffffff8111156108a2576108a16106b6565b5b6108ac8254610686565b6108b78282856107f7565b5f60209050601f8311600181146108e8575f84156108d6578287015190505b6108e08582610865565b865550610947565b601f1984166108f6866106e3565b5f5b8281101561091d578489015182556001820191506020850194506020810190506108f8565b8683101561093a5784890151610936601f891682610849565b8355505b6001600288020188555050505b50505050505056fea26469706673582212204a5bcc02f0f0094e0b6c607f15db168104528914f5a95eae139ab350a4614b9164736f6c634300081a0033",
}

// RegisteryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisteryMetaData.ABI instead.
var RegisteryABI = RegisteryMetaData.ABI

// RegisteryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegisteryMetaData.Bin instead.
var RegisteryBin = RegisteryMetaData.Bin

// DeployRegistery deploys a new Ethereum contract, binding an instance of Registery to it.
func DeployRegistery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registery, error) {
	parsed, err := RegisteryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegisteryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registery{RegisteryCaller: RegisteryCaller{contract: contract}, RegisteryTransactor: RegisteryTransactor{contract: contract}, RegisteryFilterer: RegisteryFilterer{contract: contract}}, nil
}

// Registery is an auto generated Go binding around an Ethereum contract.
type Registery struct {
	RegisteryCaller     // Read-only binding to the contract
	RegisteryTransactor // Write-only binding to the contract
	RegisteryFilterer   // Log filterer for contract events
}

// RegisteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterySession struct {
	Contract     *Registery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegisteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisteryCallerSession struct {
	Contract *RegisteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegisteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisteryTransactorSession struct {
	Contract     *RegisteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegisteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisteryRaw struct {
	Contract *Registery // Generic contract binding to access the raw methods on
}

// RegisteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisteryCallerRaw struct {
	Contract *RegisteryCaller // Generic read-only contract binding to access the raw methods on
}

// RegisteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisteryTransactorRaw struct {
	Contract *RegisteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistery creates a new instance of Registery, bound to a specific deployed contract.
func NewRegistery(address common.Address, backend bind.ContractBackend) (*Registery, error) {
	contract, err := bindRegistery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registery{RegisteryCaller: RegisteryCaller{contract: contract}, RegisteryTransactor: RegisteryTransactor{contract: contract}, RegisteryFilterer: RegisteryFilterer{contract: contract}}, nil
}

// NewRegisteryCaller creates a new read-only instance of Registery, bound to a specific deployed contract.
func NewRegisteryCaller(address common.Address, caller bind.ContractCaller) (*RegisteryCaller, error) {
	contract, err := bindRegistery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisteryCaller{contract: contract}, nil
}

// NewRegisteryTransactor creates a new write-only instance of Registery, bound to a specific deployed contract.
func NewRegisteryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisteryTransactor, error) {
	contract, err := bindRegistery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisteryTransactor{contract: contract}, nil
}

// NewRegisteryFilterer creates a new log filterer instance of Registery, bound to a specific deployed contract.
func NewRegisteryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisteryFilterer, error) {
	contract, err := bindRegistery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisteryFilterer{contract: contract}, nil
}

// bindRegistery binds a generic wrapper to an already deployed contract.
func bindRegistery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegisteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registery *RegisteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registery.Contract.RegisteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registery *RegisteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registery.Contract.RegisteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registery *RegisteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registery.Contract.RegisteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registery *RegisteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registery *RegisteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registery *RegisteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registery.Contract.contract.Transact(opts, method, params...)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(string avsName, string operatorName, address avsAddress, address operatorAddress)
func (_Registery *RegisteryCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AvsName         string
	OperatorName    string
	AvsAddress      common.Address
	OperatorAddress common.Address
}, error) {
	var out []interface{}
	err := _Registery.contract.Call(opts, &out, "events", arg0)

	outstruct := new(struct {
		AvsName         string
		OperatorName    string
		AvsAddress      common.Address
		OperatorAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvsName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.OperatorName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AvsAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OperatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(string avsName, string operatorName, address avsAddress, address operatorAddress)
func (_Registery *RegisterySession) Events(arg0 *big.Int) (struct {
	AvsName         string
	OperatorName    string
	AvsAddress      common.Address
	OperatorAddress common.Address
}, error) {
	return _Registery.Contract.Events(&_Registery.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(string avsName, string operatorName, address avsAddress, address operatorAddress)
func (_Registery *RegisteryCallerSession) Events(arg0 *big.Int) (struct {
	AvsName         string
	OperatorName    string
	AvsAddress      common.Address
	OperatorAddress common.Address
}, error) {
	return _Registery.Contract.Events(&_Registery.CallOpts, arg0)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0x70bb43e4.
//
// Solidity: function registerEvent(string avsName, string operatorName, address avsAddress, address operatorAddress) returns()
func (_Registery *RegisteryTransactor) RegisterEvent(opts *bind.TransactOpts, avsName string, operatorName string, avsAddress common.Address, operatorAddress common.Address) (*types.Transaction, error) {
	return _Registery.contract.Transact(opts, "registerEvent", avsName, operatorName, avsAddress, operatorAddress)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0x70bb43e4.
//
// Solidity: function registerEvent(string avsName, string operatorName, address avsAddress, address operatorAddress) returns()
func (_Registery *RegisterySession) RegisterEvent(avsName string, operatorName string, avsAddress common.Address, operatorAddress common.Address) (*types.Transaction, error) {
	return _Registery.Contract.RegisterEvent(&_Registery.TransactOpts, avsName, operatorName, avsAddress, operatorAddress)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0x70bb43e4.
//
// Solidity: function registerEvent(string avsName, string operatorName, address avsAddress, address operatorAddress) returns()
func (_Registery *RegisteryTransactorSession) RegisterEvent(avsName string, operatorName string, avsAddress common.Address, operatorAddress common.Address) (*types.Transaction, error) {
	return _Registery.Contract.RegisterEvent(&_Registery.TransactOpts, avsName, operatorName, avsAddress, operatorAddress)
}
