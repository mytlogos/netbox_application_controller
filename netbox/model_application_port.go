/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.1 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the ApplicationPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationPort{}

// ApplicationPort Adds support for custom fields and tags.
type ApplicationPort struct {
	Id int32 `json:"id"`
	Display string `json:"display"`
	Port int32 `json:"port"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Url string `json:"url"`
	Ipaddresses []int32 `json:"ipaddresses,omitempty"`
	ApplicationProtocols []int32 `json:"application_protocols"`
	Application int32 `json:"application"`
	// * `none` - None * `udp` - SSL * `quic` - Starttls
	TransportSecurity *string `json:"transport_security,omitempty"`
}

// NewApplicationPort instantiates a new ApplicationPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPort(id int32, display string, port int32, created NullableTime, lastUpdated NullableTime, url string, applicationProtocols []int32, application int32) *ApplicationPort {
	this := ApplicationPort{}
	this.Id = id
	this.Display = display
	this.Port = port
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Url = url
	this.ApplicationProtocols = applicationProtocols
	this.Application = application
	return &this
}

// NewApplicationPortWithDefaults instantiates a new ApplicationPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortWithDefaults() *ApplicationPort {
	this := ApplicationPort{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationPort) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationPort) SetId(v int32) {
	o.Id = v
}

// GetDisplay returns the Display field value
func (o *ApplicationPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ApplicationPort) SetDisplay(v string) {
	o.Display = v
}

// GetPort returns the Port field value
func (o *ApplicationPort) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ApplicationPort) SetPort(v int32) {
	o.Port = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ApplicationPort) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ApplicationPort) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ApplicationPort) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationPort) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationPort) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ApplicationPort) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ApplicationPort) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ApplicationPort) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ApplicationPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ApplicationPort) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationPort) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ApplicationPort) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ApplicationPort) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationPort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ApplicationPort) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetUrl returns the Url field value
func (o *ApplicationPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApplicationPort) SetUrl(v string) {
	o.Url = v
}

// GetIpaddresses returns the Ipaddresses field value if set, zero value otherwise.
func (o *ApplicationPort) GetIpaddresses() []int32 {
	if o == nil || IsNil(o.Ipaddresses) {
		var ret []int32
		return ret
	}
	return o.Ipaddresses
}

// GetIpaddressesOk returns a tuple with the Ipaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetIpaddressesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ipaddresses) {
		return nil, false
	}
	return o.Ipaddresses, true
}

// HasIpaddresses returns a boolean if a field has been set.
func (o *ApplicationPort) HasIpaddresses() bool {
	if o != nil && !IsNil(o.Ipaddresses) {
		return true
	}

	return false
}

// SetIpaddresses gets a reference to the given []int32 and assigns it to the Ipaddresses field.
func (o *ApplicationPort) SetIpaddresses(v []int32) {
	o.Ipaddresses = v
}

// GetApplicationProtocols returns the ApplicationProtocols field value
func (o *ApplicationPort) GetApplicationProtocols() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationProtocols
}

// GetApplicationProtocolsOk returns a tuple with the ApplicationProtocols field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetApplicationProtocolsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationProtocols, true
}

// SetApplicationProtocols sets field value
func (o *ApplicationPort) SetApplicationProtocols(v []int32) {
	o.ApplicationProtocols = v
}

// GetApplication returns the Application field value
func (o *ApplicationPort) GetApplication() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetApplicationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *ApplicationPort) SetApplication(v int32) {
	o.Application = v
}

// GetTransportSecurity returns the TransportSecurity field value if set, zero value otherwise.
func (o *ApplicationPort) GetTransportSecurity() string {
	if o == nil || IsNil(o.TransportSecurity) {
		var ret string
		return ret
	}
	return *o.TransportSecurity
}

// GetTransportSecurityOk returns a tuple with the TransportSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPort) GetTransportSecurityOk() (*string, bool) {
	if o == nil || IsNil(o.TransportSecurity) {
		return nil, false
	}
	return o.TransportSecurity, true
}

// HasTransportSecurity returns a boolean if a field has been set.
func (o *ApplicationPort) HasTransportSecurity() bool {
	if o != nil && !IsNil(o.TransportSecurity) {
		return true
	}

	return false
}

// SetTransportSecurity gets a reference to the given string and assigns it to the TransportSecurity field.
func (o *ApplicationPort) SetTransportSecurity(v string) {
	o.TransportSecurity = &v
}

func (o ApplicationPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["display"] = o.Display
	toSerialize["port"] = o.Port
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["url"] = o.Url
	if !IsNil(o.Ipaddresses) {
		toSerialize["ipaddresses"] = o.Ipaddresses
	}
	toSerialize["application_protocols"] = o.ApplicationProtocols
	toSerialize["application"] = o.Application
	if !IsNil(o.TransportSecurity) {
		toSerialize["transport_security"] = o.TransportSecurity
	}
	return toSerialize, nil
}

type NullableApplicationPort struct {
	value *ApplicationPort
	isSet bool
}

func (v NullableApplicationPort) Get() *ApplicationPort {
	return v.value
}

func (v *NullableApplicationPort) Set(val *ApplicationPort) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPort) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPort(val *ApplicationPort) *NullableApplicationPort {
	return &NullableApplicationPort{value: val, isSet: true}
}

func (v NullableApplicationPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


