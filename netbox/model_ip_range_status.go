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

// checks if the IPRangeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPRangeStatus{}

// IPRangeStatus struct for IPRangeStatus
type IPRangeStatus struct {
	// * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewIPRangeStatus instantiates a new IPRangeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPRangeStatus() *IPRangeStatus {
	this := IPRangeStatus{}
	return &this
}

// NewIPRangeStatusWithDefaults instantiates a new IPRangeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPRangeStatusWithDefaults() *IPRangeStatus {
	this := IPRangeStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IPRangeStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IPRangeStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IPRangeStatus) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IPRangeStatus) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPRangeStatus) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IPRangeStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IPRangeStatus) SetLabel(v string) {
	o.Label = &v
}

func (o IPRangeStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPRangeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableIPRangeStatus struct {
	value *IPRangeStatus
	isSet bool
}

func (v NullableIPRangeStatus) Get() *IPRangeStatus {
	return v.value
}

func (v *NullableIPRangeStatus) Set(val *IPRangeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIPRangeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIPRangeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPRangeStatus(val *IPRangeStatus) *NullableIPRangeStatus {
	return &NullableIPRangeStatus{value: val, isSet: true}
}

func (v NullableIPRangeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPRangeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


