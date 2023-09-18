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

// checks if the MinerAddressesInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinerAddressesInfo{}

// MinerAddressesInfo struct for MinerAddressesInfo
type MinerAddressesInfo struct {
	Addresses []AddressInfo `json:"addresses"`
}

// NewMinerAddressesInfo instantiates a new MinerAddressesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinerAddressesInfo(addresses []AddressInfo) *MinerAddressesInfo {
	this := MinerAddressesInfo{}
	this.Addresses = addresses
	return &this
}

// NewMinerAddressesInfoWithDefaults instantiates a new MinerAddressesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinerAddressesInfoWithDefaults() *MinerAddressesInfo {
	this := MinerAddressesInfo{}
	return &this
}

// GetAddresses returns the Addresses field value
func (o *MinerAddressesInfo) GetAddresses() []AddressInfo {
	if o == nil {
		var ret []AddressInfo
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *MinerAddressesInfo) GetAddressesOk() ([]AddressInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *MinerAddressesInfo) SetAddresses(v []AddressInfo) {
	o.Addresses = v
}

func (o MinerAddressesInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinerAddressesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

type NullableMinerAddressesInfo struct {
	value *MinerAddressesInfo
	isSet bool
}

func (v NullableMinerAddressesInfo) Get() *MinerAddressesInfo {
	return v.value
}

func (v *NullableMinerAddressesInfo) Set(val *MinerAddressesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMinerAddressesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMinerAddressesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinerAddressesInfo(val *MinerAddressesInfo) *NullableMinerAddressesInfo {
	return &NullableMinerAddressesInfo{value: val, isSet: true}
}

func (v NullableMinerAddressesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinerAddressesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


