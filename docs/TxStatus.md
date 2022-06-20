# TxStatus

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

### NewTxStatus

`func NewTxStatus(blockHash string, txIndex int32, chainConfirmations int32, fromGroupConfirmations int32, toGroupConfirmations int32, type_ string, ) *TxStatus`

NewTxStatus instantiates a new TxStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTxStatusWithDefaults

`func NewTxStatusWithDefaults() *TxStatus`

NewTxStatusWithDefaults instantiates a new TxStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockHash

`func (o *TxStatus) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *TxStatus) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *TxStatus) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.


### GetTxIndex

`func (o *TxStatus) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *TxStatus) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *TxStatus) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.


### GetChainConfirmations

`func (o *TxStatus) GetChainConfirmations() int32`

GetChainConfirmations returns the ChainConfirmations field if non-nil, zero value otherwise.

### GetChainConfirmationsOk

`func (o *TxStatus) GetChainConfirmationsOk() (*int32, bool)`

GetChainConfirmationsOk returns a tuple with the ChainConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainConfirmations

`func (o *TxStatus) SetChainConfirmations(v int32)`

SetChainConfirmations sets ChainConfirmations field to given value.


### GetFromGroupConfirmations

`func (o *TxStatus) GetFromGroupConfirmations() int32`

GetFromGroupConfirmations returns the FromGroupConfirmations field if non-nil, zero value otherwise.

### GetFromGroupConfirmationsOk

`func (o *TxStatus) GetFromGroupConfirmationsOk() (*int32, bool)`

GetFromGroupConfirmationsOk returns a tuple with the FromGroupConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGroupConfirmations

`func (o *TxStatus) SetFromGroupConfirmations(v int32)`

SetFromGroupConfirmations sets FromGroupConfirmations field to given value.


### GetToGroupConfirmations

`func (o *TxStatus) GetToGroupConfirmations() int32`

GetToGroupConfirmations returns the ToGroupConfirmations field if non-nil, zero value otherwise.

### GetToGroupConfirmationsOk

`func (o *TxStatus) GetToGroupConfirmationsOk() (*int32, bool)`

GetToGroupConfirmationsOk returns a tuple with the ToGroupConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToGroupConfirmations

`func (o *TxStatus) SetToGroupConfirmations(v int32)`

SetToGroupConfirmations sets ToGroupConfirmations field to given value.


### GetType

`func (o *TxStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TxStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TxStatus) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


