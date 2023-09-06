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

// checks if the TrackingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingOption{}

// TrackingOption struct for TrackingOption
type TrackingOption struct {
	// The Xero identifier for a tracking option e.g. ae777a87-5ef3-4fa0-a4f0-d10e1f13073a
	TrackingOptionID *string `json:"TrackingOptionID,omitempty"`
	// The name of the tracking option e.g. Marketing, East (max length = 100)
	Name *string `json:"Name,omitempty"`
	// The status of a tracking option
	Status *string `json:"Status,omitempty"`
	// Filter by a tracking category e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID *string `json:"TrackingCategoryID,omitempty"`
}

// NewTrackingOption instantiates a new TrackingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingOption() *TrackingOption {
	this := TrackingOption{}
	return &this
}

// NewTrackingOptionWithDefaults instantiates a new TrackingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingOptionWithDefaults() *TrackingOption {
	this := TrackingOption{}
	return &this
}

// GetTrackingOptionID returns the TrackingOptionID field value if set, zero value otherwise.
func (o *TrackingOption) GetTrackingOptionID() string {
	if o == nil || IsNil(o.TrackingOptionID) {
		var ret string
		return ret
	}
	return *o.TrackingOptionID
}

// GetTrackingOptionIDOk returns a tuple with the TrackingOptionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingOption) GetTrackingOptionIDOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingOptionID) {
		return nil, false
	}
	return o.TrackingOptionID, true
}

// HasTrackingOptionID returns a boolean if a field has been set.
func (o *TrackingOption) HasTrackingOptionID() bool {
	if o != nil && !IsNil(o.TrackingOptionID) {
		return true
	}

	return false
}

// SetTrackingOptionID gets a reference to the given string and assigns it to the TrackingOptionID field.
func (o *TrackingOption) SetTrackingOptionID(v string) {
	o.TrackingOptionID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrackingOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrackingOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrackingOption) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrackingOption) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingOption) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrackingOption) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrackingOption) SetStatus(v string) {
	o.Status = &v
}

// GetTrackingCategoryID returns the TrackingCategoryID field value if set, zero value otherwise.
func (o *TrackingOption) GetTrackingCategoryID() string {
	if o == nil || IsNil(o.TrackingCategoryID) {
		var ret string
		return ret
	}
	return *o.TrackingCategoryID
}

// GetTrackingCategoryIDOk returns a tuple with the TrackingCategoryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingOption) GetTrackingCategoryIDOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCategoryID) {
		return nil, false
	}
	return o.TrackingCategoryID, true
}

// HasTrackingCategoryID returns a boolean if a field has been set.
func (o *TrackingOption) HasTrackingCategoryID() bool {
	if o != nil && !IsNil(o.TrackingCategoryID) {
		return true
	}

	return false
}

// SetTrackingCategoryID gets a reference to the given string and assigns it to the TrackingCategoryID field.
func (o *TrackingOption) SetTrackingCategoryID(v string) {
	o.TrackingCategoryID = &v
}

func (o TrackingOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrackingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingOptionID) {
		toSerialize["TrackingOptionID"] = o.TrackingOptionID
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TrackingCategoryID) {
		toSerialize["TrackingCategoryID"] = o.TrackingCategoryID
	}
	return toSerialize, nil
}

type NullableTrackingOption struct {
	value *TrackingOption
	isSet bool
}

func (v NullableTrackingOption) Get() *TrackingOption {
	return v.value
}

func (v *NullableTrackingOption) Set(val *TrackingOption) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingOption) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingOption(val *TrackingOption) *NullableTrackingOption {
	return &NullableTrackingOption{value: val, isSet: true}
}

func (v NullableTrackingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

