# BuildTransferTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPublicKey** | **string** |  | 
**FromPublicKeyType** | Pointer to **string** |  | [optional] 
**Destinations** | [**[]Destination**](Destination.md) |  | 
**Utxos** | Pointer to [**[]OutputRef**](OutputRef.md) |  | [optional] 
**GasAmount** | Pointer to **int32** |  | [optional] 
**GasPrice** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **int32** |  | [optional] 
**TargetBlockHash** | Pointer to **string** |  | [optional] 

## Methods

### NewBuildTransferTx

`func NewBuildTransferTx(fromPublicKey string, destinations []Destination, ) *BuildTransferTx`

NewBuildTransferTx instantiates a new BuildTransferTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildTransferTxWithDefaults

`func NewBuildTransferTxWithDefaults() *BuildTransferTx`

NewBuildTransferTxWithDefaults instantiates a new BuildTransferTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPublicKey

`func (o *BuildTransferTx) GetFromPublicKey() string`

GetFromPublicKey returns the FromPublicKey field if non-nil, zero value otherwise.

### GetFromPublicKeyOk

`func (o *BuildTransferTx) GetFromPublicKeyOk() (*string, bool)`

GetFromPublicKeyOk returns a tuple with the FromPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKey

`func (o *BuildTransferTx) SetFromPublicKey(v string)`

SetFromPublicKey sets FromPublicKey field to given value.


### GetFromPublicKeyType

`func (o *BuildTransferTx) GetFromPublicKeyType() string`

GetFromPublicKeyType returns the FromPublicKeyType field if non-nil, zero value otherwise.

### GetFromPublicKeyTypeOk

`func (o *BuildTransferTx) GetFromPublicKeyTypeOk() (*string, bool)`

GetFromPublicKeyTypeOk returns a tuple with the FromPublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPublicKeyType

`func (o *BuildTransferTx) SetFromPublicKeyType(v string)`

SetFromPublicKeyType sets FromPublicKeyType field to given value.

### HasFromPublicKeyType

`func (o *BuildTransferTx) HasFromPublicKeyType() bool`

HasFromPublicKeyType returns a boolean if a field has been set.

### GetDestinations

`func (o *BuildTransferTx) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *BuildTransferTx) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *BuildTransferTx) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.


### GetUtxos

`func (o *BuildTransferTx) GetUtxos() []OutputRef`

GetUtxos returns the Utxos field if non-nil, zero value otherwise.

### GetUtxosOk

`func (o *BuildTransferTx) GetUtxosOk() (*[]OutputRef, bool)`

GetUtxosOk returns a tuple with the Utxos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxos

`func (o *BuildTransferTx) SetUtxos(v []OutputRef)`

SetUtxos sets Utxos field to given value.

### HasUtxos

`func (o *BuildTransferTx) HasUtxos() bool`

HasUtxos returns a boolean if a field has been set.

### GetGasAmount

`func (o *BuildTransferTx) GetGasAmount() int32`

GetGasAmount returns the GasAmount field if non-nil, zero value otherwise.

### GetGasAmountOk

`func (o *BuildTransferTx) GetGasAmountOk() (*int32, bool)`

GetGasAmountOk returns a tuple with the GasAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasAmount

`func (o *BuildTransferTx) SetGasAmount(v int32)`

SetGasAmount sets GasAmount field to given value.

### HasGasAmount

`func (o *BuildTransferTx) HasGasAmount() bool`

HasGasAmount returns a boolean if a field has been set.

### GetGasPrice

`func (o *BuildTransferTx) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *BuildTransferTx) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *BuildTransferTx) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *BuildTransferTx) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetGroup

`func (o *BuildTransferTx) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BuildTransferTx) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BuildTransferTx) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BuildTransferTx) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTargetBlockHash

`func (o *BuildTransferTx) GetTargetBlockHash() string`

GetTargetBlockHash returns the TargetBlockHash field if non-nil, zero value otherwise.

### GetTargetBlockHashOk

`func (o *BuildTransferTx) GetTargetBlockHashOk() (*string, bool)`

GetTargetBlockHashOk returns a tuple with the TargetBlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBlockHash

`func (o *BuildTransferTx) SetTargetBlockHash(v string)`

SetTargetBlockHash sets TargetBlockHash field to given value.

### HasTargetBlockHash

`func (o *BuildTransferTx) HasTargetBlockHash() bool`

HasTargetBlockHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


