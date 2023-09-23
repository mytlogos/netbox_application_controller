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

// checks if the WritableDeviceBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableDeviceBayRequest{}

// WritableDeviceBayRequest Adds support for custom fields and tags.
type WritableDeviceBayRequest struct {
	Device int32 `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	InstalledDevice NullableInt32 `json:"installed_device,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableDeviceBayRequest WritableDeviceBayRequest

// NewWritableDeviceBayRequest instantiates a new WritableDeviceBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableDeviceBayRequest(device int32, name string) *WritableDeviceBayRequest {
	this := WritableDeviceBayRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewWritableDeviceBayRequestWithDefaults instantiates a new WritableDeviceBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableDeviceBayRequestWithDefaults() *WritableDeviceBayRequest {
	this := WritableDeviceBayRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *WritableDeviceBayRequest) GetDevice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *WritableDeviceBayRequest) SetDevice(v int32) {
	o.Device = v
}

// GetName returns the Name field value
func (o *WritableDeviceBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableDeviceBayRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableDeviceBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableDeviceBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableDeviceBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableDeviceBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableDeviceBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableDeviceBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInstalledDevice returns the InstalledDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableDeviceBayRequest) GetInstalledDevice() int32 {
	if o == nil || IsNil(o.InstalledDevice.Get()) {
		var ret int32
		return ret
	}
	return *o.InstalledDevice.Get()
}

// GetInstalledDeviceOk returns a tuple with the InstalledDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableDeviceBayRequest) GetInstalledDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledDevice.Get(), o.InstalledDevice.IsSet()
}

// HasInstalledDevice returns a boolean if a field has been set.
func (o *WritableDeviceBayRequest) HasInstalledDevice() bool {
	if o != nil && o.InstalledDevice.IsSet() {
		return true
	}

	return false
}

// SetInstalledDevice gets a reference to the given NullableInt32 and assigns it to the InstalledDevice field.
func (o *WritableDeviceBayRequest) SetInstalledDevice(v int32) {
	o.InstalledDevice.Set(&v)
}
// SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil
func (o *WritableDeviceBayRequest) SetInstalledDeviceNil() {
	o.InstalledDevice.Set(nil)
}

// UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
func (o *WritableDeviceBayRequest) UnsetInstalledDevice() {
	o.InstalledDevice.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableDeviceBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableDeviceBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableDeviceBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableDeviceBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableDeviceBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableDeviceBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableDeviceBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableDeviceBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.InstalledDevice.IsSet() {
		toSerialize["installed_device"] = o.InstalledDevice.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableDeviceBayRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableDeviceBayRequest := _WritableDeviceBayRequest{}

	if err = json.Unmarshal(bytes, &varWritableDeviceBayRequest); err == nil {
		*o = WritableDeviceBayRequest(varWritableDeviceBayRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "installed_device")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableDeviceBayRequest struct {
	value *WritableDeviceBayRequest
	isSet bool
}

func (v NullableWritableDeviceBayRequest) Get() *WritableDeviceBayRequest {
	return v.value
}

func (v *NullableWritableDeviceBayRequest) Set(val *WritableDeviceBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDeviceBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDeviceBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDeviceBayRequest(val *WritableDeviceBayRequest) *NullableWritableDeviceBayRequest {
	return &NullableWritableDeviceBayRequest{value: val, isSet: true}
}

func (v NullableWritableDeviceBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDeviceBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


