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

// checks if the InterfaceTemplateRfRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceTemplateRfRole{}

// InterfaceTemplateRfRole struct for InterfaceTemplateRfRole
type InterfaceTemplateRfRole struct {
	// * `ap` - Access point * `station` - Station
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceTemplateRfRole InterfaceTemplateRfRole

// NewInterfaceTemplateRfRole instantiates a new InterfaceTemplateRfRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceTemplateRfRole() *InterfaceTemplateRfRole {
	this := InterfaceTemplateRfRole{}
	return &this
}

// NewInterfaceTemplateRfRoleWithDefaults instantiates a new InterfaceTemplateRfRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceTemplateRfRoleWithDefaults() *InterfaceTemplateRfRole {
	this := InterfaceTemplateRfRole{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InterfaceTemplateRfRole) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRfRole) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InterfaceTemplateRfRole) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InterfaceTemplateRfRole) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *InterfaceTemplateRfRole) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceTemplateRfRole) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *InterfaceTemplateRfRole) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *InterfaceTemplateRfRole) SetLabel(v string) {
	o.Label = &v
}

func (o InterfaceTemplateRfRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceTemplateRfRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceTemplateRfRole) UnmarshalJSON(bytes []byte) (err error) {
	varInterfaceTemplateRfRole := _InterfaceTemplateRfRole{}

	if err = json.Unmarshal(bytes, &varInterfaceTemplateRfRole); err == nil {
		*o = InterfaceTemplateRfRole(varInterfaceTemplateRfRole)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceTemplateRfRole struct {
	value *InterfaceTemplateRfRole
	isSet bool
}

func (v NullableInterfaceTemplateRfRole) Get() *InterfaceTemplateRfRole {
	return v.value
}

func (v *NullableInterfaceTemplateRfRole) Set(val *InterfaceTemplateRfRole) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplateRfRole) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplateRfRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplateRfRole(val *InterfaceTemplateRfRole) *NullableInterfaceTemplateRfRole {
	return &NullableInterfaceTemplateRfRole{value: val, isSet: true}
}

func (v NullableInterfaceTemplateRfRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplateRfRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


