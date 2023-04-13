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

// checks if the TestContractResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestContractResult{}

// TestContractResult struct for TestContractResult
type TestContractResult struct {
	Address string `json:"address"`
	CodeHash string `json:"codeHash"`
	Returns []Val `json:"returns"`
	GasUsed int32 `json:"gasUsed"`
	Contracts []ContractState `json:"contracts"`
	TxInputs []string `json:"txInputs"`
	TxOutputs []Output `json:"txOutputs"`
	Events []ContractEventByTxId `json:"events"`
	DebugMessages []DebugMessage `json:"debugMessages"`
}

// NewTestContractResult instantiates a new TestContractResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestContractResult(address string, codeHash string, returns []Val, gasUsed int32, contracts []ContractState, txInputs []string, txOutputs []Output, events []ContractEventByTxId, debugMessages []DebugMessage) *TestContractResult {
	this := TestContractResult{}
	this.Address = address
	this.CodeHash = codeHash
	this.Returns = returns
	this.GasUsed = gasUsed
	this.Contracts = contracts
	this.TxInputs = txInputs
	this.TxOutputs = txOutputs
	this.Events = events
	this.DebugMessages = debugMessages
	return &this
}

// NewTestContractResultWithDefaults instantiates a new TestContractResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestContractResultWithDefaults() *TestContractResult {
	this := TestContractResult{}
	return &this
}

// GetAddress returns the Address field value
func (o *TestContractResult) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *TestContractResult) SetAddress(v string) {
	o.Address = v
}

// GetCodeHash returns the CodeHash field value
func (o *TestContractResult) GetCodeHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetCodeHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeHash, true
}

// SetCodeHash sets field value
func (o *TestContractResult) SetCodeHash(v string) {
	o.CodeHash = v
}

// GetReturns returns the Returns field value
func (o *TestContractResult) GetReturns() []Val {
	if o == nil {
		var ret []Val
		return ret
	}

	return o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetReturnsOk() ([]Val, bool) {
	if o == nil {
		return nil, false
	}
	return o.Returns, true
}

// SetReturns sets field value
func (o *TestContractResult) SetReturns(v []Val) {
	o.Returns = v
}

// GetGasUsed returns the GasUsed field value
func (o *TestContractResult) GetGasUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GasUsed
}

// GetGasUsedOk returns a tuple with the GasUsed field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetGasUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasUsed, true
}

// SetGasUsed sets field value
func (o *TestContractResult) SetGasUsed(v int32) {
	o.GasUsed = v
}

// GetContracts returns the Contracts field value
func (o *TestContractResult) GetContracts() []ContractState {
	if o == nil {
		var ret []ContractState
		return ret
	}

	return o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetContractsOk() ([]ContractState, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contracts, true
}

// SetContracts sets field value
func (o *TestContractResult) SetContracts(v []ContractState) {
	o.Contracts = v
}

// GetTxInputs returns the TxInputs field value
func (o *TestContractResult) GetTxInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TxInputs
}

// GetTxInputsOk returns a tuple with the TxInputs field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetTxInputsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxInputs, true
}

// SetTxInputs sets field value
func (o *TestContractResult) SetTxInputs(v []string) {
	o.TxInputs = v
}

// GetTxOutputs returns the TxOutputs field value
func (o *TestContractResult) GetTxOutputs() []Output {
	if o == nil {
		var ret []Output
		return ret
	}

	return o.TxOutputs
}

// GetTxOutputsOk returns a tuple with the TxOutputs field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetTxOutputsOk() ([]Output, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxOutputs, true
}

// SetTxOutputs sets field value
func (o *TestContractResult) SetTxOutputs(v []Output) {
	o.TxOutputs = v
}

// GetEvents returns the Events field value
func (o *TestContractResult) GetEvents() []ContractEventByTxId {
	if o == nil {
		var ret []ContractEventByTxId
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetEventsOk() ([]ContractEventByTxId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *TestContractResult) SetEvents(v []ContractEventByTxId) {
	o.Events = v
}

// GetDebugMessages returns the DebugMessages field value
func (o *TestContractResult) GetDebugMessages() []DebugMessage {
	if o == nil {
		var ret []DebugMessage
		return ret
	}

	return o.DebugMessages
}

// GetDebugMessagesOk returns a tuple with the DebugMessages field value
// and a boolean to check if the value has been set.
func (o *TestContractResult) GetDebugMessagesOk() ([]DebugMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.DebugMessages, true
}

// SetDebugMessages sets field value
func (o *TestContractResult) SetDebugMessages(v []DebugMessage) {
	o.DebugMessages = v
}

func (o TestContractResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestContractResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["codeHash"] = o.CodeHash
	toSerialize["returns"] = o.Returns
	toSerialize["gasUsed"] = o.GasUsed
	toSerialize["contracts"] = o.Contracts
	toSerialize["txInputs"] = o.TxInputs
	toSerialize["txOutputs"] = o.TxOutputs
	toSerialize["events"] = o.Events
	toSerialize["debugMessages"] = o.DebugMessages
	return toSerialize, nil
}

type NullableTestContractResult struct {
	value *TestContractResult
	isSet bool
}

func (v NullableTestContractResult) Get() *TestContractResult {
	return v.value
}

func (v *NullableTestContractResult) Set(val *TestContractResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestContractResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestContractResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestContractResult(val *TestContractResult) *NullableTestContractResult {
	return &NullableTestContractResult{value: val, isSet: true}
}

func (v NullableTestContractResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestContractResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


