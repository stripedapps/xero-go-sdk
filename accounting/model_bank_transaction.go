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

// checks if the BankTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransaction{}

// BankTransaction struct for BankTransaction
type BankTransaction struct {
	// See Bank Transaction Types
	Type string `json:"Type"`
	Contact *Contact `json:"Contact,omitempty"`
	// See LineItems
	LineItems []LineItem `json:"LineItems"`
	BankAccount Account `json:"BankAccount"`
	// Boolean to show if transaction is reconciled
	IsReconciled *bool `json:"IsReconciled,omitempty"`
	// Date of transaction – YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	// Reference for the transaction. Only supported for SPEND and RECEIVE transactions.
	Reference *string `json:"Reference,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// Exchange rate to base currency when money is spent or received. e.g.0.7500 Only used for bank transactions in non base currency. If this isn’t specified for non base currency accounts then either the user-defined rate (preference) or the XE.com day rate will be used. Setting currency is only supported on overpayments.
	CurrencyRate *float64 `json:"CurrencyRate,omitempty"`
	// URL link to a source document – shown as “Go to App Name”
	Url *string `json:"Url,omitempty"`
	// See Bank Transaction Status Codes
	Status *string `json:"Status,omitempty"`
	LineAmountTypes *LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// Total of bank transaction excluding taxes
	SubTotal *float64 `json:"SubTotal,omitempty"`
	// Total tax on bank transaction
	TotalTax *float64 `json:"TotalTax,omitempty"`
	// Total of bank transaction tax inclusive
	Total *float64 `json:"Total,omitempty"`
	// Xero generated unique identifier for bank transaction
	BankTransactionID *string `json:"BankTransactionID,omitempty"`
	// Xero generated unique identifier for a Prepayment. This will be returned on BankTransactions with a Type of SPEND-PREPAYMENT or RECEIVE-PREPAYMENT
	PrepaymentID *string `json:"PrepaymentID,omitempty"`
	// Xero generated unique identifier for an Overpayment. This will be returned on BankTransactions with a Type of SPEND-OVERPAYMENT or RECEIVE-OVERPAYMENT
	OverpaymentID *string `json:"OverpaymentID,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// Boolean to indicate if a bank transaction has an attachment
	HasAttachmentsField *bool `json:"HasAttachments,omitempty"`
	// A string to indicate if a invoice status
	StatusAttributeString *string `json:"StatusAttributeString,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

// NewBankTransaction instantiates a new BankTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransaction(type_ string, lineItems []LineItem, bankAccount Account) *BankTransaction {
	this := BankTransaction{}
	this.Type = type_
	this.LineItems = lineItems
	this.BankAccount = bankAccount
	return &this
}

// NewBankTransactionWithDefaults instantiates a new BankTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransactionWithDefaults() *BankTransaction {
	this := BankTransaction{}
	return &this
}

// GetType returns the Type field value
func (o *BankTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BankTransaction) SetType(v string) {
	o.Type = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *BankTransaction) GetContact() Contact {
	if o == nil || IsNil(o.Contact) {
		var ret Contact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetContactOk() (*Contact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *BankTransaction) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given Contact and assigns it to the Contact field.
func (o *BankTransaction) SetContact(v Contact) {
	o.Contact = &v
}

// GetLineItems returns the LineItems field value
func (o *BankTransaction) GetLineItems() []LineItem {
	if o == nil {
		var ret []LineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *BankTransaction) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetBankAccount returns the BankAccount field value
func (o *BankTransaction) GetBankAccount() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetBankAccountOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *BankTransaction) SetBankAccount(v Account) {
	o.BankAccount = v
}

// GetIsReconciled returns the IsReconciled field value if set, zero value otherwise.
func (o *BankTransaction) GetIsReconciled() bool {
	if o == nil || IsNil(o.IsReconciled) {
		var ret bool
		return ret
	}
	return *o.IsReconciled
}

// GetIsReconciledOk returns a tuple with the IsReconciled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetIsReconciledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReconciled) {
		return nil, false
	}
	return o.IsReconciled, true
}

// HasIsReconciled returns a boolean if a field has been set.
func (o *BankTransaction) HasIsReconciled() bool {
	if o != nil && !IsNil(o.IsReconciled) {
		return true
	}

	return false
}

// SetIsReconciled gets a reference to the given bool and assigns it to the IsReconciled field.
func (o *BankTransaction) SetIsReconciled(v bool) {
	o.IsReconciled = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BankTransaction) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BankTransaction) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *BankTransaction) SetDate(v string) {
	o.Date = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BankTransaction) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BankTransaction) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BankTransaction) SetReference(v string) {
	o.Reference = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BankTransaction) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BankTransaction) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *BankTransaction) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *BankTransaction) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRate returns a boolean if a field has been set.
func (o *BankTransaction) HasCurrencyRate() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *BankTransaction) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BankTransaction) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BankTransaction) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BankTransaction) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BankTransaction) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BankTransaction) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BankTransaction) SetStatus(v string) {
	o.Status = &v
}

// GetLineAmountTypes returns the LineAmountTypes field value if set, zero value otherwise.
func (o *BankTransaction) GetLineAmountTypes() LineAmountTypes {
	if o == nil || IsNil(o.LineAmountTypes) {
		var ret LineAmountTypes
		return ret
	}
	return *o.LineAmountTypes
}

// GetLineAmountTypesOk returns a tuple with the LineAmountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetLineAmountTypesOk() (*LineAmountTypes, bool) {
	if o == nil || IsNil(o.LineAmountTypes) {
		return nil, false
	}
	return o.LineAmountTypes, true
}

// HasLineAmountTypes returns a boolean if a field has been set.
func (o *BankTransaction) HasLineAmountTypes() bool {
	if o != nil && !IsNil(o.LineAmountTypes) {
		return true
	}

	return false
}

// SetLineAmountTypes gets a reference to the given LineAmountTypes and assigns it to the LineAmountTypes field.
func (o *BankTransaction) SetLineAmountTypes(v LineAmountTypes) {
	o.LineAmountTypes = &v
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise.
func (o *BankTransaction) GetSubTotal() float64 {
	if o == nil || IsNil(o.SubTotal) {
		var ret float64
		return ret
	}
	return *o.SubTotal
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetSubTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.SubTotal) {
		return nil, false
	}
	return o.SubTotal, true
}

// HasSubTotal returns a boolean if a field has been set.
func (o *BankTransaction) HasSubTotal() bool {
	if o != nil && !IsNil(o.SubTotal) {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given float64 and assigns it to the SubTotal field.
func (o *BankTransaction) SetSubTotal(v float64) {
	o.SubTotal = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *BankTransaction) GetTotalTax() float64 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float64
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetTotalTaxOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *BankTransaction) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float64 and assigns it to the TotalTax field.
func (o *BankTransaction) SetTotalTax(v float64) {
	o.TotalTax = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BankTransaction) GetTotal() float64 {
	if o == nil || IsNil(o.Total) {
		var ret float64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetTotalOk() (*float64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BankTransaction) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float64 and assigns it to the Total field.
func (o *BankTransaction) SetTotal(v float64) {
	o.Total = &v
}

// GetBankTransactionID returns the BankTransactionID field value if set, zero value otherwise.
func (o *BankTransaction) GetBankTransactionID() string {
	if o == nil || IsNil(o.BankTransactionID) {
		var ret string
		return ret
	}
	return *o.BankTransactionID
}

// GetBankTransactionIDOk returns a tuple with the BankTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetBankTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.BankTransactionID) {
		return nil, false
	}
	return o.BankTransactionID, true
}

// HasBankTransactionID returns a boolean if a field has been set.
func (o *BankTransaction) HasBankTransactionID() bool {
	if o != nil && !IsNil(o.BankTransactionID) {
		return true
	}

	return false
}

// SetBankTransactionID gets a reference to the given string and assigns it to the BankTransactionID field.
func (o *BankTransaction) SetBankTransactionID(v string) {
	o.BankTransactionID = &v
}

// GetPrepaymentID returns the PrepaymentID field value if set, zero value otherwise.
func (o *BankTransaction) GetPrepaymentID() string {
	if o == nil || IsNil(o.PrepaymentID) {
		var ret string
		return ret
	}
	return *o.PrepaymentID
}

// GetPrepaymentIDOk returns a tuple with the PrepaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetPrepaymentIDOk() (*string, bool) {
	if o == nil || IsNil(o.PrepaymentID) {
		return nil, false
	}
	return o.PrepaymentID, true
}

// HasPrepaymentID returns a boolean if a field has been set.
func (o *BankTransaction) HasPrepaymentID() bool {
	if o != nil && !IsNil(o.PrepaymentID) {
		return true
	}

	return false
}

// SetPrepaymentID gets a reference to the given string and assigns it to the PrepaymentID field.
func (o *BankTransaction) SetPrepaymentID(v string) {
	o.PrepaymentID = &v
}

// GetOverpaymentID returns the OverpaymentID field value if set, zero value otherwise.
func (o *BankTransaction) GetOverpaymentID() string {
	if o == nil || IsNil(o.OverpaymentID) {
		var ret string
		return ret
	}
	return *o.OverpaymentID
}

// GetOverpaymentIDOk returns a tuple with the OverpaymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetOverpaymentIDOk() (*string, bool) {
	if o == nil || IsNil(o.OverpaymentID) {
		return nil, false
	}
	return o.OverpaymentID, true
}

// HasOverpaymentID returns a boolean if a field has been set.
func (o *BankTransaction) HasOverpaymentID() bool {
	if o != nil && !IsNil(o.OverpaymentID) {
		return true
	}

	return false
}

// SetOverpaymentID gets a reference to the given string and assigns it to the OverpaymentID field.
func (o *BankTransaction) SetOverpaymentID(v string) {
	o.OverpaymentID = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *BankTransaction) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTC returns a boolean if a field has been set.
func (o *BankTransaction) HasUpdatedDateUTC() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *BankTransaction) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetHasAttachmentsField returns the HasAttachmentsField field value if set, zero value otherwise.
func (o *BankTransaction) GetHasAttachmentsField() bool {
	if o == nil || IsNil(o.HasAttachmentsField) {
		var ret bool
		return ret
	}
	return *o.HasAttachmentsField
}

// GetHasAttachmentsFieldOk returns a tuple with the HasAttachmentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetHasAttachmentsFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachmentsField) {
		return nil, false
	}
	return o.HasAttachmentsField, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *BankTransaction) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachmentsField) {
		return true
	}

	return false
}

// SetHasAttachmentsField gets a reference to the given bool and assigns it to the HasAttachmentsField field.
func (o *BankTransaction) SetHasAttachmentsField(v bool) {
	o.HasAttachmentsField = &v
}

// GetStatusAttributeString returns the StatusAttributeString field value if set, zero value otherwise.
func (o *BankTransaction) GetStatusAttributeString() string {
	if o == nil || IsNil(o.StatusAttributeString) {
		var ret string
		return ret
	}
	return *o.StatusAttributeString
}

// GetStatusAttributeStringOk returns a tuple with the StatusAttributeString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetStatusAttributeStringOk() (*string, bool) {
	if o == nil || IsNil(o.StatusAttributeString) {
		return nil, false
	}
	return o.StatusAttributeString, true
}

// HasStatusAttributeString returns a boolean if a field has been set.
func (o *BankTransaction) HasStatusAttributeString() bool {
	if o != nil && !IsNil(o.StatusAttributeString) {
		return true
	}

	return false
}

// SetStatusAttributeString gets a reference to the given string and assigns it to the StatusAttributeString field.
func (o *BankTransaction) SetStatusAttributeString(v string) {
	o.StatusAttributeString = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *BankTransaction) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransaction) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *BankTransaction) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *BankTransaction) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o BankTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Type"] = o.Type
	if !IsNil(o.Contact) {
		toSerialize["Contact"] = o.Contact
	}
	toSerialize["LineItems"] = o.LineItems
	toSerialize["BankAccount"] = o.BankAccount
	if !IsNil(o.IsReconciled) {
		toSerialize["IsReconciled"] = o.IsReconciled
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Reference) {
		toSerialize["Reference"] = o.Reference
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["CurrencyRate"] = o.CurrencyRate
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
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
	if !IsNil(o.BankTransactionID) {
		toSerialize["BankTransactionID"] = o.BankTransactionID
	}
	if !IsNil(o.PrepaymentID) {
		toSerialize["PrepaymentID"] = o.PrepaymentID
	}
	if !IsNil(o.OverpaymentID) {
		toSerialize["OverpaymentID"] = o.OverpaymentID
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.HasAttachmentsField) {
		toSerialize["HasAttachments"] = o.HasAttachmentsField
	}
	if !IsNil(o.StatusAttributeString) {
		toSerialize["StatusAttributeString"] = o.StatusAttributeString
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableBankTransaction struct {
	value *BankTransaction
	isSet bool
}

func (v NullableBankTransaction) Get() *BankTransaction {
	return v.value
}

func (v *NullableBankTransaction) Set(val *BankTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransaction(val *BankTransaction) *NullableBankTransaction {
	return &NullableBankTransaction{value: val, isSet: true}
}

func (v NullableBankTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


