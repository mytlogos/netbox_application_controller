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

// checks if the NestedClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedClusterRequest{}

// NestedClusterRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedClusterRequest struct {
	Name string `json:"name"`
}

// NewNestedClusterRequest instantiates a new NestedClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedClusterRequest(name string) *NestedClusterRequest {
	this := NestedClusterRequest{}
	this.Name = name
	return &this
}

// NewNestedClusterRequestWithDefaults instantiates a new NestedClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedClusterRequestWithDefaults() *NestedClusterRequest {
	this := NestedClusterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedClusterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedClusterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedClusterRequest) SetName(v string) {
	o.Name = v
}

func (o NestedClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNestedClusterRequest struct {
	value *NestedClusterRequest
	isSet bool
}

func (v NullableNestedClusterRequest) Get() *NestedClusterRequest {
	return v.value
}

func (v *NullableNestedClusterRequest) Set(val *NestedClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedClusterRequest(val *NestedClusterRequest) *NullableNestedClusterRequest {
	return &NullableNestedClusterRequest{value: val, isSet: true}
}

func (v NullableNestedClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


