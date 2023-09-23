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

// checks if the DeviceTypeAirflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceTypeAirflow{}

// DeviceTypeAirflow struct for DeviceTypeAirflow
type DeviceTypeAirflow struct {
	// * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front * `left-to-right` - Left to right * `right-to-left` - Right to left * `side-to-rear` - Side to rear * `passive` - Passive * `mixed` - Mixed
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewDeviceTypeAirflow instantiates a new DeviceTypeAirflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceTypeAirflow() *DeviceTypeAirflow {
	this := DeviceTypeAirflow{}
	return &this
}

// NewDeviceTypeAirflowWithDefaults instantiates a new DeviceTypeAirflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeAirflowWithDefaults() *DeviceTypeAirflow {
	this := DeviceTypeAirflow{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceTypeAirflow) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeAirflow) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceTypeAirflow) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DeviceTypeAirflow) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceTypeAirflow) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceTypeAirflow) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceTypeAirflow) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DeviceTypeAirflow) SetLabel(v string) {
	o.Label = &v
}

func (o DeviceTypeAirflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceTypeAirflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableDeviceTypeAirflow struct {
	value *DeviceTypeAirflow
	isSet bool
}

func (v NullableDeviceTypeAirflow) Get() *DeviceTypeAirflow {
	return v.value
}

func (v *NullableDeviceTypeAirflow) Set(val *DeviceTypeAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeAirflow(val *DeviceTypeAirflow) *NullableDeviceTypeAirflow {
	return &NullableDeviceTypeAirflow{value: val, isSet: true}
}

func (v NullableDeviceTypeAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


