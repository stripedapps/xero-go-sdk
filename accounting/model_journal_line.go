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

// checks if the JournalLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JournalLine{}

// JournalLine struct for JournalLine
type JournalLine struct {
	// Xero identifier for Journal
	JournalLineID *string `json:"JournalLineID,omitempty"`
	// See Accounts
	AccountID *string `json:"AccountID,omitempty"`
	// See Accounts
	AccountCode *string `json:"AccountCode,omitempty"`
	AccountType *AccountType `json:"AccountType,omitempty"`
	// See AccountCodes
	AccountName *string `json:"AccountName,omitempty"`
	// The description from the source transaction line item. Only returned if populated.
	Description *string `json:"Description,omitempty"`
	// Net amount of journal line. This will be a positive value for a debit and negative for a credit
	NetAmount *float64 `json:"NetAmount,omitempty"`
	// Gross amount of journal line (NetAmount + TaxAmount).
	GrossAmount *float64 `json:"GrossAmount,omitempty"`
	// Total tax on a journal line
	TaxAmount *float64 `json:"TaxAmount,omitempty"`
	// The tax type from taxRates
	TaxType *string `json:"TaxType,omitempty"`
	// see TaxRates
	TaxName *string `json:"TaxName,omitempty"`
	// Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 <TrackingCategory> elements.
	TrackingCategories []TrackingCategory `json:"TrackingCategories,omitempty"`
}

// NewJournalLine instantiates a new JournalLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalLine() *JournalLine {
	this := JournalLine{}
	return &this
}

// NewJournalLineWithDefaults instantiates a new JournalLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalLineWithDefaults() *JournalLine {
	this := JournalLine{}
	return &this
}

// GetJournalLineID returns the JournalLineID field value if set, zero value otherwise.
func (o *JournalLine) GetJournalLineID() string {
	if o == nil || IsNil(o.JournalLineID) {
		var ret string
		return ret
	}
	return *o.JournalLineID
}

// GetJournalLineIDOk returns a tuple with the JournalLineID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetJournalLineIDOk() (*string, bool) {
	if o == nil || IsNil(o.JournalLineID) {
		return nil, false
	}
	return o.JournalLineID, true
}

// HasJournalLineID returns a boolean if a field has been set.
func (o *JournalLine) HasJournalLineID() bool {
	if o != nil && !IsNil(o.JournalLineID) {
		return true
	}

	return false
}

// SetJournalLineID gets a reference to the given string and assigns it to the JournalLineID field.
func (o *JournalLine) SetJournalLineID(v string) {
	o.JournalLineID = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *JournalLine) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *JournalLine) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *JournalLine) SetAccountID(v string) {
	o.AccountID = &v
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *JournalLine) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCode returns a boolean if a field has been set.
func (o *JournalLine) HasAccountCode() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *JournalLine) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *JournalLine) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *JournalLine) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *JournalLine) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *JournalLine) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *JournalLine) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *JournalLine) SetAccountName(v string) {
	o.AccountName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JournalLine) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JournalLine) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JournalLine) SetDescription(v string) {
	o.Description = &v
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise.
func (o *JournalLine) GetNetAmount() float64 {
	if o == nil || IsNil(o.NetAmount) {
		var ret float64
		return ret
	}
	return *o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetNetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.NetAmount) {
		return nil, false
	}
	return o.NetAmount, true
}

// HasNetAmount returns a boolean if a field has been set.
func (o *JournalLine) HasNetAmount() bool {
	if o != nil && !IsNil(o.NetAmount) {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given float64 and assigns it to the NetAmount field.
func (o *JournalLine) SetNetAmount(v float64) {
	o.NetAmount = &v
}

// GetGrossAmount returns the GrossAmount field value if set, zero value otherwise.
func (o *JournalLine) GetGrossAmount() float64 {
	if o == nil || IsNil(o.GrossAmount) {
		var ret float64
		return ret
	}
	return *o.GrossAmount
}

// GetGrossAmountOk returns a tuple with the GrossAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetGrossAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossAmount) {
		return nil, false
	}
	return o.GrossAmount, true
}

