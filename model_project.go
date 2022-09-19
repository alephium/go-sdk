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

// Project struct for Project
type Project struct {
	Code string `json:"code"`
	CompilerOptions *CompilerOptions `json:"compilerOptions,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(code string) *Project {
	this := Project{}
	this.Code = code
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetCode returns the Code field value
func (o *Project) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Project) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Project) SetCode(v string) {
	o.Code = v
}

// GetCompilerOptions returns the CompilerOptions field value if set, zero value otherwise.
func (o *Project) GetCompilerOptions() CompilerOptions {
	if o == nil || o.CompilerOptions == nil {
		var ret CompilerOptions
		return ret
	}
	return *o.CompilerOptions
}

// GetCompilerOptionsOk returns a tuple with the CompilerOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCompilerOptionsOk() (*CompilerOptions, bool) {
	if o == nil || o.CompilerOptions == nil {
		return nil, false
	}
	return o.CompilerOptions, true
}

// HasCompilerOptions returns a boolean if a field has been set.
func (o *Project) HasCompilerOptions() bool {
	if o != nil && o.CompilerOptions != nil {
		return true
	}

	return false
}

// SetCompilerOptions gets a reference to the given CompilerOptions and assigns it to the CompilerOptions field.
func (o *Project) SetCompilerOptions(v CompilerOptions) {
	o.CompilerOptions = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.CompilerOptions != nil {
		toSerialize["compilerOptions"] = o.CompilerOptions
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


