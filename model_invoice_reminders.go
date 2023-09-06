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

// checks if the InvoiceReminders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceReminders{}

// InvoiceReminders struct for InvoiceReminders
type InvoiceReminders struct {
	InvoiceReminders []InvoiceReminder `json:"InvoiceReminders,omitempty"`
}

// NewInvoiceReminders instantiates a new InvoiceReminders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceReminders() *InvoiceReminders {
	this := InvoiceReminders{}
	return &this
}

// NewInvoiceRemindersWithDefaults instantiates a new InvoiceReminders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRemindersWithDefaults() *InvoiceReminders {
	this := InvoiceReminders{}
	return &this
}

// GetInvoiceReminders returns the InvoiceReminders field value if set, zero value otherwise.
func (o *InvoiceReminders) GetInvoiceReminders() []InvoiceReminder {
	if o == nil || IsNil(o.InvoiceReminders) {
		var ret []InvoiceReminder
		return ret
	}
	return o.InvoiceReminders
}

// GetInvoiceRemindersOk returns a tuple with the InvoiceReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReminders) GetInvoiceRemindersOk() ([]InvoiceReminder, bool) {
	if o == nil || IsNil(o.InvoiceReminders) {
		return nil, false
	}
	return o.InvoiceReminders, true
}

// HasInvoiceReminders returns a boolean if a field has been set.
func (o *InvoiceReminders) HasInvoiceReminders() bool {
	if o != nil && !IsNil(o.InvoiceReminders) {
		return true
	}

	return false
}

// SetInvoiceReminders gets a reference to the given []InvoiceReminder and assigns it to the InvoiceReminders field.
func (o *InvoiceReminders) SetInvoiceReminders(v []InvoiceReminder) {
	o.InvoiceReminders = v
}

func (o InvoiceReminders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceReminders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceReminders) {
		toSerialize["InvoiceReminders"] = o.InvoiceReminders
	}
	return toSerialize, nil
}

type NullableInvoiceReminders struct {
	value *InvoiceReminders
	isSet bool
}

func (v NullableInvoiceReminders) Get() *InvoiceReminders {
	return v.value
}

func (v *NullableInvoiceReminders) Set(val *InvoiceReminders) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceReminders) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceReminders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceReminders(val *InvoiceReminders) *NullableInvoiceReminders {
	return &NullableInvoiceReminders{value: val, isSet: true}
}

func (v NullableInvoiceReminders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceReminders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


