/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// ContractOutput struct for ContractOutput
type ContractOutput struct {
	Hint int32 `json:"hint"`
	Key string `json:"key"`
	AttoAlphAmount string `json:"attoAlphAmount"`
	Address string `json:"address"`
	Tokens []Token `json:"tokens"`
	Type string `json:"type"`
}

// NewContractOutput instantiates a new ContractOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractOutput(hint int32, key string, attoAlphAmount string, address string, tokens []Token, type_ string) *ContractOutput {
	this := ContractOutput{}
	this.Hint = hint
	this.Key = key
	this.AttoAlphAmount = attoAlphAmount
	this.Address = address
	this.Tokens = tokens
	this.Type = type_
	return &this
}

// NewContractOutputWithDefaults instantiates a new ContractOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractOutputWithDefaults() *ContractOutput {
	this := ContractOutput{}
	return &this
}

// GetHint returns the Hint field value
func (o *ContractOutput) GetHint() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hint
}

// GetHintOk returns a tuple with the Hint field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetHintOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hint, true
}

// SetHint sets field value
func (o *ContractOutput) SetHint(v int32) {
	o.Hint = v
}

// GetKey returns the Key field value
func (o *ContractOutput) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ContractOutput) SetKey(v string) {
	o.Key = v
}

// GetAttoAlphAmount returns the AttoAlphAmount field value
func (o *ContractOutput) GetAttoAlphAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttoAlphAmount
}

// GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetAttoAlphAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttoAlphAmount, true
}

// SetAttoAlphAmount sets field value
func (o *ContractOutput) SetAttoAlphAmount(v string) {
	o.AttoAlphAmount = v
}

// GetAddress returns the Address field value
func (o *ContractOutput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContractOutput) SetAddress(v string) {
	o.Address = v
}

// GetTokens returns the Tokens field value
func (o *ContractOutput) GetTokens() []Token {
	if o == nil {
		var ret []Token
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetTokensOk() ([]Token, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *ContractOutput) SetTokens(v []Token) {
	o.Tokens = v
}

// GetType returns the Type field value
func (o *ContractOutput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContractOutput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContractOutput) SetType(v string) {
	o.Type = v
}

func (o ContractOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hint"] = o.Hint
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["attoAlphAmount"] = o.AttoAlphAmount
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableContractOutput struct {
	value *ContractOutput
	isSet bool
}

func (v NullableContractOutput) Get() *ContractOutput {
	return v.value
}

func (v *NullableContractOutput) Set(val *ContractOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableContractOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableContractOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractOutput(val *ContractOutput) *NullableContractOutput {
	return &NullableContractOutput{value: val, isSet: true}
}

func (v NullableContractOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


