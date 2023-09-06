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

// checks if the SalesTrackingCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SalesTrackingCategory{}

// SalesTrackingCategory struct for SalesTrackingCategory
type SalesTrackingCategory struct {
	// The default sales tracking category name for contacts
	TrackingCategoryName *string `json:"TrackingCategoryName,omitempty"`
	// The default purchase tracking category name for contacts
	TrackingOptionName *string `json:"TrackingOptionName,omitempty"`
}

// NewSalesTrackingCategory instantiates a new SalesTrackingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesTrackingCategory() *SalesTrackingCategory {
	this := SalesTrackingCategory{}
	return &this
}

// NewSalesTrackingCategoryWithDefaults instantiates a new SalesTrackingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesTrackingCategoryWithDefaults() *SalesTrackingCategory {
	this := SalesTrackingCategory{}
	return &this
}

// GetTrackingCategoryName returns the TrackingCategoryName field value if set, zero value otherwise.
func (o *SalesTrackingCategory) GetTrackingCategoryName() string {
	if o == nil || IsNil(o.TrackingCategoryName) {
		var ret string
		return ret
	}
	return *o.TrackingCategoryName
}

// GetTrackingCategoryNameOk returns a tuple with the TrackingCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTrackingCategory) GetTrackingCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCategoryName) {
		return nil, false
	}
	return o.TrackingCategoryName, true
}

// HasTrackingCategoryName returns a boolean if a field has been set.
func (o *SalesTrackingCategory) HasTrackingCategoryName() bool {
	if o != nil && !IsNil(o.TrackingCategoryName) {
		return true
	}

	return false
}

// SetTrackingCategoryName gets a reference to the given string and assigns it to the TrackingCategoryName field.
func (o *SalesTrackingCategory) SetTrackingCategoryName(v string) {
	o.TrackingCategoryName = &v
}

// GetTrackingOptionName returns the TrackingOptionName field value if set, zero value otherwise.
func (o *SalesTrackingCategory) GetTrackingOptionName() string {
	if o == nil || IsNil(o.TrackingOptionName) {
		var ret string
		return ret
	}
	return *o.TrackingOptionName
}

// GetTrackingOptionNameOk returns a tuple with the TrackingOptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesTrackingCategory) GetTrackingOptionNameOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingOptionName) {
		return nil, false
	}
	return o.TrackingOptionName, true
}

// HasTrackingOptionName returns a boolean if a field has been set.
func (o *SalesTrackingCategory) HasTrackingOptionName() bool {
	if o != nil && !IsNil(o.TrackingOptionName) {
		return true
	}

	return false
}

// SetTrackingOptionName gets a reference to the given string and assigns it to the TrackingOptionName field.
func (o *SalesTrackingCategory) SetTrackingOptionName(v string) {
	o.TrackingOptionName = &v
}

func (o SalesTrackingCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SalesTrackingCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingCategoryName) {
		toSerialize["TrackingCategoryName"] = o.TrackingCategoryName
	}
	if !IsNil(o.TrackingOptionName) {
		toSerialize["TrackingOptionName"] = o.TrackingOptionName
	}
	return toSerialize, nil
}

type NullableSalesTrackingCategory struct {
	value *SalesTrackingCategory
	isSet bool
}

func (v NullableSalesTrackingCategory) Get() *SalesTrackingCategory {
	return v.value
}

func (v *NullableSalesTrackingCategory) Set(val *SalesTrackingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesTrackingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesTrackingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesTrackingCategory(val *SalesTrackingCategory) *NullableSalesTrackingCategory {
	return &NullableSalesTrackingCategory{value: val, isSet: true}
}

func (v NullableSalesTrackingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesTrackingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


