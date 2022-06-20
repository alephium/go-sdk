# UnsignedTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**Version** | **int32** |  | 
**NetworkId** | **int32** |  | 
**ScriptOpt** | Pointer to **string** |  | [optional] 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**Inputs** | [**[]AssetInput**](AssetInput.md) |  | 
**FixedOutputs** | [**[]FixedAssetOutput**](FixedAssetOutput.md) |  | 

## Methods

### NewUnsignedTx

`func NewUnsignedTx(txId string, version int32, networkId int32, gasAmount int32, gasPrice string, inputs []AssetInput, fixedOutputs []FixedAssetOutput, ) *UnsignedTx`

NewUnsignedTx instantiates a new UnsignedTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsignedTxWithDefaults

`func NewUnsignedTxWithDefaults() *UnsignedTx`

NewUnsignedTxWithDefaults instantiates a new UnsignedTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *UnsignedTx) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *UnsignedTx) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *UnsignedTx) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetVersion

`func (o *UnsignedTx) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnsignedTx) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnsignedTx) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetNetworkId

`func (o *UnsignedTx) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UnsignedTx) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UnsignedTx) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.


### GetScriptOpt

`func (o *UnsignedTx) GetScriptOpt() string`

GetScriptOpt returns the ScriptOpt field if non-nil, zero value otherwise.

### GetScriptOptOk

`func (o *UnsignedTx) GetScriptOptOk() (*string, bool)`

GetScriptOptOk returns a tuple with the ScriptOpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptOpt

`func (o *UnsignedTx) SetScriptOpt(v string)`

SetScriptOpt sets ScriptOpt field to given value.

### HasScriptOpt

`func (o *UnsignedTx) HasScriptOpt() bool`

HasScriptOpt returns a boolean if a field has been set.

### GetGasAmount

`func (o *UnsignedTx) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *UnsignedTx) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *UnsignedTx) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *UnsignedTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *UnsignedTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *UnsignedTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetInputs

`func (o *UnsignedTx) GetInputs() []AssetInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *UnsignedTx) GetInputsOk() (*[]AssetInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *UnsignedTx) SetInputs(v []AssetInput)`

SetInputs sets Inputs field to given value.


### GetFixedOutputs

`func (o *UnsignedTx) GetFixedOutputs() []FixedAssetOutput`

GetFixedOutputs returns the FixedOutputs field if non-nil, zero value otherwise.

### GetFixedOutputsOk

`func (o *UnsignedTx) GetFixedOutputsOk() (*[]FixedAssetOutput, bool)`

GetFixedOutputsOk returns a tuple with the FixedOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedOutputs

`func (o *UnsignedTx) SetFixedOutputs(v []FixedAssetOutput)`

SetFixedOutputs sets FixedOutputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


