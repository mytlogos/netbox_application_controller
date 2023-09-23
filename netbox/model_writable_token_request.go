/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.1 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the WritableTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableTokenRequest{}

// WritableTokenRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableTokenRequest struct {
	User int32 `json:"user"`
	Expires NullableTime `json:"expires,omitempty"`
	LastUsed NullableTime `json:"last_used,omitempty"`
	Key *string `json:"key,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled *bool `json:"write_enabled,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewWritableTokenRequest instantiates a new WritableTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableTokenRequest(user int32) *WritableTokenRequest {
	this := WritableTokenRequest{}
	this.User = user
	return &this
}

// NewWritableTokenRequestWithDefaults instantiates a new WritableTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableTokenRequestWithDefaults() *WritableTokenRequest {
	this := WritableTokenRequest{}
	return &this
}

// GetUser returns the User field value
func (o *WritableTokenRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *WritableTokenRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *WritableTokenRequest) SetUser(v int32) {
	o.User = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableTokenRequest) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableTokenRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *WritableTokenRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *WritableTokenRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *WritableTokenRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *WritableTokenRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableTokenRequest) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableTokenRequest) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// HasLastUsed returns a boolean if a field has been set.
func (o *WritableTokenRequest) HasLastUsed() bool {
	if o != nil && o.LastUsed.IsSet() {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given NullableTime and assigns it to the LastUsed field.
func (o *WritableTokenRequest) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}
// SetLastUsedNil sets the value for LastUsed to be an explicit nil
func (o *WritableTokenRequest) SetLastUsedNil() {
	o.LastUsed.Set(nil)
}

// UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
func (o *WritableTokenRequest) UnsetLastUsed() {
	o.LastUsed.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *WritableTokenRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTokenRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *WritableTokenRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *WritableTokenRequest) SetKey(v string) {
	o.Key = &v
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *WritableTokenRequest) GetWriteEnabled() bool {
	if o == nil || IsNil(o.WriteEnabled) {
		var ret bool
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTokenRequest) GetWriteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WriteEnabled) {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *WritableTokenRequest) HasWriteEnabled() bool {
	if o != nil && !IsNil(o.WriteEnabled) {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given bool and assigns it to the WriteEnabled field.
func (o *WritableTokenRequest) SetWriteEnabled(v bool) {
	o.WriteEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableTokenRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableTokenRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableTokenRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableTokenRequest) SetDescription(v string) {
	o.Description = &v
}

func (o WritableTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.LastUsed.IsSet() {
		toSerialize["last_used"] = o.LastUsed.Get()
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.WriteEnabled) {
		toSerialize["write_enabled"] = o.WriteEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableWritableTokenRequest struct {
	value *WritableTokenRequest
	isSet bool
}

func (v NullableWritableTokenRequest) Get() *WritableTokenRequest {
	return v.value
}

func (v *NullableWritableTokenRequest) Set(val *WritableTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableTokenRequest(val *WritableTokenRequest) *NullableWritableTokenRequest {
	return &NullableWritableTokenRequest{value: val, isSet: true}
}

func (v NullableWritableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


