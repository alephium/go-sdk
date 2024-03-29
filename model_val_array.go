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

// checks if the ValArray type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValArray{}

// ValArray struct for ValArray
type ValArray struct {
	Value []Val `json:"value"`
	Type string `json:"type"`
}

// NewValArray instantiates a new ValArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValArray(value []Val, type_ string) *ValArray {
	this := ValArray{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewValArrayWithDefaults instantiates a new ValArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValArrayWithDefaults() *ValArray {
	this := ValArray{}
	return &this
}

// GetValue returns the Value field value
func (o *ValArray) GetValue() []Val {
	if o == nil {
		var ret []Val
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ValArray) GetValueOk() ([]Val, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *ValArray) SetValue(v []Val) {
	o.Value = v
}

// GetType returns the Type field value
func (o *ValArray) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ValArray) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ValArray) SetType(v string) {
	o.Type = v
}

func (o ValArray) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValArray) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableValArray struct {
	value *ValArray
	isSet bool
}

func (v NullableValArray) Get() *ValArray {
	return v.value
}

func (v *NullableValArray) Set(val *ValArray) {
	v.value = val
	v.isSet = true
}

func (v NullableValArray) IsSet() bool {
	return v.isSet
}

func (v *NullableValArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValArray(val *ValArray) *NullableValArray {
	return &NullableValArray{value: val, isSet: true}
}

func (v NullableValArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


