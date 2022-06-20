# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | [**[]Destination**](Destination.md) |  | 
**Gas** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**UtxosLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransfer

`func NewTransfer(destinations []Destination, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *Transfer) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *Transfer) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *Transfer) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetGas

`func (o *Transfer) GetGas() int32`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *Transfer) GetGasOk() (*int32, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *Transfer) SetGas(v int32)`

SetGas sets Gas field to given value.

### HasGas

`func (o *Transfer) HasGas() bool`

HasGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *Transfer) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *Transfer) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *Transfer) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *Transfer) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetUtxosLimit

`func (o *Transfer) GetUtxosLimit() int32`

GetUtxosLimit returns the UtxosLimit field if non-nil, zero value otherwise.

### GetUtxosLimitOk

`func (o *Transfer) GetUtxosLimitOk() (*int32, bool)`

GetUtxosLimitOk returns a tuple with the UtxosLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxosLimit

`func (o *Transfer) SetUtxosLimit(v int32)`

SetUtxosLimit sets UtxosLimit field to given value.

### HasUtxosLimit

`func (o *Transfer) HasUtxosLimit() bool`

HasUtxosLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


