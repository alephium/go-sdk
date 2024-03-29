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

// checks if the PeerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeerAddress{}

// PeerAddress struct for PeerAddress
type PeerAddress struct {
	Address string `json:"address"`
	RestPort int32 `json:"restPort"`
	WsPort int32 `json:"wsPort"`
	MinerApiPort int32 `json:"minerApiPort"`
}

// NewPeerAddress instantiates a new PeerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeerAddress(address string, restPort int32, wsPort int32, minerApiPort int32) *PeerAddress {
	this := PeerAddress{}
	this.Address = address
	this.RestPort = restPort
	this.WsPort = wsPort
	this.MinerApiPort = minerApiPort
	return &this
}

// NewPeerAddressWithDefaults instantiates a new PeerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeerAddressWithDefaults() *PeerAddress {
	this := PeerAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *PeerAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PeerAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PeerAddress) SetAddress(v string) {
	o.Address = v
}

// GetRestPort returns the RestPort field value
func (o *PeerAddress) GetRestPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestPort
}

// GetRestPortOk returns a tuple with the RestPort field value
// and a boolean to check if the value has been set.
func (o *PeerAddress) GetRestPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestPort, true
}

// SetRestPort sets field value
func (o *PeerAddress) SetRestPort(v int32) {
	o.RestPort = v
}

// GetWsPort returns the WsPort field value
func (o *PeerAddress) GetWsPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WsPort
}

// GetWsPortOk returns a tuple with the WsPort field value
// and a boolean to check if the value has been set.
func (o *PeerAddress) GetWsPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WsPort, true
}

// SetWsPort sets field value
func (o *PeerAddress) SetWsPort(v int32) {
	o.WsPort = v
}

// GetMinerApiPort returns the MinerApiPort field value
func (o *PeerAddress) GetMinerApiPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinerApiPort
}

// GetMinerApiPortOk returns a tuple with the MinerApiPort field value
// and a boolean to check if the value has been set.
func (o *PeerAddress) GetMinerApiPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinerApiPort, true
}

// SetMinerApiPort sets field value
func (o *PeerAddress) SetMinerApiPort(v int32) {
	o.MinerApiPort = v
}

func (o PeerAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["restPort"] = o.RestPort
	toSerialize["wsPort"] = o.WsPort
	toSerialize["minerApiPort"] = o.MinerApiPort
	return toSerialize, nil
}

type NullablePeerAddress struct {
	value *PeerAddress
	isSet bool
}

func (v NullablePeerAddress) Get() *PeerAddress {
	return v.value
}

func (v *NullablePeerAddress) Set(val *PeerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePeerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePeerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeerAddress(val *PeerAddress) *NullablePeerAddress {
	return &NullablePeerAddress{value: val, isSet: true}
}

func (v NullablePeerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


