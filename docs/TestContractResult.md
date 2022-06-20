# TestContractResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**CodeHash** | **string** |  | 
**Returns** | [**[]Val**](Val.md) |  | 
**GasUsed** | **int32** |  | 
**Contracts** | [**[]ContractState**](ContractState.md) |  | 
**TxInputs** | **[]string** |  | 
**TxOutputs** | [**[]Output**](Output.md) |  | 
**Events** | [**[]ContractEventByTxId**](ContractEventByTxId.md) |  | 

## Methods

### NewTestContractResult

`func NewTestContractResult(address string, codeHash string, returns []Val, gasUsed int32, contracts []ContractState, txInputs []string, txOutputs []Output, events []ContractEventByTxId, ) *TestContractResult`

NewTestContractResult instantiates a new TestContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestContractResultWithDefaults

`func NewTestContractResultWithDefaults() *TestContractResult`

NewTestContractResultWithDefaults instantiates a new TestContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TestContractResult) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TestContractResult) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TestContractResult) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCodeHash

`func (o *TestContractResult) GetCodeHash() string`

GetCodeHash returns the CodeHash field if non-nil, zero value otherwise.

### GetCodeHashOk

`func (o *TestContractResult) GetCodeHashOk() (*string, bool)`

GetCodeHashOk returns a tuple with the CodeHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeHash

`func (o *TestContractResult) SetCodeHash(v string)`

SetCodeHash sets CodeHash field to given value.


### GetReturns

`func (o *TestContractResult) GetReturns() []Val`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *TestContractResult) GetReturnsOk() (*[]Val, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *TestContractResult) SetReturns(v []Val)`

SetReturns sets Returns field to given value.


### GetGasUsed

`func (o *TestContractResult) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *TestContractResult) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *TestContractResult) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetContracts

`func (o *TestContractResult) GetContracts() []ContractState`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *TestContractResult) GetContractsOk() (*[]ContractState, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *TestContractResult) SetContracts(v []ContractState)`

SetContracts sets Contracts field to given value.


### GetTxInputs

`func (o *TestContractResult) GetTxInputs() []string`

GetTxInputs returns the TxInputs field if non-nil, zero value otherwise.

### GetTxInputsOk

`func (o *TestContractResult) GetTxInputsOk() (*[]string, bool)`

GetTxInputsOk returns a tuple with the TxInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxInputs

`func (o *TestContractResult) SetTxInputs(v []string)`

SetTxInputs sets TxInputs field to given value.


### GetTxOutputs

`func (o *TestContractResult) GetTxOutputs() []Output`

GetTxOutputs returns the TxOutputs field if non-nil, zero value otherwise.

### GetTxOutputsOk

`func (o *TestContractResult) GetTxOutputsOk() (*[]Output, bool)`

GetTxOutputsOk returns a tuple with the TxOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxOutputs

`func (o *TestContractResult) SetTxOutputs(v []Output)`

SetTxOutputs sets TxOutputs field to given value.


### GetEvents

`func (o *TestContractResult) GetEvents() []ContractEventByTxId`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TestContractResult) GetEventsOk() (*[]ContractEventByTxId, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TestContractResult) SetEvents(v []ContractEventByTxId)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


