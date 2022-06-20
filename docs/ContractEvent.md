# ContractEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** |  | 
**TxId** | **string** |  | 
**EventIndex** | **int32** |  | 
**Fields** | [**[]Val**](Val.md) |  | 

## Methods

### NewContractEvent

`func NewContractEvent(blockHash string, txId string, eventIndex int32, fields []Val, ) *ContractEvent`

NewContractEvent instantiates a new ContractEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventWithDefaults

`func NewContractEventWithDefaults() *ContractEvent`

NewContractEventWithDefaults instantiates a new ContractEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *ContractEvent) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractEvent) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractEvent) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetTxId

`func (o *ContractEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ContractEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ContractEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetEventIndex

`func (o *ContractEvent) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *ContractEvent) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *ContractEvent) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetFields

`func (o *ContractEvent) GetFields() []Val`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContractEvent) GetFieldsOk() (*[]Val, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContractEvent) SetFields(v []Val)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


