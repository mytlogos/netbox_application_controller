/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.1 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"os"
)

// checks if the WritableDeviceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableDeviceTypeRequest{}

// WritableDeviceTypeRequest Adds support for custom fields and tags.
type WritableDeviceTypeRequest struct {
	Manufacturer int32 `json:"manufacturer"`
	DefaultPlatform NullableInt32 `json:"default_platform,omitempty"`
	Model string `json:"model"`
	Slug string `json:"slug"`
	// Discrete part number (optional)
	PartNumber *string `json:"part_number,omitempty"`
	UHeight *float64 `json:"u_height,omitempty"`
	// Device consumes both front and rear rack faces
	IsFullDepth *bool `json:"is_full_depth,omitempty"`
	// Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child.  * `parent` - Parent * `child` - Child
	SubdeviceRole *string `json:"subdevice_role,omitempty"`
	// * `front-to-rear` - Front to rear * `rear-to-front` - Rear to front * `left-to-right` - Left to right * `right-to-left` - Right to left * `side-to-rear` - Side to rear * `passive` - Passive * `mixed` - Mixed
	Airflow *string `json:"airflow,omitempty"`
	Weight NullableFloat64 `json:"weight,omitempty"`
	// * `kg` - Kilograms * `g` - Grams * `lb` - Pounds * `oz` - Ounces
	WeightUnit *string `json:"weight_unit,omitempty"`
	FrontImage **os.File `json:"front_image,omitempty"`
	RearImage **os.File `json:"rear_image,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// NewWritableDeviceTypeRequest instantiates a new WritableDeviceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableDeviceTypeRequest(manufacturer int32, model string, slug string) *WritableDeviceTypeRequest {
	this := WritableDeviceTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	this.Slug = slug
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	return &this
}

// NewWritableDeviceTypeRequestWithDefaults instantiates a new WritableDeviceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableDeviceTypeRequestWithDefaults() *WritableDeviceTypeRequest {
	this := WritableDeviceTypeRequest{}
	var uHeight float64 = 1.0
	this.UHeight = &uHeight
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *WritableDeviceTypeRequest) GetManufacturer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetManufacturerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *WritableDeviceTypeRequest) SetManufacturer(v int32) {
	o.Manufacturer = v
}

// GetDefaultPlatform returns the DefaultPlatform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableDeviceTypeRequest) GetDefaultPlatform() int32 {
	if o == nil || IsNil(o.DefaultPlatform.Get()) {
		var ret int32
		return ret
	}
	return *o.DefaultPlatform.Get()
}

// GetDefaultPlatformOk returns a tuple with the DefaultPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableDeviceTypeRequest) GetDefaultPlatformOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPlatform.Get(), o.DefaultPlatform.IsSet()
}

// HasDefaultPlatform returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasDefaultPlatform() bool {
	if o != nil && o.DefaultPlatform.IsSet() {
		return true
	}

	return false
}

// SetDefaultPlatform gets a reference to the given NullableInt32 and assigns it to the DefaultPlatform field.
func (o *WritableDeviceTypeRequest) SetDefaultPlatform(v int32) {
	o.DefaultPlatform.Set(&v)
}
// SetDefaultPlatformNil sets the value for DefaultPlatform to be an explicit nil
func (o *WritableDeviceTypeRequest) SetDefaultPlatformNil() {
	o.DefaultPlatform.Set(nil)
}

// UnsetDefaultPlatform ensures that no value is present for DefaultPlatform, not even an explicit nil
func (o *WritableDeviceTypeRequest) UnsetDefaultPlatform() {
	o.DefaultPlatform.Unset()
}

// GetModel returns the Model field value
func (o *WritableDeviceTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *WritableDeviceTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetSlug returns the Slug field value
func (o *WritableDeviceTypeRequest) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WritableDeviceTypeRequest) SetSlug(v string) {
	o.Slug = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *WritableDeviceTypeRequest) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetUHeight() float64 {
	if o == nil || IsNil(o.UHeight) {
		var ret float64
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetUHeightOk() (*float64, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given float64 and assigns it to the UHeight field.
func (o *WritableDeviceTypeRequest) SetUHeight(v float64) {
	o.UHeight = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetIsFullDepth() bool {
	if o == nil || IsNil(o.IsFullDepth) {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullDepth) {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasIsFullDepth() bool {
	if o != nil && !IsNil(o.IsFullDepth) {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *WritableDeviceTypeRequest) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetSubdeviceRole() string {
	if o == nil || IsNil(o.SubdeviceRole) {
		var ret string
		return ret
	}
	return *o.SubdeviceRole
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetSubdeviceRoleOk() (*string, bool) {
	if o == nil || IsNil(o.SubdeviceRole) {
		return nil, false
	}
	return o.SubdeviceRole, true
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasSubdeviceRole() bool {
	if o != nil && !IsNil(o.SubdeviceRole) {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given string and assigns it to the SubdeviceRole field.
func (o *WritableDeviceTypeRequest) SetSubdeviceRole(v string) {
	o.SubdeviceRole = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetAirflow() string {
	if o == nil || IsNil(o.Airflow) {
		var ret string
		return ret
	}
	return *o.Airflow
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetAirflowOk() (*string, bool) {
	if o == nil || IsNil(o.Airflow) {
		return nil, false
	}
	return o.Airflow, true
}

// HasAirflow returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasAirflow() bool {
	if o != nil && !IsNil(o.Airflow) {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given string and assigns it to the Airflow field.
func (o *WritableDeviceTypeRequest) SetAirflow(v string) {
	o.Airflow = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableDeviceTypeRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableDeviceTypeRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *WritableDeviceTypeRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}
// SetWeightNil sets the value for Weight to be an explicit nil
func (o *WritableDeviceTypeRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *WritableDeviceTypeRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetWeightUnit() string {
	if o == nil || IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *WritableDeviceTypeRequest) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetFrontImage() *os.File {
	if o == nil || IsNil(o.FrontImage) {
		var ret *os.File
		return ret
	}
	return *o.FrontImage
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetFrontImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.FrontImage) {
		return nil, false
	}
	return o.FrontImage, true
}

// HasFrontImage returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasFrontImage() bool {
	if o != nil && !IsNil(o.FrontImage) {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given *os.File and assigns it to the FrontImage field.
func (o *WritableDeviceTypeRequest) SetFrontImage(v *os.File) {
	o.FrontImage = &v
}

// GetRearImage returns the RearImage field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetRearImage() *os.File {
	if o == nil || IsNil(o.RearImage) {
		var ret *os.File
		return ret
	}
	return *o.RearImage
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetRearImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.RearImage) {
		return nil, false
	}
	return o.RearImage, true
}

// HasRearImage returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasRearImage() bool {
	if o != nil && !IsNil(o.RearImage) {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given *os.File and assigns it to the RearImage field.
func (o *WritableDeviceTypeRequest) SetRearImage(v *os.File) {
	o.RearImage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableDeviceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableDeviceTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableDeviceTypeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableDeviceTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableDeviceTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableDeviceTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableDeviceTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableDeviceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	if o.DefaultPlatform.IsSet() {
		toSerialize["default_platform"] = o.DefaultPlatform.Get()
	}
	toSerialize["model"] = o.Model
	toSerialize["slug"] = o.Slug
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.IsFullDepth) {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if !IsNil(o.SubdeviceRole) {
		toSerialize["subdevice_role"] = o.SubdeviceRole
	}
	if !IsNil(o.Airflow) {
		toSerialize["airflow"] = o.Airflow
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
	}
	if !IsNil(o.FrontImage) {
		toSerialize["front_image"] = o.FrontImage
	}
	if !IsNil(o.RearImage) {
		toSerialize["rear_image"] = o.RearImage
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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

type NullableWritableDeviceTypeRequest struct {
	value *WritableDeviceTypeRequest
	isSet bool
}

func (v NullableWritableDeviceTypeRequest) Get() *WritableDeviceTypeRequest {
	return v.value
}

func (v *NullableWritableDeviceTypeRequest) Set(val *WritableDeviceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableDeviceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableDeviceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableDeviceTypeRequest(val *WritableDeviceTypeRequest) *NullableWritableDeviceTypeRequest {
	return &NullableWritableDeviceTypeRequest{value: val, isSet: true}
}

func (v NullableWritableDeviceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableDeviceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


