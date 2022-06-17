/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
	"fmt"
)

// DiscoveryAction - struct for DiscoveryAction
type DiscoveryAction struct {
	Reachable *Reachable
	Unreachable *Unreachable
}

// ReachableAsDiscoveryAction is a convenience function that returns Reachable wrapped in DiscoveryAction
func ReachableAsDiscoveryAction(v *Reachable) DiscoveryAction {
	return DiscoveryAction{
		Reachable: v,
	}
}

// UnreachableAsDiscoveryAction is a convenience function that returns Unreachable wrapped in DiscoveryAction
func UnreachableAsDiscoveryAction(v *Unreachable) DiscoveryAction {
	return DiscoveryAction{
		Unreachable: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DiscoveryAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Reachable
	err = newStrictDecoder(data).Decode(&dst.Reachable)
	if err == nil {
		if dst.Reachable == nil || dst.Reachable.Type != "Reachable" { // not the right type
			dst.Reachable = nil
		} else {
			match++
		}
	} else {
		dst.Reachable = nil
	}

	// try to unmarshal data into Unreachable
	err = newStrictDecoder(data).Decode(&dst.Unreachable)
	if err == nil {
		if dst.Unreachable == nil || dst.Unreachable.Type != "Unreachable" { // not the right type
			dst.Unreachable = nil
		} else {
			match++
		}
	} else {
		dst.Unreachable = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Reachable = nil
		dst.Unreachable = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DiscoveryAction)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DiscoveryAction)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DiscoveryAction) MarshalJSON() ([]byte, error) {
	if src.Reachable != nil {
		return json.Marshal(&src.Reachable)
	}

	if src.Unreachable != nil {
		return json.Marshal(&src.Unreachable)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DiscoveryAction) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Reachable != nil {
		return obj.Reachable
	}

	if obj.Unreachable != nil {
		return obj.Unreachable
	}

	// all schemas are nil
	return nil
}

type NullableDiscoveryAction struct {
	value *DiscoveryAction
	isSet bool
}

func (v NullableDiscoveryAction) Get() *DiscoveryAction {
	return v.value
}

func (v *NullableDiscoveryAction) Set(val *DiscoveryAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAction(val *DiscoveryAction) *NullableDiscoveryAction {
	return &NullableDiscoveryAction{value: val, isSet: true}
}

func (v NullableDiscoveryAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


