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

// ContractEventsByTxId struct for ContractEventsByTxId
type ContractEventsByTxId struct {
	Events []ContractEventByTxId `json:"events"`
}

// NewContractEventsByTxId instantiates a new ContractEventsByTxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractEventsByTxId(events []ContractEventByTxId) *ContractEventsByTxId {
	this := ContractEventsByTxId{}
	this.Events = events
	return &this
}

// NewContractEventsByTxIdWithDefaults instantiates a new ContractEventsByTxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractEventsByTxIdWithDefaults() *ContractEventsByTxId {
	this := ContractEventsByTxId{}
	return &this
}

// GetEvents returns the Events field value
func (o *ContractEventsByTxId) GetEvents() []ContractEventByTxId {
	if o == nil {
		var ret []ContractEventByTxId
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ContractEventsByTxId) GetEventsOk() ([]ContractEventByTxId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ContractEventsByTxId) SetEvents(v []ContractEventByTxId) {
	o.Events = v
}

func (o ContractEventsByTxId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableContractEventsByTxId struct {
	value *ContractEventsByTxId
	isSet bool
}

func (v NullableContractEventsByTxId) Get() *ContractEventsByTxId {
	return v.value
}

func (v *NullableContractEventsByTxId) Set(val *ContractEventsByTxId) {
	v.value = val
	v.isSet = true
}

func (v NullableContractEventsByTxId) IsSet() bool {
	return v.isSet
}

func (v *NullableContractEventsByTxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractEventsByTxId(val *ContractEventsByTxId) *NullableContractEventsByTxId {
	return &NullableContractEventsByTxId{value: val, isSet: true}
}

func (v NullableContractEventsByTxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractEventsByTxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


