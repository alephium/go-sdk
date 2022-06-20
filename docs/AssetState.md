# AssetState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttoAlphAmount** | **string** |  | 
**Tokens** | Pointer to [**[]Token**](Token.md) |  | [optional] 

## Methods

### NewAssetState

`func NewAssetState(attoAlphAmount string, ) *AssetState`

NewAssetState instantiates a new AssetState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetStateWithDefaults

`func NewAssetStateWithDefaults() *AssetState`

NewAssetStateWithDefaults instantiates a new AssetState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttoAlphAmount

`func (o *AssetState) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *AssetState) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *AssetState) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetTokens

`func (o *AssetState) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AssetState) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AssetState) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AssetState) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


