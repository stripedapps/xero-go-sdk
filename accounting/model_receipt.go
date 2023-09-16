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

// checks if the Receipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Receipt{}

// Receipt struct for Receipt
type Receipt struct {
	// Date of receipt – YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	Contact *Contact `json:"Contact,omitempty"`
	LineItems []LineItem `json:"LineItems,omitempty"`
	User *User `json:"User,omitempty"`
	// Additional reference number
	Reference *string `json:"Reference,omitempty"`
	LineAmountTypes *LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// Total of receipt excluding taxes
	SubTotal *float64 `json:"SubTotal,omitempty"`
	// Total tax on receipt
	TotalTax *float64 `json:"TotalTax,omitempty"`
	// Total of receipt tax inclusive (i.e. SubTotal + TotalTax)
	Total *float64 `json:"Total,omitempty"`
	// Xero generated unique identifier for receipt
	ReceiptID *string `json:"ReceiptID,omitempty"`
	// Current status of receipt – see status types
	Status *string `json:"Status,omitempty"`
	// Xero generated sequence number for receipt in current claim for a given user
	ReceiptNumber *string `json:"ReceiptNumber,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// boolean to indicate if a receipt has an attachment
	HasAttachments *bool `json:"HasAttachments,omitempty"`
	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	Url *string `json:"Url,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
	// Displays array of warning messages from the API
	Warnings []ValidationError `json:"Warnings,omitempty"`
	// Displays array of attachments from the API
	Attachments []Attachment `json:"Attachments,omitempty"`
}

// NewReceipt instantiates a new Receipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceipt() *Receipt {
	this := Receipt{}
	return &this
}

// NewReceiptWithDefaults instantiates a new Receipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptWithDefaults() *Receipt {
	this := Receipt{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Receipt) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Receipt) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Receipt) SetDate(v string) {
	o.Date = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Receipt) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Receipt) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *Receipt) SetContact(v Contact) {
	o.Contact = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Receipt) GetLineItems() []LineItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []LineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *Receipt) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []LineItem and assigns it to the LineItems field.
func (o *Receipt) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Receipt) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Receipt) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Receipt) SetUser(v User) {
	o.User = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Receipt) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Receipt) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Receipt) SetReference(v string) {
	o.Reference = &v
}

// GetLineAmountTypes returns the LineAmountTypes field value if set, zero value otherwise.
func (o *Receipt) GetLineAmountTypes() LineAmountTypes {
	if o == nil || IsNil(o.LineAmountTypes) {
		var ret LineAmountTypes
		return ret
	}
	return *o.LineAmountTypes
}

// GetLineAmountTypesOk returns a tuple with the LineAmountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetLineAmountTypesOk() (*LineAmountTypes, bool) {
	if o == nil || IsNil(o.LineAmountTypes) {
		return nil, false
	}
	return o.LineAmountTypes, true
}

// HasLineAmountTypes returns a boolean if a field has been set.
func (o *Receipt) HasLineAmountTypes() bool {
	if o != nil && !IsNil(o.LineAmountTypes) {
		return true
	}

	return false
}

// SetLineAmountTypes gets a reference to the given LineAmountTypes and assigns it to the LineAmountTypes field.
func (o *Receipt) SetLineAmountTypes(v LineAmountTypes) {
	o.LineAmountTypes = &v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *Receipt) GetSubTotal() float64 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float64
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotal returns a boolean if a field has been set.
func (o *Receipt) HasSubTotal() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float64 and assigns it to the SubTotal field.
func (o *Receipt) SetSubTotal(v float64) {
	o.SubTotal = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *Receipt) GetTotalTax() float64 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float64
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetTotalTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *Receipt) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float64 and assigns it to the TotalTax field.
func (o *Receipt) SetTotalTax(v float64) {
	o.TotalTax = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Receipt) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Receipt) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *Receipt) SetTotal(v float64) {
	o.Total = &v
}

// GetReceiptID returns the ReceiptID field value if set, zero value otherwise.
func (o *Receipt) GetReceiptID() string {
	if o == nil || IsNil(o.ReceiptID) {
		var ret string
		return ret
	}
	return *o.ReceiptID
}

