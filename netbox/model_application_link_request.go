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

// checks if the ApplicationLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLinkRequest{}

// ApplicationLinkRequest Adds support for custom fields and tags.
type ApplicationLinkRequest struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Application int32 `json:"application"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewApplicationLinkRequest instantiates a new ApplicationLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLinkRequest(name string, link string, application int32) *ApplicationLinkRequest {
	this := ApplicationLinkRequest{}
	this.Name = name
	this.Link = link
	this.Application = application
	return &this
}

// NewApplicationLinkRequestWithDefaults instantiates a new ApplicationLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLinkRequestWithDefaults() *ApplicationLinkRequest {
	this := ApplicationLinkRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationLinkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationLinkRequest) SetName(v string) {
	o.Name = v
}

// GetLink returns the Link field value
func (o *ApplicationLinkRequest) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *ApplicationLinkRequest) SetLink(v string) {
	o.Link = v
}

// GetApplication returns the Application field value
func (o *ApplicationLinkRequest) GetApplication() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetApplicationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *ApplicationLinkRequest) SetApplication(v int32) {
	o.Application = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ApplicationLinkRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ApplicationLinkRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ApplicationLinkRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationLinkRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationLinkRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ApplicationLinkRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ApplicationLinkRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinkRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ApplicationLinkRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ApplicationLinkRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ApplicationLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["link"] = o.Link
	toSerialize["application"] = o.Application
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	return toSerialize, nil
}

type NullableApplicationLinkRequest struct {
	value *ApplicationLinkRequest
	isSet bool
}

func (v NullableApplicationLinkRequest) Get() *ApplicationLinkRequest {
	return v.value
}

func (v *NullableApplicationLinkRequest) Set(val *ApplicationLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLinkRequest(val *ApplicationLinkRequest) *NullableApplicationLinkRequest {
	return &NullableApplicationLinkRequest{value: val, isSet: true}
}

func (v NullableApplicationLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


