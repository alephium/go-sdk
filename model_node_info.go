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

// checks if the NodeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInfo{}

// NodeInfo struct for NodeInfo
type NodeInfo struct {
	BuildInfo BuildInfo `json:"buildInfo"`
	Upnp bool `json:"upnp"`
	ExternalAddress *BrokerInfoAddress `json:"externalAddress,omitempty"`
}

// NewNodeInfo instantiates a new NodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInfo(buildInfo BuildInfo, upnp bool) *NodeInfo {
	this := NodeInfo{}
	this.BuildInfo = buildInfo
	this.Upnp = upnp
	return &this
}

// NewNodeInfoWithDefaults instantiates a new NodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInfoWithDefaults() *NodeInfo {
	this := NodeInfo{}
	return &this
}

// GetBuildInfo returns the BuildInfo field value
func (o *NodeInfo) GetBuildInfo() BuildInfo {
	if o == nil {
		var ret BuildInfo
		return ret
	}

	return o.BuildInfo
}

// GetBuildInfoOk returns a tuple with the BuildInfo field value
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetBuildInfoOk() (*BuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildInfo, true
}

// SetBuildInfo sets field value
func (o *NodeInfo) SetBuildInfo(v BuildInfo) {
	o.BuildInfo = v
}

// GetUpnp returns the Upnp field value
func (o *NodeInfo) GetUpnp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Upnp
}

// GetUpnpOk returns a tuple with the Upnp field value
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetUpnpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upnp, true
}

// SetUpnp sets field value
func (o *NodeInfo) SetUpnp(v bool) {
	o.Upnp = v
}

// GetExternalAddress returns the ExternalAddress field value if set, zero value otherwise.
func (o *NodeInfo) GetExternalAddress() BrokerInfoAddress {
	if o == nil || isNil(o.ExternalAddress) {
		var ret BrokerInfoAddress
		return ret
	}
	return *o.ExternalAddress
}

// GetExternalAddressOk returns a tuple with the ExternalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeInfo) GetExternalAddressOk() (*BrokerInfoAddress, bool) {
	if o == nil || isNil(o.ExternalAddress) {
		return nil, false
	}
	return o.ExternalAddress, true
}

// HasExternalAddress returns a boolean if a field has been set.
func (o *NodeInfo) HasExternalAddress() bool {
	if o != nil && !isNil(o.ExternalAddress) {
		return true
	}

	return false
}

// SetExternalAddress gets a reference to the given BrokerInfoAddress and assigns it to the ExternalAddress field.
func (o *NodeInfo) SetExternalAddress(v BrokerInfoAddress) {
	o.ExternalAddress = &v
}

func (o NodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildInfo"] = o.BuildInfo
	toSerialize["upnp"] = o.Upnp
	if !isNil(o.ExternalAddress) {
		toSerialize["externalAddress"] = o.ExternalAddress
	}
	return toSerialize, nil
}

type NullableNodeInfo struct {
	value *NodeInfo
	isSet bool
}

func (v NullableNodeInfo) Get() *NodeInfo {
	return v.value
}

func (v *NullableNodeInfo) Set(val *NodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInfo(val *NodeInfo) *NullableNodeInfo {
	return &NullableNodeInfo{value: val, isSet: true}
}

func (v NullableNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


