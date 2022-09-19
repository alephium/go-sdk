# BlockAndEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | [**BlockEntry**](BlockEntry.md) |  | 
**Events** | [**[]ContractEventByBlockHash**](ContractEventByBlockHash.md) |  | 

## Methods

### NewBlockAndEvents

`func NewBlockAndEvents(block BlockEntry, events []ContractEventByBlockHash, ) *BlockAndEvents`

NewBlockAndEvents instantiates a new BlockAndEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockAndEventsWithDefaults

`func NewBlockAndEventsWithDefaults() *BlockAndEvents`

NewBlockAndEventsWithDefaults instantiates a new BlockAndEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *BlockAndEvents) GetBlock() BlockEntry`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *BlockAndEvents) GetBlockOk() (*BlockEntry, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *BlockAndEvents) SetBlock(v BlockEntry)`

SetBlock sets Block field to given value.


### GetEvents

`func (o *BlockAndEvents) GetEvents() []ContractEventByBlockHash`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BlockAndEvents) GetEventsOk() (*[]ContractEventByBlockHash, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BlockAndEvents) SetEvents(v []ContractEventByBlockHash)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


