/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// VerifySignature struct for VerifySignature
type VerifySignature struct {
	Data string `json:"data"`
	Signature string `json:"signature"`
	PublicKey string `json:"publicKey"`
}

// NewVerifySignature instantiates a new VerifySignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifySignature(data string, signature string, publicKey string) *VerifySignature {
	this := VerifySignature{}
	this.Data = data
	this.Signature = signature
	this.PublicKey = publicKey
	return &this
}

// NewVerifySignatureWithDefaults instantiates a new VerifySignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifySignatureWithDefaults() *VerifySignature {
	this := VerifySignature{}
	return &this
}

// GetData returns the Data field value
func (o *VerifySignature) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VerifySignature) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VerifySignature) SetData(v string) {
	o.Data = v
}

// GetSignature returns the Signature field value
func (o *VerifySignature) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *VerifySignature) GetSignatureOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *VerifySignature) SetSignature(v string) {
	o.Signature = v
}

// GetPublicKey returns the PublicKey field value
func (o *VerifySignature) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *VerifySignature) GetPublicKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *VerifySignature) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o VerifySignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["signature"] = o.Signature
	}
	if true {
		toSerialize["publicKey"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableVerifySignature struct {
	value *VerifySignature
	isSet bool
}

func (v NullableVerifySignature) Get() *VerifySignature {
	return v.value
}

func (v *NullableVerifySignature) Set(val *VerifySignature) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifySignature) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifySignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifySignature(val *VerifySignature) *NullableVerifySignature {
	return &NullableVerifySignature{value: val, isSet: true}
}

func (v NullableVerifySignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifySignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


