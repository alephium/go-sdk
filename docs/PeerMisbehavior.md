# PeerMisbehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peer** | **string** |  | 
**Status** | [**PeerStatus**](PeerStatus.md) |  | 

## Methods

### NewPeerMisbehavior

`func NewPeerMisbehavior(peer string, status PeerStatus, ) *PeerMisbehavior`

NewPeerMisbehavior instantiates a new PeerMisbehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerMisbehaviorWithDefaults

`func NewPeerMisbehaviorWithDefaults() *PeerMisbehavior`

NewPeerMisbehaviorWithDefaults instantiates a new PeerMisbehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeer

`func (o *PeerMisbehavior) GetPeer() string`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *PeerMisbehavior) GetPeerOk() (*string, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *PeerMisbehavior) SetPeer(v string)`

SetPeer sets Peer field to given value.


### GetStatus

`func (o *PeerMisbehavior) GetStatus() PeerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PeerMisbehavior) GetStatusOk() (*PeerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PeerMisbehavior) SetStatus(v PeerStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


