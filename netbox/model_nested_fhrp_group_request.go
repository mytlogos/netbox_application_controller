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

// checks if the NestedFHRPGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedFHRPGroupRequest{}

// NestedFHRPGroupRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedFHRPGroupRequest struct {
	// * `vrrp2` - VRRPv2 * `vrrp3` - VRRPv3 * `carp` - CARP * `clusterxl` - ClusterXL * `hsrp` - HSRP * `glbp` - GLBP * `other` - Other
	Protocol string `json:"protocol"`
	GroupId int32 `json:"group_id"`
	AdditionalProperties map[string]interface{}
}

type _NestedFHRPGroupRequest NestedFHRPGroupRequest

// NewNestedFHRPGroupRequest instantiates a new NestedFHRPGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedFHRPGroupRequest(protocol string, groupId int32) *NestedFHRPGroupRequest {
	this := NestedFHRPGroupRequest{}
	this.Protocol = protocol
	this.GroupId = groupId
	return &this
}

// NewNestedFHRPGroupRequestWithDefaults instantiates a new NestedFHRPGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedFHRPGroupRequestWithDefaults() *NestedFHRPGroupRequest {
	this := NestedFHRPGroupRequest{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *NestedFHRPGroupRequest) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroupRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NestedFHRPGroupRequest) SetProtocol(v string) {
	o.Protocol = v
}

// GetGroupId returns the GroupId field value
func (o *NestedFHRPGroupRequest) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *NestedFHRPGroupRequest) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *NestedFHRPGroupRequest) SetGroupId(v int32) {
	o.GroupId = v
}

func (o NestedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedFHRPGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	toSerialize["group_id"] = o.GroupId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedFHRPGroupRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedFHRPGroupRequest := _NestedFHRPGroupRequest{}

	if err = json.Unmarshal(bytes, &varNestedFHRPGroupRequest); err == nil {
		*o = NestedFHRPGroupRequest(varNestedFHRPGroupRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedFHRPGroupRequest struct {
	value *NestedFHRPGroupRequest
	isSet bool
}

func (v NullableNestedFHRPGroupRequest) Get() *NestedFHRPGroupRequest {
	return v.value
}

func (v *NullableNestedFHRPGroupRequest) Set(val *NestedFHRPGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedFHRPGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedFHRPGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedFHRPGroupRequest(val *NestedFHRPGroupRequest) *NullableNestedFHRPGroupRequest {
	return &NullableNestedFHRPGroupRequest{value: val, isSet: true}
}

func (v NullableNestedFHRPGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedFHRPGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


