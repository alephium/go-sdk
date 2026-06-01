# RichContractInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | **int32** |  | 
**Key** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Address** | **string** |  | 
**Tokens** | [**[]Token**](Token.md) |  | 
**OutputRefTxId** | **string** |  | 

## Methods

### NewRichContractInput

`func NewRichContractInput(hint int32, key string, attoAlphAmount string, address string, tokens []Token, outputRefTxId string, ) *RichContractInput`

NewRichContractInput instantiates a new RichContractInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichContractInputWithDefaults

`func NewRichContractInputWithDefaults() *RichContractInput`

NewRichContractInputWithDefaults instantiates a new RichContractInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *RichContractInput) GetHint() int32`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *RichContractInput) GetHintOk() (*int32, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *RichContractInput) SetHint(v int32)`

SetHint sets Hint field to given value.


### GetKey

`func (o *RichContractInput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RichContractInput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RichContractInput) SetKey(v string)`

SetKey sets Key field to given value.


### GetAttoAlphAmount

`func (o *RichContractInput) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *RichContractInput) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *RichContractInput) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetAddress

`func (o *RichContractInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RichContractInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RichContractInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokens

`func (o *RichContractInput) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *RichContractInput) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *RichContractInput) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.


### GetOutputRefTxId

`func (o *RichContractInput) GetOutputRefTxId() string`

GetOutputRefTxId returns the OutputRefTxId field if non-nil, zero value otherwise.

### GetOutputRefTxIdOk

`func (o *RichContractInput) GetOutputRefTxIdOk() (*string, bool)`

GetOutputRefTxIdOk returns a tuple with the OutputRefTxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRefTxId

`func (o *RichContractInput) SetOutputRefTxId(v string)`

SetOutputRefTxId sets OutputRefTxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


