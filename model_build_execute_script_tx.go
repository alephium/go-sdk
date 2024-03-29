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

// checks if the BuildExecuteScriptTx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildExecuteScriptTx{}

// BuildExecuteScriptTx struct for BuildExecuteScriptTx
type BuildExecuteScriptTx struct {
	FromPublicKey string `json:"fromPublicKey"`
	FromPublicKeyType *string `json:"fromPublicKeyType,omitempty"`
	Bytecode string `json:"bytecode"`
	AttoAlphAmount *string `json:"attoAlphAmount,omitempty"`
	Tokens []Token `json:"tokens,omitempty"`
	GasAmount *int32 `json:"gasAmount,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
	TargetBlockHash *string `json:"targetBlockHash,omitempty"`
}

// NewBuildExecuteScriptTx instantiates a new BuildExecuteScriptTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildExecuteScriptTx(fromPublicKey string, bytecode string) *BuildExecuteScriptTx {
	this := BuildExecuteScriptTx{}
	this.FromPublicKey = fromPublicKey
	this.Bytecode = bytecode
	return &this
}

// NewBuildExecuteScriptTxWithDefaults instantiates a new BuildExecuteScriptTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildExecuteScriptTxWithDefaults() *BuildExecuteScriptTx {
	this := BuildExecuteScriptTx{}
	return &this
}

// GetFromPublicKey returns the FromPublicKey field value
func (o *BuildExecuteScriptTx) GetFromPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromPublicKey
}

// GetFromPublicKeyOk returns a tuple with the FromPublicKey field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetFromPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPublicKey, true
}

// SetFromPublicKey sets field value
func (o *BuildExecuteScriptTx) SetFromPublicKey(v string) {
	o.FromPublicKey = v
}

// GetFromPublicKeyType returns the FromPublicKeyType field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetFromPublicKeyType() string {
	if o == nil || isNil(o.FromPublicKeyType) {
		var ret string
		return ret
	}
	return *o.FromPublicKeyType
}

// GetFromPublicKeyTypeOk returns a tuple with the FromPublicKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetFromPublicKeyTypeOk() (*string, bool) {
	if o == nil || isNil(o.FromPublicKeyType) {
		return nil, false
	}
	return o.FromPublicKeyType, true
}

// HasFromPublicKeyType returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasFromPublicKeyType() bool {
	if o != nil && !isNil(o.FromPublicKeyType) {
		return true
	}

	return false
}

// SetFromPublicKeyType gets a reference to the given string and assigns it to the FromPublicKeyType field.
func (o *BuildExecuteScriptTx) SetFromPublicKeyType(v string) {
	o.FromPublicKeyType = &v
}

// GetBytecode returns the Bytecode field value
func (o *BuildExecuteScriptTx) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *BuildExecuteScriptTx) SetBytecode(v string) {
	o.Bytecode = v
}

// GetAttoAlphAmount returns the AttoAlphAmount field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetAttoAlphAmount() string {
	if o == nil || isNil(o.AttoAlphAmount) {
		var ret string
		return ret
	}
	return *o.AttoAlphAmount
}

// GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetAttoAlphAmountOk() (*string, bool) {
	if o == nil || isNil(o.AttoAlphAmount) {
		return nil, false
	}
	return o.AttoAlphAmount, true
}

// HasAttoAlphAmount returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasAttoAlphAmount() bool {
	if o != nil && !isNil(o.AttoAlphAmount) {
		return true
	}

	return false
}

// SetAttoAlphAmount gets a reference to the given string and assigns it to the AttoAlphAmount field.
func (o *BuildExecuteScriptTx) SetAttoAlphAmount(v string) {
	o.AttoAlphAmount = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetTokens() []Token {
	if o == nil || isNil(o.Tokens) {
		var ret []Token
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetTokensOk() ([]Token, bool) {
	if o == nil || isNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasTokens() bool {
	if o != nil && !isNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []Token and assigns it to the Tokens field.
func (o *BuildExecuteScriptTx) SetTokens(v []Token) {
	o.Tokens = v
}

// GetGasAmount returns the GasAmount field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetGasAmount() int32 {
	if o == nil || isNil(o.GasAmount) {
		var ret int32
		return ret
	}
	return *o.GasAmount
}

// GetGasAmountOk returns a tuple with the GasAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetGasAmountOk() (*int32, bool) {
	if o == nil || isNil(o.GasAmount) {
		return nil, false
	}
	return o.GasAmount, true
}

// HasGasAmount returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasGasAmount() bool {
	if o != nil && !isNil(o.GasAmount) {
		return true
	}

	return false
}

// SetGasAmount gets a reference to the given int32 and assigns it to the GasAmount field.
func (o *BuildExecuteScriptTx) SetGasAmount(v int32) {
	o.GasAmount = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetGasPrice() string {
	if o == nil || isNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetGasPriceOk() (*string, bool) {
	if o == nil || isNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasGasPrice() bool {
	if o != nil && !isNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *BuildExecuteScriptTx) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetTargetBlockHash returns the TargetBlockHash field value if set, zero value otherwise.
func (o *BuildExecuteScriptTx) GetTargetBlockHash() string {
	if o == nil || isNil(o.TargetBlockHash) {
		var ret string
		return ret
	}
	return *o.TargetBlockHash
}

// GetTargetBlockHashOk returns a tuple with the TargetBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExecuteScriptTx) GetTargetBlockHashOk() (*string, bool) {
	if o == nil || isNil(o.TargetBlockHash) {
		return nil, false
	}
	return o.TargetBlockHash, true
}

// HasTargetBlockHash returns a boolean if a field has been set.
func (o *BuildExecuteScriptTx) HasTargetBlockHash() bool {
	if o != nil && !isNil(o.TargetBlockHash) {
		return true
	}

	return false
}

// SetTargetBlockHash gets a reference to the given string and assigns it to the TargetBlockHash field.
func (o *BuildExecuteScriptTx) SetTargetBlockHash(v string) {
	o.TargetBlockHash = &v
}

func (o BuildExecuteScriptTx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildExecuteScriptTx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromPublicKey"] = o.FromPublicKey
	if !isNil(o.FromPublicKeyType) {
		toSerialize["fromPublicKeyType"] = o.FromPublicKeyType
	}
	toSerialize["bytecode"] = o.Bytecode
	if !isNil(o.AttoAlphAmount) {
		toSerialize["attoAlphAmount"] = o.AttoAlphAmount
	}
	if !isNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !isNil(o.GasAmount) {
		toSerialize["gasAmount"] = o.GasAmount
	}
	if !isNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !isNil(o.TargetBlockHash) {
		toSerialize["targetBlockHash"] = o.TargetBlockHash
	}
	return toSerialize, nil
}

type NullableBuildExecuteScriptTx struct {
	value *BuildExecuteScriptTx
	isSet bool
}

func (v NullableBuildExecuteScriptTx) Get() *BuildExecuteScriptTx {
	return v.value
}

func (v *NullableBuildExecuteScriptTx) Set(val *BuildExecuteScriptTx) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildExecuteScriptTx) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildExecuteScriptTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildExecuteScriptTx(val *BuildExecuteScriptTx) *NullableBuildExecuteScriptTx {
	return &NullableBuildExecuteScriptTx{value: val, isSet: true}
}

func (v NullableBuildExecuteScriptTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildExecuteScriptTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


