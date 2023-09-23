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

// checks if the PaginatedWirelessLinkList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedWirelessLinkList{}

// PaginatedWirelessLinkList struct for PaginatedWirelessLinkList
type PaginatedWirelessLinkList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []WirelessLink `json:"results,omitempty"`
}

// NewPaginatedWirelessLinkList instantiates a new PaginatedWirelessLinkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedWirelessLinkList() *PaginatedWirelessLinkList {
	this := PaginatedWirelessLinkList{}
	return &this
}

// NewPaginatedWirelessLinkListWithDefaults instantiates a new PaginatedWirelessLinkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedWirelessLinkListWithDefaults() *PaginatedWirelessLinkList {
	this := PaginatedWirelessLinkList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedWirelessLinkList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedWirelessLinkList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedWirelessLinkList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedWirelessLinkList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedWirelessLinkList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedWirelessLinkList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedWirelessLinkList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedWirelessLinkList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedWirelessLinkList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedWirelessLinkList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedWirelessLinkList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedWirelessLinkList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedWirelessLinkList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedWirelessLinkList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedWirelessLinkList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedWirelessLinkList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedWirelessLinkList) GetResults() []WirelessLink {
	if o == nil || IsNil(o.Results) {
		var ret []WirelessLink
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedWirelessLinkList) GetResultsOk() ([]WirelessLink, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedWirelessLinkList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WirelessLink and assigns it to the Results field.
func (o *PaginatedWirelessLinkList) SetResults(v []WirelessLink) {
	o.Results = v
}

func (o PaginatedWirelessLinkList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedWirelessLinkList) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedWirelessLinkList struct {
	value *PaginatedWirelessLinkList
	isSet bool
}

func (v NullablePaginatedWirelessLinkList) Get() *PaginatedWirelessLinkList {
	return v.value
}

func (v *NullablePaginatedWirelessLinkList) Set(val *PaginatedWirelessLinkList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedWirelessLinkList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedWirelessLinkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedWirelessLinkList(val *PaginatedWirelessLinkList) *NullablePaginatedWirelessLinkList {
	return &NullablePaginatedWirelessLinkList{value: val, isSet: true}
}

func (v NullablePaginatedWirelessLinkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedWirelessLinkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


