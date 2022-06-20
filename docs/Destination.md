# Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AttoAlphAmount** | **string** |  | 
**Tokens** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**LockTime** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewDestination

`func NewDestination(address string, attoAlphAmount string, ) *Destination`

NewDestination instantiates a new Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWithDefaults

`func NewDestinationWithDefaults() *Destination`

NewDestinationWithDefaults instantiates a new Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Destination) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Destination) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Destination) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAttoAlphAmount

`func (o *Destination) GetAttoAlphAmount() string`

GetAttoAlphAmount returns the AttoAlphAmount field if non-nil, zero value otherwise.

### GetAttoAlphAmountOk

`func (o *Destination) GetAttoAlphAmountOk() (*string, bool)`

GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttoAlphAmount

`func (o *Destination) SetAttoAlphAmount(v string)`

SetAttoAlphAmount sets AttoAlphAmount field to given value.


### GetTokens

`func (o *Destination) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Destination) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Destination) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *Destination) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetLockTime

`func (o *Destination) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *Destination) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *Destination) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *Destination) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetMessage

`func (o *Destination) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Destination) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Destination) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Destination) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


