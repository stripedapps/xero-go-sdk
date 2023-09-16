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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// Customer defined alpha numeric account code e.g 200 or SALES (max length = 10)
	Code *string `json:"Code,omitempty"`
	// Name of account (max length = 150)
	Name *string `json:"Name,omitempty"`
	// The Xero identifier for an account – specified as a string following  the endpoint name   e.g. /297c2dc5-cc47-4afd-8ec8-74990b8761e9
	AccountID *string `json:"AccountID,omitempty"`
	Type *AccountType `json:"Type,omitempty"`
	// For bank accounts only (Account Type BANK)
	BankAccountNumber *string `json:"BankAccountNumber,omitempty"`
	// Accounts with a status of ACTIVE can be updated to ARCHIVED. See Account Status Codes
	Status *string `json:"Status,omitempty"`
	// Description of the Account. Valid for all types of accounts except bank accounts (max length = 4000)
	Description *string `json:"Description,omitempty"`
	// For bank accounts only. See Bank Account types
	BankAccountType *string `json:"BankAccountType,omitempty"`
	CurrencyCode *CurrencyCode `json:"CurrencyCode,omitempty"`
	// The tax type from taxRates
	TaxType *string `json:"TaxType,omitempty"`
	// Boolean – describes whether account can have payments applied to it
	EnablePaymentsToAccount *bool `json:"EnablePaymentsToAccount,omitempty"`
	// Boolean – describes whether account code is available for use with expense claims
	ShowInExpenseClaims *bool `json:"ShowInExpenseClaims,omitempty"`
	// See Account Class Types
	Class *string `json:"Class,omitempty"`
	// If this is a system account then this element is returned. See System Account types. Note that non-system accounts may have this element set as either “” or null.
	SystemAccount *string `json:"SystemAccount,omitempty"`
	// Shown if set
	ReportingCode *string `json:"ReportingCode,omitempty"`
	// Shown if set
	ReportingCodeName *string `json:"ReportingCodeName,omitempty"`
	// boolean to indicate if an account has an attachment (read only)
	HasAttachmentsField *bool `json:"HasAttachments,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// Boolean – describes whether the account is shown in the watchlist widget on the dashboard
	AddToWatchlist *bool `json:"AddToWatchlist,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Account) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Account) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Account) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Account) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Account) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Account) SetName(v string) {
	o.Name = &v
}

// GetAccountID returns the AccountID field value if set, zero value otherwise.
func (o *Account) GetAccountID() string {
	if o == nil || IsNil(o.AccountID) {
		var ret string
		return ret
	}
	return *o.AccountID
}

// GetAccountIDOk returns a tuple with the AccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountID) {
		return nil, false
	}
	return o.AccountID, true
}

// HasAccountID returns a boolean if a field has been set.
func (o *Account) HasAccountID() bool {
	if o != nil && !IsNil(o.AccountID) {
		return true
	}

	return false
}

// SetAccountID gets a reference to the given string and assigns it to the AccountID field.
func (o *Account) SetAccountID(v string) {
	o.AccountID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Account) GetType() AccountType {
	if o == nil || IsNil(o.Type) {
		var ret AccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Account) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AccountType and assigns it to the Type field.
func (o *Account) SetType(v AccountType) {
	o.Type = &v
}

// GetBankAccountNumber returns the BankAccountNumber field value if set, zero value otherwise.
func (o *Account) GetBankAccountNumber() string {
	if o == nil || IsNil(o.BankAccountNumber) {
		var ret string
		return ret
	}
	return *o.BankAccountNumber
}

// GetBankAccountNumberOk returns a tuple with the BankAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBankAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountNumber) {
		return nil, false
	}
	return o.BankAccountNumber, true
}

// HasBankAccountNumber returns a boolean if a field has been set.
func (o *Account) HasBankAccountNumber() bool {
	if o != nil && !IsNil(o.BankAccountNumber) {
		return true
	}

	return false
}

// SetBankAccountNumber gets a reference to the given string and assigns it to the BankAccountNumber field.
func (o *Account) SetBankAccountNumber(v string) {
	o.BankAccountNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Account) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Account) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Account) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Account) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Account) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Account) SetDescription(v string) {
	o.Description = &v
}

// GetBankAccountType returns the BankAccountType field value if set, zero value otherwise.
func (o *Account) GetBankAccountType() string {
	if o == nil || IsNil(o.BankAccountType) {
		var ret string
		return ret
	}
	return *o.BankAccountType
}

