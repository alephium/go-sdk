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

// ContractEvent struct for ContractEvent
type ContractEvent struct {
	BlockHash string `json:"blockHash"`
	TxId string `json:"txId"`
	EventIndex int32 `json:"eventIndex"`
	Fields []Val `json:"fields"`
}

// NewContractEvent instantiates a new ContractEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractEvent(blockHash string, txId string, eventIndex int32, fields []Val) *ContractEvent {
	this := ContractEvent{}
	this.BlockHash = blockHash
	this.TxId = txId
	this.EventIndex = eventIndex
	this.Fields = fields
	return &this
}

// NewContractEventWithDefaults instantiates a new ContractEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractEventWithDefaults() *ContractEvent {
	this := ContractEvent{}
	return &this
}

// GetBlockHash returns the BlockHash field value
func (o *ContractEvent) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *ContractEvent) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetTxId returns the TxId field value
func (o *ContractEvent) GetTxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetTxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *ContractEvent) SetTxId(v string) {
	o.TxId = v
}

// GetEventIndex returns the EventIndex field value
func (o *ContractEvent) GetEventIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventIndex
}

// GetEventIndexOk returns a tuple with the EventIndex field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetEventIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventIndex, true
}

// SetEventIndex sets field value
func (o *ContractEvent) SetEventIndex(v int32) {
	o.EventIndex = v
}

// GetFields returns the Fields field value
func (o *ContractEvent) GetFields() []Val {
	if o == nil {
		var ret []Val
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ContractEvent) GetFieldsOk() ([]Val, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *ContractEvent) SetFields(v []Val) {
	o.Fields = v
}

func (o ContractEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockHash"] = o.BlockHash
	}
	if true {
		toSerialize["txId"] = o.TxId
	}
	if true {
		toSerialize["eventIndex"] = o.EventIndex
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableContractEvent struct {
	value *ContractEvent
	isSet bool
}

func (v NullableContractEvent) Get() *ContractEvent {
	return v.value
}

func (v *NullableContractEvent) Set(val *ContractEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableContractEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableContractEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractEvent(val *ContractEvent) *NullableContractEvent {
	return &NullableContractEvent{value: val, isSet: true}
}

func (v NullableContractEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


