# UnconfirmedTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 
**UnconfirmedTransactions** | [**[]TransactionTemplate**](TransactionTemplate.md) |  | 

## Methods

### NewUnconfirmedTransactions

`func NewUnconfirmedTransactions(fromGroup int32, toGroup int32, unconfirmedTransactions []TransactionTemplate, ) *UnconfirmedTransactions`

NewUnconfirmedTransactions instantiates a new UnconfirmedTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfirmedTransactionsWithDefaults

`func NewUnconfirmedTransactionsWithDefaults() *UnconfirmedTransactions`

NewUnconfirmedTransactionsWithDefaults instantiates a new UnconfirmedTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromGroup

`func (o *UnconfirmedTransactions) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *UnconfirmedTransactions) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *UnconfirmedTransactions) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *UnconfirmedTransactions) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *UnconfirmedTransactions) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *UnconfirmedTransactions) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetUnconfirmedTransactions

`func (o *UnconfirmedTransactions) GetUnconfirmedTransactions() []TransactionTemplate`

GetUnconfirmedTransactions returns the UnconfirmedTransactions field if non-nil, zero value otherwise.

### GetUnconfirmedTransactionsOk

`func (o *UnconfirmedTransactions) GetUnconfirmedTransactionsOk() (*[]TransactionTemplate, bool)`

GetUnconfirmedTransactionsOk returns a tuple with the UnconfirmedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmedTransactions

`func (o *UnconfirmedTransactions) SetUnconfirmedTransactions(v []TransactionTemplate)`

SetUnconfirmedTransactions sets UnconfirmedTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


