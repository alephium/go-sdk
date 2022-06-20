# AddressBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Balance** | **string** |  | 
**BalanceHint** | **string** |  | 
**LockedBalance** | **string** |  | 
**LockedBalanceHint** | **string** |  | 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressBalance

`func NewAddressBalance(address string, balance string, balanceHint string, lockedBalance string, lockedBalanceHint string, ) *AddressBalance`

NewAddressBalance instantiates a new AddressBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBalanceWithDefaults

`func NewAddressBalanceWithDefaults() *AddressBalance`

NewAddressBalanceWithDefaults instantiates a new AddressBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressBalance) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressBalance) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressBalance) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *AddressBalance) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *AddressBalance) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *AddressBalance) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetBalanceHint

`func (o *AddressBalance) GetBalanceHint() string`

GetBalanceHint returns the BalanceHint field if non-nil, zero value otherwise.

### GetBalanceHintOk

`func (o *AddressBalance) GetBalanceHintOk() (*string, bool)`

GetBalanceHintOk returns a tuple with the BalanceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceHint

`func (o *AddressBalance) SetBalanceHint(v string)`

SetBalanceHint sets BalanceHint field to given value.


### GetLockedBalance

`func (o *AddressBalance) GetLockedBalance() string`

GetLockedBalance returns the LockedBalance field if non-nil, zero value otherwise.

### GetLockedBalanceOk

`func (o *AddressBalance) GetLockedBalanceOk() (*string, bool)`

GetLockedBalanceOk returns a tuple with the LockedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBalance

`func (o *AddressBalance) SetLockedBalance(v string)`

SetLockedBalance sets LockedBalance field to given value.


### GetLockedBalanceHint

`func (o *AddressBalance) GetLockedBalanceHint() string`

GetLockedBalanceHint returns the LockedBalanceHint field if non-nil, zero value otherwise.

### GetLockedBalanceHintOk

`func (o *AddressBalance) GetLockedBalanceHintOk() (*string, bool)`

GetLockedBalanceHintOk returns a tuple with the LockedBalanceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBalanceHint

`func (o *AddressBalance) SetLockedBalanceHint(v string)`

SetLockedBalanceHint sets LockedBalanceHint field to given value.


### GetWarning

`func (o *AddressBalance) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *AddressBalance) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *AddressBalance) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *AddressBalance) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


