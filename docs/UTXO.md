# UTXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | [**OutputRef**](OutputRef.md) |  | 
**Amount** | **string** |  | 
**Tokens** | [**[]Token**](Token.md) |  | 
**LockTime** | **int64** |  | 
**AdditionalData** | **string** |  | 

## Methods

### NewUTXO

`func NewUTXO(ref OutputRef, amount string, tokens []Token, lockTime int64, additionalData string, ) *UTXO`

NewUTXO instantiates a new UTXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUTXOWithDefaults

`func NewUTXOWithDefaults() *UTXO`

NewUTXOWithDefaults instantiates a new UTXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *UTXO) GetRef() OutputRef`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *UTXO) GetRefOk() (*OutputRef, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *UTXO) SetRef(v OutputRef)`

SetRef sets Ref field to given value.


### GetAmount

`func (o *UTXO) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UTXO) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UTXO) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokens

`func (o *UTXO) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UTXO) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UTXO) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.


### GetLockTime

`func (o *UTXO) GetLockTime() int64`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *UTXO) GetLockTimeOk() (*int64, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *UTXO) SetLockTime(v int64)`

SetLockTime sets LockTime field to given value.


### GetAdditionalData

`func (o *UTXO) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *UTXO) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *UTXO) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


