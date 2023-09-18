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

// checks if the DecodeUnsignedTxResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecodeUnsignedTxResult{}

// DecodeUnsignedTxResult struct for DecodeUnsignedTxResult
type DecodeUnsignedTxResult struct {
	FromGroup int32 `json:"fromGroup"`
	ToGroup int32 `json:"toGroup"`
	UnsignedTx UnsignedTx `json:"unsignedTx"`
}

// NewDecodeUnsignedTxResult instantiates a new DecodeUnsignedTxResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeUnsignedTxResult(fromGroup int32, toGroup int32, unsignedTx UnsignedTx) *DecodeUnsignedTxResult {
	this := DecodeUnsignedTxResult{}
	this.FromGroup = fromGroup
	this.ToGroup = toGroup
	this.UnsignedTx = unsignedTx
	return &this
}

// NewDecodeUnsignedTxResultWithDefaults instantiates a new DecodeUnsignedTxResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeUnsignedTxResultWithDefaults() *DecodeUnsignedTxResult {
	this := DecodeUnsignedTxResult{}
	return &this
}

// GetFromGroup returns the FromGroup field value
func (o *DecodeUnsignedTxResult) GetFromGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromGroup
}

// GetFromGroupOk returns a tuple with the FromGroup field value
// and a boolean to check if the value has been set.
func (o *DecodeUnsignedTxResult) GetFromGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromGroup, true
}

// SetFromGroup sets field value
func (o *DecodeUnsignedTxResult) SetFromGroup(v int32) {
	o.FromGroup = v
}

// GetToGroup returns the ToGroup field value
func (o *DecodeUnsignedTxResult) GetToGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToGroup
}

// GetToGroupOk returns a tuple with the ToGroup field value
// and a boolean to check if the value has been set.
func (o *DecodeUnsignedTxResult) GetToGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToGroup, true
}

// SetToGroup sets field value
func (o *DecodeUnsignedTxResult) SetToGroup(v int32) {
	o.ToGroup = v
}

// GetUnsignedTx returns the UnsignedTx field value
func (o *DecodeUnsignedTxResult) GetUnsignedTx() UnsignedTx {
	if o == nil {
		var ret UnsignedTx
		return ret
	}

	return o.UnsignedTx
}

// GetUnsignedTxOk returns a tuple with the UnsignedTx field value
// and a boolean to check if the value has been set.
func (o *DecodeUnsignedTxResult) GetUnsignedTxOk() (*UnsignedTx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTx, true
}

// SetUnsignedTx sets field value
func (o *DecodeUnsignedTxResult) SetUnsignedTx(v UnsignedTx) {
	o.UnsignedTx = v
}

func (o DecodeUnsignedTxResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecodeUnsignedTxResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromGroup"] = o.FromGroup
	toSerialize["toGroup"] = o.ToGroup
	toSerialize["unsignedTx"] = o.UnsignedTx
	return toSerialize, nil
}

type NullableDecodeUnsignedTxResult struct {
	value *DecodeUnsignedTxResult
	isSet bool
}

func (v NullableDecodeUnsignedTxResult) Get() *DecodeUnsignedTxResult {
	return v.value
}

func (v *NullableDecodeUnsignedTxResult) Set(val *DecodeUnsignedTxResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeUnsignedTxResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeUnsignedTxResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeUnsignedTxResult(val *DecodeUnsignedTxResult) *NullableDecodeUnsignedTxResult {
	return &NullableDecodeUnsignedTxResult{value: val, isSet: true}
}

func (v NullableDecodeUnsignedTxResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeUnsignedTxResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


