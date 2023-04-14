/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// checks if the Ban type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ban{}

// Ban struct for Ban
type Ban struct {
	Peers []string `json:"peers"`
	Type string `json:"type"`
}

// NewBan instantiates a new Ban object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBan(peers []string, type_ string) *Ban {
	this := Ban{}
	this.Peers = peers
	this.Type = type_
	return &this
}

// NewBanWithDefaults instantiates a new Ban object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBanWithDefaults() *Ban {
	this := Ban{}
	return &this
}

// GetPeers returns the Peers field value
func (o *Ban) GetPeers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *Ban) GetPeersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *Ban) SetPeers(v []string) {
	o.Peers = v
}

// GetType returns the Type field value
func (o *Ban) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Ban) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Ban) SetType(v string) {
	o.Type = v
}

func (o Ban) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ban) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peers"] = o.Peers
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBan struct {
	value *Ban
	isSet bool
}

func (v NullableBan) Get() *Ban {
	return v.value
}

func (v *NullableBan) Set(val *Ban) {
	v.value = val
	v.isSet = true
}

func (v NullableBan) IsSet() bool {
	return v.isSet
}

func (v *NullableBan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBan(val *Ban) *NullableBan {
	return &NullableBan{value: val, isSet: true}
}

func (v NullableBan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


