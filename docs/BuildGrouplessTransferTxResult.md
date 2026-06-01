# BuildGrouplessTransferTxResult

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

### NewBuildGrouplessTransferTxResult

`func NewBuildGrouplessTransferTxResult(unsignedTx string, gasAmount int32, gasPrice string, txId string, fromGroup int32, toGroup int32, ) *BuildGrouplessTransferTxResult`

NewBuildGrouplessTransferTxResult instantiates a new BuildGrouplessTransferTxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildGrouplessTransferTxResultWithDefaults

`func NewBuildGrouplessTransferTxResultWithDefaults() *BuildGrouplessTransferTxResult`

NewBuildGrouplessTransferTxResultWithDefaults instantiates a new BuildGrouplessTransferTxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTx

`func (o *BuildGrouplessTransferTxResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildGrouplessTransferTxResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildGrouplessTransferTxResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildGrouplessTransferTxResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildGrouplessTransferTxResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildGrouplessTransferTxResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildGrouplessTransferTxResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildGrouplessTransferTxResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildGrouplessTransferTxResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildGrouplessTransferTxResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildGrouplessTransferTxResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildGrouplessTransferTxResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetFromGroup

`func (o *BuildGrouplessTransferTxResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildGrouplessTransferTxResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildGrouplessTransferTxResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildGrouplessTransferTxResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildGrouplessTransferTxResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildGrouplessTransferTxResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetFundingTxs

`func (o *BuildGrouplessTransferTxResult) GetFundingTxs() []BuildSimpleTransferTxResult`

GetFundingTxs returns the FundingTxs field if non-nil, zero value otherwise.

### GetFundingTxsOk

`func (o *BuildGrouplessTransferTxResult) GetFundingTxsOk() (*[]BuildSimpleTransferTxResult, bool)`

GetFundingTxsOk returns a tuple with the FundingTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTxs

`func (o *BuildGrouplessTransferTxResult) SetFundingTxs(v []BuildSimpleTransferTxResult)`

SetFundingTxs sets FundingTxs field to given value.

### HasFundingTxs

`func (o *BuildGrouplessTransferTxResult) HasFundingTxs() bool`

HasFundingTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


