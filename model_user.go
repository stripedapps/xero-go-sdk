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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	// Xero identifier
	UserID *string `json:"UserID,omitempty"`
	// Email address of user
	EmailAddress *string `json:"EmailAddress,omitempty"`
	// First name of user
	FirstName *string `json:"FirstName,omitempty"`
	// Last name of user
	LastName *string `json:"LastName,omitempty"`
	// Timestamp of last change to user
	UpdatedDateUTC *string `json:"UpdatedDateUTC,omitempty"`
	// Boolean to indicate if user is the subscriber
	IsSubscriber *bool `json:"IsSubscriber,omitempty"`
	// User role that defines permissions in Xero and via API (READONLY, INVOICEONLY, STANDARD, FINANCIALADVISER, etc)
	OrganisationRole *string `json:"OrganisationRole,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *User) GetUserID() string {
	if o == nil || IsNil(o.UserID) {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.UserID) {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *User) HasUserID() bool {
	if o != nil && !IsNil(o.UserID) {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *User) SetUserID(v string) {
	o.UserID = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *User) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *User) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *User) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *User) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName = &v
}

// GetUpdatedDateUTC returns the UpdatedDateUTC field value if set, zero value otherwise.
func (o *User) GetUpdatedDateUTC() string {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		var ret string
		return ret
	}
	return *o.UpdatedDateUTC
}

// GetUpdatedDateUTCOk returns a tuple with the UpdatedDateUTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedDateUTCOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedDateUTC) {
		return nil, false
	}
	return o.UpdatedDateUTC, true
}

// HasUpdatedDateUTC returns a boolean if a field has been set.
func (o *User) HasUpdatedDateUTC() bool {
	if o != nil && !IsNil(o.UpdatedDateUTC) {
		return true
	}

	return false
}

// SetUpdatedDateUTC gets a reference to the given string and assigns it to the UpdatedDateUTC field.
func (o *User) SetUpdatedDateUTC(v string) {
	o.UpdatedDateUTC = &v
}

// GetIsSubscriber returns the IsSubscriber field value if set, zero value otherwise.
func (o *User) GetIsSubscriber() bool {
	if o == nil || IsNil(o.IsSubscriber) {
		var ret bool
		return ret
	}
	return *o.IsSubscriber
}

// GetIsSubscriberOk returns a tuple with the IsSubscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsSubscriberOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscriber) {
		return nil, false
	}
	return o.IsSubscriber, true
}

// HasIsSubscriber returns a boolean if a field has been set.
func (o *User) HasIsSubscriber() bool {
	if o != nil && !IsNil(o.IsSubscriber) {
		return true
	}

	return false
}

// SetIsSubscriber gets a reference to the given bool and assigns it to the IsSubscriber field.
func (o *User) SetIsSubscriber(v bool) {
	o.IsSubscriber = &v
}

// GetOrganisationRole returns the OrganisationRole field value if set, zero value otherwise.
func (o *User) GetOrganisationRole() string {
	if o == nil || IsNil(o.OrganisationRole) {
		var ret string
		return ret
	}
	return *o.OrganisationRole
}

// GetOrganisationRoleOk returns a tuple with the OrganisationRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOrganisationRoleOk() (*string, bool) {
	if o == nil || IsNil(o.OrganisationRole) {
		return nil, false
	}
	return o.OrganisationRole, true
}

// HasOrganisationRole returns a boolean if a field has been set.
func (o *User) HasOrganisationRole() bool {
	if o != nil && !IsNil(o.OrganisationRole) {
		return true
	}

	return false
}

// SetOrganisationRole gets a reference to the given string and assigns it to the OrganisationRole field.
func (o *User) SetOrganisationRole(v string) {
	o.OrganisationRole = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserID) {
		toSerialize["UserID"] = o.UserID
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if !IsNil(o.FirstName) {
		toSerialize["FirstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["LastName"] = o.LastName
	}
	if !IsNil(o.UpdatedDateUTC) {
		toSerialize["UpdatedDateUTC"] = o.UpdatedDateUTC
	}
	if !IsNil(o.IsSubscriber) {
		toSerialize["IsSubscriber"] = o.IsSubscriber
	}
	if !IsNil(o.OrganisationRole) {
		toSerialize["OrganisationRole"] = o.OrganisationRole
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


