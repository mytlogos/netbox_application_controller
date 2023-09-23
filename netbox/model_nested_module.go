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

// checks if the NestedModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedModule{}

// NestedModule Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedModule struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Device NestedDevice `json:"device"`
	ModuleBay ModuleNestedModuleBay `json:"module_bay"`
	ModuleType NestedModuleType `json:"module_type"`
	AdditionalProperties map[string]interface{}
}

type _NestedModule NestedModule

// NewNestedModule instantiates a new NestedModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedModule(id int32, url string, display string, device NestedDevice, moduleBay ModuleNestedModuleBay, moduleType NestedModuleType) *NestedModule {
	this := NestedModule{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Device = device
	this.ModuleBay = moduleBay
	this.ModuleType = moduleType
	return &this
}

// NewNestedModuleWithDefaults instantiates a new NestedModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedModuleWithDefaults() *NestedModule {
	this := NestedModule{}
	return &this
}

// GetId returns the Id field value
func (o *NestedModule) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedModule) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedModule) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedModule) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedModule) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedModule) SetDisplay(v string) {
	o.Display = v
}

// GetDevice returns the Device field value
func (o *NestedModule) GetDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *NestedModule) SetDevice(v NestedDevice) {
	o.Device = v
}

// GetModuleBay returns the ModuleBay field value
func (o *NestedModule) GetModuleBay() ModuleNestedModuleBay {
	if o == nil {
		var ret ModuleNestedModuleBay
		return ret
	}

	return o.ModuleBay
}

// GetModuleBayOk returns a tuple with the ModuleBay field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetModuleBayOk() (*ModuleNestedModuleBay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleBay, true
}

// SetModuleBay sets field value
func (o *NestedModule) SetModuleBay(v ModuleNestedModuleBay) {
	o.ModuleBay = v
}

// GetModuleType returns the ModuleType field value
func (o *NestedModule) GetModuleType() NestedModuleType {
	if o == nil {
		var ret NestedModuleType
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *NestedModule) GetModuleTypeOk() (*NestedModuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *NestedModule) SetModuleType(v NestedModuleType) {
	o.ModuleType = v
}

func (o NestedModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	// skip: device is readOnly
	// skip: module_bay is readOnly
	// skip: module_type is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedModule) UnmarshalJSON(bytes []byte) (err error) {
	varNestedModule := _NestedModule{}

	if err = json.Unmarshal(bytes, &varNestedModule); err == nil {
		*o = NestedModule(varNestedModule)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module_bay")
		delete(additionalProperties, "module_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedModule struct {
	value *NestedModule
	isSet bool
}

func (v NullableNestedModule) Get() *NestedModule {
	return v.value
}

func (v *NullableNestedModule) Set(val *NestedModule) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedModule) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedModule(val *NestedModule) *NullableNestedModule {
	return &NullableNestedModule{value: val, isSet: true}
}

func (v NullableNestedModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


