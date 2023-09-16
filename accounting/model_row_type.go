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

// RowType the model 'RowType'
type RowType string

// List of RowType
const (
	ROWTYPE_HEADER RowType = "Header"
	ROWTYPE_SECTION RowType = "Section"
	ROWTYPE_ROW RowType = "Row"
	ROWTYPE_SUMMARY_ROW RowType = "SummaryRow"
)

// All allowed values of RowType enum
var AllowedRowTypeEnumValues = []RowType{
	"Header",
	"Section",
	"Row",
	"SummaryRow",
}

func (v *RowType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RowType(value)
	for _, existing := range AllowedRowTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RowType", value)
}

// NewRowTypeFromValue returns a pointer to a valid RowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRowTypeFromValue(v string) (*RowType, error) {
	ev := RowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RowType: valid values are %v", v, AllowedRowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RowType) IsValid() bool {
	for _, existing := range AllowedRowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RowType value
func (v RowType) Ptr() *RowType {
	return &v
}

type NullableRowType struct {
	value *RowType
	isSet bool
}

func (v NullableRowType) Get() *RowType {
	return v.value
}

func (v *NullableRowType) Set(val *RowType) {
	v.value = val
	v.isSet = true
}

func (v NullableRowType) IsSet() bool {
	return v.isSet
}

func (v *NullableRowType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRowType(val *RowType) *NullableRowType {
	return &NullableRowType{value: val, isSet: true}
}

func (v NullableRowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRowType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

