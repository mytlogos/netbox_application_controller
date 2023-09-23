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

// checks if the ApplicationPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationPortRequest{}

// ApplicationPortRequest Adds support for custom fields and tags.
type ApplicationPortRequest struct {
	Port int32 `json:"port"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Ipaddresses []int32 `json:"ipaddresses,omitempty"`
	ApplicationProtocols []int32 `json:"application_protocols"`
	Application int32 `json:"application"`
	// * `none` - None * `udp` - SSL * `quic` - Starttls
	TransportSecurity *string `json:"transport_security,omitempty"`
}

// NewApplicationPortRequest instantiates a new ApplicationPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPortRequest(port int32, applicationProtocols []int32, application int32) *ApplicationPortRequest {
	this := ApplicationPortRequest{}
	this.Port = port
	this.ApplicationProtocols = applicationProtocols
	this.Application = application
	return &this
}

// NewApplicationPortRequestWithDefaults instantiates a new ApplicationPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortRequestWithDefaults() *ApplicationPortRequest {
	this := ApplicationPortRequest{}
	return &this
}

// GetPort returns the Port field value
func (o *ApplicationPortRequest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ApplicationPortRequest) SetPort(v int32) {
	o.Port = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ApplicationPortRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ApplicationPortRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ApplicationPortRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetIpaddresses returns the Ipaddresses field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetIpaddresses() []int32 {
	if o == nil || IsNil(o.Ipaddresses) {
		var ret []int32
		return ret
	}
	return o.Ipaddresses
}

// GetIpaddressesOk returns a tuple with the Ipaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetIpaddressesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ipaddresses) {
		return nil, false
	}
	return o.Ipaddresses, true
}

// HasIpaddresses returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasIpaddresses() bool {
	if o != nil && !IsNil(o.Ipaddresses) {
		return true
	}

	return false
}

// SetIpaddresses gets a reference to the given []int32 and assigns it to the Ipaddresses field.
func (o *ApplicationPortRequest) SetIpaddresses(v []int32) {
	o.Ipaddresses = v
}

// GetApplicationProtocols returns the ApplicationProtocols field value
func (o *ApplicationPortRequest) GetApplicationProtocols() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationProtocols
}

// GetApplicationProtocolsOk returns a tuple with the ApplicationProtocols field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetApplicationProtocolsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationProtocols, true
}

// SetApplicationProtocols sets field value
func (o *ApplicationPortRequest) SetApplicationProtocols(v []int32) {
	o.ApplicationProtocols = v
}

// GetApplication returns the Application field value
func (o *ApplicationPortRequest) GetApplication() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetApplicationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *ApplicationPortRequest) SetApplication(v int32) {
	o.Application = v
}

// GetTransportSecurity returns the TransportSecurity field value if set, zero value otherwise.
func (o *ApplicationPortRequest) GetTransportSecurity() string {
	if o == nil || IsNil(o.TransportSecurity) {
		var ret string
		return ret
	}
	return *o.TransportSecurity
}

// GetTransportSecurityOk returns a tuple with the TransportSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortRequest) GetTransportSecurityOk() (*string, bool) {
	if o == nil || IsNil(o.TransportSecurity) {
		return nil, false
	}
	return o.TransportSecurity, true
}

// HasTransportSecurity returns a boolean if a field has been set.
func (o *ApplicationPortRequest) HasTransportSecurity() bool {
	if o != nil && !IsNil(o.TransportSecurity) {
		return true
	}

	return false
}

// SetTransportSecurity gets a reference to the given string and assigns it to the TransportSecurity field.
func (o *ApplicationPortRequest) SetTransportSecurity(v string) {
	o.TransportSecurity = &v
}

func (o ApplicationPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableApplicationPortRequest struct {
	value *ApplicationPortRequest
	isSet bool
}

func (v NullableApplicationPortRequest) Get() *ApplicationPortRequest {
	return v.value
}

func (v *NullableApplicationPortRequest) Set(val *ApplicationPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPortRequest(val *ApplicationPortRequest) *NullableApplicationPortRequest {
	return &NullableApplicationPortRequest{value: val, isSet: true}
}

func (v NullableApplicationPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


