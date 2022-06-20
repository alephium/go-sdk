# PeerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Until** | **int64** |  | 
**Type** | **string** |  | 
**Value** | **int32** |  | 

## Methods

### NewPeerStatus

`func NewPeerStatus(until int64, type_ string, value int32, ) *PeerStatus`

NewPeerStatus instantiates a new PeerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerStatusWithDefaults

`func NewPeerStatusWithDefaults() *PeerStatus`

NewPeerStatusWithDefaults instantiates a new PeerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUntil

`func (o *PeerStatus) GetUntil() int64`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *PeerStatus) GetUntilOk() (*int64, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *PeerStatus) SetUntil(v int64)`

SetUntil sets Until field to given value.


### GetType

`func (o *PeerStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PeerStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PeerStatus) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *PeerStatus) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PeerStatus) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PeerStatus) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


