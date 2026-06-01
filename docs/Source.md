# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**Destinations** | [**[]Destination**](Destination.md) |  | 
**FromPublicKeyType** | Pointer to **string** |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**Utxos** | Pointer to [**[]OutputRef**](OutputRef.md) |  | [optional] 

## Methods

### NewSource

`func NewSource(fromPublicKey string, destinations []Destination, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *Source) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *Source) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *Source) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetDestinations

`func (o *Source) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *Source) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *Source) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetFromPublicKeyType

`func (o *Source) GetFromPublicKeyType() string`

GetFromPublicKeyType returns the FromPublicKeyType field if non-nil, zero value otherwise.

### GetFromPublicKeyTypeOk

`func (o *Source) GetFromPublicKeyTypeOk() (*string, bool)`

GetFromPublicKeyTypeOk returns a tuple with the FromPublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKeyType

`func (o *Source) SetFromPublicKeyType(v string)`

SetFromPublicKeyType sets FromPublicKeyType field to given value.

### HasFromPublicKeyType

`func (o *Source) HasFromPublicKeyType() bool`

HasFromPublicKeyType returns a boolean if a field has been set.

### GetGasAmount

`func (o *Source) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *Source) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *Source) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *Source) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetUtxos

`func (o *Source) GetUtxos() []OutputRef`

GetUtxos returns the Utxos field if non-nil, zero value otherwise.

### GetUtxosOk

`func (o *Source) GetUtxosOk() (*[]OutputRef, bool)`

GetUtxosOk returns a tuple with the Utxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxos

`func (o *Source) SetUtxos(v []OutputRef)`

SetUtxos sets Utxos field to given value.

### HasUtxos

`func (o *Source) HasUtxos() bool`

HasUtxos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


