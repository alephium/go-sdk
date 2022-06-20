# AssetOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | **int32** |  | 
**Key** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Address** | **string** |  | 
**Tokens** | [**[]Token**](Token.md) |  | 
**LockTime** | **int64** |  | 
**Message** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewAssetOutput

`func NewAssetOutput(hint int32, key string, attoAlphAmount string, address string, tokens []Token, lockTime int64, message string, type_ string, ) *AssetOutput`

NewAssetOutput instantiates a new AssetOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOutputWithDefaults

`func NewAssetOutputWithDefaults() *AssetOutput`

NewAssetOutputWithDefaults instantiates a new AssetOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *AssetOutput) GetHint() int32`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *AssetOutput) GetHintOk() (*int32, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *AssetOutput) SetHint(v int32)`

SetHint sets Hint field to given value.


### GetKey

`func (o *AssetOutput) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AssetOutput) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AssetOutput) SetKey(v string)`

SetKey sets Key field to given value.


### GetAttoAlphAmount

`func (o *AssetOutput) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *AssetOutput) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *AssetOutput) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetAddress

`func (o *AssetOutput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssetOutput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssetOutput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetTokens

`func (o *AssetOutput) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AssetOutput) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AssetOutput) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.


### GetLockTime

`func (o *AssetOutput) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *AssetOutput) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *AssetOutput) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.


### GetMessage

`func (o *AssetOutput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AssetOutput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AssetOutput) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *AssetOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetOutput) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


