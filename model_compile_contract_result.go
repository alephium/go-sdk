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

// checks if the CompileContractResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompileContractResult{}

// CompileContractResult struct for CompileContractResult
type CompileContractResult struct {
	Version string `json:"version"`
	Name string `json:"name"`
	Bytecode string `json:"bytecode"`
	BytecodeDebugPatch string `json:"bytecodeDebugPatch"`
	CodeHash string `json:"codeHash"`
	CodeHashDebug string `json:"codeHashDebug"`
	Fields FieldsSig `json:"fields"`
	Functions []FunctionSig `json:"functions"`
	Constants []Constant `json:"constants"`
	Enums []Enum `json:"enums"`
	Events []EventSig `json:"events"`
	Warnings []string `json:"warnings"`
	StdInterfaceId *string `json:"stdInterfaceId,omitempty"`
}

// NewCompileContractResult instantiates a new CompileContractResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompileContractResult(version string, name string, bytecode string, bytecodeDebugPatch string, codeHash string, codeHashDebug string, fields FieldsSig, functions []FunctionSig, constants []Constant, enums []Enum, events []EventSig, warnings []string) *CompileContractResult {
	this := CompileContractResult{}
	this.Version = version
	this.Name = name
	this.Bytecode = bytecode
	this.BytecodeDebugPatch = bytecodeDebugPatch
	this.CodeHash = codeHash
	this.CodeHashDebug = codeHashDebug
	this.Fields = fields
	this.Functions = functions
	this.Constants = constants
	this.Enums = enums
	this.Events = events
	this.Warnings = warnings
	return &this
}

// NewCompileContractResultWithDefaults instantiates a new CompileContractResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompileContractResultWithDefaults() *CompileContractResult {
	this := CompileContractResult{}
	return &this
}

// GetVersion returns the Version field value
func (o *CompileContractResult) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CompileContractResult) SetVersion(v string) {
	o.Version = v
}

// GetName returns the Name field value
func (o *CompileContractResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompileContractResult) SetName(v string) {
	o.Name = v
}

// GetBytecode returns the Bytecode field value
func (o *CompileContractResult) GetBytecode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetBytecodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytecode, true
}

// SetBytecode sets field value
func (o *CompileContractResult) SetBytecode(v string) {
	o.Bytecode = v
}

// GetBytecodeDebugPatch returns the BytecodeDebugPatch field value
func (o *CompileContractResult) GetBytecodeDebugPatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BytecodeDebugPatch
}

// GetBytecodeDebugPatchOk returns a tuple with the BytecodeDebugPatch field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetBytecodeDebugPatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BytecodeDebugPatch, true
}

// SetBytecodeDebugPatch sets field value
func (o *CompileContractResult) SetBytecodeDebugPatch(v string) {
	o.BytecodeDebugPatch = v
}

// GetCodeHash returns the CodeHash field value
func (o *CompileContractResult) GetCodeHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetCodeHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeHash, true
}

// SetCodeHash sets field value
func (o *CompileContractResult) SetCodeHash(v string) {
	o.CodeHash = v
}

// GetCodeHashDebug returns the CodeHashDebug field value
func (o *CompileContractResult) GetCodeHashDebug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeHashDebug
}

// GetCodeHashDebugOk returns a tuple with the CodeHashDebug field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetCodeHashDebugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeHashDebug, true
}

// SetCodeHashDebug sets field value
func (o *CompileContractResult) SetCodeHashDebug(v string) {
	o.CodeHashDebug = v
}

