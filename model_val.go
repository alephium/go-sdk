/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
	"fmt"
)

// Val - struct for Val
type Val struct {
	ValAddress *ValAddress
	ValArray *ValArray
	ValBool *ValBool
	ValByteVec *ValByteVec
	ValI256 *ValI256
	ValU256 *ValU256
}

// ValAddressAsVal is a convenience function that returns ValAddress wrapped in Val
func ValAddressAsVal(v *ValAddress) Val {
	return Val{
		ValAddress: v,
	}
}

// ValArrayAsVal is a convenience function that returns ValArray wrapped in Val
func ValArrayAsVal(v *ValArray) Val {
	return Val{
		ValArray: v,
	}
}

// ValBoolAsVal is a convenience function that returns ValBool wrapped in Val
func ValBoolAsVal(v *ValBool) Val {
	return Val{
		ValBool: v,
	}
}

// ValByteVecAsVal is a convenience function that returns ValByteVec wrapped in Val
func ValByteVecAsVal(v *ValByteVec) Val {
	return Val{
		ValByteVec: v,
	}
}

// ValI256AsVal is a convenience function that returns ValI256 wrapped in Val
func ValI256AsVal(v *ValI256) Val {
	return Val{
		ValI256: v,
	}
}

// ValU256AsVal is a convenience function that returns ValU256 wrapped in Val
func ValU256AsVal(v *ValU256) Val {
	return Val{
		ValU256: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Val) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ValAddress
	err = newStrictDecoder(data).Decode(&dst.ValAddress)
	if err == nil {
		jsonValAddress, _ := json.Marshal(dst.ValAddress)
		if string(jsonValAddress) == "{}" { // empty struct
			dst.ValAddress = nil
		} else {
			match++
		}
	} else {
		dst.ValAddress = nil
	}

	// try to unmarshal data into ValArray
	err = newStrictDecoder(data).Decode(&dst.ValArray)
	if err == nil {
		jsonValArray, _ := json.Marshal(dst.ValArray)
		if string(jsonValArray) == "{}" { // empty struct
			dst.ValArray = nil
		} else {
			match++
		}
	} else {
		dst.ValArray = nil
	}

	// try to unmarshal data into ValBool
	err = newStrictDecoder(data).Decode(&dst.ValBool)
	if err == nil {
		jsonValBool, _ := json.Marshal(dst.ValBool)
		if string(jsonValBool) == "{}" { // empty struct
			dst.ValBool = nil
		} else {
			match++
		}
	} else {
		dst.ValBool = nil
	}

	// try to unmarshal data into ValByteVec
	err = newStrictDecoder(data).Decode(&dst.ValByteVec)
	if err == nil {
		jsonValByteVec, _ := json.Marshal(dst.ValByteVec)
		if string(jsonValByteVec) == "{}" { // empty struct
			dst.ValByteVec = nil
		} else {
			match++
		}
	} else {
		dst.ValByteVec = nil
	}

	// try to unmarshal data into ValI256
	err = newStrictDecoder(data).Decode(&dst.ValI256)
	if err == nil {
		jsonValI256, _ := json.Marshal(dst.ValI256)
		if string(jsonValI256) == "{}" { // empty struct
			dst.ValI256 = nil
		} else {
			match++
		}
	} else {
		dst.ValI256 = nil
	}

	// try to unmarshal data into ValU256
	err = newStrictDecoder(data).Decode(&dst.ValU256)
	if err == nil {
		jsonValU256, _ := json.Marshal(dst.ValU256)
		if string(jsonValU256) == "{}" { // empty struct
			dst.ValU256 = nil
		} else {
			match++
		}
	} else {
		dst.ValU256 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ValAddress = nil
		dst.ValArray = nil
		dst.ValBool = nil
		dst.ValByteVec = nil
		dst.ValI256 = nil
		dst.ValU256 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Val)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Val)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Val) MarshalJSON() ([]byte, error) {
	if src.ValAddress != nil {
		return json.Marshal(&src.ValAddress)
	}

	if src.ValArray != nil {
		return json.Marshal(&src.ValArray)
	}

	if src.ValBool != nil {
		return json.Marshal(&src.ValBool)
	}

	if src.ValByteVec != nil {
		return json.Marshal(&src.ValByteVec)
	}

	if src.ValI256 != nil {
		return json.Marshal(&src.ValI256)
	}

	if src.ValU256 != nil {
		return json.Marshal(&src.ValU256)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Val) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ValAddress != nil {
		return obj.ValAddress
	}

	if obj.ValArray != nil {
		return obj.ValArray
	}

	if obj.ValBool != nil {
		return obj.ValBool
	}

	if obj.ValByteVec != nil {
		return obj.ValByteVec
	}

	if obj.ValI256 != nil {
		return obj.ValI256
	}

	if obj.ValU256 != nil {
		return obj.ValU256
	}

	// all schemas are nil
	return nil
}

type NullableVal struct {
	value *Val
	isSet bool
}

func (v NullableVal) Get() *Val {
	return v.value
}

func (v *NullableVal) Set(val *Val) {
	v.value = val
	v.isSet = true
}

func (v NullableVal) IsSet() bool {
	return v.isSet
}

func (v *NullableVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVal(val *Val) *NullableVal {
	return &NullableVal{value: val, isSet: true}
}

func (v NullableVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


