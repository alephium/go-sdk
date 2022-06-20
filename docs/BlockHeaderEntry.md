# BlockHeaderEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Timestamp** | **int64** |  | 
**ChainFrom** | **int32** |  | 
**ChainTo** | **int32** |  | 
**Height** | **int32** |  | 
**Deps** | **[]string** |  | 

## Methods

### NewBlockHeaderEntry

`func NewBlockHeaderEntry(hash string, timestamp int64, chainFrom int32, chainTo int32, height int32, deps []string, ) *BlockHeaderEntry`

NewBlockHeaderEntry instantiates a new BlockHeaderEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockHeaderEntryWithDefaults

`func NewBlockHeaderEntryWithDefaults() *BlockHeaderEntry`

NewBlockHeaderEntryWithDefaults instantiates a new BlockHeaderEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *BlockHeaderEntry) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *BlockHeaderEntry) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *BlockHeaderEntry) SetHash(v string)`

SetHash sets Hash field to given value.


### GetTimestamp

`func (o *BlockHeaderEntry) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlockHeaderEntry) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlockHeaderEntry) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetChainFrom

`func (o *BlockHeaderEntry) GetChainFrom() int32`

GetChainFrom returns the ChainFrom field if non-nil, zero value otherwise.

### GetChainFromOk

`func (o *BlockHeaderEntry) GetChainFromOk() (*int32, bool)`

GetChainFromOk returns a tuple with the ChainFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFrom

`func (o *BlockHeaderEntry) SetChainFrom(v int32)`

SetChainFrom sets ChainFrom field to given value.


### GetChainTo

`func (o *BlockHeaderEntry) GetChainTo() int32`

GetChainTo returns the ChainTo field if non-nil, zero value otherwise.

### GetChainToOk

`func (o *BlockHeaderEntry) GetChainToOk() (*int32, bool)`

GetChainToOk returns a tuple with the ChainTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTo

`func (o *BlockHeaderEntry) SetChainTo(v int32)`

SetChainTo sets ChainTo field to given value.


### GetHeight

`func (o *BlockHeaderEntry) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BlockHeaderEntry) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BlockHeaderEntry) SetHeight(v int32)`

SetHeight sets Height field to given value.


### GetDeps

`func (o *BlockHeaderEntry) GetDeps() []string`

GetDeps returns the Deps field if non-nil, zero value otherwise.

### GetDepsOk

`func (o *BlockHeaderEntry) GetDepsOk() (*[]string, bool)`

GetDepsOk returns a tuple with the Deps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeps

`func (o *BlockHeaderEntry) SetDeps(v []string)`

SetDeps sets Deps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


