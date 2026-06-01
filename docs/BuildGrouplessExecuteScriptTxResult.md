# BuildGrouplessExecuteScriptTxResult

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

### NewBuildGrouplessExecuteScriptTxResult

`func NewBuildGrouplessExecuteScriptTxResult(fromGroup int32, toGroup int32, unsignedTx string, gasAmount int32, gasPrice string, txId string, simulationResult SimulationResult, ) *BuildGrouplessExecuteScriptTxResult`

NewBuildGrouplessExecuteScriptTxResult instantiates a new BuildGrouplessExecuteScriptTxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildGrouplessExecuteScriptTxResultWithDefaults

`func NewBuildGrouplessExecuteScriptTxResultWithDefaults() *BuildGrouplessExecuteScriptTxResult`

NewBuildGrouplessExecuteScriptTxResultWithDefaults instantiates a new BuildGrouplessExecuteScriptTxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromGroup

`func (o *BuildGrouplessExecuteScriptTxResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildGrouplessExecuteScriptTxResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildGrouplessExecuteScriptTxResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildGrouplessExecuteScriptTxResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetUnsignedTx

`func (o *BuildGrouplessExecuteScriptTxResult) GetUnsignedTx() string`

GetUnsignedTx returns the UnsignedTx field if non-nil, zero value otherwise.

### GetUnsignedTxOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetUnsignedTxOk() (*string, bool)`

GetUnsignedTxOk returns a tuple with the UnsignedTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTx

`func (o *BuildGrouplessExecuteScriptTxResult) SetUnsignedTx(v string)`

SetUnsignedTx sets UnsignedTx field to given value.


### GetGasAmount

`func (o *BuildGrouplessExecuteScriptTxResult) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildGrouplessExecuteScriptTxResult) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *BuildGrouplessExecuteScriptTxResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildGrouplessExecuteScriptTxResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetTxId

`func (o *BuildGrouplessExecuteScriptTxResult) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *BuildGrouplessExecuteScriptTxResult) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetSimulationResult

`func (o *BuildGrouplessExecuteScriptTxResult) GetSimulationResult() SimulationResult`

GetSimulationResult returns the SimulationResult field if non-nil, zero value otherwise.

### GetSimulationResultOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetSimulationResultOk() (*SimulationResult, bool)`

GetSimulationResultOk returns a tuple with the SimulationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationResult

`func (o *BuildGrouplessExecuteScriptTxResult) SetSimulationResult(v SimulationResult)`

SetSimulationResult sets SimulationResult field to given value.


### GetFundingTxs

`func (o *BuildGrouplessExecuteScriptTxResult) GetFundingTxs() []BuildSimpleTransferTxResult`

GetFundingTxs returns the FundingTxs field if non-nil, zero value otherwise.

### GetFundingTxsOk

`func (o *BuildGrouplessExecuteScriptTxResult) GetFundingTxsOk() (*[]BuildSimpleTransferTxResult, bool)`

GetFundingTxsOk returns a tuple with the FundingTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingTxs

`func (o *BuildGrouplessExecuteScriptTxResult) SetFundingTxs(v []BuildSimpleTransferTxResult)`

SetFundingTxs sets FundingTxs field to given value.

### HasFundingTxs

`func (o *BuildGrouplessExecuteScriptTxResult) HasFundingTxs() bool`

HasFundingTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


