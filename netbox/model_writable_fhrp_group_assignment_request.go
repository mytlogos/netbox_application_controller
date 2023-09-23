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

// checks if the WritableFHRPGroupAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableFHRPGroupAssignmentRequest{}

// WritableFHRPGroupAssignmentRequest Adds support for custom fields and tags.
type WritableFHRPGroupAssignmentRequest struct {
	Group int32 `json:"group"`
	InterfaceType string `json:"interface_type"`
	InterfaceId int64 `json:"interface_id"`
	Priority int32 `json:"priority"`
}

// NewWritableFHRPGroupAssignmentRequest instantiates a new WritableFHRPGroupAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableFHRPGroupAssignmentRequest(group int32, interfaceType string, interfaceId int64, priority int32) *WritableFHRPGroupAssignmentRequest {
	this := WritableFHRPGroupAssignmentRequest{}
	this.Group = group
	this.InterfaceType = interfaceType
	this.InterfaceId = interfaceId
	this.Priority = priority
	return &this
}

// NewWritableFHRPGroupAssignmentRequestWithDefaults instantiates a new WritableFHRPGroupAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableFHRPGroupAssignmentRequestWithDefaults() *WritableFHRPGroupAssignmentRequest {
	this := WritableFHRPGroupAssignmentRequest{}
	return &this
}

// GetGroup returns the Group field value
func (o *WritableFHRPGroupAssignmentRequest) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *WritableFHRPGroupAssignmentRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *WritableFHRPGroupAssignmentRequest) SetGroup(v int32) {
	o.Group = v
}

// GetInterfaceType returns the InterfaceType field value
func (o *WritableFHRPGroupAssignmentRequest) GetInterfaceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value
// and a boolean to check if the value has been set.
func (o *WritableFHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceType, true
}

// SetInterfaceType sets field value
func (o *WritableFHRPGroupAssignmentRequest) SetInterfaceType(v string) {
	o.InterfaceType = v
}

// GetInterfaceId returns the InterfaceId field value
func (o *WritableFHRPGroupAssignmentRequest) GetInterfaceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value
// and a boolean to check if the value has been set.
func (o *WritableFHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceId, true
}

// SetInterfaceId sets field value
func (o *WritableFHRPGroupAssignmentRequest) SetInterfaceId(v int64) {
	o.InterfaceId = v
}

// GetPriority returns the Priority field value
func (o *WritableFHRPGroupAssignmentRequest) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *WritableFHRPGroupAssignmentRequest) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *WritableFHRPGroupAssignmentRequest) SetPriority(v int32) {
	o.Priority = v
}

func (o WritableFHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableFHRPGroupAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["interface_type"] = o.InterfaceType
	toSerialize["interface_id"] = o.InterfaceId
	toSerialize["priority"] = o.Priority
	return toSerialize, nil
}

type NullableWritableFHRPGroupAssignmentRequest struct {
	value *WritableFHRPGroupAssignmentRequest
	isSet bool
}

func (v NullableWritableFHRPGroupAssignmentRequest) Get() *WritableFHRPGroupAssignmentRequest {
	return v.value
}

func (v *NullableWritableFHRPGroupAssignmentRequest) Set(val *WritableFHRPGroupAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableFHRPGroupAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableFHRPGroupAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableFHRPGroupAssignmentRequest(val *WritableFHRPGroupAssignmentRequest) *NullableWritableFHRPGroupAssignmentRequest {
	return &NullableWritableFHRPGroupAssignmentRequest{value: val, isSet: true}
}

func (v NullableWritableFHRPGroupAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableFHRPGroupAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


