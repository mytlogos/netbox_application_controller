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

// checks if the RackWidth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackWidth{}

// RackWidth struct for RackWidth
type RackWidth struct {
	// * `10` - 10 inches * `19` - 19 inches * `21` - 21 inches * `23` - 23 inches
	Value *int32 `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewRackWidth instantiates a new RackWidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackWidth() *RackWidth {
	this := RackWidth{}
	return &this
}

// NewRackWidthWithDefaults instantiates a new RackWidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackWidthWithDefaults() *RackWidth {
	this := RackWidth{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackWidth) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackWidth) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackWidth) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *RackWidth) SetValue(v int32) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackWidth) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackWidth) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackWidth) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *RackWidth) SetLabel(v string) {
	o.Label = &v
}

func (o RackWidth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackWidth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableRackWidth struct {
	value *RackWidth
	isSet bool
}

func (v NullableRackWidth) Get() *RackWidth {
	return v.value
}

func (v *NullableRackWidth) Set(val *RackWidth) {
	v.value = val
	v.isSet = true
}

func (v NullableRackWidth) IsSet() bool {
	return v.isSet
}

func (v *NullableRackWidth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackWidth(val *RackWidth) *NullableRackWidth {
	return &NullableRackWidth{value: val, isSet: true}
}

func (v NullableRackWidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackWidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