// GetReceiptIDOk returns a tuple with the ReceiptID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptIDOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptID) {
		return nil, false
	}
	return o.ReceiptID, true
}

// HasReceiptID returns a boolean if a field has been set.
func (o *Receipt) HasReceiptID() bool {
	if o != nil && !IsNil(o.ReceiptID) {
		return true
	}

	return false
}

// SetReceiptID gets a reference to the given string and assigns it to the ReceiptID field.
func (o *Receipt) SetReceiptID(v string) {
	o.ReceiptID = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Receipt) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Receipt) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Receipt) SetStatus(v string) {
	o.Status = &v
}

// GetReceiptNumber returns the ReceiptNumber field value if set, zero value otherwise.
func (o *Receipt) GetReceiptNumber() string {
	if o == nil || IsNil(o.ReceiptNumber) {
		var ret string
		return ret
	}
	return *o.ReceiptNumber
}

// GetReceiptNumberOk returns a tuple with the ReceiptNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetReceiptNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptNumber) {
		return nil, false
	}
	return o.ReceiptNumber, true
}

// HasReceiptNumber returns a boolean if a field has been set.
func (o *Receipt) HasReceiptNumber() bool {
	if o != nil && !IsNil(o.ReceiptNumber) {
		return true
	}

	return false
}

// SetReceiptNumber gets a reference to the given string and assigns it to the ReceiptNumber field.
func (o *Receipt) SetReceiptNumber(v string) {
	o.ReceiptNumber = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *Receipt) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTC returns a boolean if a field has been set.
func (o *Receipt) HasUpdatedDateUTC() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *Receipt) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetHasAttachments returns the HasAttachments field value if set, zero value otherwise.
func (o *Receipt) GetHasAttachments() bool {
	if o == nil || IsNil(o.HasAttachments) {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetHasAttachmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachments) {
		return nil, false
	}
	return o.HasAttachments, true
}

// HasHasAttachments returns a boolean if a field has been set.
func (o *Receipt) HasHasAttachments() bool {
	if o != nil && !IsNil(o.HasAttachments) {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *Receipt) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Receipt) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Receipt) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Receipt) SetUrl(v string) {
	o.Url = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *Receipt) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *Receipt) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *Receipt) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Receipt) GetWarnings() []ValidationError {
	if o == nil || IsNil(o.Warnings) {
		var ret []ValidationError
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetWarningsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Receipt) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []ValidationError and assigns it to the Warnings field.
func (o *Receipt) SetWarnings(v []ValidationError) {
	o.Warnings = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Receipt) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receipt) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Receipt) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *Receipt) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o Receipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Receipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Contact) {
		toSerialize["Contact"] = o.Contact
	}
	if !IsNil(o.LineItems) {
		toSerialize["LineItems"] = o.LineItems
	}
	if !IsNil(o.User) {
		toSerialize["User"] = o.User
	}
	if !IsNil(o.Reference) {
		toSerialize["Reference"] = o.Reference
	}
	if !IsNil(o.LineAmountTypes) {
		toSerialize["LineAmountTypes"] = o.LineAmountTypes
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
	if !IsNil(o.ReceiptID) {
		toSerialize["ReceiptID"] = o.ReceiptID
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.ReceiptNumber) {
		toSerialize["ReceiptNumber"] = o.ReceiptNumber
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.HasAttachments) {
		toSerialize["HasAttachments"] = o.HasAttachments
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	if !IsNil(o.Warnings) {
		toSerialize["Warnings"] = o.Warnings
	}
	if !IsNil(o.Attachments) {
		toSerialize["Attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableReceipt struct {
	value *Receipt
	isSet bool
}

func (v NullableReceipt) Get() *Receipt {
	return v.value
}

func (v *NullableReceipt) Set(val *Receipt) {
	v.value = val
	v.isSet = true
}

func (v NullableReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceipt(val *Receipt) *NullableReceipt {
	return &NullableReceipt{value: val, isSet: true}
}

func (v NullableReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

