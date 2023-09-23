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

// checks if the NestedInventoryItemRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedInventoryItemRoleRequest{}

// NestedInventoryItemRoleRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedInventoryItemRoleRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedInventoryItemRoleRequest instantiates a new NestedInventoryItemRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedInventoryItemRoleRequest(name string, slug string) *NestedInventoryItemRoleRequest {
	this := NestedInventoryItemRoleRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedInventoryItemRoleRequestWithDefaults instantiates a new NestedInventoryItemRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedInventoryItemRoleRequestWithDefaults() *NestedInventoryItemRoleRequest {
	this := NestedInventoryItemRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedInventoryItemRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedInventoryItemRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedInventoryItemRoleRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedInventoryItemRoleRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedInventoryItemRoleRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedInventoryItemRoleRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedInventoryItemRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedInventoryItemRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedInventoryItemRoleRequest struct {
	value *NestedInventoryItemRoleRequest
	isSet bool
}

func (v NullableNestedInventoryItemRoleRequest) Get() *NestedInventoryItemRoleRequest {
	return v.value
}

func (v *NullableNestedInventoryItemRoleRequest) Set(val *NestedInventoryItemRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedInventoryItemRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedInventoryItemRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedInventoryItemRoleRequest(val *NestedInventoryItemRoleRequest) *NullableNestedInventoryItemRoleRequest {
	return &NullableNestedInventoryItemRoleRequest{value: val, isSet: true}
}

func (v NullableNestedInventoryItemRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedInventoryItemRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


