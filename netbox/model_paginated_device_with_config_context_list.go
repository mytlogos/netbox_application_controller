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

// checks if the PaginatedDeviceWithConfigContextList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedDeviceWithConfigContextList{}

// PaginatedDeviceWithConfigContextList struct for PaginatedDeviceWithConfigContextList
type PaginatedDeviceWithConfigContextList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []DeviceWithConfigContext `json:"results,omitempty"`
}

// NewPaginatedDeviceWithConfigContextList instantiates a new PaginatedDeviceWithConfigContextList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDeviceWithConfigContextList() *PaginatedDeviceWithConfigContextList {
	this := PaginatedDeviceWithConfigContextList{}
	return &this
}

// NewPaginatedDeviceWithConfigContextListWithDefaults instantiates a new PaginatedDeviceWithConfigContextList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDeviceWithConfigContextListWithDefaults() *PaginatedDeviceWithConfigContextList {
	this := PaginatedDeviceWithConfigContextList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedDeviceWithConfigContextList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDeviceWithConfigContextList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedDeviceWithConfigContextList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedDeviceWithConfigContextList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDeviceWithConfigContextList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDeviceWithConfigContextList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedDeviceWithConfigContextList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedDeviceWithConfigContextList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedDeviceWithConfigContextList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedDeviceWithConfigContextList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDeviceWithConfigContextList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDeviceWithConfigContextList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedDeviceWithConfigContextList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedDeviceWithConfigContextList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedDeviceWithConfigContextList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedDeviceWithConfigContextList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedDeviceWithConfigContextList) GetResults() []DeviceWithConfigContext {
	if o == nil || IsNil(o.Results) {
		var ret []DeviceWithConfigContext
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDeviceWithConfigContextList) GetResultsOk() ([]DeviceWithConfigContext, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedDeviceWithConfigContextList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeviceWithConfigContext and assigns it to the Results field.
func (o *PaginatedDeviceWithConfigContextList) SetResults(v []DeviceWithConfigContext) {
	o.Results = v
}

func (o PaginatedDeviceWithConfigContextList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedDeviceWithConfigContextList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedDeviceWithConfigContextList struct {
	value *PaginatedDeviceWithConfigContextList
	isSet bool
}

func (v NullablePaginatedDeviceWithConfigContextList) Get() *PaginatedDeviceWithConfigContextList {
	return v.value
}

func (v *NullablePaginatedDeviceWithConfigContextList) Set(val *PaginatedDeviceWithConfigContextList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDeviceWithConfigContextList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDeviceWithConfigContextList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDeviceWithConfigContextList(val *PaginatedDeviceWithConfigContextList) *NullablePaginatedDeviceWithConfigContextList {
	return &NullablePaginatedDeviceWithConfigContextList{value: val, isSet: true}
}

func (v NullablePaginatedDeviceWithConfigContextList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDeviceWithConfigContextList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


