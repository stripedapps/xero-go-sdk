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

// checks if the Budgets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Budgets{}

// Budgets struct for Budgets
type Budgets struct {
	Budgets []Budget `json:"Budgets,omitempty"`
}

// NewBudgets instantiates a new Budgets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgets() *Budgets {
	this := Budgets{}
	return &this
}

// NewBudgetsWithDefaults instantiates a new Budgets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetsWithDefaults() *Budgets {
	this := Budgets{}
	return &this
}

// GetBudgets returns the Budgets field value if set, zero value otherwise.
func (o *Budgets) GetBudgets() []Budget {
	if o == nil || IsNil(o.Budgets) {
		var ret []Budget
		return ret
	}
	return o.Budgets
}

// GetBudgetsOk returns a tuple with the Budgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budgets) GetBudgetsOk() ([]Budget, bool) {
	if o == nil || IsNil(o.Budgets) {
		return nil, false
	}
	return o.Budgets, true
}

// HasBudgets returns a boolean if a field has been set.
func (o *Budgets) HasBudgets() bool {
	if o != nil && !IsNil(o.Budgets) {
		return true
	}

	return false
}

// SetBudgets gets a reference to the given []Budget and assigns it to the Budgets field.
func (o *Budgets) SetBudgets(v []Budget) {
	o.Budgets = v
}

func (o Budgets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Budgets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Budgets) {
		toSerialize["Budgets"] = o.Budgets
	}
	return toSerialize, nil
}

type NullableBudgets struct {
	value *Budgets
	isSet bool
}

func (v NullableBudgets) Get() *Budgets {
	return v.value
}

func (v *NullableBudgets) Set(val *Budgets) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgets) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgets(val *Budgets) *NullableBudgets {
	return &NullableBudgets{value: val, isSet: true}
}

func (v NullableBudgets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

