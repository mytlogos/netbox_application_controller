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

// checks if the DataSourceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSourceStatus{}

// DataSourceStatus struct for DataSourceStatus
type DataSourceStatus struct {
	// * `new` - New * `queued` - Queued * `syncing` - Syncing * `completed` - Completed * `failed` - Failed
	Value *string `json:"value,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataSourceStatus DataSourceStatus

// NewDataSourceStatus instantiates a new DataSourceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceStatus() *DataSourceStatus {
	this := DataSourceStatus{}
	return &this
}

// NewDataSourceStatusWithDefaults instantiates a new DataSourceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceStatusWithDefaults() *DataSourceStatus {
	this := DataSourceStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DataSourceStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DataSourceStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DataSourceStatus) SetValue(v string) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DataSourceStatus) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceStatus) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DataSourceStatus) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DataSourceStatus) SetLabel(v string) {
	o.Label = &v
}

func (o DataSourceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataSourceStatus) UnmarshalJSON(bytes []byte) (err error) {
	varDataSourceStatus := _DataSourceStatus{}

	if err = json.Unmarshal(bytes, &varDataSourceStatus); err == nil {
		*o = DataSourceStatus(varDataSourceStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataSourceStatus struct {
	value *DataSourceStatus
	isSet bool
}

func (v NullableDataSourceStatus) Get() *DataSourceStatus {
	return v.value
}

func (v *NullableDataSourceStatus) Set(val *DataSourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceStatus(val *DataSourceStatus) *NullableDataSourceStatus {
	return &NullableDataSourceStatus{value: val, isSet: true}
}

func (v NullableDataSourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


