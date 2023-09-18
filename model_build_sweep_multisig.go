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

// checks if the BuildSweepMultisig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildSweepMultisig{}

// BuildSweepMultisig struct for BuildSweepMultisig
type BuildSweepMultisig struct {
	FromAddress string `json:"fromAddress"`
	FromPublicKeys []string `json:"fromPublicKeys"`
	ToAddress string `json:"toAddress"`
	MaxAttoAlphPerUTXO *string `json:"maxAttoAlphPerUTXO,omitempty"`
	LockTime *int64 `json:"lockTime,omitempty"`
	GasAmount *int32 `json:"gasAmount,omitempty"`
	GasPrice *string `json:"gasPrice,omitempty"`
	UtxosLimit *int32 `json:"utxosLimit,omitempty"`
	TargetBlockHash *string `json:"targetBlockHash,omitempty"`
}

// NewBuildSweepMultisig instantiates a new BuildSweepMultisig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildSweepMultisig(fromAddress string, fromPublicKeys []string, toAddress string) *BuildSweepMultisig {
	this := BuildSweepMultisig{}
	this.FromAddress = fromAddress
	this.FromPublicKeys = fromPublicKeys
	this.ToAddress = toAddress
	return &this
}

// NewBuildSweepMultisigWithDefaults instantiates a new BuildSweepMultisig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildSweepMultisigWithDefaults() *BuildSweepMultisig {
	this := BuildSweepMultisig{}
	return &this
}

// GetFromAddress returns the FromAddress field value
func (o *BuildSweepMultisig) GetFromAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *BuildSweepMultisig) SetFromAddress(v string) {
	o.FromAddress = v
}

// GetFromPublicKeys returns the FromPublicKeys field value
func (o *BuildSweepMultisig) GetFromPublicKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FromPublicKeys
}

// GetFromPublicKeysOk returns a tuple with the FromPublicKeys field value
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetFromPublicKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromPublicKeys, true
}

// SetFromPublicKeys sets field value
func (o *BuildSweepMultisig) SetFromPublicKeys(v []string) {
	o.FromPublicKeys = v
}

// GetToAddress returns the ToAddress field value
func (o *BuildSweepMultisig) GetToAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetToAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *BuildSweepMultisig) SetToAddress(v string) {
	o.ToAddress = v
}

// GetMaxAttoAlphPerUTXO returns the MaxAttoAlphPerUTXO field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetMaxAttoAlphPerUTXO() string {
	if o == nil || isNil(o.MaxAttoAlphPerUTXO) {
		var ret string
		return ret
	}
	return *o.MaxAttoAlphPerUTXO
}

// GetMaxAttoAlphPerUTXOOk returns a tuple with the MaxAttoAlphPerUTXO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetMaxAttoAlphPerUTXOOk() (*string, bool) {
	if o == nil || isNil(o.MaxAttoAlphPerUTXO) {
		return nil, false
	}
	return o.MaxAttoAlphPerUTXO, true
}

// HasMaxAttoAlphPerUTXO returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasMaxAttoAlphPerUTXO() bool {
	if o != nil && !isNil(o.MaxAttoAlphPerUTXO) {
		return true
	}

	return false
}

// SetMaxAttoAlphPerUTXO gets a reference to the given string and assigns it to the MaxAttoAlphPerUTXO field.
func (o *BuildSweepMultisig) SetMaxAttoAlphPerUTXO(v string) {
	o.MaxAttoAlphPerUTXO = &v
}

// GetLockTime returns the LockTime field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetLockTime() int64 {
	if o == nil || isNil(o.LockTime) {
		var ret int64
		return ret
	}
	return *o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetLockTimeOk() (*int64, bool) {
	if o == nil || isNil(o.LockTime) {
		return nil, false
	}
	return o.LockTime, true
}

// HasLockTime returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasLockTime() bool {
	if o != nil && !isNil(o.LockTime) {
		return true
	}

	return false
}

// SetLockTime gets a reference to the given int64 and assigns it to the LockTime field.
func (o *BuildSweepMultisig) SetLockTime(v int64) {
	o.LockTime = &v
}

// GetGasAmount returns the GasAmount field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetGasAmount() int32 {
	if o == nil || isNil(o.GasAmount) {
		var ret int32
		return ret
	}
	return *o.GasAmount
}

