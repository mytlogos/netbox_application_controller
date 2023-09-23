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

// checks if the NestedContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedContactRequest{}

// NestedContactRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedContactRequest struct {
	Name string `json:"name"`
}

// NewNestedContactRequest instantiates a new NestedContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedContactRequest(name string) *NestedContactRequest {
	this := NestedContactRequest{}
	this.Name = name
	return &this
}

// NewNestedContactRequestWithDefaults instantiates a new NestedContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedContactRequestWithDefaults() *NestedContactRequest {
	this := NestedContactRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedContactRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedContactRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedContactRequest) SetName(v string) {
	o.Name = v
}

func (o NestedContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedContactRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNestedContactRequest struct {
	value *NestedContactRequest
	isSet bool
}

func (v NullableNestedContactRequest) Get() *NestedContactRequest {
	return v.value
}

func (v *NullableNestedContactRequest) Set(val *NestedContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedContactRequest(val *NestedContactRequest) *NullableNestedContactRequest {
	return &NullableNestedContactRequest{value: val, isSet: true}
}

func (v NullableNestedContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


