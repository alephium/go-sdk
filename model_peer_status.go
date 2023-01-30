/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
	"fmt"
)

// PeerStatus - struct for PeerStatus
type PeerStatus struct {
	Banned *Banned
	Penalty *Penalty
}

// BannedAsPeerStatus is a convenience function that returns Banned wrapped in PeerStatus
func BannedAsPeerStatus(v *Banned) PeerStatus {
	return PeerStatus{
		Banned: v,
	}
}

// PenaltyAsPeerStatus is a convenience function that returns Penalty wrapped in PeerStatus
func PenaltyAsPeerStatus(v *Penalty) PeerStatus {
	return PeerStatus{
		Penalty: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PeerStatus) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Banned
	err = newStrictDecoder(data).Decode(&dst.Banned)
	if err == nil {
		jsonBanned, _ := json.Marshal(dst.Banned)
		if string(jsonBanned) == "{}" { // empty struct
			dst.Banned = nil
		} else {
			match++
		}
	} else {
		dst.Banned = nil
	}

	// try to unmarshal data into Penalty
	err = newStrictDecoder(data).Decode(&dst.Penalty)
	if err == nil {
		jsonPenalty, _ := json.Marshal(dst.Penalty)
		if string(jsonPenalty) == "{}" { // empty struct
			dst.Penalty = nil
		} else {
			match++
		}
	} else {
		dst.Penalty = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Banned = nil
		dst.Penalty = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PeerStatus)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PeerStatus)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PeerStatus) MarshalJSON() ([]byte, error) {
	if src.Banned != nil {
		return json.Marshal(&src.Banned)
	}

	if src.Penalty != nil {
		return json.Marshal(&src.Penalty)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PeerStatus) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Banned != nil {
		return obj.Banned
	}

	if obj.Penalty != nil {
		return obj.Penalty
	}

	// all schemas are nil
	return nil
}

type NullablePeerStatus struct {
	value *PeerStatus
	isSet bool
}

func (v NullablePeerStatus) Get() *PeerStatus {
	return v.value
}

func (v *NullablePeerStatus) Set(val *PeerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePeerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePeerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeerStatus(val *PeerStatus) *NullablePeerStatus {
	return &NullablePeerStatus{value: val, isSet: true}
}

func (v NullablePeerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