// GetBankAccountTypeOk returns a tuple with the BankAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetBankAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountType) {
		return nil, false
	}
	return o.BankAccountType, true
}

// HasBankAccountType returns a boolean if a field has been set.
func (o *Account) HasBankAccountType() bool {
	if o != nil && !IsNil(o.BankAccountType) {
		return true
	}

	return false
}

// SetBankAccountType gets a reference to the given string and assigns it to the BankAccountType field.
func (o *Account) SetBankAccountType(v string) {
	o.BankAccountType = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Account) GetCurrencyCode() CurrencyCode {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret CurrencyCode
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Account) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given CurrencyCode and assigns it to the CurrencyCode field.
func (o *Account) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *Account) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *Account) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *Account) SetTaxType(v string) {
	o.TaxType = &v
}

// GetEnablePaymentsToAccount returns the EnablePaymentsToAccount field value if set, zero value otherwise.
func (o *Account) GetEnablePaymentsToAccount() bool {
	if o == nil || IsNil(o.EnablePaymentsToAccount) {
		var ret bool
		return ret
	}
	return *o.EnablePaymentsToAccount
}

// GetEnablePaymentsToAccountOk returns a tuple with the EnablePaymentsToAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetEnablePaymentsToAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePaymentsToAccount) {
		return nil, false
	}
	return o.EnablePaymentsToAccount, true
}

// HasEnablePaymentsToAccount returns a boolean if a field has been set.
func (o *Account) HasEnablePaymentsToAccount() bool {
	if o != nil && !IsNil(o.EnablePaymentsToAccount) {
		return true
	}

	return false
}

// SetEnablePaymentsToAccount gets a reference to the given bool and assigns it to the EnablePaymentsToAccount field.
func (o *Account) SetEnablePaymentsToAccount(v bool) {
	o.EnablePaymentsToAccount = &v
}

// GetShowInExpenseClaims returns the ShowInExpenseClaims field value if set, zero value otherwise.
func (o *Account) GetShowInExpenseClaims() bool {
	if o == nil || IsNil(o.ShowInExpenseClaims) {
		var ret bool
		return ret
	}
	return *o.ShowInExpenseClaims
}

// GetShowInExpenseClaimsOk returns a tuple with the ShowInExpenseClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetShowInExpenseClaimsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowInExpenseClaims) {
		return nil, false
	}
	return o.ShowInExpenseClaims, true
}

// HasShowInExpenseClaims returns a boolean if a field has been set.
func (o *Account) HasShowInExpenseClaims() bool {
	if o != nil && !IsNil(o.ShowInExpenseClaims) {
		return true
	}

	return false
}

// SetShowInExpenseClaims gets a reference to the given bool and assigns it to the ShowInExpenseClaims field.
func (o *Account) SetShowInExpenseClaims(v bool) {
	o.ShowInExpenseClaims = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Account) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Account) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Account) SetClass(v string) {
	o.Class = &v
}

// GetSystemAccount returns the SystemAccount field value if set, zero value otherwise.
func (o *Account) GetSystemAccount() string {
	if o == nil || IsNil(o.SystemAccount) {
		var ret string
		return ret
	}
	return *o.SystemAccount
}

// GetSystemAccountOk returns a tuple with the SystemAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetSystemAccountOk() (*string, bool) {
	if o == nil || IsNil(o.SystemAccount) {
		return nil, false
	}
	return o.SystemAccount, true
}

// HasSystemAccount returns a boolean if a field has been set.
func (o *Account) HasSystemAccount() bool {
	if o != nil && !IsNil(o.SystemAccount) {
		return true
	}

	return false
}

// SetSystemAccount gets a reference to the given string and assigns it to the SystemAccount field.
func (o *Account) SetSystemAccount(v string) {
	o.SystemAccount = &v
}

// GetReportingCode returns the ReportingCode field value if set, zero value otherwise.
func (o *Account) GetReportingCode() string {
	if o == nil || IsNil(o.ReportingCode) {
		var ret string
		return ret
	}
	return *o.ReportingCode
}

// GetReportingCodeOk returns a tuple with the ReportingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetReportingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingCode) {
		return nil, false
	}
	return o.ReportingCode, true
}

