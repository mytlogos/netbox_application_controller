/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.1 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Application{}

// Application Adds support for custom fields and tags.
type Application struct {
	Id int32 `json:"id"`
	Display string `json:"display"`
	// Name of the Application
	Name string `json:"name"`
	// Operational status of this Application  * `running` - Running * `stopped` - Stopped * `not_running` - Not running * `deprecated` - Deprecated
	Status *string `json:"status,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Url string `json:"url"`
	Device NullableInt32 `json:"device,omitempty"`
	// Maximum CPU Usage in cores or total percent (100% * number of cores) - 0 for no limit
	CpuLimit *string `json:"cpu_limit,omitempty"`
	// Maximum RAM Usage in MegaBytes - 0 for no limit
	RamLimit *float64 `json:"ram_limit,omitempty"`
	// Idle CPU Usage in cores or total percent (100% * number of cores) - 0 for unknown or none
	IdleCpu *string `json:"idle_cpu,omitempty"`
	// * `none` - None * `systemd` - Systemd * `docker` - Docker * `podman` - Podman * `kubernetes` - Kubernetes
	ApplicationManager *string `json:"application_manager,omitempty"`
	Group NullableInt32 `json:"group,omitempty"`
	// Version of the Application
	Version *string `json:"version,omitempty"`
	VersionManagerName *string `json:"version_manager_name,omitempty"`
	LatestVersion *string `json:"latest_version,omitempty"`
	ChangelogUrl *string `json:"changelog_url,omitempty"`
	LastVersionUpgrade NullableTime `json:"last_version_upgrade,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(id int32, display string, name string, created NullableTime, lastUpdated NullableTime, url string) *Application {
	this := Application{}
	this.Id = id
	this.Display = display
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.Url = url
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetId returns the Id field value
func (o *Application) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Application) SetId(v int32) {
	o.Id = v
}

// GetDisplay returns the Display field value
func (o *Application) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Application) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Application) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Application) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Application) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Application) SetStatus(v string) {
	o.Status = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Application) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Application) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Application) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Application) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Application) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Application) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Application) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Application) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Application) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Application) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Application) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Application) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Application) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetUrl returns the Url field value
func (o *Application) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Application) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Application) SetUrl(v string) {
	o.Url = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetDevice() int32 {
	if o == nil || IsNil(o.Device.Get()) {
		var ret int32
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetDeviceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *Application) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableInt32 and assigns it to the Device field.
func (o *Application) SetDevice(v int32) {
	o.Device.Set(&v)
}
// SetDeviceNil sets the value for Device to be an explicit nil
func (o *Application) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *Application) UnsetDevice() {
	o.Device.Unset()
}

// GetCpuLimit returns the CpuLimit field value if set, zero value otherwise.
func (o *Application) GetCpuLimit() string {
	if o == nil || IsNil(o.CpuLimit) {
		var ret string
		return ret
	}
	return *o.CpuLimit
}

// GetCpuLimitOk returns a tuple with the CpuLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCpuLimitOk() (*string, bool) {
	if o == nil || IsNil(o.CpuLimit) {
		return nil, false
	}
	return o.CpuLimit, true
}

// HasCpuLimit returns a boolean if a field has been set.
func (o *Application) HasCpuLimit() bool {
	if o != nil && !IsNil(o.CpuLimit) {
		return true
	}

	return false
}

// SetCpuLimit gets a reference to the given string and assigns it to the CpuLimit field.
func (o *Application) SetCpuLimit(v string) {
	o.CpuLimit = &v
}

// GetRamLimit returns the RamLimit field value if set, zero value otherwise.
func (o *Application) GetRamLimit() float64 {
	if o == nil || IsNil(o.RamLimit) {
		var ret float64
		return ret
	}
	return *o.RamLimit
}

// GetRamLimitOk returns a tuple with the RamLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetRamLimitOk() (*float64, bool) {
	if o == nil || IsNil(o.RamLimit) {
		return nil, false
	}
	return o.RamLimit, true
}

// HasRamLimit returns a boolean if a field has been set.
func (o *Application) HasRamLimit() bool {
	if o != nil && !IsNil(o.RamLimit) {
		return true
	}

	return false
}

// SetRamLimit gets a reference to the given float64 and assigns it to the RamLimit field.
func (o *Application) SetRamLimit(v float64) {
	o.RamLimit = &v
}

// GetIdleCpu returns the IdleCpu field value if set, zero value otherwise.
func (o *Application) GetIdleCpu() string {
	if o == nil || IsNil(o.IdleCpu) {
		var ret string
		return ret
	}
	return *o.IdleCpu
}

// GetIdleCpuOk returns a tuple with the IdleCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdleCpuOk() (*string, bool) {
	if o == nil || IsNil(o.IdleCpu) {
		return nil, false
	}
	return o.IdleCpu, true
}

