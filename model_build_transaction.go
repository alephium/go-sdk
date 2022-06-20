/*
Alephium API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alephium

import (
	"encoding/json"
)

// BuildTransaction struct for BuildTransaction
type BuildTransaction struct {
	FromPublicKey string `json:"fromPublicKey"`
	Destinations []Destination `json:"destinations"`
	Utxos []OutputRef `json:"utxos,omitempty"`
	GasAmount *int32 `json:"gasAmount,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
}

// NewBuildTransaction instantiates a new BuildTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildTransaction(fromPublicKey string, destinations []Destination) *BuildTransaction {
	this := BuildTransaction{}
	this.FromPublicKey = fromPublicKey
	this.Destinations = destinations
	return &this
}

// NewBuildTransactionWithDefaults instantiates a new BuildTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildTransactionWithDefaults() *BuildTransaction {
	this := BuildTransaction{}
	return &this
}

// GetFromPublicKey returns the FromPublicKey field value
func (o *BuildTransaction) GetFromPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromPublicKey
}

// GetFromPublicKeyOk returns a tuple with the FromPublicKey field value
// and a boolean to check if the value has been set.
func (o *BuildTransaction) GetFromPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPublicKey, true
}

// SetFromPublicKey sets field value
func (o *BuildTransaction) SetFromPublicKey(v string) {
	o.FromPublicKey = v
}

// GetDestinations returns the Destinations field value
func (o *BuildTransaction) GetDestinations() []Destination {
	if o == nil {
		var ret []Destination
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *BuildTransaction) GetDestinationsOk() ([]Destination, bool) {
	if o == nil {
		return nil, false
	}
	return o.Destinations, true
}

// SetDestinations sets field value
func (o *BuildTransaction) SetDestinations(v []Destination) {
	o.Destinations = v
}

// GetUtxos returns the Utxos field value if set, zero value otherwise.
func (o *BuildTransaction) GetUtxos() []OutputRef {
	if o == nil || o.Utxos == nil {
		var ret []OutputRef
		return ret
	}
	return o.Utxos
}

// GetUtxosOk returns a tuple with the Utxos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildTransaction) GetUtxosOk() ([]OutputRef, bool) {
	if o == nil || o.Utxos == nil {
		return nil, false
	}
	return o.Utxos, true
}

// HasUtxos returns a boolean if a field has been set.
func (o *BuildTransaction) HasUtxos() bool {
	if o != nil && o.Utxos != nil {
		return true
	}

	return false
}

// SetUtxos gets a reference to the given []OutputRef and assigns it to the Utxos field.
func (o *BuildTransaction) SetUtxos(v []OutputRef) {
	o.Utxos = v
}

// GetGasAmount returns the GasAmount field value if set, zero value otherwise.
func (o *BuildTransaction) GetGasAmount() int32 {
	if o == nil || o.GasAmount == nil {
		var ret int32
		return ret
	}
	return *o.GasAmount
}

// GetGasAmountOk returns a tuple with the GasAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildTransaction) GetGasAmountOk() (*int32, bool) {
	if o == nil || o.GasAmount == nil {
		return nil, false
	}
	return o.GasAmount, true
}

// HasGasAmount returns a boolean if a field has been set.
func (o *BuildTransaction) HasGasAmount() bool {
	if o != nil && o.GasAmount != nil {
		return true
	}

	return false
}

// SetGasAmount gets a reference to the given int32 and assigns it to the GasAmount field.
func (o *BuildTransaction) SetGasAmount(v int32) {
	o.GasAmount = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *BuildTransaction) GetGasPrice() string {
	if o == nil || o.GasPrice == nil {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildTransaction) GetGasPriceOk() (*string, bool) {
	if o == nil || o.GasPrice == nil {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *BuildTransaction) HasGasPrice() bool {
	if o != nil && o.GasPrice != nil {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *BuildTransaction) SetGasPrice(v string) {
	o.GasPrice = &v
}

func (o BuildTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromPublicKey"] = o.FromPublicKey
	}
	if true {
		toSerialize["destinations"] = o.Destinations
	}
	if o.Utxos != nil {
		toSerialize["utxos"] = o.Utxos
	}
	if o.GasAmount != nil {
		toSerialize["gasAmount"] = o.GasAmount
	}
	if o.GasPrice != nil {
		toSerialize["gasPrice"] = o.GasPrice
	}
	return json.Marshal(toSerialize)
}

type NullableBuildTransaction struct {
	value *BuildTransaction
	isSet bool
}

func (v NullableBuildTransaction) Get() *BuildTransaction {
	return v.value
}

func (v *NullableBuildTransaction) Set(val *BuildTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildTransaction(val *BuildTransaction) *NullableBuildTransaction {
	return &NullableBuildTransaction{value: val, isSet: true}
}

func (v NullableBuildTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


