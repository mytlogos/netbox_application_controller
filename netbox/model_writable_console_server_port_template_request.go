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

// checks if the WritableConsoleServerPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableConsoleServerPortTemplateRequest{}

// WritableConsoleServerPortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableConsoleServerPortTemplateRequest struct {
	DeviceType NullableInt32 `json:"device_type,omitempty"`
	ModuleType NullableInt32 `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// * `de-9` - DE-9 * `db-25` - DB-25 * `rj-11` - RJ-11 * `rj-12` - RJ-12 * `rj-45` - RJ-45 * `mini-din-8` - Mini-DIN 8 * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `other` - Other
	Type *string `json:"type,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableConsoleServerPortTemplateRequest WritableConsoleServerPortTemplateRequest

// NewWritableConsoleServerPortTemplateRequest instantiates a new WritableConsoleServerPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableConsoleServerPortTemplateRequest(name string) *WritableConsoleServerPortTemplateRequest {
	this := WritableConsoleServerPortTemplateRequest{}
	this.Name = name
	return &this
}

// NewWritableConsoleServerPortTemplateRequestWithDefaults instantiates a new WritableConsoleServerPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableConsoleServerPortTemplateRequestWithDefaults() *WritableConsoleServerPortTemplateRequest {
	this := WritableConsoleServerPortTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConsoleServerPortTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret int32
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConsoleServerPortTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableInt32 and assigns it to the DeviceType field.
func (o *WritableConsoleServerPortTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *WritableConsoleServerPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *WritableConsoleServerPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableConsoleServerPortTemplateRequest) GetModuleType() int32 {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret int32
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableConsoleServerPortTemplateRequest) GetModuleTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableInt32 and assigns it to the ModuleType field.
func (o *WritableConsoleServerPortTemplateRequest) SetModuleType(v int32) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *WritableConsoleServerPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *WritableConsoleServerPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *WritableConsoleServerPortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WritableConsoleServerPortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WritableConsoleServerPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplateRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplateRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WritableConsoleServerPortTemplateRequest) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableConsoleServerPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableConsoleServerPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableConsoleServerPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableConsoleServerPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o WritableConsoleServerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableConsoleServerPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableConsoleServerPortTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varWritableConsoleServerPortTemplateRequest := _WritableConsoleServerPortTemplateRequest{}

	if err = json.Unmarshal(bytes, &varWritableConsoleServerPortTemplateRequest); err == nil {
		*o = WritableConsoleServerPortTemplateRequest(varWritableConsoleServerPortTemplateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableConsoleServerPortTemplateRequest struct {
	value *WritableConsoleServerPortTemplateRequest
	isSet bool
}

func (v NullableWritableConsoleServerPortTemplateRequest) Get() *WritableConsoleServerPortTemplateRequest {
	return v.value
}

func (v *NullableWritableConsoleServerPortTemplateRequest) Set(val *WritableConsoleServerPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableConsoleServerPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableConsoleServerPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableConsoleServerPortTemplateRequest(val *WritableConsoleServerPortTemplateRequest) *NullableWritableConsoleServerPortTemplateRequest {
	return &NullableWritableConsoleServerPortTemplateRequest{value: val, isSet: true}
}

func (v NullableWritableConsoleServerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableConsoleServerPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


