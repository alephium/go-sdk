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

// Reachable struct for Reachable
type Reachable struct {
	Peers []string `json:"peers"`
	Type string `json:"type"`
}

// NewReachable instantiates a new Reachable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReachable(peers []string, type_ string) *Reachable {
	this := Reachable{}
	this.Peers = peers
	this.Type = type_
	return &this
}

// NewReachableWithDefaults instantiates a new Reachable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReachableWithDefaults() *Reachable {
	this := Reachable{}
	return &this
}

// GetPeers returns the Peers field value
func (o *Reachable) GetPeers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *Reachable) GetPeersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *Reachable) SetPeers(v []string) {
	o.Peers = v
}

// GetType returns the Type field value
func (o *Reachable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Reachable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Reachable) SetType(v string) {
	o.Type = v
}

func (o Reachable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableReachable struct {
	value *Reachable
	isSet bool
}

func (v NullableReachable) Get() *Reachable {
	return v.value
}

func (v *NullableReachable) Set(val *Reachable) {
	v.value = val
	v.isSet = true
}

func (v NullableReachable) IsSet() bool {
	return v.isSet
}

func (v *NullableReachable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachable(val *Reachable) *NullableReachable {
	return &NullableReachable{value: val, isSet: true}
}

func (v NullableReachable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


