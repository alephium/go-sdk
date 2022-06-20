# BuildMultisig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | **string** |  | 
**FromPublicKeys** | **[]string** |  | 
**Destinations** | [**[]Destination**](Destination.md) |  | 
**Gas** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildMultisig

`func NewBuildMultisig(fromAddress string, fromPublicKeys []string, destinations []Destination, ) *BuildMultisig`

NewBuildMultisig instantiates a new BuildMultisig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildMultisigWithDefaults

`func NewBuildMultisigWithDefaults() *BuildMultisig`

NewBuildMultisigWithDefaults instantiates a new BuildMultisig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *BuildMultisig) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *BuildMultisig) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *BuildMultisig) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetFromPublicKeys

`func (o *BuildMultisig) GetFromPublicKeys() []string`

GetFromPublicKeys returns the FromPublicKeys field if non-nil, zero value otherwise.

### GetFromPublicKeysOk

`func (o *BuildMultisig) GetFromPublicKeysOk() (*[]string, bool)`

GetFromPublicKeysOk returns a tuple with the FromPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKeys

`func (o *BuildMultisig) SetFromPublicKeys(v []string)`

SetFromPublicKeys sets FromPublicKeys field to given value.


### GetDestinations

`func (o *BuildMultisig) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *BuildMultisig) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *BuildMultisig) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetGas

`func (o *BuildMultisig) GetGas() int32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *BuildMultisig) GetGasOk() (*int32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *BuildMultisig) SetGas(v int32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *BuildMultisig) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildMultisig) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildMultisig) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildMultisig) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildMultisig) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


