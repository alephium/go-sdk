# BuildSweepAddressTransactionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnsignedTxs** | [**[]SweepAddressTransaction**](SweepAddressTransaction.md) |  | 
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 

## Methods

### NewBuildSweepAddressTransactionsResult

`func NewBuildSweepAddressTransactionsResult(unsignedTxs []SweepAddressTransaction, fromGroup int32, toGroup int32, ) *BuildSweepAddressTransactionsResult`

NewBuildSweepAddressTransactionsResult instantiates a new BuildSweepAddressTransactionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildSweepAddressTransactionsResultWithDefaults

`func NewBuildSweepAddressTransactionsResultWithDefaults() *BuildSweepAddressTransactionsResult`

NewBuildSweepAddressTransactionsResultWithDefaults instantiates a new BuildSweepAddressTransactionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnsignedTxs

`func (o *BuildSweepAddressTransactionsResult) GetUnsignedTxs() []SweepAddressTransaction`

GetUnsignedTxs returns the UnsignedTxs field if non-nil, zero value otherwise.

### GetUnsignedTxsOk

`func (o *BuildSweepAddressTransactionsResult) GetUnsignedTxsOk() (*[]SweepAddressTransaction, bool)`

GetUnsignedTxsOk returns a tuple with the UnsignedTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsignedTxs

`func (o *BuildSweepAddressTransactionsResult) SetUnsignedTxs(v []SweepAddressTransaction)`

SetUnsignedTxs sets UnsignedTxs field to given value.


### GetFromGroup

`func (o *BuildSweepAddressTransactionsResult) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *BuildSweepAddressTransactionsResult) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *BuildSweepAddressTransactionsResult) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *BuildSweepAddressTransactionsResult) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *BuildSweepAddressTransactionsResult) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *BuildSweepAddressTransactionsResult) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


