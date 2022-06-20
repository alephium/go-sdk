# ChainParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **int32** |  | 
**NumZerosAtLeastInHash** | **int32** |  | 
**GroupNumPerBroker** | **int32** |  | 
**Groups** | **int32** |  | 

## Methods

### NewChainParams

`func NewChainParams(networkId int32, numZerosAtLeastInHash int32, groupNumPerBroker int32, groups int32, ) *ChainParams`

NewChainParams instantiates a new ChainParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChainParamsWithDefaults

`func NewChainParamsWithDefaults() *ChainParams`

NewChainParamsWithDefaults instantiates a new ChainParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *ChainParams) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ChainParams) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ChainParams) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.


### GetNumZerosAtLeastInHash

`func (o *ChainParams) GetNumZerosAtLeastInHash() int32`

GetNumZerosAtLeastInHash returns the NumZerosAtLeastInHash field if non-nil, zero value otherwise.

### GetNumZerosAtLeastInHashOk

`func (o *ChainParams) GetNumZerosAtLeastInHashOk() (*int32, bool)`

GetNumZerosAtLeastInHashOk returns a tuple with the NumZerosAtLeastInHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumZerosAtLeastInHash

`func (o *ChainParams) SetNumZerosAtLeastInHash(v int32)`

SetNumZerosAtLeastInHash sets NumZerosAtLeastInHash field to given value.


### GetGroupNumPerBroker

`func (o *ChainParams) GetGroupNumPerBroker() int32`

GetGroupNumPerBroker returns the GroupNumPerBroker field if non-nil, zero value otherwise.

### GetGroupNumPerBrokerOk

`func (o *ChainParams) GetGroupNumPerBrokerOk() (*int32, bool)`

GetGroupNumPerBrokerOk returns a tuple with the GroupNumPerBroker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumPerBroker

`func (o *ChainParams) SetGroupNumPerBroker(v int32)`

SetGroupNumPerBroker sets GroupNumPerBroker field to given value.


### GetGroups

`func (o *ChainParams) GetGroups() int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ChainParams) GetGroupsOk() (*int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ChainParams) SetGroups(v int32)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


