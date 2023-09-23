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

// checks if the PatchedWritableExportTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableExportTemplateRequest{}

// PatchedWritableExportTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedWritableExportTemplateRequest struct {
	ContentTypes []string `json:"content_types,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Jinja2 template code. The list of objects being exported is passed as a context variable named <code>queryset</code>.
	TemplateCode *string `json:"template_code,omitempty"`
	// Defaults to <code>text/plain; charset=utf-8</code>
	MimeType *string `json:"mime_type,omitempty"`
	// Extension to append to the rendered filename
	FileExtension *string `json:"file_extension,omitempty"`
	// Download file as attachment
	AsAttachment *bool `json:"as_attachment,omitempty"`
	// Remote data source
	DataSource NullableInt32 `json:"data_source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableExportTemplateRequest PatchedWritableExportTemplateRequest

// NewPatchedWritableExportTemplateRequest instantiates a new PatchedWritableExportTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableExportTemplateRequest() *PatchedWritableExportTemplateRequest {
	this := PatchedWritableExportTemplateRequest{}
	return &this
}

// NewPatchedWritableExportTemplateRequestWithDefaults instantiates a new PatchedWritableExportTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableExportTemplateRequestWithDefaults() *PatchedWritableExportTemplateRequest {
	this := PatchedWritableExportTemplateRequest{}
	return &this
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedWritableExportTemplateRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableExportTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableExportTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTemplateCode returns the TemplateCode field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetTemplateCode() string {
	if o == nil || IsNil(o.TemplateCode) {
		var ret string
		return ret
	}
	return *o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateCode) {
		return nil, false
	}
	return o.TemplateCode, true
}

// HasTemplateCode returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasTemplateCode() bool {
	if o != nil && !IsNil(o.TemplateCode) {
		return true
	}

	return false
}

// SetTemplateCode gets a reference to the given string and assigns it to the TemplateCode field.
func (o *PatchedWritableExportTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *PatchedWritableExportTemplateRequest) SetMimeType(v string) {
	o.MimeType = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *PatchedWritableExportTemplateRequest) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetAsAttachment returns the AsAttachment field value if set, zero value otherwise.
func (o *PatchedWritableExportTemplateRequest) GetAsAttachment() bool {
	if o == nil || IsNil(o.AsAttachment) {
		var ret bool
		return ret
	}
	return *o.AsAttachment
}

// GetAsAttachmentOk returns a tuple with the AsAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableExportTemplateRequest) GetAsAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.AsAttachment) {
		return nil, false
	}
	return o.AsAttachment, true
}

// HasAsAttachment returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasAsAttachment() bool {
	if o != nil && !IsNil(o.AsAttachment) {
		return true
	}

	return false
}

// SetAsAttachment gets a reference to the given bool and assigns it to the AsAttachment field.
func (o *PatchedWritableExportTemplateRequest) SetAsAttachment(v bool) {
	o.AsAttachment = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableExportTemplateRequest) GetDataSource() int32 {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret int32
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableExportTemplateRequest) GetDataSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *PatchedWritableExportTemplateRequest) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableInt32 and assigns it to the DataSource field.
func (o *PatchedWritableExportTemplateRequest) SetDataSource(v int32) {
	o.DataSource.Set(&v)
}
// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *PatchedWritableExportTemplateRequest) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *PatchedWritableExportTemplateRequest) UnsetDataSource() {
	o.DataSource.Unset()
}

func (o PatchedWritableExportTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableExportTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TemplateCode) {
		toSerialize["template_code"] = o.TemplateCode
	}
	if !IsNil(o.MimeType) {
		toSerialize["mime_type"] = o.MimeType
	}
	if !IsNil(o.FileExtension) {
		toSerialize["file_extension"] = o.FileExtension
	}
	if !IsNil(o.AsAttachment) {
		toSerialize["as_attachment"] = o.AsAttachment
	}
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedWritableExportTemplateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedWritableExportTemplateRequest := _PatchedWritableExportTemplateRequest{}

	if err = json.Unmarshal(bytes, &varPatchedWritableExportTemplateRequest); err == nil {
		*o = PatchedWritableExportTemplateRequest(varPatchedWritableExportTemplateRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "mime_type")
		delete(additionalProperties, "file_extension")
		delete(additionalProperties, "as_attachment")
		delete(additionalProperties, "data_source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableExportTemplateRequest struct {
	value *PatchedWritableExportTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableExportTemplateRequest) Get() *PatchedWritableExportTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableExportTemplateRequest) Set(val *PatchedWritableExportTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableExportTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableExportTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableExportTemplateRequest(val *PatchedWritableExportTemplateRequest) *NullablePatchedWritableExportTemplateRequest {
	return &NullablePatchedWritableExportTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableExportTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableExportTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


