# Balances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBalance** | **string** |  | 
**TotalBalanceHint** | **string** |  | 
**Balances** | [**[]AddressBalance**](AddressBalance.md) |  | 

## Methods

### NewBalances

`func NewBalances(totalBalance string, totalBalanceHint string, balances []AddressBalance, ) *Balances`

NewBalances instantiates a new Balances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalancesWithDefaults

`func NewBalancesWithDefaults() *Balances`

NewBalancesWithDefaults instantiates a new Balances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBalance

`func (o *Balances) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *Balances) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *Balances) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.


### GetTotalBalanceHint

`func (o *Balances) GetTotalBalanceHint() string`

GetTotalBalanceHint returns the TotalBalanceHint field if non-nil, zero value otherwise.

### GetTotalBalanceHintOk

`func (o *Balances) GetTotalBalanceHintOk() (*string, bool)`

GetTotalBalanceHintOk returns a tuple with the TotalBalanceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalanceHint

`func (o *Balances) SetTotalBalanceHint(v string)`

SetTotalBalanceHint sets TotalBalanceHint field to given value.


### GetBalances

`func (o *Balances) GetBalances() []AddressBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *Balances) GetBalancesOk() (*[]AddressBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *Balances) SetBalances(v []AddressBalance)`

SetBalances sets Balances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