// HasIdleCpu returns a boolean if a field has been set.
func (o *Application) HasIdleCpu() bool {
	if o != nil && !IsNil(o.IdleCpu) {
		return true
	}

	return false
}

// SetIdleCpu gets a reference to the given string and assigns it to the IdleCpu field.
func (o *Application) SetIdleCpu(v string) {
	o.IdleCpu = &v
}

// GetApplicationManager returns the ApplicationManager field value if set, zero value otherwise.
func (o *Application) GetApplicationManager() string {
	if o == nil || IsNil(o.ApplicationManager) {
		var ret string
		return ret
	}
	return *o.ApplicationManager
}

// GetApplicationManagerOk returns a tuple with the ApplicationManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetApplicationManagerOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationManager) {
		return nil, false
	}
	return o.ApplicationManager, true
}

// HasApplicationManager returns a boolean if a field has been set.
func (o *Application) HasApplicationManager() bool {
	if o != nil && !IsNil(o.ApplicationManager) {
		return true
	}

	return false
}

// SetApplicationManager gets a reference to the given string and assigns it to the ApplicationManager field.
func (o *Application) SetApplicationManager(v string) {
	o.ApplicationManager = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Application) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *Application) SetGroup(v int32) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Application) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Application) UnsetGroup() {
	o.Group.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Application) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Application) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Application) GetVersionManagerName() string {
	if o == nil || IsNil(o.VersionManagerName) {
		var ret string
		return ret
	}
	return *o.VersionManagerName
}

// GetVersionManagerNameOk returns a tuple with the VersionManagerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetVersionManagerNameOk() (*string, bool) {
	if o == nil || IsNil(o.VersionManagerName) {
		return nil, false
	}
	return o.VersionManagerName, true
}

// HasVersionManagerName returns a boolean if a field has been set.
func (o *Application) HasVersionManagerName() bool {
	if o != nil && !IsNil(o.VersionManagerName) {
		return true
	}

	return false
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Application) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *Application) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Application) GetChangelogUrl() string {
	if o == nil || IsNil(o.ChangelogUrl) {
		var ret string
		return ret
	}
	return *o.ChangelogUrl
}

// HasChangelogUrl returns a boolean if a field has been set.
func (o *Application) HasChangelogUrl() bool {
	if o != nil && !IsNil(o.ChangelogUrl) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Application) SetVersion(v string) {
	o.Version = &v
}

// GetLastVersionUpgrade returns the LastVersionUpgrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetLastVersionUpgrade() time.Time {
	if o == nil || IsNil(o.LastVersionUpgrade.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastVersionUpgrade.Get()
}

// GetLastVersionUpgradeOk returns a tuple with the LastVersionUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetLastVersionUpgradeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastVersionUpgrade.Get(), o.LastVersionUpgrade.IsSet()
}

// HasLastVersionUpgrade returns a boolean if a field has been set.
func (o *Application) HasLastVersionUpgrade() bool {
	if o != nil && o.LastVersionUpgrade.IsSet() {
		return true
	}

	return false
}

// SetLastVersionUpgrade gets a reference to the given NullableTime and assigns it to the LastVersionUpgrade field.
func (o *Application) SetLastVersionUpgrade(v time.Time) {
	o.LastVersionUpgrade.Set(&v)
}
// SetLastVersionUpgradeNil sets the value for LastVersionUpgrade to be an explicit nil
func (o *Application) SetLastVersionUpgradeNil() {
	o.LastVersionUpgrade.Set(nil)
}

// UnsetLastVersionUpgrade ensures that no value is present for LastVersionUpgrade, not even an explicit nil
func (o *Application) UnsetLastVersionUpgrade() {
	o.LastVersionUpgrade.Unset()
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["url"] = o.Url
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if !IsNil(o.CpuLimit) {
		toSerialize["cpu_limit"] = o.CpuLimit
	}
	if !IsNil(o.RamLimit) {
		toSerialize["ram_limit"] = o.RamLimit
	}
	if !IsNil(o.IdleCpu) {
		toSerialize["idle_cpu"] = o.IdleCpu
	}
	if !IsNil(o.ApplicationManager) {
		toSerialize["application_manager"] = o.ApplicationManager
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionManagerName) {
		toSerialize["version_manager_name"] = o.VersionManagerName
	}
	if !IsNil(o.LatestVersion) {
		toSerialize["latest_version"] = o.LatestVersion
	}
	if !IsNil(o.ChangelogUrl) {
		toSerialize["changelog_url"] = o.ChangelogUrl
	}
	if o.LastVersionUpgrade.IsSet() {
		toSerialize["last_version_upgrade"] = o.LastVersionUpgrade.Get()
	}
	return toSerialize, nil
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


