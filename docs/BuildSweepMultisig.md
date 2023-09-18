# BuildSweepMultisig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | **string** |  | 
**FromPublicKeys** | **[]string** |  | 
**ToAddress** | **string** |  | 
**MaxAttoAlphPerUTXO** | Pointer to **string** |  | [optional] 
**LockTime** | Pointer to **int64** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**UtxosLimit** | Pointer to **int32** |  | [optional] 
**TargetBlockHash** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildSweepMultisig

`func NewBuildSweepMultisig(fromAddress string, fromPublicKeys []string, toAddress string, ) *BuildSweepMultisig`

NewBuildSweepMultisig instantiates a new BuildSweepMultisig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSweepMultisigWithDefaults

`func NewBuildSweepMultisigWithDefaults() *BuildSweepMultisig`

NewBuildSweepMultisigWithDefaults instantiates a new BuildSweepMultisig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *BuildSweepMultisig) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *BuildSweepMultisig) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *BuildSweepMultisig) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetFromPublicKeys

`func (o *BuildSweepMultisig) GetFromPublicKeys() []string`

GetFromPublicKeys returns the FromPublicKeys field if non-nil, zero value otherwise.

### GetFromPublicKeysOk

`func (o *BuildSweepMultisig) GetFromPublicKeysOk() (*[]string, bool)`

GetFromPublicKeysOk returns a tuple with the FromPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKeys

`func (o *BuildSweepMultisig) SetFromPublicKeys(v []string)`

SetFromPublicKeys sets FromPublicKeys field to given value.


### GetToAddress

`func (o *BuildSweepMultisig) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *BuildSweepMultisig) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *BuildSweepMultisig) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetMaxAttoAlphPerUTXO

`func (o *BuildSweepMultisig) GetMaxAttoAlphPerUTXO() string`

GetMaxAttoAlphPerUTXO returns the MaxAttoAlphPerUTXO field if non-nil, zero value otherwise.

### GetMaxAttoAlphPerUTXOOk

`func (o *BuildSweepMultisig) GetMaxAttoAlphPerUTXOOk() (*string, bool)`

GetMaxAttoAlphPerUTXOOk returns a tuple with the MaxAttoAlphPerUTXO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttoAlphPerUTXO

`func (o *BuildSweepMultisig) SetMaxAttoAlphPerUTXO(v string)`

SetMaxAttoAlphPerUTXO sets MaxAttoAlphPerUTXO field to given value.

### HasMaxAttoAlphPerUTXO

`func (o *BuildSweepMultisig) HasMaxAttoAlphPerUTXO() bool`

HasMaxAttoAlphPerUTXO returns a boolean if a field has been set.

### GetLockTime

`func (o *BuildSweepMultisig) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *BuildSweepMultisig) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *BuildSweepMultisig) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *BuildSweepMultisig) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildSweepMultisig) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildSweepMultisig) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildSweepMultisig) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildSweepMultisig) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildSweepMultisig) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildSweepMultisig) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildSweepMultisig) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildSweepMultisig) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetUtxosLimit

`func (o *BuildSweepMultisig) GetUtxosLimit() int32`

GetUtxosLimit returns the UtxosLimit field if non-nil, zero value otherwise.

### GetUtxosLimitOk

`func (o *BuildSweepMultisig) GetUtxosLimitOk() (*int32, bool)`

GetUtxosLimitOk returns a tuple with the UtxosLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxosLimit

`func (o *BuildSweepMultisig) SetUtxosLimit(v int32)`

SetUtxosLimit sets UtxosLimit field to given value.

### HasUtxosLimit

`func (o *BuildSweepMultisig) HasUtxosLimit() bool`

HasUtxosLimit returns a boolean if a field has been set.

### GetTargetBlockHash

`func (o *BuildSweepMultisig) GetTargetBlockHash() string`

GetTargetBlockHash returns the TargetBlockHash field if non-nil, zero value otherwise.

### GetTargetBlockHashOk

`func (o *BuildSweepMultisig) GetTargetBlockHashOk() (*string, bool)`

GetTargetBlockHashOk returns a tuple with the TargetBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBlockHash

`func (o *BuildSweepMultisig) SetTargetBlockHash(v string)`

SetTargetBlockHash sets TargetBlockHash field to given value.

### HasTargetBlockHash

`func (o *BuildSweepMultisig) HasTargetBlockHash() bool`

HasTargetBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


