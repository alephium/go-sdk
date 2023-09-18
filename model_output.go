/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.5.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
	"fmt"
)

// Output - struct for Output
type Output struct {
	AssetOutput *AssetOutput
	ContractOutput *ContractOutput
}

// AssetOutputAsOutput is a convenience function that returns AssetOutput wrapped in Output
func AssetOutputAsOutput(v *AssetOutput) Output {
	return Output{
		AssetOutput: v,
	}
}

// ContractOutputAsOutput is a convenience function that returns ContractOutput wrapped in Output
func ContractOutputAsOutput(v *ContractOutput) Output {
	return Output{
		ContractOutput: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Output) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AssetOutput
	err = newStrictDecoder(data).Decode(&dst.AssetOutput)
	if err == nil {
		jsonAssetOutput, _ := json.Marshal(dst.AssetOutput)
		if string(jsonAssetOutput) == "{}" { // empty struct
			dst.AssetOutput = nil
		} else {
			match++
		}
	} else {
		dst.AssetOutput = nil
	}

	// try to unmarshal data into ContractOutput
	err = newStrictDecoder(data).Decode(&dst.ContractOutput)
	if err == nil {
		jsonContractOutput, _ := json.Marshal(dst.ContractOutput)
		if string(jsonContractOutput) == "{}" { // empty struct
			dst.ContractOutput = nil
		} else {
			match++
		}
	} else {
		dst.ContractOutput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssetOutput = nil
		dst.ContractOutput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Output)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Output)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Output) MarshalJSON() ([]byte, error) {
	if src.AssetOutput != nil {
		return json.Marshal(&src.AssetOutput)
	}

	if src.ContractOutput != nil {
		return json.Marshal(&src.ContractOutput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Output) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AssetOutput != nil {
		return obj.AssetOutput
	}

	if obj.ContractOutput != nil {
		return obj.ContractOutput
	}

	// all schemas are nil
	return nil
}

type NullableOutput struct {
	value *Output
	isSet bool
}

func (v NullableOutput) Get() *Output {
	return v.value
}

func (v *NullableOutput) Set(val *Output) {
	v.value = val
	v.isSet = true
}

func (v NullableOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutput(val *Output) *NullableOutput {
	return &NullableOutput{value: val, isSet: true}
}

func (v NullableOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


