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

// checks if the MultipleCallContractResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipleCallContractResult{}

// MultipleCallContractResult struct for MultipleCallContractResult
type MultipleCallContractResult struct {
	Results []CallContractResult `json:"results"`
}

// NewMultipleCallContractResult instantiates a new MultipleCallContractResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleCallContractResult(results []CallContractResult) *MultipleCallContractResult {
	this := MultipleCallContractResult{}
	this.Results = results
	return &this
}

// NewMultipleCallContractResultWithDefaults instantiates a new MultipleCallContractResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleCallContractResultWithDefaults() *MultipleCallContractResult {
	this := MultipleCallContractResult{}
	return &this
}

// GetResults returns the Results field value
func (o *MultipleCallContractResult) GetResults() []CallContractResult {
	if o == nil {
		var ret []CallContractResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *MultipleCallContractResult) GetResultsOk() ([]CallContractResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *MultipleCallContractResult) SetResults(v []CallContractResult) {
	o.Results = v
}

func (o MultipleCallContractResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipleCallContractResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableMultipleCallContractResult struct {
	value *MultipleCallContractResult
	isSet bool
}

func (v NullableMultipleCallContractResult) Get() *MultipleCallContractResult {
	return v.value
}

func (v *NullableMultipleCallContractResult) Set(val *MultipleCallContractResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleCallContractResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleCallContractResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleCallContractResult(val *MultipleCallContractResult) *NullableMultipleCallContractResult {
	return &NullableMultipleCallContractResult{value: val, isSet: true}
}

func (v NullableMultipleCallContractResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleCallContractResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


