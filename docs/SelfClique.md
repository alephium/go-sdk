# SelfClique

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliqueId** | **string** |  | 
**Nodes** | [**[]PeerAddress**](PeerAddress.md) |  | 
**SelfReady** | **bool** |  | 
**Synced** | **bool** |  | 

## Methods

### NewSelfClique

`func NewSelfClique(cliqueId string, nodes []PeerAddress, selfReady bool, synced bool, ) *SelfClique`

NewSelfClique instantiates a new SelfClique object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfCliqueWithDefaults

`func NewSelfCliqueWithDefaults() *SelfClique`

NewSelfCliqueWithDefaults instantiates a new SelfClique object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliqueId

`func (o *SelfClique) GetCliqueId() string`

GetCliqueId returns the CliqueId field if non-nil, zero value otherwise.

### GetCliqueIdOk

`func (o *SelfClique) GetCliqueIdOk() (*string, bool)`

GetCliqueIdOk returns a tuple with the CliqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliqueId

`func (o *SelfClique) SetCliqueId(v string)`

SetCliqueId sets CliqueId field to given value.


### GetNodes

`func (o *SelfClique) GetNodes() []PeerAddress`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SelfClique) GetNodesOk() (*[]PeerAddress, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SelfClique) SetNodes(v []PeerAddress)`

SetNodes sets Nodes field to given value.


### GetSelfReady

`func (o *SelfClique) GetSelfReady() bool`

GetSelfReady returns the SelfReady field if non-nil, zero value otherwise.

### GetSelfReadyOk

`func (o *SelfClique) GetSelfReadyOk() (*bool, bool)`

GetSelfReadyOk returns a tuple with the SelfReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfReady

`func (o *SelfClique) SetSelfReady(v bool)`

SetSelfReady sets SelfReady field to given value.


### GetSynced

`func (o *SelfClique) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *SelfClique) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *SelfClique) SetSynced(v bool)`

SetSynced sets Synced field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


