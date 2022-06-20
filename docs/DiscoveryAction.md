# DiscoveryAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peers** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewDiscoveryAction

`func NewDiscoveryAction(peers []string, type_ string, ) *DiscoveryAction`

NewDiscoveryAction instantiates a new DiscoveryAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryActionWithDefaults

`func NewDiscoveryActionWithDefaults() *DiscoveryAction`

NewDiscoveryActionWithDefaults instantiates a new DiscoveryAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeers

`func (o *DiscoveryAction) GetPeers() []string`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *DiscoveryAction) GetPeersOk() (*[]string, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *DiscoveryAction) SetPeers(v []string)`

SetPeers sets Peers field to given value.


### GetType

`func (o *DiscoveryAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveryAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveryAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