// GetFields returns the Fields field value
func (o *CompileContractResult) GetFields() FieldsSig {
	if o == nil {
		var ret FieldsSig
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetFieldsOk() (*FieldsSig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *CompileContractResult) SetFields(v FieldsSig) {
	o.Fields = v
}

// GetFunctions returns the Functions field value
func (o *CompileContractResult) GetFunctions() []FunctionSig {
	if o == nil {
		var ret []FunctionSig
		return ret
	}

	return o.Functions
}

// GetFunctionsOk returns a tuple with the Functions field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetFunctionsOk() ([]FunctionSig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Functions, true
}

// SetFunctions sets field value
func (o *CompileContractResult) SetFunctions(v []FunctionSig) {
	o.Functions = v
}

// GetConstants returns the Constants field value
func (o *CompileContractResult) GetConstants() []Constant {
	if o == nil {
		var ret []Constant
		return ret
	}

	return o.Constants
}

// GetConstantsOk returns a tuple with the Constants field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetConstantsOk() ([]Constant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constants, true
}

// SetConstants sets field value
func (o *CompileContractResult) SetConstants(v []Constant) {
	o.Constants = v
}

// GetEnums returns the Enums field value
func (o *CompileContractResult) GetEnums() []Enum {
	if o == nil {
		var ret []Enum
		return ret
	}

	return o.Enums
}

// GetEnumsOk returns a tuple with the Enums field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetEnumsOk() ([]Enum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enums, true
}

// SetEnums sets field value
func (o *CompileContractResult) SetEnums(v []Enum) {
	o.Enums = v
}

// GetEvents returns the Events field value
func (o *CompileContractResult) GetEvents() []EventSig {
	if o == nil {
		var ret []EventSig
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetEventsOk() ([]EventSig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *CompileContractResult) SetEvents(v []EventSig) {
	o.Events = v
}

// GetWarnings returns the Warnings field value
func (o *CompileContractResult) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetWarningsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *CompileContractResult) SetWarnings(v []string) {
	o.Warnings = v
}

// GetStdInterfaceId returns the StdInterfaceId field value if set, zero value otherwise.
func (o *CompileContractResult) GetStdInterfaceId() string {
	if o == nil || isNil(o.StdInterfaceId) {
		var ret string
		return ret
	}
	return *o.StdInterfaceId
}

// GetStdInterfaceIdOk returns a tuple with the StdInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompileContractResult) GetStdInterfaceIdOk() (*string, bool) {
	if o == nil || isNil(o.StdInterfaceId) {
		return nil, false
	}
	return o.StdInterfaceId, true
}

// HasStdInterfaceId returns a boolean if a field has been set.
func (o *CompileContractResult) HasStdInterfaceId() bool {
	if o != nil && !isNil(o.StdInterfaceId) {
		return true
	}

	return false
}

// SetStdInterfaceId gets a reference to the given string and assigns it to the StdInterfaceId field.
func (o *CompileContractResult) SetStdInterfaceId(v string) {
	o.StdInterfaceId = &v
}

func (o CompileContractResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompileContractResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["name"] = o.Name
	toSerialize["bytecode"] = o.Bytecode
	toSerialize["bytecodeDebugPatch"] = o.BytecodeDebugPatch
	toSerialize["codeHash"] = o.CodeHash
	toSerialize["codeHashDebug"] = o.CodeHashDebug
	toSerialize["fields"] = o.Fields
	toSerialize["functions"] = o.Functions
	toSerialize["constants"] = o.Constants
	toSerialize["enums"] = o.Enums
	toSerialize["events"] = o.Events
	toSerialize["warnings"] = o.Warnings
	if !isNil(o.StdInterfaceId) {
		toSerialize["stdInterfaceId"] = o.StdInterfaceId
	}
	return toSerialize, nil
}

type NullableCompileContractResult struct {
	value *CompileContractResult
	isSet bool
}

func (v NullableCompileContractResult) Get() *CompileContractResult {
	return v.value
}

func (v *NullableCompileContractResult) Set(val *CompileContractResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCompileContractResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCompileContractResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompileContractResult(val *CompileContractResult) *NullableCompileContractResult {
	return &NullableCompileContractResult{value: val, isSet: true}
}

func (v NullableCompileContractResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompileContractResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


