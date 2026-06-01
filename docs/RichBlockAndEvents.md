# RichBlockAndEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | [**RichBlockEntry**](RichBlockEntry.md) |  | 
**Events** | [**[]ContractEventByBlockHash**](ContractEventByBlockHash.md) |  | 

## Methods

### NewRichBlockAndEvents

`func NewRichBlockAndEvents(block RichBlockEntry, events []ContractEventByBlockHash, ) *RichBlockAndEvents`

NewRichBlockAndEvents instantiates a new RichBlockAndEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichBlockAndEventsWithDefaults

`func NewRichBlockAndEventsWithDefaults() *RichBlockAndEvents`

NewRichBlockAndEventsWithDefaults instantiates a new RichBlockAndEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *RichBlockAndEvents) GetBlock() RichBlockEntry`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *RichBlockAndEvents) GetBlockOk() (*RichBlockEntry, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *RichBlockAndEvents) SetBlock(v RichBlockEntry)`

SetBlock sets Block field to given value.


### GetEvents

`func (o *RichBlockAndEvents) GetEvents() []ContractEventByBlockHash`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RichBlockAndEvents) GetEventsOk() (*[]ContractEventByBlockHash, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RichBlockAndEvents) SetEvents(v []ContractEventByBlockHash)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


