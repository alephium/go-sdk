/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// checks if the CallContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallContract{}

// CallContract struct for CallContract
type CallContract struct {
	Group int32 `json:"group"`
	WorldStateBlockHash *string `json:"worldStateBlockHash,omitempty"`
	TxId *string `json:"txId,omitempty"`
	Address string `json:"address"`
	MethodIndex int32 `json:"methodIndex"`
	Args []Val `json:"args,omitempty"`
	ExistingContracts []string `json:"existingContracts,omitempty"`
	InputAssets []TestInputAsset `json:"inputAssets,omitempty"`
}

// NewCallContract instantiates a new CallContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallContract(group int32, address string, methodIndex int32) *CallContract {
	this := CallContract{}
	this.Group = group
	this.Address = address
	this.MethodIndex = methodIndex
	return &this
}

// NewCallContractWithDefaults instantiates a new CallContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallContractWithDefaults() *CallContract {
	this := CallContract{}
	return &this
}

// GetGroup returns the Group field value
func (o *CallContract) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *CallContract) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *CallContract) SetGroup(v int32) {
	o.Group = v
}

// GetWorldStateBlockHash returns the WorldStateBlockHash field value if set, zero value otherwise.
func (o *CallContract) GetWorldStateBlockHash() string {
	if o == nil || isNil(o.WorldStateBlockHash) {
		var ret string
		return ret
	}
	return *o.WorldStateBlockHash
}

// GetWorldStateBlockHashOk returns a tuple with the WorldStateBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallContract) GetWorldStateBlockHashOk() (*string, bool) {
	if o == nil || isNil(o.WorldStateBlockHash) {
		return nil, false
	}
	return o.WorldStateBlockHash, true
}

// HasWorldStateBlockHash returns a boolean if a field has been set.
func (o *CallContract) HasWorldStateBlockHash() bool {
	if o != nil && !isNil(o.WorldStateBlockHash) {
		return true
	}

	return false
}

// SetWorldStateBlockHash gets a reference to the given string and assigns it to the WorldStateBlockHash field.
func (o *CallContract) SetWorldStateBlockHash(v string) {
	o.WorldStateBlockHash = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *CallContract) GetTxId() string {
	if o == nil || isNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallContract) GetTxIdOk() (*string, bool) {
	if o == nil || isNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *CallContract) HasTxId() bool {
	if o != nil && !isNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *CallContract) SetTxId(v string) {
	o.TxId = &v
}

// GetAddress returns the Address field value
func (o *CallContract) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CallContract) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CallContract) SetAddress(v string) {
	o.Address = v
}

// GetMethodIndex returns the MethodIndex field value
func (o *CallContract) GetMethodIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MethodIndex
}

// GetMethodIndexOk returns a tuple with the MethodIndex field value
// and a boolean to check if the value has been set.
func (o *CallContract) GetMethodIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodIndex, true
}

// SetMethodIndex sets field value
func (o *CallContract) SetMethodIndex(v int32) {
	o.MethodIndex = v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CallContract) GetArgs() []Val {
	if o == nil || isNil(o.Args) {
		var ret []Val
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallContract) GetArgsOk() ([]Val, bool) {
	if o == nil || isNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CallContract) HasArgs() bool {
	if o != nil && !isNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []Val and assigns it to the Args field.
func (o *CallContract) SetArgs(v []Val) {
	o.Args = v
}

// GetExistingContracts returns the ExistingContracts field value if set, zero value otherwise.
func (o *CallContract) GetExistingContracts() []string {
	if o == nil || isNil(o.ExistingContracts) {
		var ret []string
		return ret
	}
	return o.ExistingContracts
}

// GetExistingContractsOk returns a tuple with the ExistingContracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallContract) GetExistingContractsOk() ([]string, bool) {
	if o == nil || isNil(o.ExistingContracts) {
		return nil, false
	}
	return o.ExistingContracts, true
}

// HasExistingContracts returns a boolean if a field has been set.
func (o *CallContract) HasExistingContracts() bool {
	if o != nil && !isNil(o.ExistingContracts) {
		return true
	}

	return false
}

// SetExistingContracts gets a reference to the given []string and assigns it to the ExistingContracts field.
func (o *CallContract) SetExistingContracts(v []string) {
	o.ExistingContracts = v
}

// GetInputAssets returns the InputAssets field value if set, zero value otherwise.
func (o *CallContract) GetInputAssets() []TestInputAsset {
	if o == nil || isNil(o.InputAssets) {
		var ret []TestInputAsset
		return ret
	}
	return o.InputAssets
}

// GetInputAssetsOk returns a tuple with the InputAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallContract) GetInputAssetsOk() ([]TestInputAsset, bool) {
	if o == nil || isNil(o.InputAssets) {
		return nil, false
	}
	return o.InputAssets, true
}

// HasInputAssets returns a boolean if a field has been set.
func (o *CallContract) HasInputAssets() bool {
	if o != nil && !isNil(o.InputAssets) {
		return true
	}

	return false
}

// SetInputAssets gets a reference to the given []TestInputAsset and assigns it to the InputAssets field.
func (o *CallContract) SetInputAssets(v []TestInputAsset) {
	o.InputAssets = v
}

func (o CallContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	if !isNil(o.WorldStateBlockHash) {
		toSerialize["worldStateBlockHash"] = o.WorldStateBlockHash
	}
	if !isNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	toSerialize["address"] = o.Address
	toSerialize["methodIndex"] = o.MethodIndex
	if !isNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !isNil(o.ExistingContracts) {
		toSerialize["existingContracts"] = o.ExistingContracts
	}
	if !isNil(o.InputAssets) {
		toSerialize["inputAssets"] = o.InputAssets
	}
	return toSerialize, nil
}

type NullableCallContract struct {
	value *CallContract
	isSet bool
}

func (v NullableCallContract) Get() *CallContract {
	return v.value
}

func (v *NullableCallContract) Set(val *CallContract) {
	v.value = val
	v.isSet = true
}

func (v NullableCallContract) IsSet() bool {
	return v.isSet
}

func (v *NullableCallContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallContract(val *CallContract) *NullableCallContract {
	return &NullableCallContract{value: val, isSet: true}
}

func (v NullableCallContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


