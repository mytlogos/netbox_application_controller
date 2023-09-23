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

// checks if the InventoryItemRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItemRoleRequest{}

// InventoryItemRoleRequest Adds support for custom fields and tags.
type InventoryItemRoleRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryItemRoleRequest InventoryItemRoleRequest

// NewInventoryItemRoleRequest instantiates a new InventoryItemRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItemRoleRequest(name string, slug string) *InventoryItemRoleRequest {
	this := InventoryItemRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewInventoryItemRoleRequestWithDefaults instantiates a new InventoryItemRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemRoleRequestWithDefaults() *InventoryItemRoleRequest {
	this := InventoryItemRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *InventoryItemRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryItemRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *InventoryItemRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *InventoryItemRoleRequest) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *InventoryItemRoleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *InventoryItemRoleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *InventoryItemRoleRequest) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InventoryItemRoleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InventoryItemRoleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InventoryItemRoleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InventoryItemRoleRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InventoryItemRoleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *InventoryItemRoleRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InventoryItemRoleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItemRoleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InventoryItemRoleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InventoryItemRoleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o InventoryItemRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItemRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InventoryItemRoleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryItemRoleRequest := _InventoryItemRoleRequest{}

	if err = json.Unmarshal(bytes, &varInventoryItemRoleRequest); err == nil {
		*o = InventoryItemRoleRequest(varInventoryItemRoleRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "color")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryItemRoleRequest struct {
	value *InventoryItemRoleRequest
	isSet bool
}

func (v NullableInventoryItemRoleRequest) Get() *InventoryItemRoleRequest {
	return v.value
}

func (v *NullableInventoryItemRoleRequest) Set(val *InventoryItemRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItemRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItemRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItemRoleRequest(val *InventoryItemRoleRequest) *NullableInventoryItemRoleRequest {
	return &NullableInventoryItemRoleRequest{value: val, isSet: true}
}

func (v NullableInventoryItemRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItemRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


