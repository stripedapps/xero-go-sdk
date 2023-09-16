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

// checks if the Overpayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Overpayment{}

// Overpayment struct for Overpayment
type Overpayment struct {
	// See Overpayment Types
	Type *string `json:"Type,omitempty"`
	Contact *Contact `json:"Contact,omitempty"`
	// The date the overpayment is created YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	// See Overpayment Status Codes
	Status *string `json:"Status,omitempty"`
	LineAmountTypes *LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// See Overpayment Line Items
	LineItems []LineItem `json:"LineItems,omitempty"`
	// The subtotal of the overpayment excluding taxes
	SubTotal *float64 `json:"SubTotal,omitempty"`
	// The total tax on the overpayment
	TotalTax *float64 `json:"TotalTax,omitempty"`
	// The total of the overpayment (subtotal + total tax)
	Total *float64 `json:"Total,omitempty"`
	// UTC timestamp of last update to the overpayment
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// Xero generated unique identifier
	OverpaymentID *string `json:"OverpaymentID,omitempty"`
	// The currency rate for a multicurrency overpayment. If no rate is specified, the XE.com day rate is used
	CurrencyRate *float64 `json:"CurrencyRate,omitempty"`
	// The remaining credit balance on the overpayment
	RemainingCredit *float64 `json:"RemainingCredit,omitempty"`
	// See Allocations
	Allocations []Allocation `json:"Allocations,omitempty"`
	// The amount of applied to an invoice
	AppliedAmount *float64 `json:"AppliedAmount,omitempty"`
	// See Payments
	Payments []Payment `json:"Payments,omitempty"`
	// boolean to indicate if a overpayment has an attachment
	HasAttachmentsField *bool `json:"HasAttachments,omitempty"`
	// See Attachments
	Attachments []Attachment `json:"Attachments,omitempty"`
}

// NewOverpayment instantiates a new Overpayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOverpayment() *Overpayment {
	this := Overpayment{}
	return &this
}

// NewOverpaymentWithDefaults instantiates a new Overpayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOverpaymentWithDefaults() *Overpayment {
	this := Overpayment{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Overpayment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Overpayment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Overpayment) SetType(v string) {
	o.Type = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Overpayment) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Overpayment) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *Overpayment) SetContact(v Contact) {
	o.Contact = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Overpayment) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Overpayment) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Overpayment) SetDate(v string) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Overpayment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Overpayment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Overpayment) SetStatus(v string) {
	o.Status = &v
}

// GetLineAmountTypes returns the LineAmountTypes field value if set, zero value otherwise.
func (o *Overpayment) GetLineAmountTypes() LineAmountTypes {
	if o == nil || IsNil(o.LineAmountTypes) {
		var ret LineAmountTypes
		return ret
	}
	return *o.LineAmountTypes
}

// GetLineAmountTypesOk returns a tuple with the LineAmountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetLineAmountTypesOk() (*LineAmountTypes, bool) {
	if o == nil || IsNil(o.LineAmountTypes) {
		return nil, false
	}
	return o.LineAmountTypes, true
}

// HasLineAmountTypes returns a boolean if a field has been set.
func (o *Overpayment) HasLineAmountTypes() bool {
	if o != nil && !IsNil(o.LineAmountTypes) {
		return true
	}

	return false
}

// SetLineAmountTypes gets a reference to the given LineAmountTypes and assigns it to the LineAmountTypes field.
func (o *Overpayment) SetLineAmountTypes(v LineAmountTypes) {
	o.LineAmountTypes = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Overpayment) GetLineItems() []LineItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *Overpayment) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *Overpayment) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *Overpayment) GetSubTotal() float64 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float64
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotal returns a boolean if a field has been set.
func (o *Overpayment) HasSubTotal() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float64 and assigns it to the SubTotal field.
func (o *Overpayment) SetSubTotal(v float64) {
	o.SubTotal = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *Overpayment) GetTotalTax() float64 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float64
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetTotalTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *Overpayment) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float64 and assigns it to the TotalTax field.
func (o *Overpayment) SetTotalTax(v float64) {
	o.TotalTax = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Overpayment) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Overpayment) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *Overpayment) SetTotal(v float64) {
	o.Total = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *Overpayment) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTC returns a boolean if a field has been set.
func (o *Overpayment) HasUpdatedDateUTC() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *Overpayment) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Overpayment) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Overpayment) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *Overpayment) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetOverpaymentID returns the OverpaymentID field value if set, zero value otherwise.
func (o *Overpayment) GetOverpaymentID() string {
	if o == nil || IsNil(o.OverpaymentID) {
		var ret string
		return ret
	}
	return *o.OverpaymentID
}

// GetOverpaymentIDOk returns a tuple with the OverpaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetOverpaymentIDOk() (*string, bool) {
	if o == nil || IsNil(o.OverpaymentID) {
		return nil, false
	}
	return o.OverpaymentID, true
}

// HasOverpaymentID returns a boolean if a field has been set.
func (o *Overpayment) HasOverpaymentID() bool {
	if o != nil && !IsNil(o.OverpaymentID) {
		return true
	}

	return false
}

// SetOverpaymentID gets a reference to the given string and assigns it to the OverpaymentID field.
func (o *Overpayment) SetOverpaymentID(v string) {
	o.OverpaymentID = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *Overpayment) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRate returns a boolean if a field has been set.
func (o *Overpayment) HasCurrencyRate() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *Overpayment) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

