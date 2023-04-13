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

// checks if the Script type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Script{}

// Script struct for Script
type Script struct {
	Code string `json:"code"`
	CompilerOptions *CompilerOptions `json:"compilerOptions,omitempty"`
}

// NewScript instantiates a new Script object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScript(code string) *Script {
	this := Script{}
	this.Code = code
	return &this
}

// NewScriptWithDefaults instantiates a new Script object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptWithDefaults() *Script {
	this := Script{}
	return &this
}

// GetCode returns the Code field value
func (o *Script) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Script) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Script) SetCode(v string) {
	o.Code = v
}

// GetCompilerOptions returns the CompilerOptions field value if set, zero value otherwise.
func (o *Script) GetCompilerOptions() CompilerOptions {
	if o == nil || isNil(o.CompilerOptions) {
		var ret CompilerOptions
		return ret
	}
	return *o.CompilerOptions
}

// GetCompilerOptionsOk returns a tuple with the CompilerOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Script) GetCompilerOptionsOk() (*CompilerOptions, bool) {
	if o == nil || isNil(o.CompilerOptions) {
		return nil, false
	}
	return o.CompilerOptions, true
}

// HasCompilerOptions returns a boolean if a field has been set.
func (o *Script) HasCompilerOptions() bool {
	if o != nil && !isNil(o.CompilerOptions) {
		return true
	}

	return false
}

// SetCompilerOptions gets a reference to the given CompilerOptions and assigns it to the CompilerOptions field.
func (o *Script) SetCompilerOptions(v CompilerOptions) {
	o.CompilerOptions = &v
}

func (o Script) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Script) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !isNil(o.CompilerOptions) {
		toSerialize["compilerOptions"] = o.CompilerOptions
	}
	return toSerialize, nil
}

type NullableScript struct {
	value *Script
	isSet bool
}

func (v NullableScript) Get() *Script {
	return v.value
}

func (v *NullableScript) Set(val *Script) {
	v.value = val
	v.isSet = true
}

func (v NullableScript) IsSet() bool {
	return v.isSet
}

func (v *NullableScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScript(val *Script) *NullableScript {
	return &NullableScript{value: val, isSet: true}
}

func (v NullableScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


