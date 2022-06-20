# BlockEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Timestamp** | **int64** |  | 
**ChainFrom** | **int32** |  | 
**ChainTo** | **int32** |  | 
**Height** | **int32** |  | 
**Deps** | **[]string** |  | 
**Transactions** | [**[]Transaction**](Transaction.md) |  | 
**Nonce** | **string** |  | 
**Version** | **int32** |  | 
**DepStateHash** | **string** |  | 
**TxsHash** | **string** |  | 
**Target** | **string** |  | 

## Methods

### NewBlockEntry

`func NewBlockEntry(hash string, timestamp int64, chainFrom int32, chainTo int32, height int32, deps []string, transactions []Transaction, nonce string, version int32, depStateHash string, txsHash string, target string, ) *BlockEntry`

NewBlockEntry instantiates a new BlockEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockEntryWithDefaults

`func NewBlockEntryWithDefaults() *BlockEntry`

NewBlockEntryWithDefaults instantiates a new BlockEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *BlockEntry) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockEntry) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockEntry) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *BlockEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetChainFrom

`func (o *BlockEntry) GetChainFrom() int32`

GetChainFrom returns the ChainFrom field if non-nil, zero value otherwise.

### GetChainFromOk

`func (o *BlockEntry) GetChainFromOk() (*int32, bool)`

GetChainFromOk returns a tuple with the ChainFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFrom

`func (o *BlockEntry) SetChainFrom(v int32)`

SetChainFrom sets ChainFrom field to given value.


### GetChainTo

`func (o *BlockEntry) GetChainTo() int32`

GetChainTo returns the ChainTo field if non-nil, zero value otherwise.

### GetChainToOk

`func (o *BlockEntry) GetChainToOk() (*int32, bool)`

GetChainToOk returns a tuple with the ChainTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTo

`func (o *BlockEntry) SetChainTo(v int32)`

SetChainTo sets ChainTo field to given value.


### GetHeight

`func (o *BlockEntry) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockEntry) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockEntry) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDeps

`func (o *BlockEntry) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *BlockEntry) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *BlockEntry) SetDeps(v []string)`

SetDeps sets Deps field to given value.


### GetTransactions

`func (o *BlockEntry) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BlockEntry) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BlockEntry) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetNonce

`func (o *BlockEntry) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *BlockEntry) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *BlockEntry) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetVersion

`func (o *BlockEntry) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockEntry) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockEntry) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDepStateHash

`func (o *BlockEntry) GetDepStateHash() string`

GetDepStateHash returns the DepStateHash field if non-nil, zero value otherwise.

### GetDepStateHashOk

`func (o *BlockEntry) GetDepStateHashOk() (*string, bool)`

GetDepStateHashOk returns a tuple with the DepStateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepStateHash

`func (o *BlockEntry) SetDepStateHash(v string)`

SetDepStateHash sets DepStateHash field to given value.


### GetTxsHash

`func (o *BlockEntry) GetTxsHash() string`

GetTxsHash returns the TxsHash field if non-nil, zero value otherwise.

### GetTxsHashOk

`func (o *BlockEntry) GetTxsHashOk() (*string, bool)`

GetTxsHashOk returns a tuple with the TxsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxsHash

`func (o *BlockEntry) SetTxsHash(v string)`

SetTxsHash sets TxsHash field to given value.


### GetTarget

`func (o *BlockEntry) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BlockEntry) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BlockEntry) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


