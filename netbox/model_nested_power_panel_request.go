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

// checks if the NestedPowerPanelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedPowerPanelRequest{}

// NestedPowerPanelRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedPowerPanelRequest struct {
	Name string `json:"name"`
}

// NewNestedPowerPanelRequest instantiates a new NestedPowerPanelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedPowerPanelRequest(name string) *NestedPowerPanelRequest {
	this := NestedPowerPanelRequest{}
	this.Name = name
	return &this
}

// NewNestedPowerPanelRequestWithDefaults instantiates a new NestedPowerPanelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedPowerPanelRequestWithDefaults() *NestedPowerPanelRequest {
	this := NestedPowerPanelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NestedPowerPanelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedPowerPanelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedPowerPanelRequest) SetName(v string) {
	o.Name = v
}

func (o NestedPowerPanelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedPowerPanelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNestedPowerPanelRequest struct {
	value *NestedPowerPanelRequest
	isSet bool
}

func (v NullableNestedPowerPanelRequest) Get() *NestedPowerPanelRequest {
	return v.value
}

func (v *NullableNestedPowerPanelRequest) Set(val *NestedPowerPanelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedPowerPanelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedPowerPanelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedPowerPanelRequest(val *NestedPowerPanelRequest) *NullableNestedPowerPanelRequest {
	return &NullableNestedPowerPanelRequest{value: val, isSet: true}
}

func (v NullableNestedPowerPanelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedPowerPanelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


