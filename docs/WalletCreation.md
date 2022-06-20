# WalletCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**WalletName** | **string** |  | 
**IsMiner** | Pointer to **bool** |  | [optional] 
**MnemonicPassphrase** | Pointer to **string** |  | [optional] 
**MnemonicSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewWalletCreation

`func NewWalletCreation(password string, walletName string, ) *WalletCreation`

NewWalletCreation instantiates a new WalletCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCreationWithDefaults

`func NewWalletCreationWithDefaults() *WalletCreation`

NewWalletCreationWithDefaults instantiates a new WalletCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *WalletCreation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WalletCreation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WalletCreation) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetWalletName

`func (o *WalletCreation) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *WalletCreation) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *WalletCreation) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.


### GetIsMiner

`func (o *WalletCreation) GetIsMiner() bool`

GetIsMiner returns the IsMiner field if non-nil, zero value otherwise.

### GetIsMinerOk

`func (o *WalletCreation) GetIsMinerOk() (*bool, bool)`

GetIsMinerOk returns a tuple with the IsMiner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMiner

`func (o *WalletCreation) SetIsMiner(v bool)`

SetIsMiner sets IsMiner field to given value.

### HasIsMiner

`func (o *WalletCreation) HasIsMiner() bool`

HasIsMiner returns a boolean if a field has been set.

### GetMnemonicPassphrase

`func (o *WalletCreation) GetMnemonicPassphrase() string`

GetMnemonicPassphrase returns the MnemonicPassphrase field if non-nil, zero value otherwise.

### GetMnemonicPassphraseOk

`func (o *WalletCreation) GetMnemonicPassphraseOk() (*string, bool)`

GetMnemonicPassphraseOk returns a tuple with the MnemonicPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicPassphrase

`func (o *WalletCreation) SetMnemonicPassphrase(v string)`

SetMnemonicPassphrase sets MnemonicPassphrase field to given value.

### HasMnemonicPassphrase

`func (o *WalletCreation) HasMnemonicPassphrase() bool`

HasMnemonicPassphrase returns a boolean if a field has been set.

### GetMnemonicSize

`func (o *WalletCreation) GetMnemonicSize() int32`

GetMnemonicSize returns the MnemonicSize field if non-nil, zero value otherwise.

### GetMnemonicSizeOk

`func (o *WalletCreation) GetMnemonicSizeOk() (*int32, bool)`

GetMnemonicSizeOk returns a tuple with the MnemonicSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicSize

`func (o *WalletCreation) SetMnemonicSize(v int32)`

SetMnemonicSize sets MnemonicSize field to given value.

### HasMnemonicSize

`func (o *WalletCreation) HasMnemonicSize() bool`

HasMnemonicSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


