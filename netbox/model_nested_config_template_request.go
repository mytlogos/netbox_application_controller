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

// checks if the NestedConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedConfigTemplateRequest{}

// NestedConfigTemplateRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedConfigTemplateRequest struct {
	Name string `json:"name"`
}

// NewNestedConfigTemplateRequest instantiates a new NestedConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedConfigTemplateRequest(name string) *NestedConfigTemplateRequest {
	this := NestedConfigTemplateRequest{}
	this.Name = name
	return &this
}

// NewNestedConfigTemplateRequestWithDefaults instantiates a new NestedConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedConfigTemplateRequestWithDefaults() *NestedConfigTemplateRequest {
	this := NestedConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedConfigTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedConfigTemplateRequest) SetName(v string) {
	o.Name = v
}

func (o NestedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNestedConfigTemplateRequest struct {
	value *NestedConfigTemplateRequest
	isSet bool
}

func (v NullableNestedConfigTemplateRequest) Get() *NestedConfigTemplateRequest {
	return v.value
}

func (v *NullableNestedConfigTemplateRequest) Set(val *NestedConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedConfigTemplateRequest(val *NestedConfigTemplateRequest) *NullableNestedConfigTemplateRequest {
	return &NullableNestedConfigTemplateRequest{value: val, isSet: true}
}

func (v NullableNestedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


