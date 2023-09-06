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

// checks if the Employees type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Employees{}

// Employees struct for Employees
type Employees struct {
	Employees []Employee `json:"Employees,omitempty"`
}

// NewEmployees instantiates a new Employees object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployees() *Employees {
	this := Employees{}
	return &this
}

// NewEmployeesWithDefaults instantiates a new Employees object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeesWithDefaults() *Employees {
	this := Employees{}
	return &this
}

// GetEmployees returns the Employees field value if set, zero value otherwise.
func (o *Employees) GetEmployees() []Employee {
	if o == nil || IsNil(o.Employees) {
		var ret []Employee
		return ret
	}
	return o.Employees
}

// GetEmployeesOk returns a tuple with the Employees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employees) GetEmployeesOk() ([]Employee, bool) {
	if o == nil || IsNil(o.Employees) {
		return nil, false
	}
	return o.Employees, true
}

// HasEmployees returns a boolean if a field has been set.
func (o *Employees) HasEmployees() bool {
	if o != nil && !IsNil(o.Employees) {
		return true
	}

	return false
}

// SetEmployees gets a reference to the given []Employee and assigns it to the Employees field.
func (o *Employees) SetEmployees(v []Employee) {
	o.Employees = v
}

func (o Employees) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Employees) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Employees) {
		toSerialize["Employees"] = o.Employees
	}
	return toSerialize, nil
}

type NullableEmployees struct {
	value *Employees
	isSet bool
}

func (v NullableEmployees) Get() *Employees {
	return v.value
}

func (v *NullableEmployees) Set(val *Employees) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployees) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployees) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployees(val *Employees) *NullableEmployees {
	return &NullableEmployees{value: val, isSet: true}
}

func (v NullableEmployees) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployees) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

