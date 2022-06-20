# BuildTransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTx** | **string** |  | 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**TxId** | **string** |  | 
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 

## Methods

### NewBuildTransactionResult

`func NewBuildTransactionResult(unsignedTx string, gasAmount int32, gasPrice string, txId string, fromGroup int32, toGroup int32, ) *BuildTransactionResult`

NewBuildTransactionResult instantiates a new BuildTransactionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTransactionResultWithDefaults

`func NewBuildTransactionResultWithDefaults() *BuildTransactionResult`

NewBuildTransactionResultWithDefaults instantiates a new BuildTransactionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTx

`func (o *BuildTransactionResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildTransactionResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildTransactionResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildTransactionResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildTransactionResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildTransactionResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildTransactionResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildTransactionResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildTransactionResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildTransactionResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildTransactionResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildTransactionResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetFromGroup

`func (o *BuildTransactionResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildTransactionResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildTransactionResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildTransactionResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildTransactionResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildTransactionResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


