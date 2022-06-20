# BuildTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**Destinations** | [**[]Destination**](Destination.md) |  | 
**Utxos** | Pointer to [**[]OutputRef**](OutputRef.md) |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildTransaction

`func NewBuildTransaction(fromPublicKey string, destinations []Destination, ) *BuildTransaction`

NewBuildTransaction instantiates a new BuildTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTransactionWithDefaults

`func NewBuildTransactionWithDefaults() *BuildTransaction`

NewBuildTransactionWithDefaults instantiates a new BuildTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *BuildTransaction) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *BuildTransaction) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *BuildTransaction) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetDestinations

`func (o *BuildTransaction) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *BuildTransaction) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *BuildTransaction) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetUtxos

`func (o *BuildTransaction) GetUtxos() []OutputRef`

GetUtxos returns the Utxos field if non-nil, zero value otherwise.

### GetUtxosOk

`func (o *BuildTransaction) GetUtxosOk() (*[]OutputRef, bool)`

GetUtxosOk returns a tuple with the Utxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxos

`func (o *BuildTransaction) SetUtxos(v []OutputRef)`

SetUtxos sets Utxos field to given value.

### HasUtxos

`func (o *BuildTransaction) HasUtxos() bool`

HasUtxos returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildTransaction) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildTransaction) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildTransaction) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildTransaction) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildTransaction) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildTransaction) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildTransaction) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildTransaction) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


