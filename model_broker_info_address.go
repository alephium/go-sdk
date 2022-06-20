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

// BrokerInfoAddress struct for BrokerInfoAddress
type BrokerInfoAddress struct {
	Addr string `json:"addr"`
	Port int32 `json:"port"`
}

// NewBrokerInfoAddress instantiates a new BrokerInfoAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerInfoAddress(addr string, port int32) *BrokerInfoAddress {
	this := BrokerInfoAddress{}
	this.Addr = addr
	this.Port = port
	return &this
}

// NewBrokerInfoAddressWithDefaults instantiates a new BrokerInfoAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerInfoAddressWithDefaults() *BrokerInfoAddress {
	this := BrokerInfoAddress{}
	return &this
}

// GetAddr returns the Addr field value
func (o *BrokerInfoAddress) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *BrokerInfoAddress) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *BrokerInfoAddress) SetAddr(v string) {
	o.Addr = v
}

// GetPort returns the Port field value
func (o *BrokerInfoAddress) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *BrokerInfoAddress) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *BrokerInfoAddress) SetPort(v int32) {
	o.Port = v
}

func (o BrokerInfoAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["addr"] = o.Addr
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableBrokerInfoAddress struct {
	value *BrokerInfoAddress
	isSet bool
}

func (v NullableBrokerInfoAddress) Get() *BrokerInfoAddress {
	return v.value
}

func (v *NullableBrokerInfoAddress) Set(val *BrokerInfoAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerInfoAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerInfoAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerInfoAddress(val *BrokerInfoAddress) *NullableBrokerInfoAddress {
	return &NullableBrokerInfoAddress{value: val, isSet: true}
}

func (v NullableBrokerInfoAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerInfoAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


