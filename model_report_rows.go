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

// checks if the ReportRows type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportRows{}

// ReportRows struct for ReportRows
type ReportRows struct {
	RowType *RowType `json:"RowType,omitempty"`
	Title *string `json:"Title,omitempty"`
	Cells []ReportCell `json:"Cells,omitempty"`
	Rows []ReportRow `json:"Rows,omitempty"`
}

// NewReportRows instantiates a new ReportRows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportRows() *ReportRows {
	this := ReportRows{}
	return &this
}

// NewReportRowsWithDefaults instantiates a new ReportRows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportRowsWithDefaults() *ReportRows {
	this := ReportRows{}
	return &this
}

// GetRowType returns the RowType field value if set, zero value otherwise.
func (o *ReportRows) GetRowType() RowType {
	if o == nil || IsNil(o.RowType) {
		var ret RowType
		return ret
	}
	return *o.RowType
}

// GetRowTypeOk returns a tuple with the RowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetRowTypeOk() (*RowType, bool) {
	if o == nil || IsNil(o.RowType) {
		return nil, false
	}
	return o.RowType, true
}

// HasRowType returns a boolean if a field has been set.
func (o *ReportRows) HasRowType() bool {
	if o != nil && !IsNil(o.RowType) {
		return true
	}

	return false
}

// SetRowType gets a reference to the given RowType and assigns it to the RowType field.
func (o *ReportRows) SetRowType(v RowType) {
	o.RowType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ReportRows) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ReportRows) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ReportRows) SetTitle(v string) {
	o.Title = &v
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *ReportRows) GetCells() []ReportCell {
	if o == nil || IsNil(o.Cells) {
		var ret []ReportCell
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetCellsOk() ([]ReportCell, bool) {
	if o == nil || IsNil(o.Cells) {
		return nil, false
	}
	return o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *ReportRows) HasCells() bool {
	if o != nil && !IsNil(o.Cells) {
		return true
	}

	return false
}

// SetCells gets a reference to the given []ReportCell and assigns it to the Cells field.
func (o *ReportRows) SetCells(v []ReportCell) {
	o.Cells = v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ReportRows) GetRows() []ReportRow {
	if o == nil || IsNil(o.Rows) {
		var ret []ReportRow
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetRowsOk() ([]ReportRow, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ReportRows) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given []ReportRow and assigns it to the Rows field.
func (o *ReportRows) SetRows(v []ReportRow) {
	o.Rows = v
}

func (o ReportRows) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportRows) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RowType) {
		toSerialize["RowType"] = o.RowType
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.Cells) {
		toSerialize["Cells"] = o.Cells
	}
	if !IsNil(o.Rows) {
		toSerialize["Rows"] = o.Rows
	}
	return toSerialize, nil
}

type NullableReportRows struct {
	value *ReportRows
	isSet bool
}

func (v NullableReportRows) Get() *ReportRows {
	return v.value
}

func (v *NullableReportRows) Set(val *ReportRows) {
	v.value = val
	v.isSet = true
}

func (v NullableReportRows) IsSet() bool {
	return v.isSet
}

func (v *NullableReportRows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportRows(val *ReportRows) *NullableReportRows {
	return &NullableReportRows{value: val, isSet: true}
}

func (v NullableReportRows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportRows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


