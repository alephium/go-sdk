# WalletCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletName** | **string** |  | 
**Mnemonic** | **string** |  | 

## Methods

### NewWalletCreationResult

`func NewWalletCreationResult(walletName string, mnemonic string, ) *WalletCreationResult`

NewWalletCreationResult instantiates a new WalletCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCreationResultWithDefaults

`func NewWalletCreationResultWithDefaults() *WalletCreationResult`

NewWalletCreationResultWithDefaults instantiates a new WalletCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletName

`func (o *WalletCreationResult) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *WalletCreationResult) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *WalletCreationResult) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.


### GetMnemonic

`func (o *WalletCreationResult) GetMnemonic() string`

GetMnemonic returns the Mnemonic field if non-nil, zero value otherwise.

### GetMnemonicOk

`func (o *WalletCreationResult) GetMnemonicOk() (*string, bool)`

GetMnemonicOk returns a tuple with the Mnemonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnemonic

`func (o *WalletCreationResult) SetMnemonic(v string)`

SetMnemonic sets Mnemonic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


