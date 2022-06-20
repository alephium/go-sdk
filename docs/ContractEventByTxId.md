# ContractEventByTxId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** |  | 
**ContractAddress** | **string** |  | 
**EventIndex** | **int32** |  | 
**Fields** | [**[]Val**](Val.md) |  | 

## Methods

### NewContractEventByTxId

`func NewContractEventByTxId(blockHash string, contractAddress string, eventIndex int32, fields []Val, ) *ContractEventByTxId`

NewContractEventByTxId instantiates a new ContractEventByTxId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventByTxIdWithDefaults

`func NewContractEventByTxIdWithDefaults() *ContractEventByTxId`

NewContractEventByTxIdWithDefaults instantiates a new ContractEventByTxId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *ContractEventByTxId) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractEventByTxId) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractEventByTxId) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetContractAddress

`func (o *ContractEventByTxId) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ContractEventByTxId) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ContractEventByTxId) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetEventIndex

`func (o *ContractEventByTxId) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *ContractEventByTxId) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *ContractEventByTxId) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetFields

`func (o *ContractEventByTxId) GetFields() []Val`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContractEventByTxId) GetFieldsOk() (*[]Val, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContractEventByTxId) SetFields(v []Val)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


