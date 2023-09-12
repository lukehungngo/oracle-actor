// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimistic_oracle

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

// OptimisticOracleResult is an auto generated low-level Go binding around an user-defined struct.
type OptimisticOracleResult struct {
	TdsContractId *big.Int
	ReportId      *big.Int
	SubmittedAt   uint64
	IsDefault     bool
	Result        []byte
}

// OptimisticOracleMetaData contains all meta data concerning the OptimisticOracle contract.
var OptimisticOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TDSContractAlreadyReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TDSContractIsDefault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TDSContractReportTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TDSContractUnderReporting\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reportId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"referenceEvent\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"triggerType\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"ReportCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reportId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"ReportResultSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"}],\"name\":\"getReportIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reportId\",\"type\":\"uint256\"}],\"name\":\"getReportResultById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reportId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"submittedAt\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"internalType\":\"structOptimisticOracle.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"}],\"name\":\"isTDSContractDefault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tdsContractId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"referenceEvent\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"triggerType\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"reportId\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"submitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f8054600160a01b600160e01b0319166107e960a71b179055348015610028575f80fd5b50338061004e57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100578161005d565b506100ac565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610c8d806100b95f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c80638da5cb5b116100585780638da5cb5b14610100578063de2c072f1461011a578063e103fa6c1461013b578063f2fde38b1461015c575f80fd5b80632ac2d8331461008957806366204d40146100b2578063715018a6146100c7578063859197c3146100cf575b5f80fd5b61009c6100973660046107c8565b61016f565b6040516100a991906107df565b60405180910390f35b6100c56100c036600461087e565b6101ce565b005b6100c561041e565b6100f06100dd3660046107c8565b5f90815260056020526040902054151590565b60405190151581526020016100a9565b5f546040516001600160a01b0390911681526020016100a9565b61012d6101283660046108e0565b610431565b6040519081526020016100a9565b61014e6101493660046107c8565b6105ed565b6040516100a9929190610963565b6100c561016a3660046109f8565b61070b565b5f818152600460209081526040918290208054835181840281018401909452808452606093928301828280156101c257602002820191905f5260205f20905b8154815260200190600101908083116101ae575b50505050509050919050565b63ffffffff84165f818152600260205260408120549003610202576040516398b54f9f60e01b815260040160405180910390fd5b5f8054828252600260208190526040909220909101544291610237916001600160401b03600160a01b90920482169116610a39565b6001600160401b0316101561025f576040516324fd6fdf60e01b815260040160405180910390fd5b5f818152600360205260409020541561028b5760405163be27606b60e01b815260040160405180910390fd5b61029361074d565b5f60025f8763ffffffff1681526020019081526020015f205f015490505f4290506040518060a001604052808381526020018863ffffffff168152602001826001600160401b03168152602001871515815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92018290525093909452505063ffffffff8a16815260036020818152604092839020855181559085015160018201559184015160028301805460608701511515600160401b0268ffffffffffffffffff199091166001600160401b039093169290921791909117905560808401519192508201906103919082610afa565b5050505f828152600560205260409020541580156103ac5750855b156103c8575f82815260056020526040902063ffffffff881690555b806001600160401b03168763ffffffff16837f9695a693126c6b222f9f1131e8842bd3baf83d23888ccf94ef17e238526e739489898960405161040d93929190610bdd565b60405180910390a450505050505050565b61042661074d565b61042f5f610779565b565b5f8060015f815461044190610c01565b91829055505f898152600460209081526040808320805460018101825590845292829020909201839055815160c0810183528b8152808201849052426001600160401b0381168285015263ffffffff8c1660608301528351601f8b018490048402810184019094528984529394509160808301918a908a90819084018382808284375f92019190915250505090825250604080516020601f8901819004810282018101909252878152918101919088908890819084018382808284375f9201829052509390945250508481526002602081815260409283902085518155908501516001820155918401519082018054606086015163ffffffff16600160401b026bffffffffffffffffffffffff199091166001600160401b03909316929092179190911790556080830151909150600382019061057e9082610afa565b5060a082015160048201906105939082610afa565b50905050806001600160401b031689837e583c4772c2377a97f8f86c4f8185adb00c7c4d219ab1518cc769ba3ca37ac48b8b8b8b8b6040516105d9959493929190610c19565b60405180910390a450979650505050505050565b6040805160a0810182525f80825260208201819052918101829052606080820183905260808201525f838152600360208181526040928390208054845160a08101865281815260018301549381019390935260028201546001600160401b03811695840195909552600160401b90940460ff16151560608301529182018054931515938391608084019161068090610a74565b80601f01602080910402602001604051908101604052809291908181526020018280546106ac90610a74565b80156106f75780601f106106ce576101008083540402835291602001916106f7565b820191905f5260205f20905b8154815290600101906020018083116106da57829003601f168201915b505050505081525050905091509150915091565b61071361074d565b6001600160a01b03811661074157604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61074a81610779565b50565b5f546001600160a01b0316331461042f5760405163118cdaa760e01b8152336004820152602401610738565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156107d8575f80fd5b5035919050565b602080825282518282018190525f9190848201906040850190845b81811015610816578351835292840192918401916001016107fa565b50909695505050505050565b803563ffffffff81168114610835575f80fd5b919050565b5f8083601f84011261084a575f80fd5b5081356001600160401b03811115610860575f80fd5b602083019150836020828501011115610877575f80fd5b9250929050565b5f805f8060608587031215610891575f80fd5b61089a85610822565b9350602085013580151581146108ae575f80fd5b925060408501356001600160401b038111156108c8575f80fd5b6108d48782880161083a565b95989497509550505050565b5f805f805f80608087890312156108f5575f80fd5b8635955061090560208801610822565b945060408701356001600160401b0380821115610920575f80fd5b61092c8a838b0161083a565b90965094506060890135915080821115610944575f80fd5b5061095189828a0161083a565b979a9699509497509295939492505050565b82151581525f6020604081840152835160408401528084015160608401526001600160401b0360408501511660808401526060840151151560a0840152608084015160a060c085015280518060e08601525f5b818110156109d357828101840151868201610100015283016109b6565b5061010092505f83828701015282601f19601f83011686010193505050509392505050565b5f60208284031215610a08575f80fd5b81356001600160a01b0381168114610a1e575f80fd5b9392505050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03818116838216019080821115610a5957610a59610a25565b5092915050565b634e487b7160e01b5f52604160045260245ffd5b600181811c90821680610a8857607f821691505b602082108103610aa657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610af5575f81815260208120601f850160051c81016020861015610ad25750805b601f850160051c820191505b81811015610af157828155600101610ade565b5050505b505050565b81516001600160401b03811115610b1357610b13610a60565b610b2781610b218454610a74565b84610aac565b602080601f831160018114610b5a575f8415610b435750858301515b5f19600386901b1c1916600185901b178555610af1565b5f85815260208120601f198616915b82811015610b8857888601518255948401946001909101908401610b69565b5085821015610ba557878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8315158152604060208201525f610bf8604083018486610bb5565b95945050505050565b5f60018201610c1257610c12610a25565b5060010190565b63ffffffff86168152606060208201525f610c38606083018688610bb5565b8281036040840152610c4b818587610bb5565b9897505050505050505056fea26469706673582212202ba6daa3e74ab1ce3a6655e48376598f0ee61bdd8ab4a57cf06e626af81e868b64736f6c63430008150033",
}

// OptimisticOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimisticOracleMetaData.ABI instead.
var OptimisticOracleABI = OptimisticOracleMetaData.ABI

// OptimisticOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptimisticOracleMetaData.Bin instead.
var OptimisticOracleBin = OptimisticOracleMetaData.Bin

// DeployOptimisticOracle deploys a new Ethereum contract, binding an instance of OptimisticOracle to it.
func DeployOptimisticOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptimisticOracle, error) {
	parsed, err := OptimisticOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimisticOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimisticOracle{OptimisticOracleCaller: OptimisticOracleCaller{contract: contract}, OptimisticOracleTransactor: OptimisticOracleTransactor{contract: contract}, OptimisticOracleFilterer: OptimisticOracleFilterer{contract: contract}}, nil
}

// OptimisticOracle is an auto generated Go binding around an Ethereum contract.
type OptimisticOracle struct {
	OptimisticOracleCaller     // Read-only binding to the contract
	OptimisticOracleTransactor // Write-only binding to the contract
	OptimisticOracleFilterer   // Log filterer for contract events
}

// OptimisticOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimisticOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimisticOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimisticOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimisticOracleSession struct {
	Contract     *OptimisticOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptimisticOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimisticOracleCallerSession struct {
	Contract *OptimisticOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OptimisticOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimisticOracleTransactorSession struct {
	Contract     *OptimisticOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OptimisticOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimisticOracleRaw struct {
	Contract *OptimisticOracle // Generic contract binding to access the raw methods on
}

// OptimisticOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimisticOracleCallerRaw struct {
	Contract *OptimisticOracleCaller // Generic read-only contract binding to access the raw methods on
}

// OptimisticOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimisticOracleTransactorRaw struct {
	Contract *OptimisticOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimisticOracle creates a new instance of OptimisticOracle, bound to a specific deployed contract.
func NewOptimisticOracle(address common.Address, backend bind.ContractBackend) (*OptimisticOracle, error) {
	contract, err := bindOptimisticOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracle{OptimisticOracleCaller: OptimisticOracleCaller{contract: contract}, OptimisticOracleTransactor: OptimisticOracleTransactor{contract: contract}, OptimisticOracleFilterer: OptimisticOracleFilterer{contract: contract}}, nil
}

// NewOptimisticOracleCaller creates a new read-only instance of OptimisticOracle, bound to a specific deployed contract.
func NewOptimisticOracleCaller(address common.Address, caller bind.ContractCaller) (*OptimisticOracleCaller, error) {
	contract, err := bindOptimisticOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleCaller{contract: contract}, nil
}

// NewOptimisticOracleTransactor creates a new write-only instance of OptimisticOracle, bound to a specific deployed contract.
func NewOptimisticOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimisticOracleTransactor, error) {
	contract, err := bindOptimisticOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleTransactor{contract: contract}, nil
}

