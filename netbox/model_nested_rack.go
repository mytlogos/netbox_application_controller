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

// checks if the NestedRack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRack{}

// NestedRack Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedRack struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _NestedRack NestedRack

// NewNestedRack instantiates a new NestedRack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRack(id int32, url string, display string, name string) *NestedRack {
	this := NestedRack{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewNestedRackWithDefaults instantiates a new NestedRack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRackWithDefaults() *NestedRack {
	this := NestedRack{}
	return &this
}

// GetId returns the Id field value
func (o *NestedRack) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedRack) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedRack) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedRack) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedRack) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedRack) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedRack) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedRack) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedRack) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedRack) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedRack) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedRack) SetName(v string) {
	o.Name = v
}

func (o NestedRack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRack) UnmarshalJSON(bytes []byte) (err error) {
	varNestedRack := _NestedRack{}

	if err = json.Unmarshal(bytes, &varNestedRack); err == nil {
		*o = NestedRack(varNestedRack)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRack struct {
	value *NestedRack
	isSet bool
}

func (v NullableNestedRack) Get() *NestedRack {
	return v.value
}

func (v *NullableNestedRack) Set(val *NestedRack) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRack) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRack(val *NestedRack) *NullableNestedRack {
	return &NullableNestedRack{value: val, isSet: true}
}

func (v NullableNestedRack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


