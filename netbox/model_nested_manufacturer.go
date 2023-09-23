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

// checks if the NestedManufacturer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedManufacturer{}

// NestedManufacturer Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedManufacturer struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// NewNestedManufacturer instantiates a new NestedManufacturer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedManufacturer(id int32, url string, display string, name string, slug string) *NestedManufacturer {
	this := NestedManufacturer{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	return &this
}

// NewNestedManufacturerWithDefaults instantiates a new NestedManufacturer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedManufacturerWithDefaults() *NestedManufacturer {
	this := NestedManufacturer{}
	return &this
}

// GetId returns the Id field value
func (o *NestedManufacturer) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedManufacturer) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedManufacturer) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedManufacturer) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedManufacturer) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedManufacturer) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedManufacturer) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedManufacturer) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedManufacturer) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedManufacturer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedManufacturer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedManufacturer) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedManufacturer) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedManufacturer) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedManufacturer) SetSlug(v string) {
	o.Slug = v
}

func (o NestedManufacturer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedManufacturer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	return toSerialize, nil
}

type NullableNestedManufacturer struct {
	value *NestedManufacturer
	isSet bool
}

func (v NullableNestedManufacturer) Get() *NestedManufacturer {
	return v.value
}

func (v *NullableNestedManufacturer) Set(val *NestedManufacturer) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedManufacturer) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedManufacturer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedManufacturer(val *NestedManufacturer) *NullableNestedManufacturer {
	return &NullableNestedManufacturer{value: val, isSet: true}
}

func (v NullableNestedManufacturer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedManufacturer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


