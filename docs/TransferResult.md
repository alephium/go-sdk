# TransferResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 

## Methods

### NewTransferResult

`func NewTransferResult(txId string, fromGroup int32, toGroup int32, ) *TransferResult`

NewTransferResult instantiates a new TransferResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferResultWithDefaults

`func NewTransferResultWithDefaults() *TransferResult`

NewTransferResultWithDefaults instantiates a new TransferResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *TransferResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransferResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransferResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetFromGroup

`func (o *TransferResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *TransferResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *TransferResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *TransferResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *TransferResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *TransferResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


