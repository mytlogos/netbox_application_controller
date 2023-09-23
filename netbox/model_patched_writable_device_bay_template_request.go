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

// checks if the PatchedWritableDeviceBayTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableDeviceBayTemplateRequest{}

// PatchedWritableDeviceBayTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableDeviceBayTemplateRequest struct {
	DeviceType *int32 `json:"device_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableDeviceBayTemplateRequest PatchedWritableDeviceBayTemplateRequest

// NewPatchedWritableDeviceBayTemplateRequest instantiates a new PatchedWritableDeviceBayTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableDeviceBayTemplateRequest() *PatchedWritableDeviceBayTemplateRequest {
	this := PatchedWritableDeviceBayTemplateRequest{}
	return &this
}

// NewPatchedWritableDeviceBayTemplateRequestWithDefaults instantiates a new PatchedWritableDeviceBayTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableDeviceBayTemplateRequestWithDefaults() *PatchedWritableDeviceBayTemplateRequest {
	this := PatchedWritableDeviceBayTemplateRequest{}
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *PatchedWritableDeviceBayTemplateRequest) GetDeviceType() int32 {
	if o == nil || IsNil(o.DeviceType) {
		var ret int32
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) GetDeviceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given int32 and assigns it to the DeviceType field.
func (o *PatchedWritableDeviceBayTemplateRequest) SetDeviceType(v int32) {
	o.DeviceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableDeviceBayTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableDeviceBayTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableDeviceBayTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableDeviceBayTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableDeviceBayTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableDeviceBayTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableDeviceBayTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedWritableDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableDeviceBayTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceType) {
		toSerialize["device_type"] = o.DeviceType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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

func (o *PatchedWritableDeviceBayTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableDeviceBayTemplateRequest := _PatchedWritableDeviceBayTemplateRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableDeviceBayTemplateRequest); err == nil {
		*o = PatchedWritableDeviceBayTemplateRequest(varPatchedWritableDeviceBayTemplateRequest)
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

type NullablePatchedWritableDeviceBayTemplateRequest struct {
	value *PatchedWritableDeviceBayTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableDeviceBayTemplateRequest) Get() *PatchedWritableDeviceBayTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableDeviceBayTemplateRequest) Set(val *PatchedWritableDeviceBayTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableDeviceBayTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableDeviceBayTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableDeviceBayTemplateRequest(val *PatchedWritableDeviceBayTemplateRequest) *NullablePatchedWritableDeviceBayTemplateRequest {
	return &NullablePatchedWritableDeviceBayTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableDeviceBayTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableDeviceBayTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


