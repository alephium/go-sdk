# WalletStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletName** | **string** |  | 
**Locked** | **bool** |  | 

## Methods

### NewWalletStatus

`func NewWalletStatus(walletName string, locked bool, ) *WalletStatus`

NewWalletStatus instantiates a new WalletStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletStatusWithDefaults

`func NewWalletStatusWithDefaults() *WalletStatus`

NewWalletStatusWithDefaults instantiates a new WalletStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletName

`func (o *WalletStatus) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *WalletStatus) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *WalletStatus) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.


### GetLocked

`func (o *WalletStatus) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *WalletStatus) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *WalletStatus) SetLocked(v bool)`

SetLocked sets Locked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


