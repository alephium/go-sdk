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

// BuildMultisig struct for BuildMultisig
type BuildMultisig struct {
	FromAddress string `json:"fromAddress"`
	FromPublicKeys []string `json:"fromPublicKeys"`
	Destinations []Destination `json:"destinations"`
	Gas *int32 `json:"gas,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
}

// NewBuildMultisig instantiates a new BuildMultisig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildMultisig(fromAddress string, fromPublicKeys []string, destinations []Destination) *BuildMultisig {
	this := BuildMultisig{}
	this.FromAddress = fromAddress
	this.FromPublicKeys = fromPublicKeys
	this.Destinations = destinations
	return &this
}

// NewBuildMultisigWithDefaults instantiates a new BuildMultisig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildMultisigWithDefaults() *BuildMultisig {
	this := BuildMultisig{}
	return &this
}

// GetFromAddress returns the FromAddress field value
func (o *BuildMultisig) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *BuildMultisig) GetFromAddressOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *BuildMultisig) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetFromPublicKeys returns the FromPublicKeys field value
func (o *BuildMultisig) GetFromPublicKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FromPublicKeys
}

// GetFromPublicKeysOk returns a tuple with the FromPublicKeys field value
// and a boolean to check if the value has been set.
func (o *BuildMultisig) GetFromPublicKeysOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FromPublicKeys, true
}

// SetFromPublicKeys sets field value
func (o *BuildMultisig) SetFromPublicKeys(v []string) {
	o.FromPublicKeys = v
}

// GetDestinations returns the Destinations field value
func (o *BuildMultisig) GetDestinations() []Destination {
	if o == nil {
		var ret []Destination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *BuildMultisig) GetDestinationsOk() ([]Destination, bool) {
	if o == nil {
    return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *BuildMultisig) SetDestinations(v []Destination) {
	o.Destinations = v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *BuildMultisig) GetGas() int32 {
	if o == nil || isNil(o.Gas) {
		var ret int32
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildMultisig) GetGasOk() (*int32, bool) {
	if o == nil || isNil(o.Gas) {
    return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *BuildMultisig) HasGas() bool {
	if o != nil && !isNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given int32 and assigns it to the Gas field.
func (o *BuildMultisig) SetGas(v int32) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *BuildMultisig) GetGasPrice() string {
	if o == nil || isNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildMultisig) GetGasPriceOk() (*string, bool) {
	if o == nil || isNil(o.GasPrice) {
    return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *BuildMultisig) HasGasPrice() bool {
	if o != nil && !isNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *BuildMultisig) SetGasPrice(v string) {
	o.GasPrice = &v
}

func (o BuildMultisig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if true {
		toSerialize["fromPublicKeys"] = o.FromPublicKeys
	}
	if true {
		toSerialize["destinations"] = o.Destinations
	}
	if !isNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !isNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	return json.Marshal(toSerialize)
}

type NullableBuildMultisig struct {
	value *BuildMultisig
	isSet bool
}

func (v NullableBuildMultisig) Get() *BuildMultisig {
	return v.value
}

func (v *NullableBuildMultisig) Set(val *BuildMultisig) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildMultisig) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildMultisig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildMultisig(val *BuildMultisig) *NullableBuildMultisig {
	return &NullableBuildMultisig{value: val, isSet: true}
}

func (v NullableBuildMultisig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildMultisig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


