# Sweep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToAddress** | **string** |  | 
**LockTime** | Pointer to **int64** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**UtxosLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSweep

`func NewSweep(toAddress string, ) *Sweep`

NewSweep instantiates a new Sweep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepWithDefaults

`func NewSweepWithDefaults() *Sweep`

NewSweepWithDefaults instantiates a new Sweep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToAddress

`func (o *Sweep) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *Sweep) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *Sweep) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetLockTime

`func (o *Sweep) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *Sweep) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *Sweep) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *Sweep) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetGasAmount

`func (o *Sweep) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *Sweep) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *Sweep) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *Sweep) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *Sweep) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Sweep) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Sweep) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Sweep) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetUtxosLimit

`func (o *Sweep) GetUtxosLimit() int32`

GetUtxosLimit returns the UtxosLimit field if non-nil, zero value otherwise.

### GetUtxosLimitOk

`func (o *Sweep) GetUtxosLimitOk() (*int32, bool)`

GetUtxosLimitOk returns a tuple with the UtxosLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxosLimit

`func (o *Sweep) SetUtxosLimit(v int32)`

SetUtxosLimit sets UtxosLimit field to given value.

### HasUtxosLimit

`func (o *Sweep) HasUtxosLimit() bool`

HasUtxosLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


