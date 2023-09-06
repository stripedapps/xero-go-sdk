/*
Xero Accounting API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.40.0
Contact: api@xero.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Accounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Accounts{}

// Accounts struct for Accounts
type Accounts struct {
	Accounts []Account `json:"Accounts,omitempty"`
}

// NewAccounts instantiates a new Accounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccounts() *Accounts {
	this := Accounts{}
	return &this
}

// NewAccountsWithDefaults instantiates a new Accounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsWithDefaults() *Accounts {
	this := Accounts{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Accounts) GetAccounts() []Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accounts) GetAccountsOk() ([]Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Accounts) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Account and assigns it to the Accounts field.
func (o *Accounts) SetAccounts(v []Account) {
	o.Accounts = v
}

func (o Accounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Accounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["Accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableAccounts struct {
	value *Accounts
	isSet bool
}

func (v NullableAccounts) Get() *Accounts {
	return v.value
}

func (v *NullableAccounts) Set(val *Accounts) {
	v.value = val
	v.isSet = true
}

func (v NullableAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccounts(val *Accounts) *NullableAccounts {
	return &NullableAccounts{value: val, isSet: true}
}

func (v NullableAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


