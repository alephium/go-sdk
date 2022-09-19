/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// WalletCreation struct for WalletCreation
type WalletCreation struct {
	Password string `json:"password"`
	WalletName string `json:"walletName"`
	IsMiner *bool `json:"isMiner,omitempty"`
	MnemonicPassphrase *string `json:"mnemonicPassphrase,omitempty"`
	MnemonicSize *int32 `json:"mnemonicSize,omitempty"`
}

// NewWalletCreation instantiates a new WalletCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCreation(password string, walletName string) *WalletCreation {
	this := WalletCreation{}
	this.Password = password
	this.WalletName = walletName
	return &this
}

// NewWalletCreationWithDefaults instantiates a new WalletCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCreationWithDefaults() *WalletCreation {
	this := WalletCreation{}
	return &this
}

// GetPassword returns the Password field value
func (o *WalletCreation) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *WalletCreation) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *WalletCreation) SetPassword(v string) {
	o.Password = v
}

// GetWalletName returns the WalletName field value
func (o *WalletCreation) GetWalletName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value
// and a boolean to check if the value has been set.
func (o *WalletCreation) GetWalletNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletName, true
}

// SetWalletName sets field value
func (o *WalletCreation) SetWalletName(v string) {
	o.WalletName = v
}

// GetIsMiner returns the IsMiner field value if set, zero value otherwise.
func (o *WalletCreation) GetIsMiner() bool {
	if o == nil || o.IsMiner == nil {
		var ret bool
		return ret
	}
	return *o.IsMiner
}

// GetIsMinerOk returns a tuple with the IsMiner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreation) GetIsMinerOk() (*bool, bool) {
	if o == nil || o.IsMiner == nil {
		return nil, false
	}
	return o.IsMiner, true
}

// HasIsMiner returns a boolean if a field has been set.
func (o *WalletCreation) HasIsMiner() bool {
	if o != nil && o.IsMiner != nil {
		return true
	}

	return false
}

// SetIsMiner gets a reference to the given bool and assigns it to the IsMiner field.
func (o *WalletCreation) SetIsMiner(v bool) {
	o.IsMiner = &v
}

// GetMnemonicPassphrase returns the MnemonicPassphrase field value if set, zero value otherwise.
func (o *WalletCreation) GetMnemonicPassphrase() string {
	if o == nil || o.MnemonicPassphrase == nil {
		var ret string
		return ret
	}
	return *o.MnemonicPassphrase
}

// GetMnemonicPassphraseOk returns a tuple with the MnemonicPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreation) GetMnemonicPassphraseOk() (*string, bool) {
	if o == nil || o.MnemonicPassphrase == nil {
		return nil, false
	}
	return o.MnemonicPassphrase, true
}

// HasMnemonicPassphrase returns a boolean if a field has been set.
func (o *WalletCreation) HasMnemonicPassphrase() bool {
	if o != nil && o.MnemonicPassphrase != nil {
		return true
	}

	return false
}

// SetMnemonicPassphrase gets a reference to the given string and assigns it to the MnemonicPassphrase field.
func (o *WalletCreation) SetMnemonicPassphrase(v string) {
	o.MnemonicPassphrase = &v
}

// GetMnemonicSize returns the MnemonicSize field value if set, zero value otherwise.
func (o *WalletCreation) GetMnemonicSize() int32 {
	if o == nil || o.MnemonicSize == nil {
		var ret int32
		return ret
	}
	return *o.MnemonicSize
}

// GetMnemonicSizeOk returns a tuple with the MnemonicSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCreation) GetMnemonicSizeOk() (*int32, bool) {
	if o == nil || o.MnemonicSize == nil {
		return nil, false
	}
	return o.MnemonicSize, true
}

// HasMnemonicSize returns a boolean if a field has been set.
func (o *WalletCreation) HasMnemonicSize() bool {
	if o != nil && o.MnemonicSize != nil {
		return true
	}

	return false
}

// SetMnemonicSize gets a reference to the given int32 and assigns it to the MnemonicSize field.
func (o *WalletCreation) SetMnemonicSize(v int32) {
	o.MnemonicSize = &v
}

func (o WalletCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["walletName"] = o.WalletName
	}
	if o.IsMiner != nil {
		toSerialize["isMiner"] = o.IsMiner
	}
	if o.MnemonicPassphrase != nil {
		toSerialize["mnemonicPassphrase"] = o.MnemonicPassphrase
	}
	if o.MnemonicSize != nil {
		toSerialize["mnemonicSize"] = o.MnemonicSize
	}
	return json.Marshal(toSerialize)
}

type NullableWalletCreation struct {
	value *WalletCreation
	isSet bool
}

func (v NullableWalletCreation) Get() *WalletCreation {
	return v.value
}

func (v *NullableWalletCreation) Set(val *WalletCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCreation(val *WalletCreation) *NullableWalletCreation {
	return &NullableWalletCreation{value: val, isSet: true}
}

func (v NullableWalletCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


