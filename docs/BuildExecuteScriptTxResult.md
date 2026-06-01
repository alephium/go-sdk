# BuildExecuteScriptTxResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 
**UnsignedTx** | **string** |  | 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**TxId** | **string** |  | 
**SimulationResult** | [**SimulationResult**](SimulationResult.md) |  | 
**FundingTxs** | Pointer to [**[]BuildSimpleTransferTxResult**](BuildSimpleTransferTxResult.md) |  | [optional] 

## Methods

### NewBuildExecuteScriptTxResult

`func NewBuildExecuteScriptTxResult(fromGroup int32, toGroup int32, unsignedTx string, gasAmount int32, gasPrice string, txId string, simulationResult SimulationResult, ) *BuildExecuteScriptTxResult`

NewBuildExecuteScriptTxResult instantiates a new BuildExecuteScriptTxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildExecuteScriptTxResultWithDefaults

`func NewBuildExecuteScriptTxResultWithDefaults() *BuildExecuteScriptTxResult`

NewBuildExecuteScriptTxResultWithDefaults instantiates a new BuildExecuteScriptTxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromGroup

`func (o *BuildExecuteScriptTxResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildExecuteScriptTxResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildExecuteScriptTxResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildExecuteScriptTxResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildExecuteScriptTxResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildExecuteScriptTxResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetUnsignedTx

`func (o *BuildExecuteScriptTxResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildExecuteScriptTxResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildExecuteScriptTxResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildExecuteScriptTxResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildExecuteScriptTxResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildExecuteScriptTxResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildExecuteScriptTxResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildExecuteScriptTxResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildExecuteScriptTxResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildExecuteScriptTxResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildExecuteScriptTxResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildExecuteScriptTxResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetSimulationResult

`func (o *BuildExecuteScriptTxResult) GetSimulationResult() SimulationResult`

GetSimulationResult returns the SimulationResult field if non-nil, zero value otherwise.

### GetSimulationResultOk

`func (o *BuildExecuteScriptTxResult) GetSimulationResultOk() (*SimulationResult, bool)`

GetSimulationResultOk returns a tuple with the SimulationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationResult

`func (o *BuildExecuteScriptTxResult) SetSimulationResult(v SimulationResult)`

SetSimulationResult sets SimulationResult field to given value.


### GetFundingTxs

`func (o *BuildExecuteScriptTxResult) GetFundingTxs() []BuildSimpleTransferTxResult`

GetFundingTxs returns the FundingTxs field if non-nil, zero value otherwise.

### GetFundingTxsOk

`func (o *BuildExecuteScriptTxResult) GetFundingTxsOk() (*[]BuildSimpleTransferTxResult, bool)`

GetFundingTxsOk returns a tuple with the FundingTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTxs

`func (o *BuildExecuteScriptTxResult) SetFundingTxs(v []BuildSimpleTransferTxResult)`

SetFundingTxs sets FundingTxs field to given value.

### HasFundingTxs

`func (o *BuildExecuteScriptTxResult) HasFundingTxs() bool`

HasFundingTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


