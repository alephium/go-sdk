# NodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildInfo** | [**BuildInfo**](BuildInfo.md) |  | 
**Upnp** | **bool** |  | 
**ExternalAddress** | Pointer to [**BrokerInfoAddress**](BrokerInfoAddress.md) |  | [optional] 

## Methods

### NewNodeInfo

`func NewNodeInfo(buildInfo BuildInfo, upnp bool, ) *NodeInfo`

NewNodeInfo instantiates a new NodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInfoWithDefaults

`func NewNodeInfoWithDefaults() *NodeInfo`

NewNodeInfoWithDefaults instantiates a new NodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildInfo

`func (o *NodeInfo) GetBuildInfo() BuildInfo`

GetBuildInfo returns the BuildInfo field if non-nil, zero value otherwise.

### GetBuildInfoOk

`func (o *NodeInfo) GetBuildInfoOk() (*BuildInfo, bool)`

GetBuildInfoOk returns a tuple with the BuildInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildInfo

`func (o *NodeInfo) SetBuildInfo(v BuildInfo)`

SetBuildInfo sets BuildInfo field to given value.


### GetUpnp

`func (o *NodeInfo) GetUpnp() bool`

GetUpnp returns the Upnp field if non-nil, zero value otherwise.

### GetUpnpOk

`func (o *NodeInfo) GetUpnpOk() (*bool, bool)`

GetUpnpOk returns a tuple with the Upnp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpnp

`func (o *NodeInfo) SetUpnp(v bool)`

SetUpnp sets Upnp field to given value.


### GetExternalAddress

`func (o *NodeInfo) GetExternalAddress() BrokerInfoAddress`

GetExternalAddress returns the ExternalAddress field if non-nil, zero value otherwise.

### GetExternalAddressOk

`func (o *NodeInfo) GetExternalAddressOk() (*BrokerInfoAddress, bool)`

GetExternalAddressOk returns a tuple with the ExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddress

`func (o *NodeInfo) SetExternalAddress(v BrokerInfoAddress)`

SetExternalAddress sets ExternalAddress field to given value.

### HasExternalAddress

`func (o *NodeInfo) HasExternalAddress() bool`

HasExternalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


