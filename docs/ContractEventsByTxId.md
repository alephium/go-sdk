# ContractEventsByTxId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]ContractEventByTxId**](ContractEventByTxId.md) |  | 
**NextStart** | **int32** |  | 

## Methods

### NewContractEventsByTxId

`func NewContractEventsByTxId(events []ContractEventByTxId, nextStart int32, ) *ContractEventsByTxId`

NewContractEventsByTxId instantiates a new ContractEventsByTxId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractEventsByTxIdWithDefaults

`func NewContractEventsByTxIdWithDefaults() *ContractEventsByTxId`

NewContractEventsByTxIdWithDefaults instantiates a new ContractEventsByTxId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ContractEventsByTxId) GetEvents() []ContractEventByTxId`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ContractEventsByTxId) GetEventsOk() (*[]ContractEventByTxId, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ContractEventsByTxId) SetEvents(v []ContractEventByTxId)`

SetEvents sets Events field to given value.


### GetNextStart

`func (o *ContractEventsByTxId) GetNextStart() int32`

GetNextStart returns the NextStart field if non-nil, zero value otherwise.

### GetNextStartOk

`func (o *ContractEventsByTxId) GetNextStartOk() (*int32, bool)`

GetNextStartOk returns a tuple with the NextStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStart

`func (o *ContractEventsByTxId) SetNextStart(v int32)`

SetNextStart sets NextStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