// HasReportingCode returns a boolean if a field has been set.
func (o *Account) HasReportingCode() bool {
	if o != nil && !IsNil(o.ReportingCode) {
		return true
	}

	return false
}

// SetReportingCode gets a reference to the given string and assigns it to the ReportingCode field.
func (o *Account) SetReportingCode(v string) {
	o.ReportingCode = &v
}

// GetReportingCodeName returns the ReportingCodeName field value if set, zero value otherwise.
func (o *Account) GetReportingCodeName() string {
	if o == nil || IsNil(o.ReportingCodeName) {
		var ret string
		return ret
	}
	return *o.ReportingCodeName
}

// GetReportingCodeNameOk returns a tuple with the ReportingCodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetReportingCodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingCodeName) {
		return nil, false
	}
	return o.ReportingCodeName, true
}

// HasReportingCodeName returns a boolean if a field has been set.
func (o *Account) HasReportingCodeName() bool {
	if o != nil && !IsNil(o.ReportingCodeName) {
		return true
	}

	return false
}

// SetReportingCodeName gets a reference to the given string and assigns it to the ReportingCodeName field.
func (o *Account) SetReportingCodeName(v string) {
	o.ReportingCodeName = &v
}

// GetHasAttachmentsField returns the HasAttachmentsField field value if set, zero value otherwise.
func (o *Account) GetHasAttachmentsField() bool {
	if o == nil || IsNil(o.HasAttachmentsField) {
		var ret bool
		return ret
	}
	return *o.HasAttachmentsField
}

// GetHasAttachmentsFieldOk returns a tuple with the HasAttachmentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetHasAttachmentsFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachmentsField) {
		return nil, false
	}
	return o.HasAttachmentsField, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *Account) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachmentsField) {
		return true
	}

	return false
}

// SetHasAttachmentsField gets a reference to the given bool and assigns it to the HasAttachmentsField field.
func (o *Account) SetHasAttachmentsField(v bool) {
	o.HasAttachmentsField = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *Account) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTC returns a boolean if a field has been set.
func (o *Account) HasUpdatedDateUTC() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *Account) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetAddToWatchlist returns the AddToWatchlist field value if set, zero value otherwise.
func (o *Account) GetAddToWatchlist() bool {
	if o == nil || IsNil(o.AddToWatchlist) {
		var ret bool
		return ret
	}
	return *o.AddToWatchlist
}

// GetAddToWatchlistOk returns a tuple with the AddToWatchlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAddToWatchlistOk() (*bool, bool) {
	if o == nil || IsNil(o.AddToWatchlist) {
		return nil, false
	}
	return o.AddToWatchlist, true
}

// HasAddToWatchlist returns a boolean if a field has been set.
func (o *Account) HasAddToWatchlist() bool {
	if o != nil && !IsNil(o.AddToWatchlist) {
		return true
	}

	return false
}

// SetAddToWatchlist gets a reference to the given bool and assigns it to the AddToWatchlist field.
func (o *Account) SetAddToWatchlist(v bool) {
	o.AddToWatchlist = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *Account) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *Account) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *Account) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.AccountID) {
		toSerialize["AccountID"] = o.AccountID
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.BankAccountNumber) {
		toSerialize["BankAccountNumber"] = o.BankAccountNumber
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.BankAccountType) {
		toSerialize["BankAccountType"] = o.BankAccountType
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.TaxType) {
		toSerialize["TaxType"] = o.TaxType
	}
	if !IsNil(o.EnablePaymentsToAccount) {
		toSerialize["EnablePaymentsToAccount"] = o.EnablePaymentsToAccount
	}
	if !IsNil(o.ShowInExpenseClaims) {
		toSerialize["ShowInExpenseClaims"] = o.ShowInExpenseClaims
	}
	if !IsNil(o.Class) {
		toSerialize["Class"] = o.Class
	}
	if !IsNil(o.SystemAccount) {
		toSerialize["SystemAccount"] = o.SystemAccount
	}
	if !IsNil(o.ReportingCode) {
		toSerialize["ReportingCode"] = o.ReportingCode
	}
	if !IsNil(o.ReportingCodeName) {
		toSerialize["ReportingCodeName"] = o.ReportingCodeName
	}
	if !IsNil(o.HasAttachmentsField) {
		toSerialize["HasAttachments"] = o.HasAttachmentsField
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.AddToWatchlist) {
		toSerialize["AddToWatchlist"] = o.AddToWatchlist
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


