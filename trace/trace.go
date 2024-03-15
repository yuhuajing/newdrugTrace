// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trace

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

// RockPaperScissorMintStageInfo is an auto generated low-level Go binding around an user-defined struct.
type RockPaperScissorMintStageInfo struct {
	Traceid    string
	Prodinfo   string
	Storeinfo  string
	Logisinfo  string
	Salestring string
	Batchid    string
}

// TraceMetaData contains all meta data concerning the Trace contract.
var TraceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_producer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_storer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_logostic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sales\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"addOrupdateLogisData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"addOrupdateProdData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"addOrupdateSaleData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"addOrupdateStoreData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tracedata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"traceid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"prodinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storeinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logisinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"salestring\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchid\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"tracedataById\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"traceid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"prodinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"storeinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logisinfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"salestring\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"batchid\",\"type\":\"string\"}],\"internalType\":\"structRockPaperScissor.MintStageInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620012d8380380620012d88339818101604052810190620000369190620001a5565b8360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505062000214565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200016f8262000144565b9050919050565b620001818162000163565b81146200018c575f80fd5b50565b5f815190506200019f8162000176565b92915050565b5f805f8060808587031215620001c057620001bf62000140565b5b5f620001cf878288016200018f565b9450506020620001e2878288016200018f565b9350506040620001f5878288016200018f565b925050606062000208878288016200018f565b91505092959194509250565b6110b680620002225f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806320037dd3146100645780634e9fc3981461008057806359374c40146100b05780635c7e9495146100cc5780637f1a8f02146101015780637f6d888b1461011d575b5f80fd5b61007e60048036038101906100799190610a8d565b610139565b005b61009a60048036038101906100959190610b03565b61016b565b6040516100a79190610c6d565b60405180910390f35b6100ca60048036038101906100c59190610a8d565b610501565b005b6100e660048036038101906100e19190610b03565b610533565b6040516100f896959493929190610cd5565b60405180910390f35b61011b60048036038101906101169190610a8d565b6108a6565b005b61013760048036038101906101329190610a8d565b6108d8565b005b805f836040516101499190610d98565b908152602001604051809103902060020190816101669190610fb1565b505050565b61017361090a565b5f826040516101829190610d98565b90815260200160405180910390206040518060c00160405290815f820180546101aa90610ddb565b80601f01602080910402602001604051908101604052809291908181526020018280546101d690610ddb565b80156102215780601f106101f857610100808354040283529160200191610221565b820191905f5260205f20905b81548152906001019060200180831161020457829003601f168201915b5050505050815260200160018201805461023a90610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461026690610ddb565b80156102b15780601f10610288576101008083540402835291602001916102b1565b820191905f5260205f20905b81548152906001019060200180831161029457829003601f168201915b505050505081526020016002820180546102ca90610ddb565b80601f01602080910402602001604051908101604052809291908181526020018280546102f690610ddb565b80156103415780601f1061031857610100808354040283529160200191610341565b820191905f5260205f20905b81548152906001019060200180831161032457829003601f168201915b5050505050815260200160038201805461035a90610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461038690610ddb565b80156103d15780601f106103a8576101008083540402835291602001916103d1565b820191905f5260205f20905b8154815290600101906020018083116103b457829003601f168201915b505050505081526020016004820180546103ea90610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461041690610ddb565b80156104615780601f1061043857610100808354040283529160200191610461565b820191905f5260205f20905b81548152906001019060200180831161044457829003601f168201915b5050505050815260200160058201805461047a90610ddb565b80601f01602080910402602001604051908101604052809291908181526020018280546104a690610ddb565b80156104f15780601f106104c8576101008083540402835291602001916104f1565b820191905f5260205f20905b8154815290600101906020018083116104d457829003601f168201915b5050505050815250509050919050565b805f836040516105119190610d98565b9081526020016040518091039020600101908161052e9190610fb1565b505050565b5f818051602081018201805184825260208301602085012081835280955050505050505f91509050805f01805461056990610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461059590610ddb565b80156105e05780601f106105b7576101008083540402835291602001916105e0565b820191905f5260205f20905b8154815290600101906020018083116105c357829003601f168201915b5050505050908060010180546105f590610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461062190610ddb565b801561066c5780601f106106435761010080835404028352916020019161066c565b820191905f5260205f20905b81548152906001019060200180831161064f57829003601f168201915b50505050509080600201805461068190610ddb565b80601f01602080910402602001604051908101604052809291908181526020018280546106ad90610ddb565b80156106f85780601f106106cf576101008083540402835291602001916106f8565b820191905f5260205f20905b8154815290600101906020018083116106db57829003601f168201915b50505050509080600301805461070d90610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461073990610ddb565b80156107845780601f1061075b57610100808354040283529160200191610784565b820191905f5260205f20905b81548152906001019060200180831161076757829003601f168201915b50505050509080600401805461079990610ddb565b80601f01602080910402602001604051908101604052809291908181526020018280546107c590610ddb565b80156108105780601f106107e757610100808354040283529160200191610810565b820191905f5260205f20905b8154815290600101906020018083116107f357829003601f168201915b50505050509080600501805461082590610ddb565b80601f016020809104026020016040519081016040528092919081815260200182805461085190610ddb565b801561089c5780601f106108735761010080835404028352916020019161089c565b820191905f5260205f20905b81548152906001019060200180831161087f57829003601f168201915b5050505050905086565b805f836040516108b69190610d98565b908152602001604051809103902060030190816108d39190610fb1565b505050565b805f836040516108e89190610d98565b908152602001604051809103902060040190816109059190610fb1565b505050565b6040518060c001604052806060815260200160608152602001606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61099f82610959565b810181811067ffffffffffffffff821117156109be576109bd610969565b5b80604052505050565b5f6109d0610940565b90506109dc8282610996565b919050565b5f67ffffffffffffffff8211156109fb576109fa610969565b5b610a0482610959565b9050602081019050919050565b828183375f83830152505050565b5f610a31610a2c846109e1565b6109c7565b905082815260208101848484011115610a4d57610a4c610955565b5b610a58848285610a11565b509392505050565b5f82601f830112610a7457610a73610951565b5b8135610a84848260208601610a1f565b91505092915050565b5f8060408385031215610aa357610aa2610949565b5b5f83013567ffffffffffffffff811115610ac057610abf61094d565b5b610acc85828601610a60565b925050602083013567ffffffffffffffff811115610aed57610aec61094d565b5b610af985828601610a60565b9150509250929050565b5f60208284031215610b1857610b17610949565b5b5f82013567ffffffffffffffff811115610b3557610b3461094d565b5b610b4184828501610a60565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610b81578082015181840152602081019050610b66565b5f8484015250505050565b5f610b9682610b4a565b610ba08185610b54565b9350610bb0818560208601610b64565b610bb981610959565b840191505092915050565b5f60c083015f8301518482035f860152610bde8282610b8c565b91505060208301518482036020860152610bf88282610b8c565b91505060408301518482036040860152610c128282610b8c565b91505060608301518482036060860152610c2c8282610b8c565b91505060808301518482036080860152610c468282610b8c565b91505060a083015184820360a0860152610c608282610b8c565b9150508091505092915050565b5f6020820190508181035f830152610c858184610bc4565b905092915050565b5f82825260208201905092915050565b5f610ca782610b4a565b610cb18185610c8d565b9350610cc1818560208601610b64565b610cca81610959565b840191505092915050565b5f60c0820190508181035f830152610ced8189610c9d565b90508181036020830152610d018188610c9d565b90508181036040830152610d158187610c9d565b90508181036060830152610d298186610c9d565b90508181036080830152610d3d8185610c9d565b905081810360a0830152610d518184610c9d565b9050979650505050505050565b5f81905092915050565b5f610d7282610b4a565b610d7c8185610d5e565b9350610d8c818560208601610b64565b80840191505092915050565b5f610da38284610d68565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610df257607f821691505b602082108103610e0557610e04610dae565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610e677fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e2c565b610e718683610e2c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610eb5610eb0610eab84610e89565b610e92565b610e89565b9050919050565b5f819050919050565b610ece83610e9b565b610ee2610eda82610ebc565b848454610e38565b825550505050565b5f90565b610ef6610eea565b610f01818484610ec5565b505050565b5b81811015610f2457610f195f82610eee565b600181019050610f07565b5050565b601f821115610f6957610f3a81610e0b565b610f4384610e1d565b81016020851015610f52578190505b610f66610f5e85610e1d565b830182610f06565b50505b505050565b5f82821c905092915050565b5f610f895f1984600802610f6e565b1980831691505092915050565b5f610fa18383610f7a565b9150826002028217905092915050565b610fba82610b4a565b67ffffffffffffffff811115610fd357610fd2610969565b5b610fdd8254610ddb565b610fe8828285610f28565b5f60209050601f831160018114611019575f8415611007578287015190505b6110118582610f96565b865550611078565b601f19841661102786610e0b565b5f5b8281101561104e57848901518255600182019150602085019450602081019050611029565b8683101561106b5784890151611067601f891682610f7a565b8355505b6001600288020188555050505b50505050505056fea26469706673582212209aefc189851ef7c5b9bbb93afff2045b18b3a23c06950bc260a5edbe9c60e3cf64736f6c63430008180033",
}

