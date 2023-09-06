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

// checks if the BrandingThemes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingThemes{}

// BrandingThemes struct for BrandingThemes
type BrandingThemes struct {
	BrandingThemes []BrandingTheme `json:"BrandingThemes,omitempty"`
}

// NewBrandingThemes instantiates a new BrandingThemes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingThemes() *BrandingThemes {
	this := BrandingThemes{}
	return &this
}

// NewBrandingThemesWithDefaults instantiates a new BrandingThemes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingThemesWithDefaults() *BrandingThemes {
	this := BrandingThemes{}
	return &this
}

// GetBrandingThemes returns the BrandingThemes field value if set, zero value otherwise.
func (o *BrandingThemes) GetBrandingThemes() []BrandingTheme {
	if o == nil || IsNil(o.BrandingThemes) {
		var ret []BrandingTheme
		return ret
	}
	return o.BrandingThemes
}

// GetBrandingThemesOk returns a tuple with the BrandingThemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingThemes) GetBrandingThemesOk() ([]BrandingTheme, bool) {
	if o == nil || IsNil(o.BrandingThemes) {
		return nil, false
	}
	return o.BrandingThemes, true
}

// HasBrandingThemes returns a boolean if a field has been set.
func (o *BrandingThemes) HasBrandingThemes() bool {
	if o != nil && !IsNil(o.BrandingThemes) {
		return true
	}

	return false
}

// SetBrandingThemes gets a reference to the given []BrandingTheme and assigns it to the BrandingThemes field.
func (o *BrandingThemes) SetBrandingThemes(v []BrandingTheme) {
	o.BrandingThemes = v
}

func (o BrandingThemes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingThemes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandingThemes) {
		toSerialize["BrandingThemes"] = o.BrandingThemes
	}
	return toSerialize, nil
}

type NullableBrandingThemes struct {
	value *BrandingThemes
	isSet bool
}

func (v NullableBrandingThemes) Get() *BrandingThemes {
	return v.value
}

func (v *NullableBrandingThemes) Set(val *BrandingThemes) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingThemes) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingThemes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingThemes(val *BrandingThemes) *NullableBrandingThemes {
	return &NullableBrandingThemes{value: val, isSet: true}
}

func (v NullableBrandingThemes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingThemes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


