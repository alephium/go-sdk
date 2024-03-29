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

// checks if the Penalty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Penalty{}

// Penalty struct for Penalty
type Penalty struct {
	Value int32 `json:"value"`
	Type string `json:"type"`
}

// NewPenalty instantiates a new Penalty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPenalty(value int32, type_ string) *Penalty {
	this := Penalty{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewPenaltyWithDefaults instantiates a new Penalty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPenaltyWithDefaults() *Penalty {
	this := Penalty{}
	return &this
}

// GetValue returns the Value field value
func (o *Penalty) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Penalty) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Penalty) SetValue(v int32) {
	o.Value = v
}

// GetType returns the Type field value
func (o *Penalty) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Penalty) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Penalty) SetType(v string) {
	o.Type = v
}

func (o Penalty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Penalty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePenalty struct {
	value *Penalty
	isSet bool
}

func (v NullablePenalty) Get() *Penalty {
	return v.value
}

func (v *NullablePenalty) Set(val *Penalty) {
	v.value = val
	v.isSet = true
}

func (v NullablePenalty) IsSet() bool {
	return v.isSet
}

func (v *NullablePenalty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePenalty(val *Penalty) *NullablePenalty {
	return &NullablePenalty{value: val, isSet: true}
}

func (v NullablePenalty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePenalty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


