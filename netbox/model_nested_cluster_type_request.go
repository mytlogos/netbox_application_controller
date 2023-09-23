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

// checks if the NestedClusterTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedClusterTypeRequest{}

// NestedClusterTypeRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedClusterTypeRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	AdditionalProperties map[string]interface{}
}

type _NestedClusterTypeRequest NestedClusterTypeRequest

// NewNestedClusterTypeRequest instantiates a new NestedClusterTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedClusterTypeRequest(name string, slug string) *NestedClusterTypeRequest {
	this := NestedClusterTypeRequest{}
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedClusterTypeRequestWithDefaults instantiates a new NestedClusterTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterTypeRequestWithDefaults() *NestedClusterTypeRequest {
	this := NestedClusterTypeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedClusterTypeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedClusterTypeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedClusterTypeRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedClusterTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedClusterTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedClusterTypeRequest) SetSlug(v string) {
	o.Slug = v
}

func (o NestedClusterTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedClusterTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedClusterTypeRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedClusterTypeRequest := _NestedClusterTypeRequest{}

	if err = json.Unmarshal(bytes, &varNestedClusterTypeRequest); err == nil {
		*o = NestedClusterTypeRequest(varNestedClusterTypeRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedClusterTypeRequest struct {
	value *NestedClusterTypeRequest
	isSet bool
}

func (v NullableNestedClusterTypeRequest) Get() *NestedClusterTypeRequest {
	return v.value
}

func (v *NullableNestedClusterTypeRequest) Set(val *NestedClusterTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedClusterTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedClusterTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedClusterTypeRequest(val *NestedClusterTypeRequest) *NullableNestedClusterTypeRequest {
	return &NullableNestedClusterTypeRequest{value: val, isSet: true}
}

func (v NullableNestedClusterTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedClusterTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

