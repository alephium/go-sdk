# BuildSweepAddressTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**FromPublicKeyType** | Pointer to **string** |  | [optional] 
**ToAddress** | **string** |  | 
**MaxAttoAlphPerUTXO** | Pointer to **string** |  | [optional] 
**LockTime** | Pointer to **int64** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**TargetBlockHash** | Pointer to **string** |  | [optional] 
**UtxosLimit** | Pointer to **int32** |  | [optional] 
**Group** | Pointer to **int32** |  | [optional] 
**SweepAlphOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewBuildSweepAddressTransactions

`func NewBuildSweepAddressTransactions(fromPublicKey string, toAddress string, ) *BuildSweepAddressTransactions`

NewBuildSweepAddressTransactions instantiates a new BuildSweepAddressTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSweepAddressTransactionsWithDefaults

`func NewBuildSweepAddressTransactionsWithDefaults() *BuildSweepAddressTransactions`

NewBuildSweepAddressTransactionsWithDefaults instantiates a new BuildSweepAddressTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *BuildSweepAddressTransactions) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *BuildSweepAddressTransactions) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *BuildSweepAddressTransactions) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetFromPublicKeyType

`func (o *BuildSweepAddressTransactions) GetFromPublicKeyType() string`

GetFromPublicKeyType returns the FromPublicKeyType field if non-nil, zero value otherwise.

### GetFromPublicKeyTypeOk

`func (o *BuildSweepAddressTransactions) GetFromPublicKeyTypeOk() (*string, bool)`

GetFromPublicKeyTypeOk returns a tuple with the FromPublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKeyType

`func (o *BuildSweepAddressTransactions) SetFromPublicKeyType(v string)`

SetFromPublicKeyType sets FromPublicKeyType field to given value.

### HasFromPublicKeyType

`func (o *BuildSweepAddressTransactions) HasFromPublicKeyType() bool`

HasFromPublicKeyType returns a boolean if a field has been set.

### GetToAddress

`func (o *BuildSweepAddressTransactions) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *BuildSweepAddressTransactions) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *BuildSweepAddressTransactions) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetMaxAttoAlphPerUTXO

`func (o *BuildSweepAddressTransactions) GetMaxAttoAlphPerUTXO() string`

GetMaxAttoAlphPerUTXO returns the MaxAttoAlphPerUTXO field if non-nil, zero value otherwise.

### GetMaxAttoAlphPerUTXOOk

`func (o *BuildSweepAddressTransactions) GetMaxAttoAlphPerUTXOOk() (*string, bool)`

GetMaxAttoAlphPerUTXOOk returns a tuple with the MaxAttoAlphPerUTXO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttoAlphPerUTXO

`func (o *BuildSweepAddressTransactions) SetMaxAttoAlphPerUTXO(v string)`

SetMaxAttoAlphPerUTXO sets MaxAttoAlphPerUTXO field to given value.

### HasMaxAttoAlphPerUTXO

`func (o *BuildSweepAddressTransactions) HasMaxAttoAlphPerUTXO() bool`

HasMaxAttoAlphPerUTXO returns a boolean if a field has been set.

### GetLockTime

`func (o *BuildSweepAddressTransactions) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *BuildSweepAddressTransactions) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *BuildSweepAddressTransactions) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *BuildSweepAddressTransactions) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildSweepAddressTransactions) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildSweepAddressTransactions) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildSweepAddressTransactions) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildSweepAddressTransactions) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildSweepAddressTransactions) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildSweepAddressTransactions) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildSweepAddressTransactions) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildSweepAddressTransactions) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetTargetBlockHash

`func (o *BuildSweepAddressTransactions) GetTargetBlockHash() string`

GetTargetBlockHash returns the TargetBlockHash field if non-nil, zero value otherwise.

### GetTargetBlockHashOk

`func (o *BuildSweepAddressTransactions) GetTargetBlockHashOk() (*string, bool)`

GetTargetBlockHashOk returns a tuple with the TargetBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBlockHash

`func (o *BuildSweepAddressTransactions) SetTargetBlockHash(v string)`

SetTargetBlockHash sets TargetBlockHash field to given value.

### HasTargetBlockHash

`func (o *BuildSweepAddressTransactions) HasTargetBlockHash() bool`

HasTargetBlockHash returns a boolean if a field has been set.

### GetUtxosLimit

`func (o *BuildSweepAddressTransactions) GetUtxosLimit() int32`

GetUtxosLimit returns the UtxosLimit field if non-nil, zero value otherwise.

### GetUtxosLimitOk

`func (o *BuildSweepAddressTransactions) GetUtxosLimitOk() (*int32, bool)`

GetUtxosLimitOk returns a tuple with the UtxosLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxosLimit

`func (o *BuildSweepAddressTransactions) SetUtxosLimit(v int32)`

SetUtxosLimit sets UtxosLimit field to given value.

### HasUtxosLimit

`func (o *BuildSweepAddressTransactions) HasUtxosLimit() bool`

HasUtxosLimit returns a boolean if a field has been set.

### GetGroup

`func (o *BuildSweepAddressTransactions) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BuildSweepAddressTransactions) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BuildSweepAddressTransactions) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BuildSweepAddressTransactions) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSweepAlphOnly

`func (o *BuildSweepAddressTransactions) GetSweepAlphOnly() bool`

GetSweepAlphOnly returns the SweepAlphOnly field if non-nil, zero value otherwise.

### GetSweepAlphOnlyOk

`func (o *BuildSweepAddressTransactions) GetSweepAlphOnlyOk() (*bool, bool)`

GetSweepAlphOnlyOk returns a tuple with the SweepAlphOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSweepAlphOnly

`func (o *BuildSweepAddressTransactions) SetSweepAlphOnly(v bool)`

SetSweepAlphOnly sets SweepAlphOnly field to given value.

### HasSweepAlphOnly

`func (o *BuildSweepAddressTransactions) HasSweepAlphOnly() bool`

HasSweepAlphOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


