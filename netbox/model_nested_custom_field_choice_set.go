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

// checks if the NestedCustomFieldChoiceSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedCustomFieldChoiceSet{}

// NestedCustomFieldChoiceSet Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedCustomFieldChoiceSet struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	ChoicesCount string `json:"choices_count"`
}

// NewNestedCustomFieldChoiceSet instantiates a new NestedCustomFieldChoiceSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedCustomFieldChoiceSet(id int32, url string, display string, name string, choicesCount string) *NestedCustomFieldChoiceSet {
	this := NestedCustomFieldChoiceSet{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.ChoicesCount = choicesCount
	return &this
}

// NewNestedCustomFieldChoiceSetWithDefaults instantiates a new NestedCustomFieldChoiceSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedCustomFieldChoiceSetWithDefaults() *NestedCustomFieldChoiceSet {
	this := NestedCustomFieldChoiceSet{}
	return &this
}

// GetId returns the Id field value
func (o *NestedCustomFieldChoiceSet) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSet) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedCustomFieldChoiceSet) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedCustomFieldChoiceSet) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSet) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedCustomFieldChoiceSet) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedCustomFieldChoiceSet) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSet) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedCustomFieldChoiceSet) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedCustomFieldChoiceSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedCustomFieldChoiceSet) SetName(v string) {
	o.Name = v
}

// GetChoicesCount returns the ChoicesCount field value
func (o *NestedCustomFieldChoiceSet) GetChoicesCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChoicesCount
}

// GetChoicesCountOk returns a tuple with the ChoicesCount field value
// and a boolean to check if the value has been set.
func (o *NestedCustomFieldChoiceSet) GetChoicesCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoicesCount, true
}

// SetChoicesCount sets field value
func (o *NestedCustomFieldChoiceSet) SetChoicesCount(v string) {
	o.ChoicesCount = v
}

func (o NestedCustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedCustomFieldChoiceSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["choices_count"] = o.ChoicesCount
	return toSerialize, nil
}

type NullableNestedCustomFieldChoiceSet struct {
	value *NestedCustomFieldChoiceSet
	isSet bool
}

func (v NullableNestedCustomFieldChoiceSet) Get() *NestedCustomFieldChoiceSet {
	return v.value
}

func (v *NullableNestedCustomFieldChoiceSet) Set(val *NestedCustomFieldChoiceSet) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedCustomFieldChoiceSet) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedCustomFieldChoiceSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedCustomFieldChoiceSet(val *NestedCustomFieldChoiceSet) *NullableNestedCustomFieldChoiceSet {
	return &NullableNestedCustomFieldChoiceSet{value: val, isSet: true}
}

func (v NullableNestedCustomFieldChoiceSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedCustomFieldChoiceSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


