# InterCliquePeerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliqueId** | **string** |  | 
**BrokerId** | **int32** |  | 
**GroupNumPerBroker** | **int32** |  | 
**Address** | [**BrokerInfoAddress**](BrokerInfoAddress.md) |  | 
**IsSynced** | **bool** |  | 
**ClientVersion** | **string** |  | 

## Methods

### NewInterCliquePeerInfo

`func NewInterCliquePeerInfo(cliqueId string, brokerId int32, groupNumPerBroker int32, address BrokerInfoAddress, isSynced bool, clientVersion string, ) *InterCliquePeerInfo`

NewInterCliquePeerInfo instantiates a new InterCliquePeerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterCliquePeerInfoWithDefaults

`func NewInterCliquePeerInfoWithDefaults() *InterCliquePeerInfo`

NewInterCliquePeerInfoWithDefaults instantiates a new InterCliquePeerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliqueId

`func (o *InterCliquePeerInfo) GetCliqueId() string`

GetCliqueId returns the CliqueId field if non-nil, zero value otherwise.

### GetCliqueIdOk

`func (o *InterCliquePeerInfo) GetCliqueIdOk() (*string, bool)`

GetCliqueIdOk returns a tuple with the CliqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliqueId

`func (o *InterCliquePeerInfo) SetCliqueId(v string)`

SetCliqueId sets CliqueId field to given value.


### GetBrokerId

`func (o *InterCliquePeerInfo) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *InterCliquePeerInfo) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *InterCliquePeerInfo) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetGroupNumPerBroker

`func (o *InterCliquePeerInfo) GetGroupNumPerBroker() int32`

GetGroupNumPerBroker returns the GroupNumPerBroker field if non-nil, zero value otherwise.

### GetGroupNumPerBrokerOk

`func (o *InterCliquePeerInfo) GetGroupNumPerBrokerOk() (*int32, bool)`

GetGroupNumPerBrokerOk returns a tuple with the GroupNumPerBroker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNumPerBroker

`func (o *InterCliquePeerInfo) SetGroupNumPerBroker(v int32)`

SetGroupNumPerBroker sets GroupNumPerBroker field to given value.


### GetAddress

`func (o *InterCliquePeerInfo) GetAddress() BrokerInfoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InterCliquePeerInfo) GetAddressOk() (*BrokerInfoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InterCliquePeerInfo) SetAddress(v BrokerInfoAddress)`

SetAddress sets Address field to given value.


### GetIsSynced

`func (o *InterCliquePeerInfo) GetIsSynced() bool`

GetIsSynced returns the IsSynced field if non-nil, zero value otherwise.

### GetIsSyncedOk

`func (o *InterCliquePeerInfo) GetIsSyncedOk() (*bool, bool)`

GetIsSyncedOk returns a tuple with the IsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynced

`func (o *InterCliquePeerInfo) SetIsSynced(v bool)`

SetIsSynced sets IsSynced field to given value.


### GetClientVersion

`func (o *InterCliquePeerInfo) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *InterCliquePeerInfo) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *InterCliquePeerInfo) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


