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

// checks if the Addresses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Addresses{}

// Addresses struct for Addresses
type Addresses struct {
	ActiveAddress string `json:"activeAddress"`
	Addresses []AddressInfo `json:"addresses"`
}

// NewAddresses instantiates a new Addresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddresses(activeAddress string, addresses []AddressInfo) *Addresses {
	this := Addresses{}
	this.ActiveAddress = activeAddress
	this.Addresses = addresses
	return &this
}

// NewAddressesWithDefaults instantiates a new Addresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressesWithDefaults() *Addresses {
	this := Addresses{}
	return &this
}

// GetActiveAddress returns the ActiveAddress field value
func (o *Addresses) GetActiveAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveAddress
}

// GetActiveAddressOk returns a tuple with the ActiveAddress field value
// and a boolean to check if the value has been set.
func (o *Addresses) GetActiveAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveAddress, true
}

// SetActiveAddress sets field value
func (o *Addresses) SetActiveAddress(v string) {
	o.ActiveAddress = v
}

// GetAddresses returns the Addresses field value
func (o *Addresses) GetAddresses() []AddressInfo {
	if o == nil {
		var ret []AddressInfo
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *Addresses) GetAddressesOk() ([]AddressInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *Addresses) SetAddresses(v []AddressInfo) {
	o.Addresses = v
}

func (o Addresses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Addresses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeAddress"] = o.ActiveAddress
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

type NullableAddresses struct {
	value *Addresses
	isSet bool
}

func (v NullableAddresses) Get() *Addresses {
	return v.value
}

func (v *NullableAddresses) Set(val *Addresses) {
	v.value = val
	v.isSet = true
}

func (v NullableAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddresses(val *Addresses) *NullableAddresses {
	return &NullableAddresses{value: val, isSet: true}
}

func (v NullableAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


