/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.6.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// TestInputAsset struct for TestInputAsset
type TestInputAsset struct {
	Address string `json:"address"`
	Asset AssetState `json:"asset"`
}

// NewTestInputAsset instantiates a new TestInputAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInputAsset(address string, asset AssetState) *TestInputAsset {
	this := TestInputAsset{}
	this.Address = address
	this.Asset = asset
	return &this
}

// NewTestInputAssetWithDefaults instantiates a new TestInputAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInputAssetWithDefaults() *TestInputAsset {
	this := TestInputAsset{}
	return &this
}

// GetAddress returns the Address field value
func (o *TestInputAsset) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *TestInputAsset) GetAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *TestInputAsset) SetAddress(v string) {
	o.Address = v
}

// GetAsset returns the Asset field value
func (o *TestInputAsset) GetAsset() AssetState {
	if o == nil {
		var ret AssetState
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *TestInputAsset) GetAssetOk() (*AssetState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *TestInputAsset) SetAsset(v AssetState) {
	o.Asset = v
}

func (o TestInputAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullableTestInputAsset struct {
	value *TestInputAsset
	isSet bool
}

func (v NullableTestInputAsset) Get() *TestInputAsset {
	return v.value
}

func (v *NullableTestInputAsset) Set(val *TestInputAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInputAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInputAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInputAsset(val *TestInputAsset) *NullableTestInputAsset {
	return &NullableTestInputAsset{value: val, isSet: true}
}

func (v NullableTestInputAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInputAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


