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

// ContractState struct for ContractState
type ContractState struct {
	Address string `json:"address"`
	Bytecode string `json:"bytecode"`
	CodeHash string `json:"codeHash"`
	InitialStateHash *string `json:"initialStateHash,omitempty"`
	Fields []Val `json:"fields"`
	Asset AssetState `json:"asset"`
}

// NewContractState instantiates a new ContractState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractState(address string, bytecode string, codeHash string, fields []Val, asset AssetState) *ContractState {
	this := ContractState{}
	this.Address = address
	this.Bytecode = bytecode
	this.CodeHash = codeHash
	this.Fields = fields
	this.Asset = asset
	return &this
}

// NewContractStateWithDefaults instantiates a new ContractState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractStateWithDefaults() *ContractState {
	this := ContractState{}
	return &this
}

// GetAddress returns the Address field value
func (o *ContractState) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContractState) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContractState) SetAddress(v string) {
	o.Address = v
}

// GetBytecode returns the Bytecode field value
func (o *ContractState) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *ContractState) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *ContractState) SetBytecode(v string) {
	o.Bytecode = v
}

// GetCodeHash returns the CodeHash field value
func (o *ContractState) GetCodeHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value
// and a boolean to check if the value has been set.
func (o *ContractState) GetCodeHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeHash, true
}

// SetCodeHash sets field value
func (o *ContractState) SetCodeHash(v string) {
	o.CodeHash = v
}

// GetInitialStateHash returns the InitialStateHash field value if set, zero value otherwise.
func (o *ContractState) GetInitialStateHash() string {
	if o == nil || o.InitialStateHash == nil {
		var ret string
		return ret
	}
	return *o.InitialStateHash
}

// GetInitialStateHashOk returns a tuple with the InitialStateHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractState) GetInitialStateHashOk() (*string, bool) {
	if o == nil || o.InitialStateHash == nil {
		return nil, false
	}
	return o.InitialStateHash, true
}

// HasInitialStateHash returns a boolean if a field has been set.
func (o *ContractState) HasInitialStateHash() bool {
	if o != nil && o.InitialStateHash != nil {
		return true
	}

	return false
}

// SetInitialStateHash gets a reference to the given string and assigns it to the InitialStateHash field.
func (o *ContractState) SetInitialStateHash(v string) {
	o.InitialStateHash = &v
}

// GetFields returns the Fields field value
func (o *ContractState) GetFields() []Val {
	if o == nil {
		var ret []Val
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ContractState) GetFieldsOk() ([]Val, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *ContractState) SetFields(v []Val) {
	o.Fields = v
}

// GetAsset returns the Asset field value
func (o *ContractState) GetAsset() AssetState {
	if o == nil {
		var ret AssetState
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *ContractState) GetAssetOk() (*AssetState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *ContractState) SetAsset(v AssetState) {
	o.Asset = v
}

func (o ContractState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["bytecode"] = o.Bytecode
	}
	if true {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.InitialStateHash != nil {
		toSerialize["initialStateHash"] = o.InitialStateHash
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullableContractState struct {
	value *ContractState
	isSet bool
}

func (v NullableContractState) Get() *ContractState {
	return v.value
}

func (v *NullableContractState) Set(val *ContractState) {
	v.value = val
	v.isSet = true
}

func (v NullableContractState) IsSet() bool {
	return v.isSet
}

func (v *NullableContractState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractState(val *ContractState) *NullableContractState {
	return &NullableContractState{value: val, isSet: true}
}

func (v NullableContractState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


