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

// checks if the WirelessLANAuthCipher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLANAuthCipher{}

// WirelessLANAuthCipher struct for WirelessLANAuthCipher
type WirelessLANAuthCipher struct {
	// * `auto` - Auto * `tkip` - TKIP * `aes` - AES
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewWirelessLANAuthCipher instantiates a new WirelessLANAuthCipher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLANAuthCipher() *WirelessLANAuthCipher {
	this := WirelessLANAuthCipher{}
	return &this
}

// NewWirelessLANAuthCipherWithDefaults instantiates a new WirelessLANAuthCipher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLANAuthCipherWithDefaults() *WirelessLANAuthCipher {
	this := WirelessLANAuthCipher{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WirelessLANAuthCipher) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLANAuthCipher) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WirelessLANAuthCipher) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WirelessLANAuthCipher) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WirelessLANAuthCipher) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLANAuthCipher) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WirelessLANAuthCipher) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WirelessLANAuthCipher) SetLabel(v string) {
	o.Label = &v
}

func (o WirelessLANAuthCipher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLANAuthCipher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableWirelessLANAuthCipher struct {
	value *WirelessLANAuthCipher
	isSet bool
}

func (v NullableWirelessLANAuthCipher) Get() *WirelessLANAuthCipher {
	return v.value
}

func (v *NullableWirelessLANAuthCipher) Set(val *WirelessLANAuthCipher) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLANAuthCipher) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLANAuthCipher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLANAuthCipher(val *WirelessLANAuthCipher) *NullableWirelessLANAuthCipher {
	return &NullableWirelessLANAuthCipher{value: val, isSet: true}
}

func (v NullableWirelessLANAuthCipher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLANAuthCipher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