// TraceABI is the input ABI used to generate the binding from.
// Deprecated: Use TraceMetaData.ABI instead.
var TraceABI = TraceMetaData.ABI

// TraceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TraceMetaData.Bin instead.
var TraceBin = TraceMetaData.Bin

// DeployTrace deploys a new Ethereum contract, binding an instance of Trace to it.
func DeployTrace(auth *bind.TransactOpts, backend bind.ContractBackend, _producer common.Address, _storer common.Address, _logostic common.Address, _sales common.Address) (common.Address, *types.Transaction, *Trace, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TraceBin), backend, _producer, _storer, _logostic, _sales)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// Trace is an auto generated Go binding around an Ethereum contract.
type Trace struct {
	TraceCaller     // Read-only binding to the contract
	TraceTransactor // Write-only binding to the contract
	TraceFilterer   // Log filterer for contract events
}

// TraceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraceSession struct {
	Contract     *Trace            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraceCallerSession struct {
	Contract *TraceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TraceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraceTransactorSession struct {
	Contract     *TraceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraceRaw struct {
	Contract *Trace // Generic contract binding to access the raw methods on
}

// TraceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraceCallerRaw struct {
	Contract *TraceCaller // Generic read-only contract binding to access the raw methods on
}

// TraceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraceTransactorRaw struct {
	Contract *TraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrace creates a new instance of Trace, bound to a specific deployed contract.
func NewTrace(address common.Address, backend bind.ContractBackend) (*Trace, error) {
	contract, err := bindTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trace{TraceCaller: TraceCaller{contract: contract}, TraceTransactor: TraceTransactor{contract: contract}, TraceFilterer: TraceFilterer{contract: contract}}, nil
}

// NewTraceCaller creates a new read-only instance of Trace, bound to a specific deployed contract.
func NewTraceCaller(address common.Address, caller bind.ContractCaller) (*TraceCaller, error) {
	contract, err := bindTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraceCaller{contract: contract}, nil
}

// NewTraceTransactor creates a new write-only instance of Trace, bound to a specific deployed contract.
func NewTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*TraceTransactor, error) {
	contract, err := bindTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraceTransactor{contract: contract}, nil
}

