# ContractOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | **int32** |  | 
**Key** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Address** | **string** |  | 
**Tokens** | [**[]Token**](Token.md) |  | 
**Type** | **string** |  | 

## Methods

### NewContractOutput

`func NewContractOutput(hint int32, key string, attoAlphAmount string, address string, tokens []Token, type_ string, ) *ContractOutput`

NewContractOutput instantiates a new ContractOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractOutputWithDefaults

`func NewContractOutputWithDefaults() *ContractOutput`

NewContractOutputWithDefaults instantiates a new ContractOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *ContractOutput) GetHint() int32`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ContractOutput) GetHintOk() (*int32, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ContractOutput) SetHint(v int32)`

SetHint sets Hint field to given value.


### GetKey

`func (o *ContractOutput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ContractOutput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ContractOutput) SetKey(v string)`

SetKey sets Key field to given value.


### GetAttoAlphAmount

`func (o *ContractOutput) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *ContractOutput) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *ContractOutput) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetAddress

`func (o *ContractOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractOutput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokens

`func (o *ContractOutput) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ContractOutput) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ContractOutput) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.


### GetType

`func (o *ContractOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractOutput) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


