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

// checks if the Phone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Phone{}

// Phone struct for Phone
type Phone struct {
	PhoneType *string `json:"PhoneType,omitempty"`
	// max length = 50
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// max length = 10
	PhoneAreaCode *string `json:"PhoneAreaCode,omitempty"`
	// max length = 20
	PhoneCountryCode *string `json:"PhoneCountryCode,omitempty"`
}

// NewPhone instantiates a new Phone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhone() *Phone {
	this := Phone{}
	return &this
}

// NewPhoneWithDefaults instantiates a new Phone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneWithDefaults() *Phone {
	this := Phone{}
	return &this
}

// GetPhoneType returns the PhoneType field value if set, zero value otherwise.
func (o *Phone) GetPhoneType() string {
	if o == nil || IsNil(o.PhoneType) {
		var ret string
		return ret
	}
	return *o.PhoneType
}

// GetPhoneTypeOk returns a tuple with the PhoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetPhoneTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneType) {
		return nil, false
	}
	return o.PhoneType, true
}

// HasPhoneType returns a boolean if a field has been set.
func (o *Phone) HasPhoneType() bool {
	if o != nil && !IsNil(o.PhoneType) {
		return true
	}

	return false
}

// SetPhoneType gets a reference to the given string and assigns it to the PhoneType field.
func (o *Phone) SetPhoneType(v string) {
	o.PhoneType = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Phone) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Phone) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Phone) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneAreaCode returns the PhoneAreaCode field value if set, zero value otherwise.
func (o *Phone) GetPhoneAreaCode() string {
	if o == nil || IsNil(o.PhoneAreaCode) {
		var ret string
		return ret
	}
	return *o.PhoneAreaCode
}

// GetPhoneAreaCodeOk returns a tuple with the PhoneAreaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetPhoneAreaCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneAreaCode) {
		return nil, false
	}
	return o.PhoneAreaCode, true
}

// HasPhoneAreaCode returns a boolean if a field has been set.
func (o *Phone) HasPhoneAreaCode() bool {
	if o != nil && !IsNil(o.PhoneAreaCode) {
		return true
	}

	return false
}

// SetPhoneAreaCode gets a reference to the given string and assigns it to the PhoneAreaCode field.
func (o *Phone) SetPhoneAreaCode(v string) {
	o.PhoneAreaCode = &v
}

// GetPhoneCountryCode returns the PhoneCountryCode field value if set, zero value otherwise.
func (o *Phone) GetPhoneCountryCode() string {
	if o == nil || IsNil(o.PhoneCountryCode) {
		var ret string
		return ret
	}
	return *o.PhoneCountryCode
}

// GetPhoneCountryCodeOk returns a tuple with the PhoneCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Phone) GetPhoneCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneCountryCode) {
		return nil, false
	}
	return o.PhoneCountryCode, true
}

// HasPhoneCountryCode returns a boolean if a field has been set.
func (o *Phone) HasPhoneCountryCode() bool {
	if o != nil && !IsNil(o.PhoneCountryCode) {
		return true
	}

	return false
}

// SetPhoneCountryCode gets a reference to the given string and assigns it to the PhoneCountryCode field.
func (o *Phone) SetPhoneCountryCode(v string) {
	o.PhoneCountryCode = &v
}

func (o Phone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Phone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneType) {
		toSerialize["PhoneType"] = o.PhoneType
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["PhoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PhoneAreaCode) {
		toSerialize["PhoneAreaCode"] = o.PhoneAreaCode
	}
	if !IsNil(o.PhoneCountryCode) {
		toSerialize["PhoneCountryCode"] = o.PhoneCountryCode
	}
	return toSerialize, nil
}

type NullablePhone struct {
	value *Phone
	isSet bool
}

func (v NullablePhone) Get() *Phone {
	return v.value
}

func (v *NullablePhone) Set(val *Phone) {
	v.value = val
	v.isSet = true
}

func (v NullablePhone) IsSet() bool {
	return v.isSet
}

func (v *NullablePhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhone(val *Phone) *NullablePhone {
	return &NullablePhone{value: val, isSet: true}
}

func (v NullablePhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


