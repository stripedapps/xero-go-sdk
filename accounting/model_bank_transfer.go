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

// checks if the BankTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankTransfer{}

// BankTransfer struct for BankTransfer
type BankTransfer struct {
	FromBankAccount Account `json:"FromBankAccount"`
	ToBankAccount Account `json:"ToBankAccount"`
	// amount of the transaction
	Amount float64 `json:"Amount"`
	// The date of the Transfer YYYY-MM-DD
	Date *string `json:"Date,omitempty"`
	// The identifier of the Bank Transfer
	BankTransferID *string `json:"BankTransferID,omitempty"`
	// The currency rate
	CurrencyRate *float64 `json:"CurrencyRate,omitempty"`
	// The Bank Transaction ID for the source account
	FromBankTransactionID *string `json:"FromBankTransactionID,omitempty"`
	// The Bank Transaction ID for the destination account
	ToBankTransactionID *string `json:"ToBankTransactionID,omitempty"`
	// The Bank Transaction boolean to show if it is reconciled for the source account
	FromIsReconciled *bool `json:"FromIsReconciled,omitempty"`
	// The Bank Transaction boolean to show if it is reconciled for the destination account
	ToIsReconciled *bool `json:"ToIsReconciled,omitempty"`
	// Reference for the transactions.
	Reference *string `json:"Reference,omitempty"`
	// Boolean to indicate if a Bank Transfer has an attachment
	HasAttachmentsField *bool `json:"HasAttachments,omitempty"`
	// UTC timestamp of creation date of bank transfer
	CreatedDateUTC *string `json:"CreatedDateUTC,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}

// NewBankTransfer instantiates a new BankTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransfer(fromBankAccount Account, toBankAccount Account, amount float64) *BankTransfer {
	this := BankTransfer{}
	this.FromBankAccount = fromBankAccount
	this.ToBankAccount = toBankAccount
	this.Amount = amount
	var fromIsReconciled bool = false
	this.FromIsReconciled = &fromIsReconciled
	var toIsReconciled bool = false
	this.ToIsReconciled = &toIsReconciled
	return &this
}

// NewBankTransferWithDefaults instantiates a new BankTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferWithDefaults() *BankTransfer {
	this := BankTransfer{}
	var fromIsReconciled bool = false
	this.FromIsReconciled = &fromIsReconciled
	var toIsReconciled bool = false
	this.ToIsReconciled = &toIsReconciled
	return &this
}

// GetFromBankAccount returns the FromBankAccount field value
func (o *BankTransfer) GetFromBankAccount() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.FromBankAccount
}

// GetFromBankAccountOk returns a tuple with the FromBankAccount field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetFromBankAccountOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromBankAccount, true
}

// SetFromBankAccount sets field value
func (o *BankTransfer) SetFromBankAccount(v Account) {
	o.FromBankAccount = v
}

// GetToBankAccount returns the ToBankAccount field value
func (o *BankTransfer) GetToBankAccount() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.ToBankAccount
}

// GetToBankAccountOk returns a tuple with the ToBankAccount field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetToBankAccountOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToBankAccount, true
}

// SetToBankAccount sets field value
func (o *BankTransfer) SetToBankAccount(v Account) {
	o.ToBankAccount = v
}

// GetAmount returns the Amount field value
func (o *BankTransfer) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BankTransfer) SetAmount(v float64) {
	o.Amount = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BankTransfer) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BankTransfer) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *BankTransfer) SetDate(v string) {
	o.Date = &v
}

// GetBankTransferID returns the BankTransferID field value if set, zero value otherwise.
func (o *BankTransfer) GetBankTransferID() string {
	if o == nil || IsNil(o.BankTransferID) {
		var ret string
		return ret
	}
	return *o.BankTransferID
}

// GetBankTransferIDOk returns a tuple with the BankTransferID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetBankTransferIDOk() (*string, bool) {
	if o == nil || IsNil(o.BankTransferID) {
		return nil, false
	}
	return o.BankTransferID, true
}

// HasBankTransferID returns a boolean if a field has been set.
func (o *BankTransfer) HasBankTransferID() bool {
	if o != nil && !IsNil(o.BankTransferID) {
		return true
	}

	return false
}

// SetBankTransferID gets a reference to the given string and assigns it to the BankTransferID field.
func (o *BankTransfer) SetBankTransferID(v string) {
	o.BankTransferID = &v
}

// GetCurrencyRate returns the CurrencyRate field value if set, zero value otherwise.
func (o *BankTransfer) GetCurrencyRate() float64 {
	if o == nil || IsNil(o.CurrencyRate) {
		var ret float64
		return ret
	}
	return *o.CurrencyRate
}

// GetCurrencyRateOk returns a tuple with the CurrencyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetCurrencyRateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrencyRate) {
		return nil, false
	}
	return o.CurrencyRate, true
}

// HasCurrencyRate returns a boolean if a field has been set.
func (o *BankTransfer) HasCurrencyRate() bool {
	if o != nil && !IsNil(o.CurrencyRate) {
		return true
	}

	return false
}

// SetCurrencyRate gets a reference to the given float64 and assigns it to the CurrencyRate field.
func (o *BankTransfer) SetCurrencyRate(v float64) {
	o.CurrencyRate = &v
}

// GetFromBankTransactionID returns the FromBankTransactionID field value if set, zero value otherwise.
func (o *BankTransfer) GetFromBankTransactionID() string {
	if o == nil || IsNil(o.FromBankTransactionID) {
		var ret string
		return ret
	}
	return *o.FromBankTransactionID
}

// GetFromBankTransactionIDOk returns a tuple with the FromBankTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetFromBankTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.FromBankTransactionID) {
		return nil, false
	}
	return o.FromBankTransactionID, true
}

// HasFromBankTransactionID returns a boolean if a field has been set.
func (o *BankTransfer) HasFromBankTransactionID() bool {
	if o != nil && !IsNil(o.FromBankTransactionID) {
		return true
	}

	return false
}

// SetFromBankTransactionID gets a reference to the given string and assigns it to the FromBankTransactionID field.
func (o *BankTransfer) SetFromBankTransactionID(v string) {
	o.FromBankTransactionID = &v
}

// GetToBankTransactionID returns the ToBankTransactionID field value if set, zero value otherwise.
func (o *BankTransfer) GetToBankTransactionID() string {
	if o == nil || IsNil(o.ToBankTransactionID) {
		var ret string
		return ret
	}
	return *o.ToBankTransactionID
}

// GetToBankTransactionIDOk returns a tuple with the ToBankTransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetToBankTransactionIDOk() (*string, bool) {
	if o == nil || IsNil(o.ToBankTransactionID) {
		return nil, false
	}
	return o.ToBankTransactionID, true
}

// HasToBankTransactionID returns a boolean if a field has been set.
func (o *BankTransfer) HasToBankTransactionID() bool {
	if o != nil && !IsNil(o.ToBankTransactionID) {
		return true
	}

	return false
}

// SetToBankTransactionID gets a reference to the given string and assigns it to the ToBankTransactionID field.
func (o *BankTransfer) SetToBankTransactionID(v string) {
	o.ToBankTransactionID = &v
}

// GetFromIsReconciled returns the FromIsReconciled field value if set, zero value otherwise.
func (o *BankTransfer) GetFromIsReconciled() bool {
	if o == nil || IsNil(o.FromIsReconciled) {
		var ret bool
		return ret
	}
	return *o.FromIsReconciled
}

// GetFromIsReconciledOk returns a tuple with the FromIsReconciled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetFromIsReconciledOk() (*bool, bool) {
	if o == nil || IsNil(o.FromIsReconciled) {
		return nil, false
	}
	return o.FromIsReconciled, true
}

// HasFromIsReconciled returns a boolean if a field has been set.
func (o *BankTransfer) HasFromIsReconciled() bool {
	if o != nil && !IsNil(o.FromIsReconciled) {
		return true
	}

	return false
}

// SetFromIsReconciled gets a reference to the given bool and assigns it to the FromIsReconciled field.
func (o *BankTransfer) SetFromIsReconciled(v bool) {
	o.FromIsReconciled = &v
}

// GetToIsReconciled returns the ToIsReconciled field value if set, zero value otherwise.
func (o *BankTransfer) GetToIsReconciled() bool {
	if o == nil || IsNil(o.ToIsReconciled) {
		var ret bool
		return ret
	}
	return *o.ToIsReconciled
}

// GetToIsReconciledOk returns a tuple with the ToIsReconciled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetToIsReconciledOk() (*bool, bool) {
	if o == nil || IsNil(o.ToIsReconciled) {
		return nil, false
	}
	return o.ToIsReconciled, true
}

// HasToIsReconciled returns a boolean if a field has been set.
func (o *BankTransfer) HasToIsReconciled() bool {
	if o != nil && !IsNil(o.ToIsReconciled) {
		return true
	}

	return false
}

// SetToIsReconciled gets a reference to the given bool and assigns it to the ToIsReconciled field.
func (o *BankTransfer) SetToIsReconciled(v bool) {
	o.ToIsReconciled = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BankTransfer) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BankTransfer) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BankTransfer) SetReference(v string) {
	o.Reference = &v
}

// GetHasAttachmentsField returns the HasAttachmentsField field value if set, zero value otherwise.
func (o *BankTransfer) GetHasAttachmentsField() bool {
	if o == nil || IsNil(o.HasAttachmentsField) {
		var ret bool
		return ret
	}
	return *o.HasAttachmentsField
}

// GetHasAttachmentsFieldOk returns a tuple with the HasAttachmentsField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetHasAttachmentsFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAttachmentsField) {
		return nil, false
	}
	return o.HasAttachmentsField, true
}

// HasHasAttachmentsField returns a boolean if a field has been set.
func (o *BankTransfer) HasHasAttachmentsField() bool {
	if o != nil && !IsNil(o.HasAttachmentsField) {
		return true
	}

	return false
}

// SetHasAttachmentsField gets a reference to the given bool and assigns it to the HasAttachmentsField field.
func (o *BankTransfer) SetHasAttachmentsField(v bool) {
	o.HasAttachmentsField = &v
}

// GetCreatedDateUTC returns the CreatedDateUTC field value if set, zero value otherwise.
func (o *BankTransfer) GetCreatedDateUTC() string {
	if o == nil || IsNil(o.CreatedDateUTC) {
		var ret string
		return ret
	}
	return *o.CreatedDateUTC
}

// GetCreatedDateUTCOk returns a tuple with the CreatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetCreatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateUTC) {
		return nil, false
	}
	return o.CreatedDateUTC, true
}

// HasCreatedDateUTC returns a boolean if a field has been set.
func (o *BankTransfer) HasCreatedDateUTC() bool {
	if o != nil && !IsNil(o.CreatedDateUTC) {
		return true
	}

	return false
}

// SetCreatedDateUTC gets a reference to the given string and assigns it to the CreatedDateUTC field.
func (o *BankTransfer) SetCreatedDateUTC(v string) {
	o.CreatedDateUTC = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *BankTransfer) GetValidationErrors() []ValidationError {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []ValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransfer) GetValidationErrorsOk() ([]ValidationError, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *BankTransfer) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []ValidationError and assigns it to the ValidationErrors field.
func (o *BankTransfer) SetValidationErrors(v []ValidationError) {
	o.ValidationErrors = v
}

func (o BankTransfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FromBankAccount"] = o.FromBankAccount
	toSerialize["ToBankAccount"] = o.ToBankAccount
	toSerialize["Amount"] = o.Amount
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.BankTransferID) {
		toSerialize["BankTransferID"] = o.BankTransferID
	}
	if !IsNil(o.CurrencyRate) {
		toSerialize["CurrencyRate"] = o.CurrencyRate
	}
	if !IsNil(o.FromBankTransactionID) {
		toSerialize["FromBankTransactionID"] = o.FromBankTransactionID
	}
	if !IsNil(o.ToBankTransactionID) {
		toSerialize["ToBankTransactionID"] = o.ToBankTransactionID
	}
	if !IsNil(o.FromIsReconciled) {
		toSerialize["FromIsReconciled"] = o.FromIsReconciled
	}
	if !IsNil(o.ToIsReconciled) {
		toSerialize["ToIsReconciled"] = o.ToIsReconciled
	}
	if !IsNil(o.Reference) {
		toSerialize["Reference"] = o.Reference
	}
	if !IsNil(o.HasAttachmentsField) {
		toSerialize["HasAttachments"] = o.HasAttachmentsField
	}
	if !IsNil(o.CreatedDateUTC) {
		toSerialize["CreatedDateUTC"] = o.CreatedDateUTC
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["ValidationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableBankTransfer struct {
	value *BankTransfer
	isSet bool
}

func (v NullableBankTransfer) Get() *BankTransfer {
	return v.value
}

func (v *NullableBankTransfer) Set(val *BankTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransfer(val *BankTransfer) *NullableBankTransfer {
	return &NullableBankTransfer{value: val, isSet: true}
}

func (v NullableBankTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


