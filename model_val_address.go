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

// checks if the ValAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValAddress{}

// ValAddress struct for ValAddress
type ValAddress struct {
	Value string `json:"value"`
	Type string `json:"type"`
}

// NewValAddress instantiates a new ValAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValAddress(value string, type_ string) *ValAddress {
	this := ValAddress{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewValAddressWithDefaults instantiates a new ValAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValAddressWithDefaults() *ValAddress {
	this := ValAddress{}
	return &this
}

// GetValue returns the Value field value
func (o *ValAddress) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValAddress) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ValAddress) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *ValAddress) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ValAddress) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ValAddress) SetType(v string) {
	o.Type = v
}

func (o ValAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableValAddress struct {
	value *ValAddress
	isSet bool
}

func (v NullableValAddress) Get() *ValAddress {
	return v.value
}

func (v *NullableValAddress) Set(val *ValAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableValAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableValAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValAddress(val *ValAddress) *NullableValAddress {
	return &NullableValAddress{value: val, isSet: true}
}

func (v NullableValAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


