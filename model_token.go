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

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Token{}

// Token struct for Token
type Token struct {
	Id string `json:"id"`
	Amount string `json:"amount"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(id string, amount string) *Token {
	this := Token{}
	this.Id = id
	this.Amount = amount
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetId returns the Id field value
func (o *Token) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Token) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Token) SetId(v string) {
	o.Id = v
}

// GetAmount returns the Amount field value
func (o *Token) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Token) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Token) SetAmount(v string) {
	o.Amount = v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


