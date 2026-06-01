# BuildTransferTxResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTx** | **string** |  | 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**TxId** | **string** |  | 
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 
**FundingTxs** | Pointer to [**[]BuildSimpleTransferTxResult**](BuildSimpleTransferTxResult.md) |  | [optional] 

## Methods

### NewBuildTransferTxResult

`func NewBuildTransferTxResult(unsignedTx string, gasAmount int32, gasPrice string, txId string, fromGroup int32, toGroup int32, ) *BuildTransferTxResult`

NewBuildTransferTxResult instantiates a new BuildTransferTxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTransferTxResultWithDefaults

`func NewBuildTransferTxResultWithDefaults() *BuildTransferTxResult`

NewBuildTransferTxResultWithDefaults instantiates a new BuildTransferTxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTx

`func (o *BuildTransferTxResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildTransferTxResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildTransferTxResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildTransferTxResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildTransferTxResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildTransferTxResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildTransferTxResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildTransferTxResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildTransferTxResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildTransferTxResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildTransferTxResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildTransferTxResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetFromGroup

`func (o *BuildTransferTxResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildTransferTxResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildTransferTxResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildTransferTxResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildTransferTxResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildTransferTxResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetFundingTxs

`func (o *BuildTransferTxResult) GetFundingTxs() []BuildSimpleTransferTxResult`

GetFundingTxs returns the FundingTxs field if non-nil, zero value otherwise.

### GetFundingTxsOk

`func (o *BuildTransferTxResult) GetFundingTxsOk() (*[]BuildSimpleTransferTxResult, bool)`

GetFundingTxsOk returns a tuple with the FundingTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTxs

`func (o *BuildTransferTxResult) SetFundingTxs(v []BuildSimpleTransferTxResult)`

SetFundingTxs sets FundingTxs field to given value.

### HasFundingTxs

`func (o *BuildTransferTxResult) HasFundingTxs() bool`

HasFundingTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


