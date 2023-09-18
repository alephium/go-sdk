# CallContractSucceeded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Returns** | [**[]Val**](Val.md) |  | 
**GasUsed** | **int32** |  | 
**Contracts** | [**[]ContractState**](ContractState.md) |  | 
**TxInputs** | **[]string** |  | 
**TxOutputs** | [**[]Output**](Output.md) |  | 
**Events** | [**[]ContractEventByTxId**](ContractEventByTxId.md) |  | 
**Type** | **string** |  | 

## Methods

### NewCallContractSucceeded

`func NewCallContractSucceeded(returns []Val, gasUsed int32, contracts []ContractState, txInputs []string, txOutputs []Output, events []ContractEventByTxId, type_ string, ) *CallContractSucceeded`

NewCallContractSucceeded instantiates a new CallContractSucceeded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallContractSucceededWithDefaults

`func NewCallContractSucceededWithDefaults() *CallContractSucceeded`

NewCallContractSucceededWithDefaults instantiates a new CallContractSucceeded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturns

`func (o *CallContractSucceeded) GetReturns() []Val`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *CallContractSucceeded) GetReturnsOk() (*[]Val, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *CallContractSucceeded) SetReturns(v []Val)`

SetReturns sets Returns field to given value.


### GetGasUsed

`func (o *CallContractSucceeded) GetGasUsed() int32`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *CallContractSucceeded) GetGasUsedOk() (*int32, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *CallContractSucceeded) SetGasUsed(v int32)`

SetGasUsed sets GasUsed field to given value.


### GetContracts

`func (o *CallContractSucceeded) GetContracts() []ContractState`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *CallContractSucceeded) GetContractsOk() (*[]ContractState, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *CallContractSucceeded) SetContracts(v []ContractState)`

SetContracts sets Contracts field to given value.


### GetTxInputs

`func (o *CallContractSucceeded) GetTxInputs() []string`

GetTxInputs returns the TxInputs field if non-nil, zero value otherwise.

### GetTxInputsOk

`func (o *CallContractSucceeded) GetTxInputsOk() (*[]string, bool)`

GetTxInputsOk returns a tuple with the TxInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxInputs

`func (o *CallContractSucceeded) SetTxInputs(v []string)`

SetTxInputs sets TxInputs field to given value.


### GetTxOutputs

`func (o *CallContractSucceeded) GetTxOutputs() []Output`

GetTxOutputs returns the TxOutputs field if non-nil, zero value otherwise.

### GetTxOutputsOk

`func (o *CallContractSucceeded) GetTxOutputsOk() (*[]Output, bool)`

GetTxOutputsOk returns a tuple with the TxOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxOutputs

`func (o *CallContractSucceeded) SetTxOutputs(v []Output)`

SetTxOutputs sets TxOutputs field to given value.


### GetEvents

`func (o *CallContractSucceeded) GetEvents() []ContractEventByTxId`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CallContractSucceeded) GetEventsOk() (*[]ContractEventByTxId, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CallContractSucceeded) SetEvents(v []ContractEventByTxId)`

SetEvents sets Events field to given value.


### GetType

`func (o *CallContractSucceeded) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallContractSucceeded) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallContractSucceeded) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


