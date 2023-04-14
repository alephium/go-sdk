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

// checks if the DecodeUnsignedTx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecodeUnsignedTx{}

// DecodeUnsignedTx struct for DecodeUnsignedTx
type DecodeUnsignedTx struct {
	UnsignedTx string `json:"unsignedTx"`
}

// NewDecodeUnsignedTx instantiates a new DecodeUnsignedTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodeUnsignedTx(unsignedTx string) *DecodeUnsignedTx {
	this := DecodeUnsignedTx{}
	this.UnsignedTx = unsignedTx
	return &this
}

// NewDecodeUnsignedTxWithDefaults instantiates a new DecodeUnsignedTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodeUnsignedTxWithDefaults() *DecodeUnsignedTx {
	this := DecodeUnsignedTx{}
	return &this
}

// GetUnsignedTx returns the UnsignedTx field value
func (o *DecodeUnsignedTx) GetUnsignedTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedTx
}

// GetUnsignedTxOk returns a tuple with the UnsignedTx field value
// and a boolean to check if the value has been set.
func (o *DecodeUnsignedTx) GetUnsignedTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTx, true
}

// SetUnsignedTx sets field value
func (o *DecodeUnsignedTx) SetUnsignedTx(v string) {
	o.UnsignedTx = v
}

func (o DecodeUnsignedTx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecodeUnsignedTx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unsignedTx"] = o.UnsignedTx
	return toSerialize, nil
}

type NullableDecodeUnsignedTx struct {
	value *DecodeUnsignedTx
	isSet bool
}

func (v NullableDecodeUnsignedTx) Get() *DecodeUnsignedTx {
	return v.value
}

func (v *NullableDecodeUnsignedTx) Set(val *DecodeUnsignedTx) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodeUnsignedTx) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodeUnsignedTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodeUnsignedTx(val *DecodeUnsignedTx) *NullableDecodeUnsignedTx {
	return &NullableDecodeUnsignedTx{value: val, isSet: true}
}

func (v NullableDecodeUnsignedTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodeUnsignedTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


