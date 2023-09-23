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

// checks if the ModuleBayNestedModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleBayNestedModuleRequest{}

// ModuleBayNestedModuleRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type ModuleBayNestedModuleRequest struct {
	Serial *string `json:"serial,omitempty"`
}

// NewModuleBayNestedModuleRequest instantiates a new ModuleBayNestedModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBayNestedModuleRequest() *ModuleBayNestedModuleRequest {
	this := ModuleBayNestedModuleRequest{}
	return &this
}

// NewModuleBayNestedModuleRequestWithDefaults instantiates a new ModuleBayNestedModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBayNestedModuleRequestWithDefaults() *ModuleBayNestedModuleRequest {
	this := ModuleBayNestedModuleRequest{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ModuleBayNestedModuleRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayNestedModuleRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ModuleBayNestedModuleRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ModuleBayNestedModuleRequest) SetSerial(v string) {
	o.Serial = &v
}

func (o ModuleBayNestedModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleBayNestedModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableModuleBayNestedModuleRequest struct {
	value *ModuleBayNestedModuleRequest
	isSet bool
}

func (v NullableModuleBayNestedModuleRequest) Get() *ModuleBayNestedModuleRequest {
	return v.value
}

func (v *NullableModuleBayNestedModuleRequest) Set(val *ModuleBayNestedModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBayNestedModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBayNestedModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBayNestedModuleRequest(val *ModuleBayNestedModuleRequest) *NullableModuleBayNestedModuleRequest {
	return &NullableModuleBayNestedModuleRequest{value: val, isSet: true}
}

func (v NullableModuleBayNestedModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBayNestedModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


