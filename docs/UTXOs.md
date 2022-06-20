# UTXOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Utxos** | [**[]UTXO**](UTXO.md) |  | 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewUTXOs

`func NewUTXOs(utxos []UTXO, ) *UTXOs`

NewUTXOs instantiates a new UTXOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUTXOsWithDefaults

`func NewUTXOsWithDefaults() *UTXOs`

NewUTXOsWithDefaults instantiates a new UTXOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtxos

`func (o *UTXOs) GetUtxos() []UTXO`

GetUtxos returns the Utxos field if non-nil, zero value otherwise.

### GetUtxosOk

`func (o *UTXOs) GetUtxosOk() (*[]UTXO, bool)`

GetUtxosOk returns a tuple with the Utxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxos

`func (o *UTXOs) SetUtxos(v []UTXO)`

SetUtxos sets Utxos field to given value.


### GetWarning

`func (o *UTXOs) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *UTXOs) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *UTXOs) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *UTXOs) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


