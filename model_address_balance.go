/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.5.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// checks if the AddressBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressBalance{}

// AddressBalance struct for AddressBalance
type AddressBalance struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	BalanceHint string `json:"balanceHint"`
	LockedBalance string `json:"lockedBalance"`
	LockedBalanceHint string `json:"lockedBalanceHint"`
	Warning *string `json:"warning,omitempty"`
}

// NewAddressBalance instantiates a new AddressBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalance(address string, balance string, balanceHint string, lockedBalance string, lockedBalanceHint string) *AddressBalance {
	this := AddressBalance{}
	this.Address = address
	this.Balance = balance
	this.BalanceHint = balanceHint
	this.LockedBalance = lockedBalance
	this.LockedBalanceHint = lockedBalanceHint
	return &this
}

// NewAddressBalanceWithDefaults instantiates a new AddressBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceWithDefaults() *AddressBalance {
	this := AddressBalance{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddressBalance) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressBalance) SetAddress(v string) {
	o.Address = v
}

// GetBalance returns the Balance field value
func (o *AddressBalance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *AddressBalance) SetBalance(v string) {
	o.Balance = v
}

// GetBalanceHint returns the BalanceHint field value
func (o *AddressBalance) GetBalanceHint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BalanceHint
}

// GetBalanceHintOk returns a tuple with the BalanceHint field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetBalanceHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceHint, true
}

// SetBalanceHint sets field value
func (o *AddressBalance) SetBalanceHint(v string) {
	o.BalanceHint = v
}

// GetLockedBalance returns the LockedBalance field value
func (o *AddressBalance) GetLockedBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedBalance
}

// GetLockedBalanceOk returns a tuple with the LockedBalance field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetLockedBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedBalance, true
}

// SetLockedBalance sets field value
func (o *AddressBalance) SetLockedBalance(v string) {
	o.LockedBalance = v
}

// GetLockedBalanceHint returns the LockedBalanceHint field value
func (o *AddressBalance) GetLockedBalanceHint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockedBalanceHint
}

// GetLockedBalanceHintOk returns a tuple with the LockedBalanceHint field value
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetLockedBalanceHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedBalanceHint, true
}

// SetLockedBalanceHint sets field value
func (o *AddressBalance) SetLockedBalanceHint(v string) {
	o.LockedBalanceHint = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *AddressBalance) GetWarning() string {
	if o == nil || isNil(o.Warning) {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetWarningOk() (*string, bool) {
	if o == nil || isNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *AddressBalance) HasWarning() bool {
	if o != nil && !isNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *AddressBalance) SetWarning(v string) {
	o.Warning = &v
}

func (o AddressBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["balance"] = o.Balance
	toSerialize["balanceHint"] = o.BalanceHint
	toSerialize["lockedBalance"] = o.LockedBalance
	toSerialize["lockedBalanceHint"] = o.LockedBalanceHint
	if !isNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableAddressBalance struct {
	value *AddressBalance
	isSet bool
}

func (v NullableAddressBalance) Get() *AddressBalance {
	return v.value
}

func (v *NullableAddressBalance) Set(val *AddressBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalance(val *AddressBalance) *NullableAddressBalance {
	return &NullableAddressBalance{value: val, isSet: true}
}

func (v NullableAddressBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


