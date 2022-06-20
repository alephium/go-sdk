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

// BlockHeaderEntry struct for BlockHeaderEntry
type BlockHeaderEntry struct {
	Hash string `json:"hash"`
	Timestamp int64 `json:"timestamp"`
	ChainFrom int32 `json:"chainFrom"`
	ChainTo int32 `json:"chainTo"`
	Height int32 `json:"height"`
	Deps []string `json:"deps"`
}

// NewBlockHeaderEntry instantiates a new BlockHeaderEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeaderEntry(hash string, timestamp int64, chainFrom int32, chainTo int32, height int32, deps []string) *BlockHeaderEntry {
	this := BlockHeaderEntry{}
	this.Hash = hash
	this.Timestamp = timestamp
	this.ChainFrom = chainFrom
	this.ChainTo = chainTo
	this.Height = height
	this.Deps = deps
	return &this
}

// NewBlockHeaderEntryWithDefaults instantiates a new BlockHeaderEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeaderEntryWithDefaults() *BlockHeaderEntry {
	this := BlockHeaderEntry{}
	return &this
}

// GetHash returns the Hash field value
func (o *BlockHeaderEntry) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockHeaderEntry) SetHash(v string) {
	o.Hash = v
}

// GetTimestamp returns the Timestamp field value
func (o *BlockHeaderEntry) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BlockHeaderEntry) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetChainFrom returns the ChainFrom field value
func (o *BlockHeaderEntry) GetChainFrom() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainFrom
}

// GetChainFromOk returns a tuple with the ChainFrom field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetChainFromOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainFrom, true
}

// SetChainFrom sets field value
func (o *BlockHeaderEntry) SetChainFrom(v int32) {
	o.ChainFrom = v
}

// GetChainTo returns the ChainTo field value
func (o *BlockHeaderEntry) GetChainTo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainTo
}

// GetChainToOk returns a tuple with the ChainTo field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetChainToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainTo, true
}

// SetChainTo sets field value
func (o *BlockHeaderEntry) SetChainTo(v int32) {
	o.ChainTo = v
}

// GetHeight returns the Height field value
func (o *BlockHeaderEntry) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockHeaderEntry) SetHeight(v int32) {
	o.Height = v
}

// GetDeps returns the Deps field value
func (o *BlockHeaderEntry) GetDeps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Deps
}

// GetDepsOk returns a tuple with the Deps field value
// and a boolean to check if the value has been set.
func (o *BlockHeaderEntry) GetDepsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deps, true
}

// SetDeps sets field value
func (o *BlockHeaderEntry) SetDeps(v []string) {
	o.Deps = v
}

func (o BlockHeaderEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["chainFrom"] = o.ChainFrom
	}
	if true {
		toSerialize["chainTo"] = o.ChainTo
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["deps"] = o.Deps
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeaderEntry struct {
	value *BlockHeaderEntry
	isSet bool
}

func (v NullableBlockHeaderEntry) Get() *BlockHeaderEntry {
	return v.value
}

func (v *NullableBlockHeaderEntry) Set(val *BlockHeaderEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeaderEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeaderEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeaderEntry(val *BlockHeaderEntry) *NullableBlockHeaderEntry {
	return &NullableBlockHeaderEntry{value: val, isSet: true}
}

func (v NullableBlockHeaderEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeaderEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


