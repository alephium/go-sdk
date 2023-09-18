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

// checks if the TransferResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferResults{}

// TransferResults struct for TransferResults
type TransferResults struct {
	Results []TransferResult `json:"results"`
}

// NewTransferResults instantiates a new TransferResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferResults(results []TransferResult) *TransferResults {
	this := TransferResults{}
	this.Results = results
	return &this
}

// NewTransferResultsWithDefaults instantiates a new TransferResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferResultsWithDefaults() *TransferResults {
	this := TransferResults{}
	return &this
}

// GetResults returns the Results field value
func (o *TransferResults) GetResults() []TransferResult {
	if o == nil {
		var ret []TransferResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *TransferResults) GetResultsOk() ([]TransferResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *TransferResults) SetResults(v []TransferResult) {
	o.Results = v
}

func (o TransferResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableTransferResults struct {
	value *TransferResults
	isSet bool
}

func (v NullableTransferResults) Get() *TransferResults {
	return v.value
}

func (v *NullableTransferResults) Set(val *TransferResults) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferResults) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferResults(val *TransferResults) *NullableTransferResults {
	return &NullableTransferResults{value: val, isSet: true}
}

func (v NullableTransferResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