// GetRemainingCredit returns the RemainingCredit field value if set, zero value otherwise.
func (o *Overpayment) GetRemainingCredit() float64 {
	if o == nil || IsNil(o.RemainingCredit) {
		var ret float64
		return ret
	}
	return *o.RemainingCredit
}

// GetRemainingCreditOk returns a tuple with the RemainingCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetRemainingCreditOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingCredit) {
		return nil, false
	}
	return o.RemainingCredit, true
}

// HasRemainingCredit returns a boolean if a field has been set.
func (o *Overpayment) HasRemainingCredit() bool {
	if o != nil && !IsNil(o.RemainingCredit) {
		return true
	}

	return false
}

// SetRemainingCredit gets a reference to the given float64 and assigns it to the RemainingCredit field.
func (o *Overpayment) SetRemainingCredit(v float64) {
	o.RemainingCredit = &v
}

// GetAllocations returns the Allocations field value if set, zero value otherwise.
func (o *Overpayment) GetAllocations() []Allocation {
	if o == nil || IsNil(o.Allocations) {
		var ret []Allocation
		return ret
	}
	return o.Allocations
}

// GetAllocationsOk returns a tuple with the Allocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetAllocationsOk() ([]Allocation, bool) {
	if o == nil || IsNil(o.Allocations) {
		return nil, false
	}
	return o.Allocations, true
}

// HasAllocations returns a boolean if a field has been set.
func (o *Overpayment) HasAllocations() bool {
	if o != nil && !IsNil(o.Allocations) {
		return true
	}

	return false
}

// SetAllocations gets a reference to the given []Allocation and assigns it to the Allocations field.
func (o *Overpayment) SetAllocations(v []Allocation) {
	o.Allocations = v
}

// GetAppliedAmount returns the AppliedAmount field value if set, zero value otherwise.
func (o *Overpayment) GetAppliedAmount() float64 {
	if o == nil || IsNil(o.AppliedAmount) {
		var ret float64
		return ret
	}
	return *o.AppliedAmount
}

// GetAppliedAmountOk returns a tuple with the AppliedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetAppliedAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.AppliedAmount) {
		return nil, false
	}
	return o.AppliedAmount, true
}

// HasAppliedAmount returns a boolean if a field has been set.
func (o *Overpayment) HasAppliedAmount() bool {
	if o != nil && !IsNil(o.AppliedAmount) {
		return true
	}

	return false
}

// SetAppliedAmount gets a reference to the given float64 and assigns it to the AppliedAmount field.
func (o *Overpayment) SetAppliedAmount(v float64) {
	o.AppliedAmount = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *Overpayment) GetPayments() []Payment {
	if o == nil || IsNil(o.Payments) {
		var ret []Payment
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetPaymentsOk() ([]Payment, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *Overpayment) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []Payment and assigns it to the Payments field.
func (o *Overpayment) SetPayments(v []Payment) {
	o.Payments = v
}

// GetHasAttachmentsField returns the HasAttachmentsField field value if set, zero value otherwise.
func (o *Overpayment) GetHasAttachmentsField() bool {
	if o == nil || IsNil(o.HasAttachmentsField) {
		var ret bool
		return ret
	}
	return *o.HasAttachmentsField
}

// GetHasAttachmentsFieldOk returns a tuple with the HasAttachmentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetHasAttachmentsFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachmentsField) {
		return nil, false
	}
	return o.HasAttachmentsField, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *Overpayment) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachmentsField) {
		return true
	}

	return false
}

// SetHasAttachmentsField gets a reference to the given bool and assigns it to the HasAttachmentsField field.
func (o *Overpayment) SetHasAttachmentsField(v bool) {
	o.HasAttachmentsField = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Overpayment) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Overpayment) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Overpayment) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *Overpayment) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o Overpayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Overpayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Contact) {
		toSerialize["Contact"] = o.Contact
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.LineAmountTypes) {
		toSerialize["LineAmountTypes"] = o.LineAmountTypes
	}
	if !IsNil(o.LineItems) {
		toSerialize["LineItems"] = o.LineItems
	}
	if !IsNil(o.SubTotal) {
		toSerialize["SubTotal"] = o.SubTotal
	}
	if !IsNil(o.TotalTax) {
		toSerialize["TotalTax"] = o.TotalTax
	}
	if !IsNil(o.Total) {
		toSerialize["Total"] = o.Total
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.OverpaymentID) {
		toSerialize["OverpaymentID"] = o.OverpaymentID
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["CurrencyRate"] = o.CurrencyRate
	}
	if !IsNil(o.RemainingCredit) {
		toSerialize["RemainingCredit"] = o.RemainingCredit
	}
	if !IsNil(o.Allocations) {
		toSerialize["Allocations"] = o.Allocations
	}
	if !IsNil(o.AppliedAmount) {
		toSerialize["AppliedAmount"] = o.AppliedAmount
	}
	if !IsNil(o.Payments) {
		toSerialize["Payments"] = o.Payments
	}
	if !IsNil(o.HasAttachmentsField) {
		toSerialize["HasAttachments"] = o.HasAttachmentsField
	}
	if !IsNil(o.Attachments) {
		toSerialize["Attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableOverpayment struct {
	value *Overpayment
	isSet bool
}

func (v NullableOverpayment) Get() *Overpayment {
	return v.value
}

func (v *NullableOverpayment) Set(val *Overpayment) {
	v.value = val
	v.isSet = true
}

func (v NullableOverpayment) IsSet() bool {
	return v.isSet
}

func (v *NullableOverpayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverpayment(val *Overpayment) *NullableOverpayment {
	return &NullableOverpayment{value: val, isSet: true}
}

func (v NullableOverpayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverpayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


