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

// checks if the NestedDataSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedDataSourceRequest{}

// NestedDataSourceRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedDataSourceRequest struct {
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedDataSourceRequest NestedDataSourceRequest

// NewNestedDataSourceRequest instantiates a new NestedDataSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDataSourceRequest(name string) *NestedDataSourceRequest {
	this := NestedDataSourceRequest{}
	this.Name = name
	return &this
}

// NewNestedDataSourceRequestWithDefaults instantiates a new NestedDataSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDataSourceRequestWithDefaults() *NestedDataSourceRequest {
	this := NestedDataSourceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedDataSourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedDataSourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedDataSourceRequest) SetName(v string) {
	o.Name = v
}

func (o NestedDataSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedDataSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedDataSourceRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedDataSourceRequest := _NestedDataSourceRequest{}

	if err = json.Unmarshal(bytes, &varNestedDataSourceRequest); err == nil {
		*o = NestedDataSourceRequest(varNestedDataSourceRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedDataSourceRequest struct {
	value *NestedDataSourceRequest
	isSet bool
}

func (v NullableNestedDataSourceRequest) Get() *NestedDataSourceRequest {
	return v.value
}

func (v *NullableNestedDataSourceRequest) Set(val *NestedDataSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDataSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDataSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDataSourceRequest(val *NestedDataSourceRequest) *NullableNestedDataSourceRequest {
	return &NullableNestedDataSourceRequest{value: val, isSet: true}
}

func (v NullableNestedDataSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDataSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

