# WalletUnlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**MnemonicPassphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewWalletUnlock

`func NewWalletUnlock(password string, ) *WalletUnlock`

NewWalletUnlock instantiates a new WalletUnlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletUnlockWithDefaults

`func NewWalletUnlockWithDefaults() *WalletUnlock`

NewWalletUnlockWithDefaults instantiates a new WalletUnlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *WalletUnlock) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WalletUnlock) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WalletUnlock) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetMnemonicPassphrase

`func (o *WalletUnlock) GetMnemonicPassphrase() string`

GetMnemonicPassphrase returns the MnemonicPassphrase field if non-nil, zero value otherwise.

### GetMnemonicPassphraseOk

`func (o *WalletUnlock) GetMnemonicPassphraseOk() (*string, bool)`

GetMnemonicPassphraseOk returns a tuple with the MnemonicPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonicPassphrase

`func (o *WalletUnlock) SetMnemonicPassphrase(v string)`

SetMnemonicPassphrase sets MnemonicPassphrase field to given value.

### HasMnemonicPassphrase

`func (o *WalletUnlock) HasMnemonicPassphrase() bool`

HasMnemonicPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


