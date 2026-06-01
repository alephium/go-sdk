# RichUnsignedTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxId** | **string** |  | 
**Version** | **int32** |  | 
**NetworkId** | **int32** |  | 
**ScriptOpt** | Pointer to **string** |  | [optional] 
**GasAmount** | **int32** |  | 
**GasPrice** | **string** |  | 
**Inputs** | [**[]RichAssetInput**](RichAssetInput.md) |  | 
**FixedOutputs** | [**[]FixedAssetOutput**](FixedAssetOutput.md) |  | 

## Methods

### NewRichUnsignedTx

`func NewRichUnsignedTx(txId string, version int32, networkId int32, gasAmount int32, gasPrice string, inputs []RichAssetInput, fixedOutputs []FixedAssetOutput, ) *RichUnsignedTx`

NewRichUnsignedTx instantiates a new RichUnsignedTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichUnsignedTxWithDefaults

`func NewRichUnsignedTxWithDefaults() *RichUnsignedTx`

NewRichUnsignedTxWithDefaults instantiates a new RichUnsignedTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxId

`func (o *RichUnsignedTx) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *RichUnsignedTx) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *RichUnsignedTx) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetVersion

`func (o *RichUnsignedTx) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RichUnsignedTx) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RichUnsignedTx) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetNetworkId

`func (o *RichUnsignedTx) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *RichUnsignedTx) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *RichUnsignedTx) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.


### GetScriptOpt

`func (o *RichUnsignedTx) GetScriptOpt() string`

GetScriptOpt returns the ScriptOpt field if non-nil, zero value otherwise.

### GetScriptOptOk

`func (o *RichUnsignedTx) GetScriptOptOk() (*string, bool)`

GetScriptOptOk returns a tuple with the ScriptOpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptOpt

`func (o *RichUnsignedTx) SetScriptOpt(v string)`

SetScriptOpt sets ScriptOpt field to given value.

### HasScriptOpt

`func (o *RichUnsignedTx) HasScriptOpt() bool`

HasScriptOpt returns a boolean if a field has been set.

### GetGasAmount

`func (o *RichUnsignedTx) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *RichUnsignedTx) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *RichUnsignedTx) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.


### GetGasPrice

`func (o *RichUnsignedTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *RichUnsignedTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *RichUnsignedTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.


### GetInputs

`func (o *RichUnsignedTx) GetInputs() []RichAssetInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *RichUnsignedTx) GetInputsOk() (*[]RichAssetInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *RichUnsignedTx) SetInputs(v []RichAssetInput)`

SetInputs sets Inputs field to given value.


### GetFixedOutputs

`func (o *RichUnsignedTx) GetFixedOutputs() []FixedAssetOutput`

GetFixedOutputs returns the FixedOutputs field if non-nil, zero value otherwise.

### GetFixedOutputsOk

`func (o *RichUnsignedTx) GetFixedOutputsOk() (*[]FixedAssetOutput, bool)`

GetFixedOutputsOk returns a tuple with the FixedOutputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedOutputs

`func (o *RichUnsignedTx) SetFixedOutputs(v []FixedAssetOutput)`

SetFixedOutputs sets FixedOutputs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


