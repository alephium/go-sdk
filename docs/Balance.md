# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** |  | 
**BalanceHint** | **string** |  | 
**LockedBalance** | **string** |  | 
**LockedBalanceHint** | **string** |  | 
**TokenBalances** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**LockedTokenBalances** | Pointer to [**[]Token**](Token.md) |  | [optional] 
**UtxoNum** | **int32** |  | 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewBalance

`func NewBalance(balance string, balanceHint string, lockedBalance string, lockedBalanceHint string, utxoNum int32, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Balance) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Balance) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Balance) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetBalanceHint

`func (o *Balance) GetBalanceHint() string`

GetBalanceHint returns the BalanceHint field if non-nil, zero value otherwise.

### GetBalanceHintOk

`func (o *Balance) GetBalanceHintOk() (*string, bool)`

GetBalanceHintOk returns a tuple with the BalanceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceHint

`func (o *Balance) SetBalanceHint(v string)`

SetBalanceHint sets BalanceHint field to given value.


### GetLockedBalance

`func (o *Balance) GetLockedBalance() string`

GetLockedBalance returns the LockedBalance field if non-nil, zero value otherwise.

### GetLockedBalanceOk

`func (o *Balance) GetLockedBalanceOk() (*string, bool)`

GetLockedBalanceOk returns a tuple with the LockedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBalance

`func (o *Balance) SetLockedBalance(v string)`

SetLockedBalance sets LockedBalance field to given value.


### GetLockedBalanceHint

`func (o *Balance) GetLockedBalanceHint() string`

GetLockedBalanceHint returns the LockedBalanceHint field if non-nil, zero value otherwise.

### GetLockedBalanceHintOk

`func (o *Balance) GetLockedBalanceHintOk() (*string, bool)`

GetLockedBalanceHintOk returns a tuple with the LockedBalanceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBalanceHint

`func (o *Balance) SetLockedBalanceHint(v string)`

SetLockedBalanceHint sets LockedBalanceHint field to given value.


### GetTokenBalances

`func (o *Balance) GetTokenBalances() []Token`

GetTokenBalances returns the TokenBalances field if non-nil, zero value otherwise.

### GetTokenBalancesOk

`func (o *Balance) GetTokenBalancesOk() (*[]Token, bool)`

GetTokenBalancesOk returns a tuple with the TokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBalances

`func (o *Balance) SetTokenBalances(v []Token)`

SetTokenBalances sets TokenBalances field to given value.

### HasTokenBalances

`func (o *Balance) HasTokenBalances() bool`

HasTokenBalances returns a boolean if a field has been set.

### GetLockedTokenBalances

`func (o *Balance) GetLockedTokenBalances() []Token`

GetLockedTokenBalances returns the LockedTokenBalances field if non-nil, zero value otherwise.

### GetLockedTokenBalancesOk

`func (o *Balance) GetLockedTokenBalancesOk() (*[]Token, bool)`

GetLockedTokenBalancesOk returns a tuple with the LockedTokenBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedTokenBalances

`func (o *Balance) SetLockedTokenBalances(v []Token)`

SetLockedTokenBalances sets LockedTokenBalances field to given value.

### HasLockedTokenBalances

`func (o *Balance) HasLockedTokenBalances() bool`

HasLockedTokenBalances returns a boolean if a field has been set.

### GetUtxoNum

`func (o *Balance) GetUtxoNum() int32`

GetUtxoNum returns the UtxoNum field if non-nil, zero value otherwise.

### GetUtxoNumOk

`func (o *Balance) GetUtxoNumOk() (*int32, bool)`

GetUtxoNumOk returns a tuple with the UtxoNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtxoNum

`func (o *Balance) SetUtxoNum(v int32)`

SetUtxoNum sets UtxoNum field to given value.


### GetWarning

`func (o *Balance) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *Balance) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *Balance) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *Balance) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


