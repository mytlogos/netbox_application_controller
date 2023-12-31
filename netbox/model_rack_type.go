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

// checks if the RackType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackType{}

// RackType struct for RackType
type RackType struct {
	// * `2-post-frame` - 2-post frame * `4-post-frame` - 4-post frame * `4-post-cabinet` - 4-post cabinet * `wall-frame` - Wall-mounted frame * `wall-frame-vertical` - Wall-mounted frame (vertical) * `wall-cabinet` - Wall-mounted cabinet * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewRackType instantiates a new RackType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackType() *RackType {
	this := RackType{}
	return &this
}

// NewRackTypeWithDefaults instantiates a new RackType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackTypeWithDefaults() *RackType {
	this := RackType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackType) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RackType) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackType) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackType) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RackType) SetLabel(v string) {
	o.Label = &v
}

func (o RackType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableRackType struct {
	value *RackType
	isSet bool
}

func (v NullableRackType) Get() *RackType {
	return v.value
}

func (v *NullableRackType) Set(val *RackType) {
	v.value = val
	v.isSet = true
}

func (v NullableRackType) IsSet() bool {
	return v.isSet
}

func (v *NullableRackType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackType(val *RackType) *NullableRackType {
	return &NullableRackType{value: val, isSet: true}
}

func (v NullableRackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


