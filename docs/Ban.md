# Ban

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Peers** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewBan

`func NewBan(peers []string, type_ string, ) *Ban`

NewBan instantiates a new Ban object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBanWithDefaults

`func NewBanWithDefaults() *Ban`

NewBanWithDefaults instantiates a new Ban object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeers

`func (o *Ban) GetPeers() []string`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *Ban) GetPeersOk() (*[]string, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *Ban) SetPeers(v []string)`

SetPeers sets Peers field to given value.


### GetType

`func (o *Ban) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Ban) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Ban) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