// GetGasAmountOk returns a tuple with the GasAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetGasAmountOk() (*int32, bool) {
	if o == nil || isNil(o.GasAmount) {
		return nil, false
	}
	return o.GasAmount, true
}

// HasGasAmount returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasGasAmount() bool {
	if o != nil && !isNil(o.GasAmount) {
		return true
	}

	return false
}

// SetGasAmount gets a reference to the given int32 and assigns it to the GasAmount field.
func (o *BuildSweepMultisig) SetGasAmount(v int32) {
	o.GasAmount = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetGasPrice() string {
	if o == nil || isNil(o.GasPrice) {
		var ret string
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetGasPriceOk() (*string, bool) {
	if o == nil || isNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasGasPrice() bool {
	if o != nil && !isNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given string and assigns it to the GasPrice field.
func (o *BuildSweepMultisig) SetGasPrice(v string) {
	o.GasPrice = &v
}

// GetUtxosLimit returns the UtxosLimit field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetUtxosLimit() int32 {
	if o == nil || isNil(o.UtxosLimit) {
		var ret int32
		return ret
	}
	return *o.UtxosLimit
}

// GetUtxosLimitOk returns a tuple with the UtxosLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetUtxosLimitOk() (*int32, bool) {
	if o == nil || isNil(o.UtxosLimit) {
		return nil, false
	}
	return o.UtxosLimit, true
}

// HasUtxosLimit returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasUtxosLimit() bool {
	if o != nil && !isNil(o.UtxosLimit) {
		return true
	}

	return false
}

// SetUtxosLimit gets a reference to the given int32 and assigns it to the UtxosLimit field.
func (o *BuildSweepMultisig) SetUtxosLimit(v int32) {
	o.UtxosLimit = &v
}

// GetTargetBlockHash returns the TargetBlockHash field value if set, zero value otherwise.
func (o *BuildSweepMultisig) GetTargetBlockHash() string {
	if o == nil || isNil(o.TargetBlockHash) {
		var ret string
		return ret
	}
	return *o.TargetBlockHash
}

// GetTargetBlockHashOk returns a tuple with the TargetBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSweepMultisig) GetTargetBlockHashOk() (*string, bool) {
	if o == nil || isNil(o.TargetBlockHash) {
		return nil, false
	}
	return o.TargetBlockHash, true
}

// HasTargetBlockHash returns a boolean if a field has been set.
func (o *BuildSweepMultisig) HasTargetBlockHash() bool {
	if o != nil && !isNil(o.TargetBlockHash) {
		return true
	}

	return false
}

// SetTargetBlockHash gets a reference to the given string and assigns it to the TargetBlockHash field.
func (o *BuildSweepMultisig) SetTargetBlockHash(v string) {
	o.TargetBlockHash = &v
}

func (o BuildSweepMultisig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildSweepMultisig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["fromPublicKeys"] = o.FromPublicKeys
	toSerialize["toAddress"] = o.ToAddress
	if !isNil(o.MaxAttoAlphPerUTXO) {
		toSerialize["maxAttoAlphPerUTXO"] = o.MaxAttoAlphPerUTXO
	}
	if !isNil(o.LockTime) {
		toSerialize["lockTime"] = o.LockTime
	}
	if !isNil(o.GasAmount) {
		toSerialize["gasAmount"] = o.GasAmount
	}
	if !isNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !isNil(o.UtxosLimit) {
		toSerialize["utxosLimit"] = o.UtxosLimit
	}
	if !isNil(o.TargetBlockHash) {
		toSerialize["targetBlockHash"] = o.TargetBlockHash
	}
	return toSerialize, nil
}

type NullableBuildSweepMultisig struct {
	value *BuildSweepMultisig
	isSet bool
}

func (v NullableBuildSweepMultisig) Get() *BuildSweepMultisig {
	return v.value
}

func (v *NullableBuildSweepMultisig) Set(val *BuildSweepMultisig) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildSweepMultisig) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildSweepMultisig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildSweepMultisig(val *BuildSweepMultisig) *NullableBuildSweepMultisig {
	return &NullableBuildSweepMultisig{value: val, isSet: true}
}

func (v NullableBuildSweepMultisig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildSweepMultisig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


