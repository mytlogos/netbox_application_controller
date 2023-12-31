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

// checks if the NestedProviderAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedProviderAccountRequest{}

// NestedProviderAccountRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedProviderAccountRequest struct {
	Name *string `json:"name,omitempty"`
	Account string `json:"account"`
}

// NewNestedProviderAccountRequest instantiates a new NestedProviderAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedProviderAccountRequest(account string) *NestedProviderAccountRequest {
	this := NestedProviderAccountRequest{}
	this.Account = account
	return &this
}

// NewNestedProviderAccountRequestWithDefaults instantiates a new NestedProviderAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedProviderAccountRequestWithDefaults() *NestedProviderAccountRequest {
	this := NestedProviderAccountRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NestedProviderAccountRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedProviderAccountRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NestedProviderAccountRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NestedProviderAccountRequest) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value
func (o *NestedProviderAccountRequest) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NestedProviderAccountRequest) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NestedProviderAccountRequest) SetAccount(v string) {
	o.Account = v
}

func (o NestedProviderAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedProviderAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account
	return toSerialize, nil
}

type NullableNestedProviderAccountRequest struct {
	value *NestedProviderAccountRequest
	isSet bool
}

func (v NullableNestedProviderAccountRequest) Get() *NestedProviderAccountRequest {
	return v.value
}

func (v *NullableNestedProviderAccountRequest) Set(val *NestedProviderAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedProviderAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedProviderAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedProviderAccountRequest(val *NestedProviderAccountRequest) *NullableNestedProviderAccountRequest {
	return &NullableNestedProviderAccountRequest{value: val, isSet: true}
}

func (v NullableNestedProviderAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedProviderAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


