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

// checks if the BlocksPerTimeStampRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlocksPerTimeStampRange{}

// BlocksPerTimeStampRange struct for BlocksPerTimeStampRange
type BlocksPerTimeStampRange struct {
	Blocks [][]BlockEntry `json:"blocks"`
}

// NewBlocksPerTimeStampRange instantiates a new BlocksPerTimeStampRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlocksPerTimeStampRange(blocks [][]BlockEntry) *BlocksPerTimeStampRange {
	this := BlocksPerTimeStampRange{}
	this.Blocks = blocks
	return &this
}

// NewBlocksPerTimeStampRangeWithDefaults instantiates a new BlocksPerTimeStampRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlocksPerTimeStampRangeWithDefaults() *BlocksPerTimeStampRange {
	this := BlocksPerTimeStampRange{}
	return &this
}

// GetBlocks returns the Blocks field value
func (o *BlocksPerTimeStampRange) GetBlocks() [][]BlockEntry {
	if o == nil {
		var ret [][]BlockEntry
		return ret
	}

	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value
// and a boolean to check if the value has been set.
func (o *BlocksPerTimeStampRange) GetBlocksOk() ([][]BlockEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blocks, true
}

// SetBlocks sets field value
func (o *BlocksPerTimeStampRange) SetBlocks(v [][]BlockEntry) {
	o.Blocks = v
}

func (o BlocksPerTimeStampRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlocksPerTimeStampRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blocks"] = o.Blocks
	return toSerialize, nil
}

type NullableBlocksPerTimeStampRange struct {
	value *BlocksPerTimeStampRange
	isSet bool
}

func (v NullableBlocksPerTimeStampRange) Get() *BlocksPerTimeStampRange {
	return v.value
}

func (v *NullableBlocksPerTimeStampRange) Set(val *BlocksPerTimeStampRange) {
	v.value = val
	v.isSet = true
}

func (v NullableBlocksPerTimeStampRange) IsSet() bool {
	return v.isSet
}

func (v *NullableBlocksPerTimeStampRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlocksPerTimeStampRange(val *BlocksPerTimeStampRange) *NullableBlocksPerTimeStampRange {
	return &NullableBlocksPerTimeStampRange{value: val, isSet: true}
}

func (v NullableBlocksPerTimeStampRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlocksPerTimeStampRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


