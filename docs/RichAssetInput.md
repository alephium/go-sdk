# RichAssetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | **int32** |  | 
**Key** | **string** |  | 
**UnlockScript** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Address** | **string** |  | 
**Tokens** | [**[]Token**](Token.md) |  | 
**OutputRefTxId** | **string** |  | 

## Methods

### NewRichAssetInput

`func NewRichAssetInput(hint int32, key string, unlockScript string, attoAlphAmount string, address string, tokens []Token, outputRefTxId string, ) *RichAssetInput`

NewRichAssetInput instantiates a new RichAssetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichAssetInputWithDefaults

`func NewRichAssetInputWithDefaults() *RichAssetInput`

NewRichAssetInputWithDefaults instantiates a new RichAssetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *RichAssetInput) GetHint() int32`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *RichAssetInput) GetHintOk() (*int32, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *RichAssetInput) SetHint(v int32)`

SetHint sets Hint field to given value.


### GetKey

`func (o *RichAssetInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RichAssetInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RichAssetInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetUnlockScript

`func (o *RichAssetInput) GetUnlockScript() string`

GetUnlockScript returns the UnlockScript field if non-nil, zero value otherwise.

### GetUnlockScriptOk

`func (o *RichAssetInput) GetUnlockScriptOk() (*string, bool)`

GetUnlockScriptOk returns a tuple with the UnlockScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockScript

`func (o *RichAssetInput) SetUnlockScript(v string)`

SetUnlockScript sets UnlockScript field to given value.


### GetAttoAlphAmount

`func (o *RichAssetInput) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *RichAssetInput) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *RichAssetInput) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetAddress

`func (o *RichAssetInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RichAssetInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RichAssetInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokens

`func (o *RichAssetInput) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *RichAssetInput) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *RichAssetInput) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.


### GetOutputRefTxId

`func (o *RichAssetInput) GetOutputRefTxId() string`

GetOutputRefTxId returns the OutputRefTxId field if non-nil, zero value otherwise.

### GetOutputRefTxIdOk

`func (o *RichAssetInput) GetOutputRefTxIdOk() (*string, bool)`

GetOutputRefTxIdOk returns a tuple with the OutputRefTxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRefTxId

`func (o *RichAssetInput) SetOutputRefTxId(v string)`

SetOutputRefTxId sets OutputRefTxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


