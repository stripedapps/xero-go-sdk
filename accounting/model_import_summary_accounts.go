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

// checks if the ImportSummaryAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportSummaryAccounts{}

// ImportSummaryAccounts A summary of the accounts changes
type ImportSummaryAccounts struct {
	// The total number of accounts in the org
	Total *int32 `json:"Total,omitempty"`
	// The number of new accounts created
	New *int32 `json:"New,omitempty"`
	// The number of accounts updated
	Updated *int32 `json:"Updated,omitempty"`
	// The number of accounts deleted
	Deleted *int32 `json:"Deleted,omitempty"`
	// The number of locked accounts
	Locked *int32 `json:"Locked,omitempty"`
	// The number of system accounts
	System *int32 `json:"System,omitempty"`
	// The number of accounts that had an error
	Errored *int32 `json:"Errored,omitempty"`
	Present *bool `json:"Present,omitempty"`
	// The number of new or updated accounts
	NewOrUpdated *int32 `json:"NewOrUpdated,omitempty"`
}

// NewImportSummaryAccounts instantiates a new ImportSummaryAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSummaryAccounts() *ImportSummaryAccounts {
	this := ImportSummaryAccounts{}
	return &this
}

// NewImportSummaryAccountsWithDefaults instantiates a new ImportSummaryAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSummaryAccountsWithDefaults() *ImportSummaryAccounts {
	this := ImportSummaryAccounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ImportSummaryAccounts) SetTotal(v int32) {
	o.Total = &v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasNew() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *ImportSummaryAccounts) SetNew(v int32) {
	o.New = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetUpdated() int32 {
	if o == nil || IsNil(o.Updated) {
		var ret int32
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetUpdatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given int32 and assigns it to the Updated field.
func (o *ImportSummaryAccounts) SetUpdated(v int32) {
	o.Updated = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetDeleted() int32 {
	if o == nil || IsNil(o.Deleted) {
		var ret int32
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int32 and assigns it to the Deleted field.
func (o *ImportSummaryAccounts) SetDeleted(v int32) {
	o.Deleted = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetLocked() int32 {
	if o == nil || IsNil(o.Locked) {
		var ret int32
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetLockedOk() (*int32, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given int32 and assigns it to the Locked field.
func (o *ImportSummaryAccounts) SetLocked(v int32) {
	o.Locked = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetSystem() int32 {
	if o == nil || IsNil(o.System) {
		var ret int32
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetSystemOk() (*int32, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given int32 and assigns it to the System field.
func (o *ImportSummaryAccounts) SetSystem(v int32) {
	o.System = &v
}

// GetErrored returns the Errored field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetErrored() int32 {
	if o == nil || IsNil(o.Errored) {
		var ret int32
		return ret
	}
	return *o.Errored
}

// GetErroredOk returns a tuple with the Errored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetErroredOk() (*int32, bool) {
	if o == nil || IsNil(o.Errored) {
		return nil, false
	}
	return o.Errored, true
}

// HasErrored returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasErrored() bool {
	if o != nil && !IsNil(o.Errored) {
		return true
	}

	return false
}

// SetErrored gets a reference to the given int32 and assigns it to the Errored field.
func (o *ImportSummaryAccounts) SetErrored(v int32) {
	o.Errored = &v
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetPresent() bool {
	if o == nil || IsNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Present) {
		return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasPresent() bool {
	if o != nil && !IsNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *ImportSummaryAccounts) SetPresent(v bool) {
	o.Present = &v
}

// GetNewOrUpdated returns the NewOrUpdated field value if set, zero value otherwise.
func (o *ImportSummaryAccounts) GetNewOrUpdated() int32 {
	if o == nil || IsNil(o.NewOrUpdated) {
		var ret int32
		return ret
	}
	return *o.NewOrUpdated
}

// GetNewOrUpdatedOk returns a tuple with the NewOrUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSummaryAccounts) GetNewOrUpdatedOk() (*int32, bool) {
	if o == nil || IsNil(o.NewOrUpdated) {
		return nil, false
	}
	return o.NewOrUpdated, true
}

// HasNewOrUpdated returns a boolean if a field has been set.
func (o *ImportSummaryAccounts) HasNewOrUpdated() bool {
	if o != nil && !IsNil(o.NewOrUpdated) {
		return true
	}

	return false
}

// SetNewOrUpdated gets a reference to the given int32 and assigns it to the NewOrUpdated field.
func (o *ImportSummaryAccounts) SetNewOrUpdated(v int32) {
	o.NewOrUpdated = &v
}

func (o ImportSummaryAccounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportSummaryAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}
	if !IsNil(o.New) {
		toSerialize["New"] = o.New
	}
	if !IsNil(o.Updated) {
		toSerialize["Updated"] = o.Updated
	}
	if !IsNil(o.Deleted) {
		toSerialize["Deleted"] = o.Deleted
	}
	if !IsNil(o.Locked) {
		toSerialize["Locked"] = o.Locked
	}
	if !IsNil(o.System) {
		toSerialize["System"] = o.System
	}
	if !IsNil(o.Errored) {
		toSerialize["Errored"] = o.Errored
	}
	if !IsNil(o.Present) {
		toSerialize["Present"] = o.Present
	}
	if !IsNil(o.NewOrUpdated) {
		toSerialize["NewOrUpdated"] = o.NewOrUpdated
	}
	return toSerialize, nil
}

type NullableImportSummaryAccounts struct {
	value *ImportSummaryAccounts
	isSet bool
}

func (v NullableImportSummaryAccounts) Get() *ImportSummaryAccounts {
	return v.value
}

func (v *NullableImportSummaryAccounts) Set(val *ImportSummaryAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSummaryAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSummaryAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSummaryAccounts(val *ImportSummaryAccounts) *NullableImportSummaryAccounts {
	return &NullableImportSummaryAccounts{value: val, isSet: true}
}

func (v NullableImportSummaryAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSummaryAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

