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
	"fmt"
)

// LineAmountTypes Line amounts are exclusive of tax by default if you don’t specify this element. See Line Amount Types
type LineAmountTypes string

// List of LineAmountTypes
const (
	LINEAMOUNTTYPES_EXCLUSIVE LineAmountTypes = "Exclusive"
	LINEAMOUNTTYPES_INCLUSIVE LineAmountTypes = "Inclusive"
	LINEAMOUNTTYPES_NO_TAX LineAmountTypes = "NoTax"
)

// All allowed values of LineAmountTypes enum
var AllowedLineAmountTypesEnumValues = []LineAmountTypes{
	"Exclusive",
	"Inclusive",
	"NoTax",
}

func (v *LineAmountTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LineAmountTypes(value)
	for _, existing := range AllowedLineAmountTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LineAmountTypes", value)
}

// NewLineAmountTypesFromValue returns a pointer to a valid LineAmountTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLineAmountTypesFromValue(v string) (*LineAmountTypes, error) {
	ev := LineAmountTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LineAmountTypes: valid values are %v", v, AllowedLineAmountTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LineAmountTypes) IsValid() bool {
	for _, existing := range AllowedLineAmountTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LineAmountTypes value
func (v LineAmountTypes) Ptr() *LineAmountTypes {
	return &v
}

type NullableLineAmountTypes struct {
	value *LineAmountTypes
	isSet bool
}

func (v NullableLineAmountTypes) Get() *LineAmountTypes {
	return v.value
}

func (v *NullableLineAmountTypes) Set(val *LineAmountTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableLineAmountTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableLineAmountTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineAmountTypes(val *LineAmountTypes) *NullableLineAmountTypes {
	return &NullableLineAmountTypes{value: val, isSet: true}
}

func (v NullableLineAmountTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineAmountTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

