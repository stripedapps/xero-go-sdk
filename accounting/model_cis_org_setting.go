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

// checks if the CISOrgSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CISOrgSetting{}

// CISOrgSetting struct for CISOrgSetting
type CISOrgSetting struct {
	// true or false - Boolean that describes if the organisation is a CIS Contractor
	CISContractorEnabled *bool `json:"CISContractorEnabled,omitempty"`
	// true or false - Boolean that describes if the organisation is a CIS SubContractor
	CISSubContractorEnabled *bool `json:"CISSubContractorEnabled,omitempty"`
	// CIS Deduction rate for the organisation
	Rate *float64 `json:"Rate,omitempty"`
}

// NewCISOrgSetting instantiates a new CISOrgSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCISOrgSetting() *CISOrgSetting {
	this := CISOrgSetting{}
	return &this
}

// NewCISOrgSettingWithDefaults instantiates a new CISOrgSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCISOrgSettingWithDefaults() *CISOrgSetting {
	this := CISOrgSetting{}
	return &this
}

// GetCISContractorEnabled returns the CISContractorEnabled field value if set, zero value otherwise.
func (o *CISOrgSetting) GetCISContractorEnabled() bool {
	if o == nil || IsNil(o.CISContractorEnabled) {
		var ret bool
		return ret
	}
	return *o.CISContractorEnabled
}

// GetCISContractorEnabledOk returns a tuple with the CISContractorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISOrgSetting) GetCISContractorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CISContractorEnabled) {
		return nil, false
	}
	return o.CISContractorEnabled, true
}

// HasCISContractorEnabled returns a boolean if a field has been set.
func (o *CISOrgSetting) HasCISContractorEnabled() bool {
	if o != nil && !IsNil(o.CISContractorEnabled) {
		return true
	}

	return false
}

// SetCISContractorEnabled gets a reference to the given bool and assigns it to the CISContractorEnabled field.
func (o *CISOrgSetting) SetCISContractorEnabled(v bool) {
	o.CISContractorEnabled = &v
}

// GetCISSubContractorEnabled returns the CISSubContractorEnabled field value if set, zero value otherwise.
func (o *CISOrgSetting) GetCISSubContractorEnabled() bool {
	if o == nil || IsNil(o.CISSubContractorEnabled) {
		var ret bool
		return ret
	}
	return *o.CISSubContractorEnabled
}

// GetCISSubContractorEnabledOk returns a tuple with the CISSubContractorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISOrgSetting) GetCISSubContractorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CISSubContractorEnabled) {
		return nil, false
	}
	return o.CISSubContractorEnabled, true
}

// HasCISSubContractorEnabled returns a boolean if a field has been set.
func (o *CISOrgSetting) HasCISSubContractorEnabled() bool {
	if o != nil && !IsNil(o.CISSubContractorEnabled) {
		return true
	}

	return false
}

// SetCISSubContractorEnabled gets a reference to the given bool and assigns it to the CISSubContractorEnabled field.
func (o *CISOrgSetting) SetCISSubContractorEnabled(v bool) {
	o.CISSubContractorEnabled = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CISOrgSetting) GetRate() float64 {
	if o == nil || IsNil(o.Rate) {
		var ret float64
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CISOrgSetting) GetRateOk() (*float64, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CISOrgSetting) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float64 and assigns it to the Rate field.
func (o *CISOrgSetting) SetRate(v float64) {
	o.Rate = &v
}

func (o CISOrgSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CISOrgSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CISContractorEnabled) {
		toSerialize["CISContractorEnabled"] = o.CISContractorEnabled
	}
	if !IsNil(o.CISSubContractorEnabled) {
		toSerialize["CISSubContractorEnabled"] = o.CISSubContractorEnabled
	}
	if !IsNil(o.Rate) {
		toSerialize["Rate"] = o.Rate
	}
	return toSerialize, nil
}

type NullableCISOrgSetting struct {
	value *CISOrgSetting
	isSet bool
}

func (v NullableCISOrgSetting) Get() *CISOrgSetting {
	return v.value
}

func (v *NullableCISOrgSetting) Set(val *CISOrgSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableCISOrgSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableCISOrgSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCISOrgSetting(val *CISOrgSetting) *NullableCISOrgSetting {
	return &NullableCISOrgSetting{value: val, isSet: true}
}

func (v NullableCISOrgSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCISOrgSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

