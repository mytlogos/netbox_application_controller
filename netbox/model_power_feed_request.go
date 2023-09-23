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

// checks if the PowerFeedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFeedRequest{}

// PowerFeedRequest Adds support for custom fields and tags.
type PowerFeedRequest struct {
	PowerPanel NestedPowerPanelRequest `json:"power_panel"`
	Rack NullableNestedRackRequest `json:"rack,omitempty"`
	Name string `json:"name"`
	// * `offline` - Offline * `active` - Active * `planned` - Planned * `failed` - Failed
	Status *string `json:"status,omitempty"`
	// * `primary` - Primary * `redundant` - Redundant
	Type *string `json:"type,omitempty"`
	// * `ac` - AC * `dc` - DC
	Supply *string `json:"supply,omitempty"`
	// * `single-phase` - Single phase * `three-phase` - Three-phase
	Phase *string `json:"phase,omitempty"`
	Voltage *int32 `json:"voltage,omitempty"`
	Amperage *int32 `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization *int32 `json:"max_utilization,omitempty"`
	// Treat as if a cable is connected
	MarkConnected *bool `json:"mark_connected,omitempty"`
	Description *string `json:"description,omitempty"`
	Tenant NullableNestedTenantRequest `json:"tenant,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTagRequest `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerFeedRequest PowerFeedRequest

// NewPowerFeedRequest instantiates a new PowerFeedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFeedRequest(powerPanel NestedPowerPanelRequest, name string) *PowerFeedRequest {
	this := PowerFeedRequest{}
	this.PowerPanel = powerPanel
	this.Name = name
	return &this
}

// NewPowerFeedRequestWithDefaults instantiates a new PowerFeedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFeedRequestWithDefaults() *PowerFeedRequest {
	this := PowerFeedRequest{}
	return &this
}

// GetPowerPanel returns the PowerPanel field value
func (o *PowerFeedRequest) GetPowerPanel() NestedPowerPanelRequest {
	if o == nil {
		var ret NestedPowerPanelRequest
		return ret
	}

	return o.PowerPanel
}

// GetPowerPanelOk returns a tuple with the PowerPanel field value
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetPowerPanelOk() (*NestedPowerPanelRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerPanel, true
}

// SetPowerPanel sets field value
func (o *PowerFeedRequest) SetPowerPanel(v NestedPowerPanelRequest) {
	o.PowerPanel = v
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerFeedRequest) GetRack() NestedRackRequest {
	if o == nil || IsNil(o.Rack.Get()) {
		var ret NestedRackRequest
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeedRequest) GetRackOk() (*NestedRackRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableNestedRackRequest and assigns it to the Rack field.
func (o *PowerFeedRequest) SetRack(v NestedRackRequest) {
	o.Rack.Set(&v)
}
// SetRackNil sets the value for Rack to be an explicit nil
func (o *PowerFeedRequest) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *PowerFeedRequest) UnsetRack() {
	o.Rack.Unset()
}

// GetName returns the Name field value
func (o *PowerFeedRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PowerFeedRequest) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PowerFeedRequest) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PowerFeedRequest) SetType(v string) {
	o.Type = &v
}

// GetSupply returns the Supply field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetSupply() string {
	if o == nil || IsNil(o.Supply) {
		var ret string
		return ret
	}
	return *o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetSupplyOk() (*string, bool) {
	if o == nil || IsNil(o.Supply) {
		return nil, false
	}
	return o.Supply, true
}

// HasSupply returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasSupply() bool {
	if o != nil && !IsNil(o.Supply) {
		return true
	}

	return false
}

// SetSupply gets a reference to the given string and assigns it to the Supply field.
func (o *PowerFeedRequest) SetSupply(v string) {
	o.Supply = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetPhase() string {
	if o == nil || IsNil(o.Phase) {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *PowerFeedRequest) SetPhase(v string) {
	o.Phase = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetVoltage() int32 {
	if o == nil || IsNil(o.Voltage) {
		var ret int32
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetVoltageOk() (*int32, bool) {
	if o == nil || IsNil(o.Voltage) {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasVoltage() bool {
	if o != nil && !IsNil(o.Voltage) {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given int32 and assigns it to the Voltage field.
func (o *PowerFeedRequest) SetVoltage(v int32) {
	o.Voltage = &v
}

// GetAmperage returns the Amperage field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetAmperage() int32 {
	if o == nil || IsNil(o.Amperage) {
		var ret int32
		return ret
	}
	return *o.Amperage
}

// GetAmperageOk returns a tuple with the Amperage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetAmperageOk() (*int32, bool) {
	if o == nil || IsNil(o.Amperage) {
		return nil, false
	}
	return o.Amperage, true
}

// HasAmperage returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasAmperage() bool {
	if o != nil && !IsNil(o.Amperage) {
		return true
	}

	return false
}

// SetAmperage gets a reference to the given int32 and assigns it to the Amperage field.
func (o *PowerFeedRequest) SetAmperage(v int32) {
	o.Amperage = &v
}

// GetMaxUtilization returns the MaxUtilization field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetMaxUtilization() int32 {
	if o == nil || IsNil(o.MaxUtilization) {
		var ret int32
		return ret
	}
	return *o.MaxUtilization
}

// GetMaxUtilizationOk returns a tuple with the MaxUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetMaxUtilizationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUtilization) {
		return nil, false
	}
	return o.MaxUtilization, true
}

// HasMaxUtilization returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasMaxUtilization() bool {
	if o != nil && !IsNil(o.MaxUtilization) {
		return true
	}

	return false
}

// SetMaxUtilization gets a reference to the given int32 and assigns it to the MaxUtilization field.
func (o *PowerFeedRequest) SetMaxUtilization(v int32) {
	o.MaxUtilization = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *PowerFeedRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PowerFeedRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PowerFeedRequest) GetTenant() NestedTenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PowerFeedRequest) GetTenantOk() (*NestedTenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenantRequest and assigns it to the Tenant field.
func (o *PowerFeedRequest) SetTenant(v NestedTenantRequest) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PowerFeedRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PowerFeedRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PowerFeedRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PowerFeedRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PowerFeedRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFeedRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PowerFeedRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PowerFeedRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PowerFeedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFeedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["power_panel"] = o.PowerPanel
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Supply) {
		toSerialize["supply"] = o.Supply
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.Voltage) {
		toSerialize["voltage"] = o.Voltage
	}
	if !IsNil(o.Amperage) {
		toSerialize["amperage"] = o.Amperage
	}
	if !IsNil(o.MaxUtilization) {
		toSerialize["max_utilization"] = o.MaxUtilization
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PowerFeedRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPowerFeedRequest := _PowerFeedRequest{}

	if err = json.Unmarshal(bytes, &varPowerFeedRequest); err == nil {
		*o = PowerFeedRequest(varPowerFeedRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "power_panel")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "supply")
		delete(additionalProperties, "phase")
		delete(additionalProperties, "voltage")
		delete(additionalProperties, "amperage")
		delete(additionalProperties, "max_utilization")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerFeedRequest struct {
	value *PowerFeedRequest
	isSet bool
}

func (v NullablePowerFeedRequest) Get() *PowerFeedRequest {
	return v.value
}

func (v *NullablePowerFeedRequest) Set(val *PowerFeedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedRequest(val *PowerFeedRequest) *NullablePowerFeedRequest {
	return &NullablePowerFeedRequest{value: val, isSet: true}
}

func (v NullablePowerFeedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


