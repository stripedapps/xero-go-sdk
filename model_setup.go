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

// checks if the Setup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Setup{}

// Setup struct for Setup
type Setup struct {
	ConversionDate *ConversionDate `json:"ConversionDate,omitempty"`
	// Balance supplied for each account that has a value as at the conversion date.
	ConversionBalances []ConversionBalances `json:"ConversionBalances,omitempty"`
	Accounts []Account `json:"Accounts,omitempty"`
}

// NewSetup instantiates a new Setup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetup() *Setup {
	this := Setup{}
	return &this
}

// NewSetupWithDefaults instantiates a new Setup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupWithDefaults() *Setup {
	this := Setup{}
	return &this
}

// GetConversionDate returns the ConversionDate field value if set, zero value otherwise.
func (o *Setup) GetConversionDate() ConversionDate {
	if o == nil || IsNil(o.ConversionDate) {
		var ret ConversionDate
		return ret
	}
	return *o.ConversionDate
}

// GetConversionDateOk returns a tuple with the ConversionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Setup) GetConversionDateOk() (*ConversionDate, bool) {
	if o == nil || IsNil(o.ConversionDate) {
		return nil, false
	}
	return o.ConversionDate, true
}

// HasConversionDate returns a boolean if a field has been set.
func (o *Setup) HasConversionDate() bool {
	if o != nil && !IsNil(o.ConversionDate) {
		return true
	}

	return false
}

// SetConversionDate gets a reference to the given ConversionDate and assigns it to the ConversionDate field.
func (o *Setup) SetConversionDate(v ConversionDate) {
	o.ConversionDate = &v
}

// GetConversionBalances returns the ConversionBalances field value if set, zero value otherwise.
func (o *Setup) GetConversionBalances() []ConversionBalances {
	if o == nil || IsNil(o.ConversionBalances) {
		var ret []ConversionBalances
		return ret
	}
	return o.ConversionBalances
}

// GetConversionBalancesOk returns a tuple with the ConversionBalances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Setup) GetConversionBalancesOk() ([]ConversionBalances, bool) {
	if o == nil || IsNil(o.ConversionBalances) {
		return nil, false
	}
	return o.ConversionBalances, true
}

// HasConversionBalances returns a boolean if a field has been set.
func (o *Setup) HasConversionBalances() bool {
	if o != nil && !IsNil(o.ConversionBalances) {
		return true
	}

	return false
}

// SetConversionBalances gets a reference to the given []ConversionBalances and assigns it to the ConversionBalances field.
func (o *Setup) SetConversionBalances(v []ConversionBalances) {
	o.ConversionBalances = v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *Setup) GetAccounts() []Account {
	if o == nil || IsNil(o.Accounts) {
		var ret []Account
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Setup) GetAccountsOk() ([]Account, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *Setup) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []Account and assigns it to the Accounts field.
func (o *Setup) SetAccounts(v []Account) {
	o.Accounts = v
}

func (o Setup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Setup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversionDate) {
		toSerialize["ConversionDate"] = o.ConversionDate
	}
	if !IsNil(o.ConversionBalances) {
		toSerialize["ConversionBalances"] = o.ConversionBalances
	}
	if !IsNil(o.Accounts) {
		toSerialize["Accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableSetup struct {
	value *Setup
	isSet bool
}

func (v NullableSetup) Get() *Setup {
	return v.value
}

func (v *NullableSetup) Set(val *Setup) {
	v.value = val
	v.isSet = true
}

func (v NullableSetup) IsSet() bool {
	return v.isSet
}

func (v *NullableSetup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetup(val *Setup) *NullableSetup {
	return &NullableSetup{value: val, isSet: true}
}

func (v NullableSetup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


