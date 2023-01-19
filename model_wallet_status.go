/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// WalletStatus struct for WalletStatus
type WalletStatus struct {
	WalletName string `json:"walletName"`
	Locked bool `json:"locked"`
}

// NewWalletStatus instantiates a new WalletStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletStatus(walletName string, locked bool) *WalletStatus {
	this := WalletStatus{}
	this.WalletName = walletName
	this.Locked = locked
	return &this
}

// NewWalletStatusWithDefaults instantiates a new WalletStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletStatusWithDefaults() *WalletStatus {
	this := WalletStatus{}
	return &this
}

// GetWalletName returns the WalletName field value
func (o *WalletStatus) GetWalletName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value
// and a boolean to check if the value has been set.
func (o *WalletStatus) GetWalletNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.WalletName, true
}

// SetWalletName sets field value
func (o *WalletStatus) SetWalletName(v string) {
	o.WalletName = v
}

// GetLocked returns the Locked field value
func (o *WalletStatus) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *WalletStatus) GetLockedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *WalletStatus) SetLocked(v bool) {
	o.Locked = v
}

func (o WalletStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["walletName"] = o.WalletName
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	return json.Marshal(toSerialize)
}

type NullableWalletStatus struct {
	value *WalletStatus
	isSet bool
}

func (v NullableWalletStatus) Get() *WalletStatus {
	return v.value
}

func (v *NullableWalletStatus) Set(val *WalletStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletStatus(val *WalletStatus) *NullableWalletStatus {
	return &NullableWalletStatus{value: val, isSet: true}
}

func (v NullableWalletStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


