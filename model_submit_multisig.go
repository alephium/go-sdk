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

// checks if the SubmitMultisig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitMultisig{}

// SubmitMultisig struct for SubmitMultisig
type SubmitMultisig struct {
	UnsignedTx string `json:"unsignedTx"`
	Signatures []string `json:"signatures"`
}

// NewSubmitMultisig instantiates a new SubmitMultisig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitMultisig(unsignedTx string, signatures []string) *SubmitMultisig {
	this := SubmitMultisig{}
	this.UnsignedTx = unsignedTx
	this.Signatures = signatures
	return &this
}

// NewSubmitMultisigWithDefaults instantiates a new SubmitMultisig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitMultisigWithDefaults() *SubmitMultisig {
	this := SubmitMultisig{}
	return &this
}

// GetUnsignedTx returns the UnsignedTx field value
func (o *SubmitMultisig) GetUnsignedTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnsignedTx
}

// GetUnsignedTxOk returns a tuple with the UnsignedTx field value
// and a boolean to check if the value has been set.
func (o *SubmitMultisig) GetUnsignedTxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnsignedTx, true
}

// SetUnsignedTx sets field value
func (o *SubmitMultisig) SetUnsignedTx(v string) {
	o.UnsignedTx = v
}

// GetSignatures returns the Signatures field value
func (o *SubmitMultisig) GetSignatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value
// and a boolean to check if the value has been set.
func (o *SubmitMultisig) GetSignaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures, true
}

// SetSignatures sets field value
func (o *SubmitMultisig) SetSignatures(v []string) {
	o.Signatures = v
}

func (o SubmitMultisig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitMultisig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unsignedTx"] = o.UnsignedTx
	toSerialize["signatures"] = o.Signatures
	return toSerialize, nil
}

type NullableSubmitMultisig struct {
	value *SubmitMultisig
	isSet bool
}

func (v NullableSubmitMultisig) Get() *SubmitMultisig {
	return v.value
}

func (v *NullableSubmitMultisig) Set(val *SubmitMultisig) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitMultisig) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitMultisig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitMultisig(val *SubmitMultisig) *NullableSubmitMultisig {
	return &NullableSubmitMultisig{value: val, isSet: true}
}

func (v NullableSubmitMultisig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitMultisig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