// NewTraceFilterer creates a new log filterer instance of Trace, bound to a specific deployed contract.
func NewTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*TraceFilterer, error) {
	contract, err := bindTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraceFilterer{contract: contract}, nil
}

// bindTrace binds a generic wrapper to an already deployed contract.
func bindTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.TraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.TraceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trace *TraceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trace *TraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trace *TraceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trace.Contract.contract.Transact(opts, method, params...)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string traceid, string prodinfo, string storeinfo, string logisinfo, string salestring, string batchid)
func (_Trace *TraceCaller) Tracedata(opts *bind.CallOpts, arg0 string) (struct {
	Traceid    string
	Prodinfo   string
	Storeinfo  string
	Logisinfo  string
	Salestring string
	Batchid    string
}, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedata", arg0)

	outstruct := new(struct {
		Traceid    string
		Prodinfo   string
		Storeinfo  string
		Logisinfo  string
		Salestring string
		Batchid    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Traceid = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Prodinfo = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Storeinfo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Logisinfo = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Salestring = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Batchid = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string traceid, string prodinfo, string storeinfo, string logisinfo, string salestring, string batchid)
func (_Trace *TraceSession) Tracedata(arg0 string) (struct {
	Traceid    string
	Prodinfo   string
	Storeinfo  string
	Logisinfo  string
	Salestring string
	Batchid    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// Tracedata is a free data retrieval call binding the contract method 0x5c7e9495.
//
// Solidity: function tracedata(string ) view returns(string traceid, string prodinfo, string storeinfo, string logisinfo, string salestring, string batchid)
func (_Trace *TraceCallerSession) Tracedata(arg0 string) (struct {
	Traceid    string
	Prodinfo   string
	Storeinfo  string
	Logisinfo  string
	Salestring string
	Batchid    string
}, error) {
	return _Trace.Contract.Tracedata(&_Trace.CallOpts, arg0)
}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string,string,string))
func (_Trace *TraceCaller) TracedataById(opts *bind.CallOpts, id string) (RockPaperScissorMintStageInfo, error) {
	var out []interface{}
	err := _Trace.contract.Call(opts, &out, "tracedataById", id)

	if err != nil {
		return *new(RockPaperScissorMintStageInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RockPaperScissorMintStageInfo)).(*RockPaperScissorMintStageInfo)

	return out0, err

}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string,string,string))
func (_Trace *TraceSession) TracedataById(id string) (RockPaperScissorMintStageInfo, error) {
	return _Trace.Contract.TracedataById(&_Trace.CallOpts, id)
}

