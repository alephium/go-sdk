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

// checks if the InterCliquePeerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterCliquePeerInfo{}

// InterCliquePeerInfo struct for InterCliquePeerInfo
type InterCliquePeerInfo struct {
	CliqueId string `json:"cliqueId"`
	BrokerId int32 `json:"brokerId"`
	GroupNumPerBroker int32 `json:"groupNumPerBroker"`
	Address BrokerInfoAddress `json:"address"`
	IsSynced bool `json:"isSynced"`
	ClientVersion string `json:"clientVersion"`
}

// NewInterCliquePeerInfo instantiates a new InterCliquePeerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterCliquePeerInfo(cliqueId string, brokerId int32, groupNumPerBroker int32, address BrokerInfoAddress, isSynced bool, clientVersion string) *InterCliquePeerInfo {
	this := InterCliquePeerInfo{}
	this.CliqueId = cliqueId
	this.BrokerId = brokerId
	this.GroupNumPerBroker = groupNumPerBroker
	this.Address = address
	this.IsSynced = isSynced
	this.ClientVersion = clientVersion
	return &this
}

// NewInterCliquePeerInfoWithDefaults instantiates a new InterCliquePeerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterCliquePeerInfoWithDefaults() *InterCliquePeerInfo {
	this := InterCliquePeerInfo{}
	return &this
}

// GetCliqueId returns the CliqueId field value
func (o *InterCliquePeerInfo) GetCliqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CliqueId
}

// GetCliqueIdOk returns a tuple with the CliqueId field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetCliqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CliqueId, true
}

// SetCliqueId sets field value
func (o *InterCliquePeerInfo) SetCliqueId(v string) {
	o.CliqueId = v
}

// GetBrokerId returns the BrokerId field value
func (o *InterCliquePeerInfo) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *InterCliquePeerInfo) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetGroupNumPerBroker returns the GroupNumPerBroker field value
func (o *InterCliquePeerInfo) GetGroupNumPerBroker() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupNumPerBroker
}

// GetGroupNumPerBrokerOk returns a tuple with the GroupNumPerBroker field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetGroupNumPerBrokerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupNumPerBroker, true
}

// SetGroupNumPerBroker sets field value
func (o *InterCliquePeerInfo) SetGroupNumPerBroker(v int32) {
	o.GroupNumPerBroker = v
}

// GetAddress returns the Address field value
func (o *InterCliquePeerInfo) GetAddress() BrokerInfoAddress {
	if o == nil {
		var ret BrokerInfoAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetAddressOk() (*BrokerInfoAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *InterCliquePeerInfo) SetAddress(v BrokerInfoAddress) {
	o.Address = v
}

// GetIsSynced returns the IsSynced field value
func (o *InterCliquePeerInfo) GetIsSynced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSynced
}

// GetIsSyncedOk returns a tuple with the IsSynced field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetIsSyncedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSynced, true
}

// SetIsSynced sets field value
func (o *InterCliquePeerInfo) SetIsSynced(v bool) {
	o.IsSynced = v
}

// GetClientVersion returns the ClientVersion field value
func (o *InterCliquePeerInfo) GetClientVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientVersion
}

// GetClientVersionOk returns a tuple with the ClientVersion field value
// and a boolean to check if the value has been set.
func (o *InterCliquePeerInfo) GetClientVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientVersion, true
}

// SetClientVersion sets field value
func (o *InterCliquePeerInfo) SetClientVersion(v string) {
	o.ClientVersion = v
}

func (o InterCliquePeerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterCliquePeerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cliqueId"] = o.CliqueId
	toSerialize["brokerId"] = o.BrokerId
	toSerialize["groupNumPerBroker"] = o.GroupNumPerBroker
	toSerialize["address"] = o.Address
	toSerialize["isSynced"] = o.IsSynced
	toSerialize["clientVersion"] = o.ClientVersion
	return toSerialize, nil
}

type NullableInterCliquePeerInfo struct {
	value *InterCliquePeerInfo
	isSet bool
}

func (v NullableInterCliquePeerInfo) Get() *InterCliquePeerInfo {
	return v.value
}

func (v *NullableInterCliquePeerInfo) Set(val *InterCliquePeerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInterCliquePeerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInterCliquePeerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterCliquePeerInfo(val *InterCliquePeerInfo) *NullableInterCliquePeerInfo {
	return &NullableInterCliquePeerInfo{value: val, isSet: true}
}

func (v NullableInterCliquePeerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterCliquePeerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


