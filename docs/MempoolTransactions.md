# MempoolTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromGroup** | **int32** |  | 
**ToGroup** | **int32** |  | 
**Transactions** | [**[]TransactionTemplate**](TransactionTemplate.md) |  | 

## Methods

### NewMempoolTransactions

`func NewMempoolTransactions(fromGroup int32, toGroup int32, transactions []TransactionTemplate, ) *MempoolTransactions`

NewMempoolTransactions instantiates a new MempoolTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTransactionsWithDefaults

`func NewMempoolTransactionsWithDefaults() *MempoolTransactions`

NewMempoolTransactionsWithDefaults instantiates a new MempoolTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromGroup

`func (o *MempoolTransactions) GetFromGroup() int32`

GetFromGroup returns the FromGroup field if non-nil, zero value otherwise.

### GetFromGroupOk

`func (o *MempoolTransactions) GetFromGroupOk() (*int32, bool)`

GetFromGroupOk returns a tuple with the FromGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroup

`func (o *MempoolTransactions) SetFromGroup(v int32)`

SetFromGroup sets FromGroup field to given value.


### GetToGroup

`func (o *MempoolTransactions) GetToGroup() int32`

GetToGroup returns the ToGroup field if non-nil, zero value otherwise.

### GetToGroupOk

`func (o *MempoolTransactions) GetToGroupOk() (*int32, bool)`

GetToGroupOk returns a tuple with the ToGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroup

`func (o *MempoolTransactions) SetToGroup(v int32)`

SetToGroup sets ToGroup field to given value.


### GetTransactions

`func (o *MempoolTransactions) GetTransactions() []TransactionTemplate`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MempoolTransactions) GetTransactionsOk() (*[]TransactionTemplate, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MempoolTransactions) SetTransactions(v []TransactionTemplate)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


