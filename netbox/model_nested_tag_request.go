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

// checks if the NestedTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTagRequest{}

// NestedTagRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color *string `json:"color,omitempty"`
}

// NewNestedTagRequest instantiates a new NestedTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTagRequest(name string, slug string) *NestedTagRequest {
	this := NestedTagRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedTagRequestWithDefaults instantiates a new NestedTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTagRequestWithDefaults() *NestedTagRequest {
	this := NestedTagRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedTagRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTagRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTagRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTagRequest) SetSlug(v string) {
	o.Slug = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *NestedTagRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedTagRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *NestedTagRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *NestedTagRequest) SetColor(v string) {
	o.Color = &v
}

func (o NestedTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

type NullableNestedTagRequest struct {
	value *NestedTagRequest
	isSet bool
}

func (v NullableNestedTagRequest) Get() *NestedTagRequest {
	return v.value
}

func (v *NullableNestedTagRequest) Set(val *NestedTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTagRequest(val *NestedTagRequest) *NullableNestedTagRequest {
	return &NullableNestedTagRequest{value: val, isSet: true}
}

func (v NullableNestedTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


