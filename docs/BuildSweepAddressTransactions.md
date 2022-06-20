# BuildSweepAddressTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**ToAddress** | **string** |  | 
**LockTime** | Pointer to **int64** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


