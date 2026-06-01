# AddressAssetState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Tokens** | Pointer to [**[]Token**](Token.md) |  | [optional] 

## Methods

### NewAddressAssetState

`func NewAddressAssetState(address string, attoAlphAmount string, ) *AddressAssetState`

NewAddressAssetState instantiates a new AddressAssetState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressAssetStateWithDefaults

`func NewAddressAssetStateWithDefaults() *AddressAssetState`

NewAddressAssetStateWithDefaults instantiates a new AddressAssetState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressAssetState) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressAssetState) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressAssetState) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAttoAlphAmount

`func (o *AddressAssetState) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *AddressAssetState) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *AddressAssetState) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetTokens

`func (o *AddressAssetState) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *AddressAssetState) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *AddressAssetState) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *AddressAssetState) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


