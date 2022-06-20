# PeerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**RestPort** | **int32** |  | 
**WsPort** | **int32** |  | 
**MinerApiPort** | **int32** |  | 

## Methods

### NewPeerAddress

`func NewPeerAddress(address string, restPort int32, wsPort int32, minerApiPort int32, ) *PeerAddress`

NewPeerAddress instantiates a new PeerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerAddressWithDefaults

`func NewPeerAddressWithDefaults() *PeerAddress`

NewPeerAddressWithDefaults instantiates a new PeerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PeerAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PeerAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PeerAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetRestPort

`func (o *PeerAddress) GetRestPort() int32`

GetRestPort returns the RestPort field if non-nil, zero value otherwise.

### GetRestPortOk

`func (o *PeerAddress) GetRestPortOk() (*int32, bool)`

GetRestPortOk returns a tuple with the RestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPort

`func (o *PeerAddress) SetRestPort(v int32)`

SetRestPort sets RestPort field to given value.


### GetWsPort

`func (o *PeerAddress) GetWsPort() int32`

GetWsPort returns the WsPort field if non-nil, zero value otherwise.

### GetWsPortOk

`func (o *PeerAddress) GetWsPortOk() (*int32, bool)`

GetWsPortOk returns a tuple with the WsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsPort

`func (o *PeerAddress) SetWsPort(v int32)`

SetWsPort sets WsPort field to given value.


### GetMinerApiPort

`func (o *PeerAddress) GetMinerApiPort() int32`

GetMinerApiPort returns the MinerApiPort field if non-nil, zero value otherwise.

### GetMinerApiPortOk

`func (o *PeerAddress) GetMinerApiPortOk() (*int32, bool)`

GetMinerApiPortOk returns a tuple with the MinerApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinerApiPort

`func (o *PeerAddress) SetMinerApiPort(v int32)`

SetMinerApiPort sets MinerApiPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