// HasGrossAmount returns a boolean if a field has been set.
func (o *JournalLine) HasGrossAmount() bool {
	if o != nil && !IsNil(o.GrossAmount) {
		return true
	}

	return false
}

// SetGrossAmount gets a reference to the given float64 and assigns it to the GrossAmount field.
func (o *JournalLine) SetGrossAmount(v float64) {
	o.GrossAmount = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *JournalLine) GetTaxAmount() float64 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float64
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetTaxAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *JournalLine) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float64 and assigns it to the TaxAmount field.
func (o *JournalLine) SetTaxAmount(v float64) {
	o.TaxAmount = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *JournalLine) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *JournalLine) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *JournalLine) SetTaxType(v string) {
	o.TaxType = &v
}

// GetTaxName returns the TaxName field value if set, zero value otherwise.
func (o *JournalLine) GetTaxName() string {
	if o == nil || IsNil(o.TaxName) {
		var ret string
		return ret
	}
	return *o.TaxName
}

// GetTaxNameOk returns a tuple with the TaxName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetTaxNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaxName) {
		return nil, false
	}
	return o.TaxName, true
}

// HasTaxName returns a boolean if a field has been set.
func (o *JournalLine) HasTaxName() bool {
	if o != nil && !IsNil(o.TaxName) {
		return true
	}

	return false
}

// SetTaxName gets a reference to the given string and assigns it to the TaxName field.
func (o *JournalLine) SetTaxName(v string) {
	o.TaxName = &v
}

// GetTrackingCategories returns the TrackingCategories field value if set, zero value otherwise.
func (o *JournalLine) GetTrackingCategories() []TrackingCategory {
	if o == nil || IsNil(o.TrackingCategories) {
		var ret []TrackingCategory
		return ret
	}
	return o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetTrackingCategoriesOk() ([]TrackingCategory, bool) {
	if o == nil || IsNil(o.TrackingCategories) {
		return nil, false
	}
	return o.TrackingCategories, true
}

// HasTrackingCategories returns a boolean if a field has been set.
func (o *JournalLine) HasTrackingCategories() bool {
	if o != nil && !IsNil(o.TrackingCategories) {
		return true
	}

	return false
}

// SetTrackingCategories gets a reference to the given []TrackingCategory and assigns it to the TrackingCategories field.
func (o *JournalLine) SetTrackingCategories(v []TrackingCategory) {
	o.TrackingCategories = v
}

func (o JournalLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JournalLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JournalLineID) {
		toSerialize["JournalLineID"] = o.JournalLineID
	}
	if !IsNil(o.AccountID) {
		toSerialize["AccountID"] = o.AccountID
	}
	if !IsNil(o.AccountCode) {
		toSerialize["AccountCode"] = o.AccountCode
	}
	if !IsNil(o.AccountType) {
		toSerialize["AccountType"] = o.AccountType
	}
	if !IsNil(o.AccountName) {
		toSerialize["AccountName"] = o.AccountName
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.NetAmount) {
		toSerialize["NetAmount"] = o.NetAmount
	}
	if !IsNil(o.GrossAmount) {
		toSerialize["GrossAmount"] = o.GrossAmount
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["TaxAmount"] = o.TaxAmount
	}
	if !IsNil(o.TaxType) {
		toSerialize["TaxType"] = o.TaxType
	}
	if !IsNil(o.TaxName) {
		toSerialize["TaxName"] = o.TaxName
	}
	if !IsNil(o.TrackingCategories) {
		toSerialize["TrackingCategories"] = o.TrackingCategories
	}
	return toSerialize, nil
}

type NullableJournalLine struct {
	value *JournalLine
	isSet bool
}

func (v NullableJournalLine) Get() *JournalLine {
	return v.value
}

func (v *NullableJournalLine) Set(val *JournalLine) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalLine) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalLine(val *JournalLine) *NullableJournalLine {
	return &NullableJournalLine{value: val, isSet: true}
}

func (v NullableJournalLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

