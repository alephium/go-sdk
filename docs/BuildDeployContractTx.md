# BuildDeployContractTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**Bytecode** | **string** |  | 
**InitialAttoAlphAmount** | Pointer to **string** |  | [optional] 
**InitialTokenAmounts** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**IssueTokenAmount** | Pointer to **string** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildDeployContractTx

`func NewBuildDeployContractTx(fromPublicKey string, bytecode string, ) *BuildDeployContractTx`

NewBuildDeployContractTx instantiates a new BuildDeployContractTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildDeployContractTxWithDefaults

`func NewBuildDeployContractTxWithDefaults() *BuildDeployContractTx`

NewBuildDeployContractTxWithDefaults instantiates a new BuildDeployContractTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *BuildDeployContractTx) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *BuildDeployContractTx) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *BuildDeployContractTx) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetBytecode

`func (o *BuildDeployContractTx) GetBytecode() string`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *BuildDeployContractTx) GetBytecodeOk() (*string, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *BuildDeployContractTx) SetBytecode(v string)`

SetBytecode sets Bytecode field to given value.


### GetInitialAttoAlphAmount

`func (o *BuildDeployContractTx) GetInitialAttoAlphAmount() string`

GetInitialAttoAlphAmount returns the InitialAttoAlphAmount field if non-nil, zero value otherwise.

### GetInitialAttoAlphAmountOk

`func (o *BuildDeployContractTx) GetInitialAttoAlphAmountOk() (*string, bool)`

GetInitialAttoAlphAmountOk returns a tuple with the InitialAttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAttoAlphAmount

`func (o *BuildDeployContractTx) SetInitialAttoAlphAmount(v string)`

SetInitialAttoAlphAmount sets InitialAttoAlphAmount field to given value.

### HasInitialAttoAlphAmount

`func (o *BuildDeployContractTx) HasInitialAttoAlphAmount() bool`

HasInitialAttoAlphAmount returns a boolean if a field has been set.

### GetInitialTokenAmounts

`func (o *BuildDeployContractTx) GetInitialTokenAmounts() []Token`

GetInitialTokenAmounts returns the InitialTokenAmounts field if non-nil, zero value otherwise.

### GetInitialTokenAmountsOk

`func (o *BuildDeployContractTx) GetInitialTokenAmountsOk() (*[]Token, bool)`

GetInitialTokenAmountsOk returns a tuple with the InitialTokenAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialTokenAmounts

`func (o *BuildDeployContractTx) SetInitialTokenAmounts(v []Token)`

SetInitialTokenAmounts sets InitialTokenAmounts field to given value.

### HasInitialTokenAmounts

`func (o *BuildDeployContractTx) HasInitialTokenAmounts() bool`

HasInitialTokenAmounts returns a boolean if a field has been set.

### GetIssueTokenAmount

`func (o *BuildDeployContractTx) GetIssueTokenAmount() string`

GetIssueTokenAmount returns the IssueTokenAmount field if non-nil, zero value otherwise.

### GetIssueTokenAmountOk

`func (o *BuildDeployContractTx) GetIssueTokenAmountOk() (*string, bool)`

GetIssueTokenAmountOk returns a tuple with the IssueTokenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTokenAmount

`func (o *BuildDeployContractTx) SetIssueTokenAmount(v string)`

SetIssueTokenAmount sets IssueTokenAmount field to given value.

### HasIssueTokenAmount

`func (o *BuildDeployContractTx) HasIssueTokenAmount() bool`

HasIssueTokenAmount returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildDeployContractTx) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildDeployContractTx) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildDeployContractTx) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildDeployContractTx) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildDeployContractTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildDeployContractTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildDeployContractTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildDeployContractTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


