# BuildExecuteScriptTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**Bytecode** | **string** |  | 
**AttoAlphAmount** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildExecuteScriptTx

`func NewBuildExecuteScriptTx(fromPublicKey string, bytecode string, ) *BuildExecuteScriptTx`

NewBuildExecuteScriptTx instantiates a new BuildExecuteScriptTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildExecuteScriptTxWithDefaults

`func NewBuildExecuteScriptTxWithDefaults() *BuildExecuteScriptTx`

NewBuildExecuteScriptTxWithDefaults instantiates a new BuildExecuteScriptTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *BuildExecuteScriptTx) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *BuildExecuteScriptTx) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *BuildExecuteScriptTx) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetBytecode

`func (o *BuildExecuteScriptTx) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *BuildExecuteScriptTx) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *BuildExecuteScriptTx) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetAttoAlphAmount

`func (o *BuildExecuteScriptTx) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *BuildExecuteScriptTx) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *BuildExecuteScriptTx) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.

### HasAttoAlphAmount

`func (o *BuildExecuteScriptTx) HasAttoAlphAmount() bool`

HasAttoAlphAmount returns a boolean if a field has been set.

### GetTokens

`func (o *BuildExecuteScriptTx) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *BuildExecuteScriptTx) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *BuildExecuteScriptTx) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *BuildExecuteScriptTx) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildExecuteScriptTx) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildExecuteScriptTx) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildExecuteScriptTx) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildExecuteScriptTx) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildExecuteScriptTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildExecuteScriptTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildExecuteScriptTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildExecuteScriptTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