// NewOptimisticOracleFilterer creates a new log filterer instance of OptimisticOracle, bound to a specific deployed contract.
func NewOptimisticOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimisticOracleFilterer, error) {
	contract, err := bindOptimisticOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleFilterer{contract: contract}, nil
}

// bindOptimisticOracle binds a generic wrapper to an already deployed contract.
func bindOptimisticOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimisticOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticOracle *OptimisticOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticOracle.Contract.OptimisticOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticOracle *OptimisticOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.OptimisticOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticOracle *OptimisticOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.OptimisticOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticOracle *OptimisticOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticOracle *OptimisticOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticOracle *OptimisticOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.contract.Transact(opts, method, params...)
}

// GetReportIds is a free data retrieval call binding the contract method 0x2ac2d833.
//
// Solidity: function getReportIds(uint256 tdsContractId) view returns(uint256[])
func (_OptimisticOracle *OptimisticOracleCaller) GetReportIds(opts *bind.CallOpts, tdsContractId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _OptimisticOracle.contract.Call(opts, &out, "getReportIds", tdsContractId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetReportIds is a free data retrieval call binding the contract method 0x2ac2d833.
//
// Solidity: function getReportIds(uint256 tdsContractId) view returns(uint256[])
func (_OptimisticOracle *OptimisticOracleSession) GetReportIds(tdsContractId *big.Int) ([]*big.Int, error) {
	return _OptimisticOracle.Contract.GetReportIds(&_OptimisticOracle.CallOpts, tdsContractId)
}

// GetReportIds is a free data retrieval call binding the contract method 0x2ac2d833.
//
// Solidity: function getReportIds(uint256 tdsContractId) view returns(uint256[])
func (_OptimisticOracle *OptimisticOracleCallerSession) GetReportIds(tdsContractId *big.Int) ([]*big.Int, error) {
	return _OptimisticOracle.Contract.GetReportIds(&_OptimisticOracle.CallOpts, tdsContractId)
}

// GetReportResultById is a free data retrieval call binding the contract method 0xe103fa6c.
//
// Solidity: function getReportResultById(uint256 reportId) view returns(bool, (uint256,uint256,uint64,bool,bytes))
func (_OptimisticOracle *OptimisticOracleCaller) GetReportResultById(opts *bind.CallOpts, reportId *big.Int) (bool, OptimisticOracleResult, error) {
	var out []interface{}
	err := _OptimisticOracle.contract.Call(opts, &out, "getReportResultById", reportId)

	if err != nil {
		return *new(bool), *new(OptimisticOracleResult), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(OptimisticOracleResult)).(*OptimisticOracleResult)

	return out0, out1, err

}

// GetReportResultById is a free data retrieval call binding the contract method 0xe103fa6c.
//
// Solidity: function getReportResultById(uint256 reportId) view returns(bool, (uint256,uint256,uint64,bool,bytes))
func (_OptimisticOracle *OptimisticOracleSession) GetReportResultById(reportId *big.Int) (bool, OptimisticOracleResult, error) {
	return _OptimisticOracle.Contract.GetReportResultById(&_OptimisticOracle.CallOpts, reportId)
}

// GetReportResultById is a free data retrieval call binding the contract method 0xe103fa6c.
//
// Solidity: function getReportResultById(uint256 reportId) view returns(bool, (uint256,uint256,uint64,bool,bytes))
func (_OptimisticOracle *OptimisticOracleCallerSession) GetReportResultById(reportId *big.Int) (bool, OptimisticOracleResult, error) {
	return _OptimisticOracle.Contract.GetReportResultById(&_OptimisticOracle.CallOpts, reportId)
}

// IsTDSContractDefault is a free data retrieval call binding the contract method 0x859197c3.
//
// Solidity: function isTDSContractDefault(uint256 tdsContractId) view returns(bool)
func (_OptimisticOracle *OptimisticOracleCaller) IsTDSContractDefault(opts *bind.CallOpts, tdsContractId *big.Int) (bool, error) {
	var out []interface{}
	err := _OptimisticOracle.contract.Call(opts, &out, "isTDSContractDefault", tdsContractId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTDSContractDefault is a free data retrieval call binding the contract method 0x859197c3.
//
// Solidity: function isTDSContractDefault(uint256 tdsContractId) view returns(bool)
func (_OptimisticOracle *OptimisticOracleSession) IsTDSContractDefault(tdsContractId *big.Int) (bool, error) {
	return _OptimisticOracle.Contract.IsTDSContractDefault(&_OptimisticOracle.CallOpts, tdsContractId)
}

// IsTDSContractDefault is a free data retrieval call binding the contract method 0x859197c3.
//
// Solidity: function isTDSContractDefault(uint256 tdsContractId) view returns(bool)
func (_OptimisticOracle *OptimisticOracleCallerSession) IsTDSContractDefault(tdsContractId *big.Int) (bool, error) {
	return _OptimisticOracle.Contract.IsTDSContractDefault(&_OptimisticOracle.CallOpts, tdsContractId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptimisticOracle *OptimisticOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimisticOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptimisticOracle *OptimisticOracleSession) Owner() (common.Address, error) {
	return _OptimisticOracle.Contract.Owner(&_OptimisticOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OptimisticOracle *OptimisticOracleCallerSession) Owner() (common.Address, error) {
	return _OptimisticOracle.Contract.Owner(&_OptimisticOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptimisticOracle *OptimisticOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptimisticOracle *OptimisticOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _OptimisticOracle.Contract.RenounceOwnership(&_OptimisticOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OptimisticOracle *OptimisticOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OptimisticOracle.Contract.RenounceOwnership(&_OptimisticOracle.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0xde2c072f.
//
// Solidity: function report(uint256 tdsContractId, uint32 referenceEvent, bytes triggerType, bytes proof) returns(uint256)
func (_OptimisticOracle *OptimisticOracleTransactor) Report(opts *bind.TransactOpts, tdsContractId *big.Int, referenceEvent uint32, triggerType []byte, proof []byte) (*types.Transaction, error) {
	return _OptimisticOracle.contract.Transact(opts, "report", tdsContractId, referenceEvent, triggerType, proof)
}

// Report is a paid mutator transaction binding the contract method 0xde2c072f.
//
// Solidity: function report(uint256 tdsContractId, uint32 referenceEvent, bytes triggerType, bytes proof) returns(uint256)
func (_OptimisticOracle *OptimisticOracleSession) Report(tdsContractId *big.Int, referenceEvent uint32, triggerType []byte, proof []byte) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.Report(&_OptimisticOracle.TransactOpts, tdsContractId, referenceEvent, triggerType, proof)
}

// Report is a paid mutator transaction binding the contract method 0xde2c072f.
//
// Solidity: function report(uint256 tdsContractId, uint32 referenceEvent, bytes triggerType, bytes proof) returns(uint256)
func (_OptimisticOracle *OptimisticOracleTransactorSession) Report(tdsContractId *big.Int, referenceEvent uint32, triggerType []byte, proof []byte) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.Report(&_OptimisticOracle.TransactOpts, tdsContractId, referenceEvent, triggerType, proof)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x66204d40.
//
// Solidity: function submitResult(uint32 reportId, bool isDefault, bytes result) returns()
func (_OptimisticOracle *OptimisticOracleTransactor) SubmitResult(opts *bind.TransactOpts, reportId uint32, isDefault bool, result []byte) (*types.Transaction, error) {
	return _OptimisticOracle.contract.Transact(opts, "submitResult", reportId, isDefault, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x66204d40.
//
// Solidity: function submitResult(uint32 reportId, bool isDefault, bytes result) returns()
func (_OptimisticOracle *OptimisticOracleSession) SubmitResult(reportId uint32, isDefault bool, result []byte) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.SubmitResult(&_OptimisticOracle.TransactOpts, reportId, isDefault, result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0x66204d40.
//
// Solidity: function submitResult(uint32 reportId, bool isDefault, bytes result) returns()
func (_OptimisticOracle *OptimisticOracleTransactorSession) SubmitResult(reportId uint32, isDefault bool, result []byte) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.SubmitResult(&_OptimisticOracle.TransactOpts, reportId, isDefault, result)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptimisticOracle *OptimisticOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OptimisticOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptimisticOracle *OptimisticOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.TransferOwnership(&_OptimisticOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OptimisticOracle *OptimisticOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OptimisticOracle.Contract.TransferOwnership(&_OptimisticOracle.TransactOpts, newOwner)
}

// OptimisticOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OptimisticOracle contract.
type OptimisticOracleOwnershipTransferredIterator struct {
	Event *OptimisticOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OptimisticOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimisticOracleOwnershipTransferred)
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
		it.Event = new(OptimisticOracleOwnershipTransferred)
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
func (it *OptimisticOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimisticOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimisticOracleOwnershipTransferred represents a OwnershipTransferred event raised by the OptimisticOracle contract.
type OptimisticOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OptimisticOracle *OptimisticOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OptimisticOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OptimisticOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleOwnershipTransferredIterator{contract: _OptimisticOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OptimisticOracle *OptimisticOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimisticOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OptimisticOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimisticOracleOwnershipTransferred)
				if err := _OptimisticOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OptimisticOracle *OptimisticOracleFilterer) ParseOwnershipTransferred(log types.Log) (*OptimisticOracleOwnershipTransferred, error) {
	event := new(OptimisticOracleOwnershipTransferred)
	if err := _OptimisticOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimisticOracleReportCreatedIterator is returned from FilterReportCreated and is used to iterate over the raw logs and unpacked data for ReportCreated events raised by the OptimisticOracle contract.
type OptimisticOracleReportCreatedIterator struct {
	Event *OptimisticOracleReportCreated // Event containing the contract specifics and raw log

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
func (it *OptimisticOracleReportCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimisticOracleReportCreated)
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
		it.Event = new(OptimisticOracleReportCreated)
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
func (it *OptimisticOracleReportCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimisticOracleReportCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimisticOracleReportCreated represents a ReportCreated event raised by the OptimisticOracle contract.
type OptimisticOracleReportCreated struct {
	TdsContractId  *big.Int
	ReportId       *big.Int
	StartTime      uint64
	ReferenceEvent uint32
	TriggerType    []byte
	Proof          []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReportCreated is a free log retrieval operation binding the contract event 0x00583c4772c2377a97f8f86c4f8185adb00c7c4d219ab1518cc769ba3ca37ac4.
//
// Solidity: event ReportCreated(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed startTime, uint32 referenceEvent, bytes triggerType, bytes proof)
func (_OptimisticOracle *OptimisticOracleFilterer) FilterReportCreated(opts *bind.FilterOpts, tdsContractId []*big.Int, reportId []*big.Int, startTime []uint64) (*OptimisticOracleReportCreatedIterator, error) {

	var tdsContractIdRule []interface{}
	for _, tdsContractIdItem := range tdsContractId {
		tdsContractIdRule = append(tdsContractIdRule, tdsContractIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}
	var startTimeRule []interface{}
	for _, startTimeItem := range startTime {
		startTimeRule = append(startTimeRule, startTimeItem)
	}

	logs, sub, err := _OptimisticOracle.contract.FilterLogs(opts, "ReportCreated", tdsContractIdRule, reportIdRule, startTimeRule)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleReportCreatedIterator{contract: _OptimisticOracle.contract, event: "ReportCreated", logs: logs, sub: sub}, nil
}

// WatchReportCreated is a free log subscription operation binding the contract event 0x00583c4772c2377a97f8f86c4f8185adb00c7c4d219ab1518cc769ba3ca37ac4.
//
// Solidity: event ReportCreated(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed startTime, uint32 referenceEvent, bytes triggerType, bytes proof)
func (_OptimisticOracle *OptimisticOracleFilterer) WatchReportCreated(opts *bind.WatchOpts, sink chan<- *OptimisticOracleReportCreated, tdsContractId []*big.Int, reportId []*big.Int, startTime []uint64) (event.Subscription, error) {

	var tdsContractIdRule []interface{}
	for _, tdsContractIdItem := range tdsContractId {
		tdsContractIdRule = append(tdsContractIdRule, tdsContractIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}
	var startTimeRule []interface{}
	for _, startTimeItem := range startTime {
		startTimeRule = append(startTimeRule, startTimeItem)
	}

	logs, sub, err := _OptimisticOracle.contract.WatchLogs(opts, "ReportCreated", tdsContractIdRule, reportIdRule, startTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimisticOracleReportCreated)
				if err := _OptimisticOracle.contract.UnpackLog(event, "ReportCreated", log); err != nil {
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

// ParseReportCreated is a log parse operation binding the contract event 0x00583c4772c2377a97f8f86c4f8185adb00c7c4d219ab1518cc769ba3ca37ac4.
//
// Solidity: event ReportCreated(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed startTime, uint32 referenceEvent, bytes triggerType, bytes proof)
func (_OptimisticOracle *OptimisticOracleFilterer) ParseReportCreated(log types.Log) (*OptimisticOracleReportCreated, error) {
	event := new(OptimisticOracleReportCreated)
	if err := _OptimisticOracle.contract.UnpackLog(event, "ReportCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptimisticOracleReportResultSubmittedIterator is returned from FilterReportResultSubmitted and is used to iterate over the raw logs and unpacked data for ReportResultSubmitted events raised by the OptimisticOracle contract.
type OptimisticOracleReportResultSubmittedIterator struct {
	Event *OptimisticOracleReportResultSubmitted // Event containing the contract specifics and raw log

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
func (it *OptimisticOracleReportResultSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimisticOracleReportResultSubmitted)
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
		it.Event = new(OptimisticOracleReportResultSubmitted)
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
func (it *OptimisticOracleReportResultSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OptimisticOracleReportResultSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OptimisticOracleReportResultSubmitted represents a ReportResultSubmitted event raised by the OptimisticOracle contract.
type OptimisticOracleReportResultSubmitted struct {
	TdsContractId *big.Int
	ReportId      *big.Int
	SubmittedAt   uint64
	IsDefault     bool
	Result        []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReportResultSubmitted is a free log retrieval operation binding the contract event 0x9695a693126c6b222f9f1131e8842bd3baf83d23888ccf94ef17e238526e7394.
//
// Solidity: event ReportResultSubmitted(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed submittedAt, bool isDefault, bytes result)
func (_OptimisticOracle *OptimisticOracleFilterer) FilterReportResultSubmitted(opts *bind.FilterOpts, tdsContractId []*big.Int, reportId []*big.Int, submittedAt []uint64) (*OptimisticOracleReportResultSubmittedIterator, error) {

	var tdsContractIdRule []interface{}
	for _, tdsContractIdItem := range tdsContractId {
		tdsContractIdRule = append(tdsContractIdRule, tdsContractIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}
	var submittedAtRule []interface{}
	for _, submittedAtItem := range submittedAt {
		submittedAtRule = append(submittedAtRule, submittedAtItem)
	}

	logs, sub, err := _OptimisticOracle.contract.FilterLogs(opts, "ReportResultSubmitted", tdsContractIdRule, reportIdRule, submittedAtRule)
	if err != nil {
		return nil, err
	}
	return &OptimisticOracleReportResultSubmittedIterator{contract: _OptimisticOracle.contract, event: "ReportResultSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportResultSubmitted is a free log subscription operation binding the contract event 0x9695a693126c6b222f9f1131e8842bd3baf83d23888ccf94ef17e238526e7394.
//
// Solidity: event ReportResultSubmitted(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed submittedAt, bool isDefault, bytes result)
func (_OptimisticOracle *OptimisticOracleFilterer) WatchReportResultSubmitted(opts *bind.WatchOpts, sink chan<- *OptimisticOracleReportResultSubmitted, tdsContractId []*big.Int, reportId []*big.Int, submittedAt []uint64) (event.Subscription, error) {

	var tdsContractIdRule []interface{}
	for _, tdsContractIdItem := range tdsContractId {
		tdsContractIdRule = append(tdsContractIdRule, tdsContractIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}
	var submittedAtRule []interface{}
	for _, submittedAtItem := range submittedAt {
		submittedAtRule = append(submittedAtRule, submittedAtItem)
	}

	logs, sub, err := _OptimisticOracle.contract.WatchLogs(opts, "ReportResultSubmitted", tdsContractIdRule, reportIdRule, submittedAtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OptimisticOracleReportResultSubmitted)
				if err := _OptimisticOracle.contract.UnpackLog(event, "ReportResultSubmitted", log); err != nil {
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

// ParseReportResultSubmitted is a log parse operation binding the contract event 0x9695a693126c6b222f9f1131e8842bd3baf83d23888ccf94ef17e238526e7394.
//
// Solidity: event ReportResultSubmitted(uint256 indexed tdsContractId, uint256 indexed reportId, uint64 indexed submittedAt, bool isDefault, bytes result)
func (_OptimisticOracle *OptimisticOracleFilterer) ParseReportResultSubmitted(log types.Log) (*OptimisticOracleReportResultSubmitted, error) {
	event := new(OptimisticOracleReportResultSubmitted)
	if err := _OptimisticOracle.contract.UnpackLog(event, "ReportResultSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
