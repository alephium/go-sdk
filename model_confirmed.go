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

// Confirmed struct for Confirmed
type Confirmed struct {
	BlockHash string `json:"blockHash"`
	TxIndex int32 `json:"txIndex"`
	ChainConfirmations int32 `json:"chainConfirmations"`
	FromGroupConfirmations int32 `json:"fromGroupConfirmations"`
	ToGroupConfirmations int32 `json:"toGroupConfirmations"`
	Type string `json:"type"`
}

// NewConfirmed instantiates a new Confirmed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmed(blockHash string, txIndex int32, chainConfirmations int32, fromGroupConfirmations int32, toGroupConfirmations int32, type_ string) *Confirmed {
	this := Confirmed{}
	this.BlockHash = blockHash
	this.TxIndex = txIndex
	this.ChainConfirmations = chainConfirmations
	this.FromGroupConfirmations = fromGroupConfirmations
	this.ToGroupConfirmations = toGroupConfirmations
	this.Type = type_
	return &this
}

// NewConfirmedWithDefaults instantiates a new Confirmed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmedWithDefaults() *Confirmed {
	this := Confirmed{}
	return &this
}

// GetBlockHash returns the BlockHash field value
func (o *Confirmed) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *Confirmed) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetTxIndex returns the TxIndex field value
func (o *Confirmed) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *Confirmed) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetChainConfirmations returns the ChainConfirmations field value
func (o *Confirmed) GetChainConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainConfirmations
}

// GetChainConfirmationsOk returns a tuple with the ChainConfirmations field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetChainConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainConfirmations, true
}

// SetChainConfirmations sets field value
func (o *Confirmed) SetChainConfirmations(v int32) {
	o.ChainConfirmations = v
}

// GetFromGroupConfirmations returns the FromGroupConfirmations field value
func (o *Confirmed) GetFromGroupConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromGroupConfirmations
}

// GetFromGroupConfirmationsOk returns a tuple with the FromGroupConfirmations field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetFromGroupConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromGroupConfirmations, true
}

// SetFromGroupConfirmations sets field value
func (o *Confirmed) SetFromGroupConfirmations(v int32) {
	o.FromGroupConfirmations = v
}

// GetToGroupConfirmations returns the ToGroupConfirmations field value
func (o *Confirmed) GetToGroupConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToGroupConfirmations
}

// GetToGroupConfirmationsOk returns a tuple with the ToGroupConfirmations field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetToGroupConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToGroupConfirmations, true
}

// SetToGroupConfirmations sets field value
func (o *Confirmed) SetToGroupConfirmations(v int32) {
	o.ToGroupConfirmations = v
}

// GetType returns the Type field value
func (o *Confirmed) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Confirmed) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Confirmed) SetType(v string) {
	o.Type = v
}

func (o Confirmed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blockHash"] = o.BlockHash
	}
	if true {
		toSerialize["txIndex"] = o.TxIndex
	}
	if true {
		toSerialize["chainConfirmations"] = o.ChainConfirmations
	}
	if true {
		toSerialize["fromGroupConfirmations"] = o.FromGroupConfirmations
	}
	if true {
		toSerialize["toGroupConfirmations"] = o.ToGroupConfirmations
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableConfirmed struct {
	value *Confirmed
	isSet bool
}

func (v NullableConfirmed) Get() *Confirmed {
	return v.value
}

func (v *NullableConfirmed) Set(val *Confirmed) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmed) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmed(val *Confirmed) *NullableConfirmed {
	return &NullableConfirmed{value: val, isSet: true}
}

func (v NullableConfirmed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


