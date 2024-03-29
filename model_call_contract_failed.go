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

// checks if the CallContractFailed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallContractFailed{}

// CallContractFailed struct for CallContractFailed
type CallContractFailed struct {
	Error string `json:"error"`
	Type string `json:"type"`
}

// NewCallContractFailed instantiates a new CallContractFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallContractFailed(error_ string, type_ string) *CallContractFailed {
	this := CallContractFailed{}
	this.Error = error_
	this.Type = type_
	return &this
}

// NewCallContractFailedWithDefaults instantiates a new CallContractFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallContractFailedWithDefaults() *CallContractFailed {
	this := CallContractFailed{}
	return &this
}

// GetError returns the Error field value
func (o *CallContractFailed) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *CallContractFailed) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *CallContractFailed) SetError(v string) {
	o.Error = v
}

// GetType returns the Type field value
func (o *CallContractFailed) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CallContractFailed) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CallContractFailed) SetType(v string) {
	o.Type = v
}

func (o CallContractFailed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallContractFailed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCallContractFailed struct {
	value *CallContractFailed
	isSet bool
}

func (v NullableCallContractFailed) Get() *CallContractFailed {
	return v.value
}

func (v *NullableCallContractFailed) Set(val *CallContractFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableCallContractFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableCallContractFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallContractFailed(val *CallContractFailed) *NullableCallContractFailed {
	return &NullableCallContractFailed{value: val, isSet: true}
}

func (v NullableCallContractFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallContractFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


