# BrokerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliqueId** | **string** |  | 
**BrokerId** | **int32** |  | 
**BrokerNum** | **int32** |  | 
**Address** | [**BrokerInfoAddress**](BrokerInfoAddress.md) |  | 

## Methods

### NewBrokerInfo

`func NewBrokerInfo(cliqueId string, brokerId int32, brokerNum int32, address BrokerInfoAddress, ) *BrokerInfo`

NewBrokerInfo instantiates a new BrokerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerInfoWithDefaults

`func NewBrokerInfoWithDefaults() *BrokerInfo`

NewBrokerInfoWithDefaults instantiates a new BrokerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliqueId

`func (o *BrokerInfo) GetCliqueId() string`

GetCliqueId returns the CliqueId field if non-nil, zero value otherwise.

### GetCliqueIdOk

`func (o *BrokerInfo) GetCliqueIdOk() (*string, bool)`

GetCliqueIdOk returns a tuple with the CliqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliqueId

`func (o *BrokerInfo) SetCliqueId(v string)`

SetCliqueId sets CliqueId field to given value.


### GetBrokerId

`func (o *BrokerInfo) GetBrokerId() int32`

GetBrokerId returns the BrokerId field if non-nil, zero value otherwise.

### GetBrokerIdOk

`func (o *BrokerInfo) GetBrokerIdOk() (*int32, bool)`

GetBrokerIdOk returns a tuple with the BrokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerId

`func (o *BrokerInfo) SetBrokerId(v int32)`

SetBrokerId sets BrokerId field to given value.


### GetBrokerNum

`func (o *BrokerInfo) GetBrokerNum() int32`

GetBrokerNum returns the BrokerNum field if non-nil, zero value otherwise.

### GetBrokerNumOk

`func (o *BrokerInfo) GetBrokerNumOk() (*int32, bool)`

GetBrokerNumOk returns a tuple with the BrokerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerNum

`func (o *BrokerInfo) SetBrokerNum(v int32)`

SetBrokerNum sets BrokerNum field to given value.


### GetAddress

`func (o *BrokerInfo) GetAddress() BrokerInfoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BrokerInfo) GetAddressOk() (*BrokerInfoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BrokerInfo) SetAddress(v BrokerInfoAddress)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


