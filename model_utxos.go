/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// checks if the UTXOs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UTXOs{}

// UTXOs struct for UTXOs
type UTXOs struct {
	Utxos []UTXO `json:"utxos"`
	Warning *string `json:"warning,omitempty"`
}

// NewUTXOs instantiates a new UTXOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUTXOs(utxos []UTXO) *UTXOs {
	this := UTXOs{}
	this.Utxos = utxos
	return &this
}

// NewUTXOsWithDefaults instantiates a new UTXOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUTXOsWithDefaults() *UTXOs {
	this := UTXOs{}
	return &this
}

// GetUtxos returns the Utxos field value
func (o *UTXOs) GetUtxos() []UTXO {
	if o == nil {
		var ret []UTXO
		return ret
	}

	return o.Utxos
}

// GetUtxosOk returns a tuple with the Utxos field value
// and a boolean to check if the value has been set.
func (o *UTXOs) GetUtxosOk() ([]UTXO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Utxos, true
}

// SetUtxos sets field value
func (o *UTXOs) SetUtxos(v []UTXO) {
	o.Utxos = v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *UTXOs) GetWarning() string {
	if o == nil || isNil(o.Warning) {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UTXOs) GetWarningOk() (*string, bool) {
	if o == nil || isNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *UTXOs) HasWarning() bool {
	if o != nil && !isNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *UTXOs) SetWarning(v string) {
	o.Warning = &v
}

func (o UTXOs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UTXOs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["utxos"] = o.Utxos
	if !isNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableUTXOs struct {
	value *UTXOs
	isSet bool
}

func (v NullableUTXOs) Get() *UTXOs {
	return v.value
}

func (v *NullableUTXOs) Set(val *UTXOs) {
	v.value = val
	v.isSet = true
}

func (v NullableUTXOs) IsSet() bool {
	return v.isSet
}

func (v *NullableUTXOs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUTXOs(val *UTXOs) *NullableUTXOs {
	return &NullableUTXOs{value: val, isSet: true}
}

func (v NullableUTXOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUTXOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


