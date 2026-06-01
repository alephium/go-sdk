# RichBlockEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Timestamp** | **int64** |  | 
**ChainFrom** | **int32** |  | 
**ChainTo** | **int32** |  | 
**Height** | **int32** |  | 
**Deps** | **[]string** |  | 
**Transactions** | [**[]RichTransaction**](RichTransaction.md) |  | 
**Nonce** | **string** |  | 
**Version** | **int32** |  | 
**DepStateHash** | **string** |  | 
**TxsHash** | **string** |  | 
**Target** | **string** |  | 
**GhostUncles** | [**[]GhostUncleBlockEntry**](GhostUncleBlockEntry.md) |  | 
**ConflictedTxs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRichBlockEntry

`func NewRichBlockEntry(hash string, timestamp int64, chainFrom int32, chainTo int32, height int32, deps []string, transactions []RichTransaction, nonce string, version int32, depStateHash string, txsHash string, target string, ghostUncles []GhostUncleBlockEntry, ) *RichBlockEntry`

NewRichBlockEntry instantiates a new RichBlockEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichBlockEntryWithDefaults

`func NewRichBlockEntryWithDefaults() *RichBlockEntry`

NewRichBlockEntryWithDefaults instantiates a new RichBlockEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *RichBlockEntry) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *RichBlockEntry) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *RichBlockEntry) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *RichBlockEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RichBlockEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RichBlockEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetChainFrom

`func (o *RichBlockEntry) GetChainFrom() int32`

GetChainFrom returns the ChainFrom field if non-nil, zero value otherwise.

### GetChainFromOk

`func (o *RichBlockEntry) GetChainFromOk() (*int32, bool)`

GetChainFromOk returns a tuple with the ChainFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFrom

`func (o *RichBlockEntry) SetChainFrom(v int32)`

SetChainFrom sets ChainFrom field to given value.


### GetChainTo

`func (o *RichBlockEntry) GetChainTo() int32`

GetChainTo returns the ChainTo field if non-nil, zero value otherwise.

### GetChainToOk

`func (o *RichBlockEntry) GetChainToOk() (*int32, bool)`

GetChainToOk returns a tuple with the ChainTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTo

`func (o *RichBlockEntry) SetChainTo(v int32)`

SetChainTo sets ChainTo field to given value.


### GetHeight

`func (o *RichBlockEntry) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *RichBlockEntry) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *RichBlockEntry) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDeps

`func (o *RichBlockEntry) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *RichBlockEntry) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *RichBlockEntry) SetDeps(v []string)`

SetDeps sets Deps field to given value.


### GetTransactions

`func (o *RichBlockEntry) GetTransactions() []RichTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *RichBlockEntry) GetTransactionsOk() (*[]RichTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *RichBlockEntry) SetTransactions(v []RichTransaction)`

SetTransactions sets Transactions field to given value.


### GetNonce

`func (o *RichBlockEntry) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *RichBlockEntry) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *RichBlockEntry) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetVersion

`func (o *RichBlockEntry) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RichBlockEntry) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RichBlockEntry) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetDepStateHash

`func (o *RichBlockEntry) GetDepStateHash() string`

GetDepStateHash returns the DepStateHash field if non-nil, zero value otherwise.

### GetDepStateHashOk

`func (o *RichBlockEntry) GetDepStateHashOk() (*string, bool)`

GetDepStateHashOk returns a tuple with the DepStateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepStateHash

`func (o *RichBlockEntry) SetDepStateHash(v string)`

SetDepStateHash sets DepStateHash field to given value.


### GetTxsHash

`func (o *RichBlockEntry) GetTxsHash() string`

GetTxsHash returns the TxsHash field if non-nil, zero value otherwise.

### GetTxsHashOk

`func (o *RichBlockEntry) GetTxsHashOk() (*string, bool)`

GetTxsHashOk returns a tuple with the TxsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxsHash

`func (o *RichBlockEntry) SetTxsHash(v string)`

SetTxsHash sets TxsHash field to given value.


### GetTarget

`func (o *RichBlockEntry) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RichBlockEntry) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RichBlockEntry) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetGhostUncles

`func (o *RichBlockEntry) GetGhostUncles() []GhostUncleBlockEntry`

GetGhostUncles returns the GhostUncles field if non-nil, zero value otherwise.

### GetGhostUnclesOk

`func (o *RichBlockEntry) GetGhostUnclesOk() (*[]GhostUncleBlockEntry, bool)`

GetGhostUnclesOk returns a tuple with the GhostUncles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhostUncles

`func (o *RichBlockEntry) SetGhostUncles(v []GhostUncleBlockEntry)`

SetGhostUncles sets GhostUncles field to given value.


### GetConflictedTxs

`func (o *RichBlockEntry) GetConflictedTxs() []string`

GetConflictedTxs returns the ConflictedTxs field if non-nil, zero value otherwise.

### GetConflictedTxsOk

`func (o *RichBlockEntry) GetConflictedTxsOk() (*[]string, bool)`

GetConflictedTxsOk returns a tuple with the ConflictedTxs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictedTxs

`func (o *RichBlockEntry) SetConflictedTxs(v []string)`

SetConflictedTxs sets ConflictedTxs field to given value.

### HasConflictedTxs

`func (o *RichBlockEntry) HasConflictedTxs() bool`

HasConflictedTxs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


