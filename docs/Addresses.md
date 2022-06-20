# Addresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAddress** | **string** |  | 
**Addresses** | [**[]AddressInfo**](AddressInfo.md) |  | 

## Methods

### NewAddresses

`func NewAddresses(activeAddress string, addresses []AddressInfo, ) *Addresses`

NewAddresses instantiates a new Addresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressesWithDefaults

`func NewAddressesWithDefaults() *Addresses`

NewAddressesWithDefaults instantiates a new Addresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAddress

`func (o *Addresses) GetActiveAddress() string`

GetActiveAddress returns the ActiveAddress field if non-nil, zero value otherwise.

### GetActiveAddressOk

`func (o *Addresses) GetActiveAddressOk() (*string, bool)`

GetActiveAddressOk returns a tuple with the ActiveAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAddress

`func (o *Addresses) SetActiveAddress(v string)`

SetActiveAddress sets ActiveAddress field to given value.


### GetAddresses

`func (o *Addresses) GetAddresses() []AddressInfo`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Addresses) GetAddressesOk() (*[]AddressInfo, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Addresses) SetAddresses(v []AddressInfo)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


