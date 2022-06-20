# Confirmed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHash** | **string** |  | 
**TxIndex** | **int32** |  | 
**ChainConfirmations** | **int32** |  | 
**FromGroupConfirmations** | **int32** |  | 
**ToGroupConfirmations** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewConfirmed

`func NewConfirmed(blockHash string, txIndex int32, chainConfirmations int32, fromGroupConfirmations int32, toGroupConfirmations int32, type_ string, ) *Confirmed`

NewConfirmed instantiates a new Confirmed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmedWithDefaults

`func NewConfirmedWithDefaults() *Confirmed`

NewConfirmedWithDefaults instantiates a new Confirmed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *Confirmed) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *Confirmed) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *Confirmed) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetTxIndex

`func (o *Confirmed) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *Confirmed) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *Confirmed) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetChainConfirmations

`func (o *Confirmed) GetChainConfirmations() int32`

GetChainConfirmations returns the ChainConfirmations field if non-nil, zero value otherwise.

### GetChainConfirmationsOk

`func (o *Confirmed) GetChainConfirmationsOk() (*int32, bool)`

GetChainConfirmationsOk returns a tuple with the ChainConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainConfirmations

`func (o *Confirmed) SetChainConfirmations(v int32)`

SetChainConfirmations sets ChainConfirmations field to given value.


### GetFromGroupConfirmations

`func (o *Confirmed) GetFromGroupConfirmations() int32`

GetFromGroupConfirmations returns the FromGroupConfirmations field if non-nil, zero value otherwise.

### GetFromGroupConfirmationsOk

`func (o *Confirmed) GetFromGroupConfirmationsOk() (*int32, bool)`

GetFromGroupConfirmationsOk returns a tuple with the FromGroupConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroupConfirmations

`func (o *Confirmed) SetFromGroupConfirmations(v int32)`

SetFromGroupConfirmations sets FromGroupConfirmations field to given value.


### GetToGroupConfirmations

`func (o *Confirmed) GetToGroupConfirmations() int32`

GetToGroupConfirmations returns the ToGroupConfirmations field if non-nil, zero value otherwise.

### GetToGroupConfirmationsOk

`func (o *Confirmed) GetToGroupConfirmationsOk() (*int32, bool)`

GetToGroupConfirmationsOk returns a tuple with the ToGroupConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroupConfirmations

`func (o *Confirmed) SetToGroupConfirmations(v int32)`

SetToGroupConfirmations sets ToGroupConfirmations field to given value.


### GetType

`func (o *Confirmed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Confirmed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Confirmed) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


