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

// checks if the NestedIPAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedIPAddressRequest{}

// NestedIPAddressRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedIPAddressRequest struct {
	Address string `json:"address"`
	AdditionalProperties map[string]interface{}
}

type _NestedIPAddressRequest NestedIPAddressRequest

// NewNestedIPAddressRequest instantiates a new NestedIPAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedIPAddressRequest(address string) *NestedIPAddressRequest {
	this := NestedIPAddressRequest{}
	this.Address = address
	return &this
}

// NewNestedIPAddressRequestWithDefaults instantiates a new NestedIPAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedIPAddressRequestWithDefaults() *NestedIPAddressRequest {
	this := NestedIPAddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *NestedIPAddressRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NestedIPAddressRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NestedIPAddressRequest) SetAddress(v string) {
	o.Address = v
}

func (o NestedIPAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedIPAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedIPAddressRequest) UnmarshalJSON(bytes []byte) (err error) {
	varNestedIPAddressRequest := _NestedIPAddressRequest{}

	if err = json.Unmarshal(bytes, &varNestedIPAddressRequest); err == nil {
		*o = NestedIPAddressRequest(varNestedIPAddressRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedIPAddressRequest struct {
	value *NestedIPAddressRequest
	isSet bool
}

func (v NullableNestedIPAddressRequest) Get() *NestedIPAddressRequest {
	return v.value
}

func (v *NullableNestedIPAddressRequest) Set(val *NestedIPAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedIPAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedIPAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedIPAddressRequest(val *NestedIPAddressRequest) *NullableNestedIPAddressRequest {
	return &NullableNestedIPAddressRequest{value: val, isSet: true}
}

func (v NullableNestedIPAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedIPAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

