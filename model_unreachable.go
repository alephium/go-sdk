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

// Unreachable struct for Unreachable
type Unreachable struct {
	Peers []string `json:"peers"`
	Type string `json:"type"`
}

// NewUnreachable instantiates a new Unreachable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnreachable(peers []string, type_ string) *Unreachable {
	this := Unreachable{}
	this.Peers = peers
	this.Type = type_
	return &this
}

// NewUnreachableWithDefaults instantiates a new Unreachable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnreachableWithDefaults() *Unreachable {
	this := Unreachable{}
	return &this
}

// GetPeers returns the Peers field value
func (o *Unreachable) GetPeers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *Unreachable) GetPeersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *Unreachable) SetPeers(v []string) {
	o.Peers = v
}

// GetType returns the Type field value
func (o *Unreachable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Unreachable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Unreachable) SetType(v string) {
	o.Type = v
}

func (o Unreachable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUnreachable struct {
	value *Unreachable
	isSet bool
}

func (v NullableUnreachable) Get() *Unreachable {
	return v.value
}

func (v *NullableUnreachable) Set(val *Unreachable) {
	v.value = val
	v.isSet = true
}

func (v NullableUnreachable) IsSet() bool {
	return v.isSet
}

func (v *NullableUnreachable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnreachable(val *Unreachable) *NullableUnreachable {
	return &NullableUnreachable{value: val, isSet: true}
}

func (v NullableUnreachable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnreachable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


