/*
Xero OAuth 2 Identity Service API

These endpoints are related to managing authentication tokens and identity for Xero API

API version: 2.40.0
Contact: api@xero.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Connection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connection{}

// Connection struct for Connection
type Connection struct {
	// Xero identifier
	Id *string `json:"id,omitempty"`
	// Xero identifier of organisation
	TenantId *string `json:"tenantId,omitempty"`
	// Identifier shared across connections authorised at the same time
	AuthEventId *string `json:"authEventId,omitempty"`
	// Xero tenant type (i.e. ORGANISATION, PRACTICE)
	TenantType *string `json:"tenantType,omitempty"`
	// Xero tenant name
	TenantName *string `json:"tenantName,omitempty"`
	// The date when the user connected this tenant to your app
	CreatedDateUtc *time.Time `json:"createdDateUtc,omitempty"`
	// The date when the user most recently connected this tenant to your app. May differ to the created date if the user has disconnected and subsequently reconnected this tenant to your app.
	UpdatedDateUtc *time.Time `json:"updatedDateUtc,omitempty"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection() *Connection {
	this := Connection{}
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Connection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Connection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Connection) SetId(v string) {
	o.Id = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *Connection) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *Connection) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *Connection) SetTenantId(v string) {
	o.TenantId = &v
}

// GetAuthEventId returns the AuthEventId field value if set, zero value otherwise.
func (o *Connection) GetAuthEventId() string {
	if o == nil || IsNil(o.AuthEventId) {
		var ret string
		return ret
	}
	return *o.AuthEventId
}

// GetAuthEventIdOk returns a tuple with the AuthEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetAuthEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthEventId) {
		return nil, false
	}
	return o.AuthEventId, true
}

// HasAuthEventId returns a boolean if a field has been set.
func (o *Connection) HasAuthEventId() bool {
	if o != nil && !IsNil(o.AuthEventId) {
		return true
	}

	return false
}

// SetAuthEventId gets a reference to the given string and assigns it to the AuthEventId field.
func (o *Connection) SetAuthEventId(v string) {
	o.AuthEventId = &v
}

// GetTenantType returns the TenantType field value if set, zero value otherwise.
func (o *Connection) GetTenantType() string {
	if o == nil || IsNil(o.TenantType) {
		var ret string
		return ret
	}
	return *o.TenantType
}

// GetTenantTypeOk returns a tuple with the TenantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTenantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TenantType) {
		return nil, false
	}
	return o.TenantType, true
}

// HasTenantType returns a boolean if a field has been set.
func (o *Connection) HasTenantType() bool {
	if o != nil && !IsNil(o.TenantType) {
		return true
	}

	return false
}

// SetTenantType gets a reference to the given string and assigns it to the TenantType field.
func (o *Connection) SetTenantType(v string) {
	o.TenantType = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *Connection) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *Connection) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *Connection) SetTenantName(v string) {
	o.TenantName = &v
}

// GetCreatedDateUtc returns the CreatedDateUtc field value if set, zero value otherwise.
func (o *Connection) GetCreatedDateUtc() time.Time {
	if o == nil || IsNil(o.CreatedDateUtc) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateUtc
}

// GetCreatedDateUtcOk returns a tuple with the CreatedDateUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetCreatedDateUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDateUtc) {
		return nil, false
	}
	return o.CreatedDateUtc, true
}

// HasCreatedDateUtc returns a boolean if a field has been set.
func (o *Connection) HasCreatedDateUtc() bool {
	if o != nil && !IsNil(o.CreatedDateUtc) {
		return true
	}

	return false
}

// SetCreatedDateUtc gets a reference to the given time.Time and assigns it to the CreatedDateUtc field.
func (o *Connection) SetCreatedDateUtc(v time.Time) {
	o.CreatedDateUtc = &v
}

// GetUpdatedDateUtc returns the UpdatedDateUtc field value if set, zero value otherwise.
func (o *Connection) GetUpdatedDateUtc() time.Time {
	if o == nil || IsNil(o.UpdatedDateUtc) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDateUtc
}

// GetUpdatedDateUtcOk returns a tuple with the UpdatedDateUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetUpdatedDateUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDateUtc) {
		return nil, false
	}
	return o.UpdatedDateUtc, true
}

// HasUpdatedDateUtc returns a boolean if a field has been set.
func (o *Connection) HasUpdatedDateUtc() bool {
	if o != nil && !IsNil(o.UpdatedDateUtc) {
		return true
	}

	return false
}

// SetUpdatedDateUtc gets a reference to the given time.Time and assigns it to the UpdatedDateUtc field.
func (o *Connection) SetUpdatedDateUtc(v time.Time) {
	o.UpdatedDateUtc = &v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.AuthEventId) {
		toSerialize["authEventId"] = o.AuthEventId
	}
	if !IsNil(o.TenantType) {
		toSerialize["tenantType"] = o.TenantType
	}
	if !IsNil(o.TenantName) {
		toSerialize["tenantName"] = o.TenantName
	}
	if !IsNil(o.CreatedDateUtc) {
		toSerialize["createdDateUtc"] = o.CreatedDateUtc
	}
	if !IsNil(o.UpdatedDateUtc) {
		toSerialize["updatedDateUtc"] = o.UpdatedDateUtc
	}
	return toSerialize, nil
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

