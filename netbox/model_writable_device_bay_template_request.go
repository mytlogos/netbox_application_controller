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

// checks if the WritableDeviceBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableDeviceBayTemplateRequest{}

// WritableDeviceBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableDeviceBayTemplateRequest struct {
	DeviceType int32 `json:"device_type"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableDeviceBayTemplateRequest WritableDeviceBayTemplateRequest

// NewWritableDeviceBayTemplateRequest instantiates a new WritableDeviceBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableDeviceBayTemplateRequest(deviceType int32, name string) *WritableDeviceBayTemplateRequest {
	this := WritableDeviceBayTemplateRequest{}
	this.DeviceType = deviceType
	this.Name = name
	return &this
}

// NewWritableDeviceBayTemplateRequestWithDefaults instantiates a new WritableDeviceBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableDeviceBayTemplateRequestWithDefaults() *WritableDeviceBayTemplateRequest {
	this := WritableDeviceBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value
func (o *WritableDeviceBayTemplateRequest) GetDeviceType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *WritableDeviceBayTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType = v
}

// GetName returns the Name field value
func (o *WritableDeviceBayTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableDeviceBayTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableDeviceBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableDeviceBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableDeviceBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableDeviceBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableDeviceBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableDeviceBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o WritableDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableDeviceBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device_type"] = o.DeviceType
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableDeviceBayTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableDeviceBayTemplateRequest := _WritableDeviceBayTemplateRequest{}

	if err = json.Unmarshal(bytes, &varWritableDeviceBayTemplateRequest); err == nil {
		*o = WritableDeviceBayTemplateRequest(varWritableDeviceBayTemplateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableDeviceBayTemplateRequest struct {
	value *WritableDeviceBayTemplateRequest
	isSet bool
}

func (v NullableWritableDeviceBayTemplateRequest) Get() *WritableDeviceBayTemplateRequest {
	return v.value
}

func (v *NullableWritableDeviceBayTemplateRequest) Set(val *WritableDeviceBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDeviceBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDeviceBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDeviceBayTemplateRequest(val *WritableDeviceBayTemplateRequest) *NullableWritableDeviceBayTemplateRequest {
	return &NullableWritableDeviceBayTemplateRequest{value: val, isSet: true}
}

func (v NullableWritableDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDeviceBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


