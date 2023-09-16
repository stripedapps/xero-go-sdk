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

// checks if the HistoryRecords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoryRecords{}

// HistoryRecords struct for HistoryRecords
type HistoryRecords struct {
	HistoryRecords []HistoryRecord `json:"HistoryRecords,omitempty"`
}

// NewHistoryRecords instantiates a new HistoryRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryRecords() *HistoryRecords {
	this := HistoryRecords{}
	return &this
}

// NewHistoryRecordsWithDefaults instantiates a new HistoryRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryRecordsWithDefaults() *HistoryRecords {
	this := HistoryRecords{}
	return &this
}

// GetHistoryRecords returns the HistoryRecords field value if set, zero value otherwise.
func (o *HistoryRecords) GetHistoryRecords() []HistoryRecord {
	if o == nil || IsNil(o.HistoryRecords) {
		var ret []HistoryRecord
		return ret
	}
	return o.HistoryRecords
}

// GetHistoryRecordsOk returns a tuple with the HistoryRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryRecords) GetHistoryRecordsOk() ([]HistoryRecord, bool) {
	if o == nil || IsNil(o.HistoryRecords) {
		return nil, false
	}
	return o.HistoryRecords, true
}

// HasHistoryRecords returns a boolean if a field has been set.
func (o *HistoryRecords) HasHistoryRecords() bool {
	if o != nil && !IsNil(o.HistoryRecords) {
		return true
	}

	return false
}

// SetHistoryRecords gets a reference to the given []HistoryRecord and assigns it to the HistoryRecords field.
func (o *HistoryRecords) SetHistoryRecords(v []HistoryRecord) {
	o.HistoryRecords = v
}

func (o HistoryRecords) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoryRecords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HistoryRecords) {
		toSerialize["HistoryRecords"] = o.HistoryRecords
	}
	return toSerialize, nil
}

type NullableHistoryRecords struct {
	value *HistoryRecords
	isSet bool
}

func (v NullableHistoryRecords) Get() *HistoryRecords {
	return v.value
}

func (v *NullableHistoryRecords) Set(val *HistoryRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryRecords(val *HistoryRecords) *NullableHistoryRecords {
	return &NullableHistoryRecords{value: val, isSet: true}
}

func (v NullableHistoryRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