// TracedataById is a free data retrieval call binding the contract method 0x4e9fc398.
//
// Solidity: function tracedataById(string id) view returns((string,string,string,string,string,string))
func (_Trace *TraceCallerSession) TracedataById(id string) (RockPaperScissorMintStageInfo, error) {
	return _Trace.Contract.TracedataById(&_Trace.CallOpts, id)
}

// AddOrupdateLogisData is a paid mutator transaction binding the contract method 0x7f1a8f02.
//
// Solidity: function addOrupdateLogisData(string id, string data) returns()
func (_Trace *TraceTransactor) AddOrupdateLogisData(opts *bind.TransactOpts, id string, data string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateLogisData", id, data)
}

// AddOrupdateLogisData is a paid mutator transaction binding the contract method 0x7f1a8f02.
//
// Solidity: function addOrupdateLogisData(string id, string data) returns()
func (_Trace *TraceSession) AddOrupdateLogisData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateLogisData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateLogisData is a paid mutator transaction binding the contract method 0x7f1a8f02.
//
// Solidity: function addOrupdateLogisData(string id, string data) returns()
func (_Trace *TraceTransactorSession) AddOrupdateLogisData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateLogisData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0x59374c40.
//
// Solidity: function addOrupdateProdData(string id, string data) returns()
func (_Trace *TraceTransactor) AddOrupdateProdData(opts *bind.TransactOpts, id string, data string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateProdData", id, data)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0x59374c40.
//
// Solidity: function addOrupdateProdData(string id, string data) returns()
func (_Trace *TraceSession) AddOrupdateProdData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateProdData is a paid mutator transaction binding the contract method 0x59374c40.
//
// Solidity: function addOrupdateProdData(string id, string data) returns()
func (_Trace *TraceTransactorSession) AddOrupdateProdData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateProdData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateSaleData is a paid mutator transaction binding the contract method 0x7f6d888b.
//
// Solidity: function addOrupdateSaleData(string id, string data) returns()
func (_Trace *TraceTransactor) AddOrupdateSaleData(opts *bind.TransactOpts, id string, data string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateSaleData", id, data)
}

// AddOrupdateSaleData is a paid mutator transaction binding the contract method 0x7f6d888b.
//
// Solidity: function addOrupdateSaleData(string id, string data) returns()
func (_Trace *TraceSession) AddOrupdateSaleData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateSaleData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateSaleData is a paid mutator transaction binding the contract method 0x7f6d888b.
//
// Solidity: function addOrupdateSaleData(string id, string data) returns()
func (_Trace *TraceTransactorSession) AddOrupdateSaleData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateSaleData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x20037dd3.
//
// Solidity: function addOrupdateStoreData(string id, string data) returns()
func (_Trace *TraceTransactor) AddOrupdateStoreData(opts *bind.TransactOpts, id string, data string) (*types.Transaction, error) {
	return _Trace.contract.Transact(opts, "addOrupdateStoreData", id, data)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x20037dd3.
//
// Solidity: function addOrupdateStoreData(string id, string data) returns()
func (_Trace *TraceSession) AddOrupdateStoreData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateStoreData(&_Trace.TransactOpts, id, data)
}

// AddOrupdateStoreData is a paid mutator transaction binding the contract method 0x20037dd3.
//
// Solidity: function addOrupdateStoreData(string id, string data) returns()
func (_Trace *TraceTransactorSession) AddOrupdateStoreData(id string, data string) (*types.Transaction, error) {
	return _Trace.Contract.AddOrupdateStoreData(&_Trace.TransactOpts, id, data)
}