# WalletRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**Mnemonic** | **string** |  | 
**WalletName** | **string** |  | 
**IsMiner** | Pointer to **bool** |  | [optional] 
**MnemonicPassphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewWalletRestore

`func NewWalletRestore(password string, mnemonic string, walletName string, ) *WalletRestore`

NewWalletRestore instantiates a new WalletRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletRestoreWithDefaults

`func NewWalletRestoreWithDefaults() *WalletRestore`

NewWalletRestoreWithDefaults instantiates a new WalletRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *WalletRestore) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WalletRestore) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WalletRestore) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetMnemonic

`func (o *WalletRestore) GetMnemonic() string`

GetMnemonic returns the Mnemonic field if non-nil, zero value otherwise.

### GetMnemonicOk

`func (o *WalletRestore) GetMnemonicOk() (*string, bool)`

GetMnemonicOk returns a tuple with the Mnemonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonic

`func (o *WalletRestore) SetMnemonic(v string)`

SetMnemonic sets Mnemonic field to given value.


### GetWalletName

`func (o *WalletRestore) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *WalletRestore) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *WalletRestore) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.


### GetIsMiner

`func (o *WalletRestore) GetIsMiner() bool`

GetIsMiner returns the IsMiner field if non-nil, zero value otherwise.

### GetIsMinerOk

`func (o *WalletRestore) GetIsMinerOk() (*bool, bool)`

GetIsMinerOk returns a tuple with the IsMiner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMiner

`func (o *WalletRestore) SetIsMiner(v bool)`

SetIsMiner sets IsMiner field to given value.

### HasIsMiner

`func (o *WalletRestore) HasIsMiner() bool`

HasIsMiner returns a boolean if a field has been set.

### GetMnemonicPassphrase

`func (o *WalletRestore) GetMnemonicPassphrase() string`

GetMnemonicPassphrase returns the MnemonicPassphrase field if non-nil, zero value otherwise.

### GetMnemonicPassphraseOk

`func (o *WalletRestore) GetMnemonicPassphraseOk() (*string, bool)`

GetMnemonicPassphraseOk returns a tuple with the MnemonicPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicPassphrase

`func (o *WalletRestore) SetMnemonicPassphrase(v string)`

SetMnemonicPassphrase sets MnemonicPassphrase field to given value.

### HasMnemonicPassphrase

`func (o *WalletRestore) HasMnemonicPassphrase() bool`

HasMnemonicPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


