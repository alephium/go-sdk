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

// checks if the ChangeActiveAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeActiveAddress{}

// ChangeActiveAddress struct for ChangeActiveAddress
type ChangeActiveAddress struct {
	Address string `json:"address"`
}

// NewChangeActiveAddress instantiates a new ChangeActiveAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeActiveAddress(address string) *ChangeActiveAddress {
	this := ChangeActiveAddress{}
	this.Address = address
	return &this
}

// NewChangeActiveAddressWithDefaults instantiates a new ChangeActiveAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeActiveAddressWithDefaults() *ChangeActiveAddress {
	this := ChangeActiveAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *ChangeActiveAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ChangeActiveAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ChangeActiveAddress) SetAddress(v string) {
	o.Address = v
}

func (o ChangeActiveAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeActiveAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableChangeActiveAddress struct {
	value *ChangeActiveAddress
	isSet bool
}

func (v NullableChangeActiveAddress) Get() *ChangeActiveAddress {
	return v.value
}

func (v *NullableChangeActiveAddress) Set(val *ChangeActiveAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeActiveAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeActiveAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeActiveAddress(val *ChangeActiveAddress) *NullableChangeActiveAddress {
	return &NullableChangeActiveAddress{value: val, isSet: true}
}

func (v NullableChangeActiveAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeActiveAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


