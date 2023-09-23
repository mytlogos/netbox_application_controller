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

// checks if the NestedTenantGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedTenantGroup{}

// NestedTenantGroup Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedTenantGroup struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Depth int32 `json:"_depth"`
	AdditionalProperties map[string]interface{}
}

type _NestedTenantGroup NestedTenantGroup

// NewNestedTenantGroup instantiates a new NestedTenantGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedTenantGroup(id int32, url string, display string, name string, slug string, depth int32) *NestedTenantGroup {
	this := NestedTenantGroup{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Slug = slug
	this.Depth = depth
	return &this
}

// NewNestedTenantGroupWithDefaults instantiates a new NestedTenantGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedTenantGroupWithDefaults() *NestedTenantGroup {
	this := NestedTenantGroup{}
	return &this
}

// GetId returns the Id field value
func (o *NestedTenantGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedTenantGroup) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *NestedTenantGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedTenantGroup) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *NestedTenantGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *NestedTenantGroup) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *NestedTenantGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NestedTenantGroup) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *NestedTenantGroup) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NestedTenantGroup) SetSlug(v string) {
	o.Slug = v
}

// GetDepth returns the Depth field value
func (o *NestedTenantGroup) GetDepth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *NestedTenantGroup) GetDepthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *NestedTenantGroup) SetDepth(v int32) {
	o.Depth = v
}

func (o NestedTenantGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedTenantGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	// skip: _depth is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedTenantGroup) UnmarshalJSON(bytes []byte) (err error) {
	varNestedTenantGroup := _NestedTenantGroup{}

	if err = json.Unmarshal(bytes, &varNestedTenantGroup); err == nil {
		*o = NestedTenantGroup(varNestedTenantGroup)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "_depth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedTenantGroup struct {
	value *NestedTenantGroup
	isSet bool
}

func (v NullableNestedTenantGroup) Get() *NestedTenantGroup {
	return v.value
}

func (v *NullableNestedTenantGroup) Set(val *NestedTenantGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedTenantGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedTenantGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedTenantGroup(val *NestedTenantGroup) *NullableNestedTenantGroup {
	return &NullableNestedTenantGroup{value: val, isSet: true}
}

func (v NullableNestedTenantGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedTenantGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


