# CallContractResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Returns** | [**[]Val**](Val.md) |  | 
**GasUsed** | **int32** |  | 
**Contracts** | [**[]ContractState**](ContractState.md) |  | 
**TxInputs** | **[]string** |  | 
**TxOutputs** | [**[]Output**](Output.md) |  | 
**Events** | [**[]ContractEventByTxId**](ContractEventByTxId.md) |  | 

## Methods

### NewCallContractResult

`func NewCallContractResult(returns []Val, gasUsed int32, contracts []ContractState, txInputs []string, txOutputs []Output, events []ContractEventByTxId, ) *CallContractResult`

NewCallContractResult instantiates a new CallContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallContractResultWithDefaults

`func NewCallContractResultWithDefaults() *CallContractResult`

NewCallContractResultWithDefaults instantiates a new CallContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturns

`func (o *CallContractResult) GetReturns() []Val`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *CallContractResult) GetReturnsOk() (*[]Val, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *CallContractResult) SetReturns(v []Val)`

SetReturns sets Returns field to given value.


### GetGasUsed

`func (o *CallContractResult) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *CallContractResult) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *CallContractResult) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetContracts

`func (o *CallContractResult) GetContracts() []ContractState`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *CallContractResult) GetContractsOk() (*[]ContractState, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *CallContractResult) SetContracts(v []ContractState)`

SetContracts sets Contracts field to given value.


### GetTxInputs

`func (o *CallContractResult) GetTxInputs() []string`

GetTxInputs returns the TxInputs field if non-nil, zero value otherwise.

### GetTxInputsOk

`func (o *CallContractResult) GetTxInputsOk() (*[]string, bool)`

GetTxInputsOk returns a tuple with the TxInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxInputs

`func (o *CallContractResult) SetTxInputs(v []string)`

SetTxInputs sets TxInputs field to given value.


### GetTxOutputs

`func (o *CallContractResult) GetTxOutputs() []Output`

GetTxOutputs returns the TxOutputs field if non-nil, zero value otherwise.

### GetTxOutputsOk

`func (o *CallContractResult) GetTxOutputsOk() (*[]Output, bool)`

GetTxOutputsOk returns a tuple with the TxOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxOutputs

`func (o *CallContractResult) SetTxOutputs(v []Output)`

SetTxOutputs sets TxOutputs field to given value.


### GetEvents

`func (o *CallContractResult) GetEvents() []ContractEventByTxId`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CallContractResult) GetEventsOk() (*[]ContractEventByTxId, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CallContractResult) SetEvents(v []ContractEventByTxId)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


