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

// checks if the WalletRestore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletRestore{}

// WalletRestore struct for WalletRestore
type WalletRestore struct {
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	WalletName string `json:"walletName"`
	IsMiner *bool `json:"isMiner,omitempty"`
	MnemonicPassphrase *string `json:"mnemonicPassphrase,omitempty"`
}

// NewWalletRestore instantiates a new WalletRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletRestore(password string, mnemonic string, walletName string) *WalletRestore {
	this := WalletRestore{}
	this.Password = password
	this.Mnemonic = mnemonic
	this.WalletName = walletName
	return &this
}

// NewWalletRestoreWithDefaults instantiates a new WalletRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletRestoreWithDefaults() *WalletRestore {
	this := WalletRestore{}
	return &this
}

// GetPassword returns the Password field value
func (o *WalletRestore) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *WalletRestore) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *WalletRestore) SetPassword(v string) {
	o.Password = v
}

// GetMnemonic returns the Mnemonic field value
func (o *WalletRestore) GetMnemonic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnemonic
}

// GetMnemonicOk returns a tuple with the Mnemonic field value
// and a boolean to check if the value has been set.
func (o *WalletRestore) GetMnemonicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnemonic, true
}

// SetMnemonic sets field value
func (o *WalletRestore) SetMnemonic(v string) {
	o.Mnemonic = v
}

// GetWalletName returns the WalletName field value
func (o *WalletRestore) GetWalletName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletName
}

// GetWalletNameOk returns a tuple with the WalletName field value
// and a boolean to check if the value has been set.
func (o *WalletRestore) GetWalletNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalletName, true
}

// SetWalletName sets field value
func (o *WalletRestore) SetWalletName(v string) {
	o.WalletName = v
}

// GetIsMiner returns the IsMiner field value if set, zero value otherwise.
func (o *WalletRestore) GetIsMiner() bool {
	if o == nil || isNil(o.IsMiner) {
		var ret bool
		return ret
	}
	return *o.IsMiner
}

// GetIsMinerOk returns a tuple with the IsMiner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRestore) GetIsMinerOk() (*bool, bool) {
	if o == nil || isNil(o.IsMiner) {
		return nil, false
	}
	return o.IsMiner, true
}

// HasIsMiner returns a boolean if a field has been set.
func (o *WalletRestore) HasIsMiner() bool {
	if o != nil && !isNil(o.IsMiner) {
		return true
	}

	return false
}

// SetIsMiner gets a reference to the given bool and assigns it to the IsMiner field.
func (o *WalletRestore) SetIsMiner(v bool) {
	o.IsMiner = &v
}

// GetMnemonicPassphrase returns the MnemonicPassphrase field value if set, zero value otherwise.
func (o *WalletRestore) GetMnemonicPassphrase() string {
	if o == nil || isNil(o.MnemonicPassphrase) {
		var ret string
		return ret
	}
	return *o.MnemonicPassphrase
}

// GetMnemonicPassphraseOk returns a tuple with the MnemonicPassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletRestore) GetMnemonicPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.MnemonicPassphrase) {
		return nil, false
	}
	return o.MnemonicPassphrase, true
}

// HasMnemonicPassphrase returns a boolean if a field has been set.
func (o *WalletRestore) HasMnemonicPassphrase() bool {
	if o != nil && !isNil(o.MnemonicPassphrase) {
		return true
	}

	return false
}

// SetMnemonicPassphrase gets a reference to the given string and assigns it to the MnemonicPassphrase field.
func (o *WalletRestore) SetMnemonicPassphrase(v string) {
	o.MnemonicPassphrase = &v
}

func (o WalletRestore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletRestore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["mnemonic"] = o.Mnemonic
	toSerialize["walletName"] = o.WalletName
	if !isNil(o.IsMiner) {
		toSerialize["isMiner"] = o.IsMiner
	}
	if !isNil(o.MnemonicPassphrase) {
		toSerialize["mnemonicPassphrase"] = o.MnemonicPassphrase
	}
	return toSerialize, nil
}

type NullableWalletRestore struct {
	value *WalletRestore
	isSet bool
}

func (v NullableWalletRestore) Get() *WalletRestore {
	return v.value
}

func (v *NullableWalletRestore) Set(val *WalletRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletRestore(val *WalletRestore) *NullableWalletRestore {
	return &NullableWalletRestore{value: val, isSet: true}
}

func (v NullableWalletRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


