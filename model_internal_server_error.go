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

// checks if the InternalServerError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalServerError{}

// InternalServerError struct for InternalServerError
type InternalServerError struct {
	Detail string `json:"detail"`
}

// NewInternalServerError instantiates a new InternalServerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalServerError(detail string) *InternalServerError {
	this := InternalServerError{}
	this.Detail = detail
	return &this
}

// NewInternalServerErrorWithDefaults instantiates a new InternalServerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalServerErrorWithDefaults() *InternalServerError {
	this := InternalServerError{}
	return &this
}

// GetDetail returns the Detail field value
func (o *InternalServerError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *InternalServerError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *InternalServerError) SetDetail(v string) {
	o.Detail = v
}

func (o InternalServerError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalServerError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	return toSerialize, nil
}

type NullableInternalServerError struct {
	value *InternalServerError
	isSet bool
}

func (v NullableInternalServerError) Get() *InternalServerError {
	return v.value
}

func (v *NullableInternalServerError) Set(val *InternalServerError) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalServerError) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalServerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalServerError(val *InternalServerError) *NullableInternalServerError {
	return &NullableInternalServerError{value: val, isSet: true}
}

func (v NullableInternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalServerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


