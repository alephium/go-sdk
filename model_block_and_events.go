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

// checks if the BlockAndEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockAndEvents{}

// BlockAndEvents struct for BlockAndEvents
type BlockAndEvents struct {
	Block BlockEntry `json:"block"`
	Events []ContractEventByBlockHash `json:"events"`
}

// NewBlockAndEvents instantiates a new BlockAndEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockAndEvents(block BlockEntry, events []ContractEventByBlockHash) *BlockAndEvents {
	this := BlockAndEvents{}
	this.Block = block
	this.Events = events
	return &this
}

// NewBlockAndEventsWithDefaults instantiates a new BlockAndEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockAndEventsWithDefaults() *BlockAndEvents {
	this := BlockAndEvents{}
	return &this
}

// GetBlock returns the Block field value
func (o *BlockAndEvents) GetBlock() BlockEntry {
	if o == nil {
		var ret BlockEntry
		return ret
	}

	return o.Block
}

// GetBlockOk returns a tuple with the Block field value
// and a boolean to check if the value has been set.
func (o *BlockAndEvents) GetBlockOk() (*BlockEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Block, true
}

// SetBlock sets field value
func (o *BlockAndEvents) SetBlock(v BlockEntry) {
	o.Block = v
}

// GetEvents returns the Events field value
func (o *BlockAndEvents) GetEvents() []ContractEventByBlockHash {
	if o == nil {
		var ret []ContractEventByBlockHash
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *BlockAndEvents) GetEventsOk() ([]ContractEventByBlockHash, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *BlockAndEvents) SetEvents(v []ContractEventByBlockHash) {
	o.Events = v
}

func (o BlockAndEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockAndEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["block"] = o.Block
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

type NullableBlockAndEvents struct {
	value *BlockAndEvents
	isSet bool
}

func (v NullableBlockAndEvents) Get() *BlockAndEvents {
	return v.value
}

func (v *NullableBlockAndEvents) Set(val *BlockAndEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockAndEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockAndEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockAndEvents(val *BlockAndEvents) *NullableBlockAndEvents {
	return &NullableBlockAndEvents{value: val, isSet: true}
}

func (v NullableBlockAndEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockAndEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


