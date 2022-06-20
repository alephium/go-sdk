/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// Banned struct for Banned
type Banned struct {
	Until int64 `json:"until"`
	Type string `json:"type"`
}

// NewBanned instantiates a new Banned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBanned(until int64, type_ string) *Banned {
	this := Banned{}
	this.Until = until
	this.Type = type_
	return &this
}

// NewBannedWithDefaults instantiates a new Banned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBannedWithDefaults() *Banned {
	this := Banned{}
	return &this
}

// GetUntil returns the Until field value
func (o *Banned) GetUntil() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Until
}

// GetUntilOk returns a tuple with the Until field value
// and a boolean to check if the value has been set.
func (o *Banned) GetUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Until, true
}

// SetUntil sets field value
func (o *Banned) SetUntil(v int64) {
	o.Until = v
}

// GetType returns the Type field value
func (o *Banned) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Banned) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Banned) SetType(v string) {
	o.Type = v
}

func (o Banned) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["until"] = o.Until
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBanned struct {
	value *Banned
	isSet bool
}

func (v NullableBanned) Get() *Banned {
	return v.value
}

func (v *NullableBanned) Set(val *Banned) {
	v.value = val
	v.isSet = true
}

func (v NullableBanned) IsSet() bool {
	return v.isSet
}

func (v *NullableBanned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBanned(val *Banned) *NullableBanned {
	return &NullableBanned{value: val, isSet: true}
}

func (v NullableBanned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBanned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


