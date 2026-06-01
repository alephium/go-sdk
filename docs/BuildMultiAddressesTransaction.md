# BuildMultiAddressesTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**[]Source**](Source.md) |  | 
**GasPrice** | Pointer to **string** |  | [optional] 
**TargetBlockHash** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildMultiAddressesTransaction

`func NewBuildMultiAddressesTransaction(from []Source, ) *BuildMultiAddressesTransaction`

NewBuildMultiAddressesTransaction instantiates a new BuildMultiAddressesTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildMultiAddressesTransactionWithDefaults

`func NewBuildMultiAddressesTransactionWithDefaults() *BuildMultiAddressesTransaction`

NewBuildMultiAddressesTransactionWithDefaults instantiates a new BuildMultiAddressesTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *BuildMultiAddressesTransaction) GetFrom() []Source`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BuildMultiAddressesTransaction) GetFromOk() (*[]Source, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BuildMultiAddressesTransaction) SetFrom(v []Source)`

SetFrom sets From field to given value.


### GetGasPrice

`func (o *BuildMultiAddressesTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildMultiAddressesTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildMultiAddressesTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildMultiAddressesTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetTargetBlockHash

`func (o *BuildMultiAddressesTransaction) GetTargetBlockHash() string`

GetTargetBlockHash returns the TargetBlockHash field if non-nil, zero value otherwise.

### GetTargetBlockHashOk

`func (o *BuildMultiAddressesTransaction) GetTargetBlockHashOk() (*string, bool)`

GetTargetBlockHashOk returns a tuple with the TargetBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBlockHash

`func (o *BuildMultiAddressesTransaction) SetTargetBlockHash(v string)`

SetTargetBlockHash sets TargetBlockHash field to given value.

### HasTargetBlockHash

`func (o *BuildMultiAddressesTransaction) HasTargetBlockHash() bool`

HasTargetBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


