/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.1 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the WritableASNRangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableASNRangeRequest{}

// WritableASNRangeRequest Adds support for custom fields and tags.
type WritableASNRangeRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Rir int32 `json:"rir"`
	Start int64 `json:"start"`
	End int64 `json:"end"`
	Tenant NullableInt32 `json:"tenant,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableASNRangeRequest instantiates a new WritableASNRangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableASNRangeRequest(name string, slug string, rir int32, start int64, end int64) *WritableASNRangeRequest {
	this := WritableASNRangeRequest{}
	this.Name = name
	this.Slug = slug
	this.Rir = rir
	this.Start = start
	this.End = end
	return &this
}

// NewWritableASNRangeRequestWithDefaults instantiates a new WritableASNRangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableASNRangeRequestWithDefaults() *WritableASNRangeRequest {
	this := WritableASNRangeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WritableASNRangeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableASNRangeRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WritableASNRangeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableASNRangeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetRir returns the Rir field value
func (o *WritableASNRangeRequest) GetRir() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rir
}

// GetRirOk returns a tuple with the Rir field value
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetRirOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rir, true
}

// SetRir sets field value
func (o *WritableASNRangeRequest) SetRir(v int32) {
	o.Rir = v
}

// GetStart returns the Start field value
func (o *WritableASNRangeRequest) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *WritableASNRangeRequest) SetStart(v int64) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *WritableASNRangeRequest) GetEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *WritableASNRangeRequest) SetEnd(v int64) {
	o.End = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableASNRangeRequest) GetTenant() int32 {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret int32
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableASNRangeRequest) GetTenantOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WritableASNRangeRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableInt32 and assigns it to the Tenant field.
func (o *WritableASNRangeRequest) SetTenant(v int32) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WritableASNRangeRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WritableASNRangeRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableASNRangeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableASNRangeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableASNRangeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableASNRangeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableASNRangeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableASNRangeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableASNRangeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableASNRangeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableASNRangeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableASNRangeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableASNRangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableASNRangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["rir"] = o.Rir
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableWritableASNRangeRequest struct {
	value *WritableASNRangeRequest
	isSet bool
}

func (v NullableWritableASNRangeRequest) Get() *WritableASNRangeRequest {
	return v.value
}

func (v *NullableWritableASNRangeRequest) Set(val *WritableASNRangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableASNRangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableASNRangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableASNRangeRequest(val *WritableASNRangeRequest) *NullableWritableASNRangeRequest {
	return &NullableWritableASNRangeRequest{value: val, isSet: true}
}

func (v NullableWritableASNRangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableASNRangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


