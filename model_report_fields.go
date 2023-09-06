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

// checks if the ReportFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportFields{}

// ReportFields struct for ReportFields
type ReportFields struct {
	FieldID *string `json:"FieldID,omitempty"`
	Description *string `json:"Description,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// NewReportFields instantiates a new ReportFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportFields() *ReportFields {
	this := ReportFields{}
	return &this
}

// NewReportFieldsWithDefaults instantiates a new ReportFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportFieldsWithDefaults() *ReportFields {
	this := ReportFields{}
	return &this
}

// GetFieldID returns the FieldID field value if set, zero value otherwise.
func (o *ReportFields) GetFieldID() string {
	if o == nil || IsNil(o.FieldID) {
		var ret string
		return ret
	}
	return *o.FieldID
}

// GetFieldIDOk returns a tuple with the FieldID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportFields) GetFieldIDOk() (*string, bool) {
	if o == nil || IsNil(o.FieldID) {
		return nil, false
	}
	return o.FieldID, true
}

// HasFieldID returns a boolean if a field has been set.
func (o *ReportFields) HasFieldID() bool {
	if o != nil && !IsNil(o.FieldID) {
		return true
	}

	return false
}

// SetFieldID gets a reference to the given string and assigns it to the FieldID field.
func (o *ReportFields) SetFieldID(v string) {
	o.FieldID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReportFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReportFields) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ReportFields) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportFields) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ReportFields) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ReportFields) SetValue(v string) {
	o.Value = &v
}

func (o ReportFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldID) {
		toSerialize["FieldID"] = o.FieldID
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableReportFields struct {
	value *ReportFields
	isSet bool
}

func (v NullableReportFields) Get() *ReportFields {
	return v.value
}

func (v *NullableReportFields) Set(val *ReportFields) {
	v.value = val
	v.isSet = true
}

func (v NullableReportFields) IsSet() bool {
	return v.isSet
}

func (v *NullableReportFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportFields(val *ReportFields) *NullableReportFields {
	return &NullableReportFields{value: val, isSet: true}
}

func (v NullableReportFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


