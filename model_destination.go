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

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination struct for Destination
type Destination struct {
	Address string `json:"address"`
	AttoAlphAmount string `json:"attoAlphAmount"`
	Tokens []Token `json:"tokens,omitempty"`
	LockTime *int64 `json:"lockTime,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination(address string, attoAlphAmount string) *Destination {
	this := Destination{}
	this.Address = address
	this.AttoAlphAmount = attoAlphAmount
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	return &this
}

// GetAddress returns the Address field value
func (o *Destination) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Destination) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Destination) SetAddress(v string) {
	o.Address = v
}

// GetAttoAlphAmount returns the AttoAlphAmount field value
func (o *Destination) GetAttoAlphAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttoAlphAmount
}

// GetAttoAlphAmountOk returns a tuple with the AttoAlphAmount field value
// and a boolean to check if the value has been set.
func (o *Destination) GetAttoAlphAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttoAlphAmount, true
}

// SetAttoAlphAmount sets field value
func (o *Destination) SetAttoAlphAmount(v string) {
	o.AttoAlphAmount = v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *Destination) GetTokens() []Token {
	if o == nil || isNil(o.Tokens) {
		var ret []Token
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetTokensOk() ([]Token, bool) {
	if o == nil || isNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *Destination) HasTokens() bool {
	if o != nil && !isNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []Token and assigns it to the Tokens field.
func (o *Destination) SetTokens(v []Token) {
	o.Tokens = v
}

// GetLockTime returns the LockTime field value if set, zero value otherwise.
func (o *Destination) GetLockTime() int64 {
	if o == nil || isNil(o.LockTime) {
		var ret int64
		return ret
	}
	return *o.LockTime
}

// GetLockTimeOk returns a tuple with the LockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetLockTimeOk() (*int64, bool) {
	if o == nil || isNil(o.LockTime) {
		return nil, false
	}
	return o.LockTime, true
}

// HasLockTime returns a boolean if a field has been set.
func (o *Destination) HasLockTime() bool {
	if o != nil && !isNil(o.LockTime) {
		return true
	}

	return false
}

// SetLockTime gets a reference to the given int64 and assigns it to the LockTime field.
func (o *Destination) SetLockTime(v int64) {
	o.LockTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Destination) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Destination) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Destination) SetMessage(v string) {
	o.Message = &v
}

func (o Destination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["attoAlphAmount"] = o.AttoAlphAmount
	if !isNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !isNil(o.LockTime) {
		toSerialize["lockTime"] = o.LockTime
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


