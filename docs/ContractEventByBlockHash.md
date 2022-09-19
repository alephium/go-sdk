# ContractEventByBlockHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**ContractAddress** | **string** |  | 
**EventIndex** | **int32** |  | 
**Fields** | [**[]Val**](Val.md) |  | 

## Methods

### NewContractEventByBlockHash

`func NewContractEventByBlockHash(txId string, contractAddress string, eventIndex int32, fields []Val, ) *ContractEventByBlockHash`

NewContractEventByBlockHash instantiates a new ContractEventByBlockHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventByBlockHashWithDefaults

`func NewContractEventByBlockHashWithDefaults() *ContractEventByBlockHash`

NewContractEventByBlockHashWithDefaults instantiates a new ContractEventByBlockHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *ContractEventByBlockHash) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *ContractEventByBlockHash) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *ContractEventByBlockHash) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetContractAddress

`func (o *ContractEventByBlockHash) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *ContractEventByBlockHash) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *ContractEventByBlockHash) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetEventIndex

`func (o *ContractEventByBlockHash) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *ContractEventByBlockHash) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *ContractEventByBlockHash) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.


### GetFields

`func (o *ContractEventByBlockHash) GetFields() []Val`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ContractEventByBlockHash) GetFieldsOk() (*[]Val, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ContractEventByBlockHash) SetFields(v []Val)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


