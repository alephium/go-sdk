/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// checks if the FieldsSig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldsSig{}

// FieldsSig struct for FieldsSig
type FieldsSig struct {
	Names []string `json:"names"`
	Types []string `json:"types"`
	IsMutable []bool `json:"isMutable"`
}

// NewFieldsSig instantiates a new FieldsSig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldsSig(names []string, types []string, isMutable []bool) *FieldsSig {
	this := FieldsSig{}
	this.Names = names
	this.Types = types
	this.IsMutable = isMutable
	return &this
}

// NewFieldsSigWithDefaults instantiates a new FieldsSig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldsSigWithDefaults() *FieldsSig {
	this := FieldsSig{}
	return &this
}

// GetNames returns the Names field value
func (o *FieldsSig) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *FieldsSig) GetNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *FieldsSig) SetNames(v []string) {
	o.Names = v
}

// GetTypes returns the Types field value
func (o *FieldsSig) GetTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *FieldsSig) GetTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *FieldsSig) SetTypes(v []string) {
	o.Types = v
}

// GetIsMutable returns the IsMutable field value
func (o *FieldsSig) GetIsMutable() []bool {
	if o == nil {
		var ret []bool
		return ret
	}

	return o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value
// and a boolean to check if the value has been set.
func (o *FieldsSig) GetIsMutableOk() ([]bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// SetIsMutable sets field value
func (o *FieldsSig) SetIsMutable(v []bool) {
	o.IsMutable = v
}

func (o FieldsSig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldsSig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["names"] = o.Names
	toSerialize["types"] = o.Types
	toSerialize["isMutable"] = o.IsMutable
	return toSerialize, nil
}

type NullableFieldsSig struct {
	value *FieldsSig
	isSet bool
}

func (v NullableFieldsSig) Get() *FieldsSig {
	return v.value
}

func (v *NullableFieldsSig) Set(val *FieldsSig) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldsSig) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldsSig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldsSig(val *FieldsSig) *NullableFieldsSig {
	return &NullableFieldsSig{value: val, isSet: true}
}

func (v NullableFieldsSig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldsSig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


