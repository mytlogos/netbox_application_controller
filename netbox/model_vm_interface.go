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

// checks if the VMInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInterface{}

// VMInterface Adds support for custom fields and tags.
type VMInterface struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Display string `json:"display"`
	VirtualMachine NestedVirtualMachine `json:"virtual_machine"`
	Name string `json:"name"`
	Enabled *bool `json:"enabled,omitempty"`
	Parent NullableNestedVMInterface `json:"parent,omitempty"`
	Bridge NullableNestedVMInterface `json:"bridge,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	MacAddress NullableString `json:"mac_address,omitempty"`
	Description *string `json:"description,omitempty"`
	Mode *InterfaceMode `json:"mode,omitempty"`
	UntaggedVlan NullableNestedVLAN `json:"untagged_vlan,omitempty"`
	TaggedVlans []int32 `json:"tagged_vlans,omitempty"`
	Vrf NullableNestedVRF `json:"vrf,omitempty"`
	L2vpnTermination NullableNestedL2VPNTermination `json:"l2vpn_termination"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	CountIpaddresses int32 `json:"count_ipaddresses"`
	CountFhrpGroups int32 `json:"count_fhrp_groups"`
	AdditionalProperties map[string]interface{}
}

type _VMInterface VMInterface

// NewVMInterface instantiates a new VMInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInterface(id int32, url string, display string, virtualMachine NestedVirtualMachine, name string, l2vpnTermination NullableNestedL2VPNTermination, created NullableTime, lastUpdated NullableTime, countIpaddresses int32, countFhrpGroups int32) *VMInterface {
	this := VMInterface{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.VirtualMachine = virtualMachine
	this.Name = name
	this.L2vpnTermination = l2vpnTermination
	this.Created = created
	this.LastUpdated = lastUpdated
	this.CountIpaddresses = countIpaddresses
	this.CountFhrpGroups = countFhrpGroups
	return &this
}

// NewVMInterfaceWithDefaults instantiates a new VMInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInterfaceWithDefaults() *VMInterface {
	this := VMInterface{}
	return &this
}

// GetId returns the Id field value
func (o *VMInterface) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMInterface) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VMInterface) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VMInterface) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *VMInterface) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VMInterface) SetDisplay(v string) {
	o.Display = v
}

// GetVirtualMachine returns the VirtualMachine field value
func (o *VMInterface) GetVirtualMachine() NestedVirtualMachine {
	if o == nil {
		var ret NestedVirtualMachine
		return ret
	}

	return o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetVirtualMachineOk() (*NestedVirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualMachine, true
}

// SetVirtualMachine sets field value
func (o *VMInterface) SetVirtualMachine(v NestedVirtualMachine) {
	o.VirtualMachine = v
}

// GetName returns the Name field value
func (o *VMInterface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VMInterface) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VMInterface) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VMInterface) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VMInterface) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetParent() NestedVMInterface {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret NestedVMInterface
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetParentOk() (*NestedVMInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *VMInterface) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableNestedVMInterface and assigns it to the Parent field.
func (o *VMInterface) SetParent(v NestedVMInterface) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *VMInterface) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *VMInterface) UnsetParent() {
	o.Parent.Unset()
}

// GetBridge returns the Bridge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetBridge() NestedVMInterface {
	if o == nil || IsNil(o.Bridge.Get()) {
		var ret NestedVMInterface
		return ret
	}
	return *o.Bridge.Get()
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetBridgeOk() (*NestedVMInterface, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bridge.Get(), o.Bridge.IsSet()
}

// HasBridge returns a boolean if a field has been set.
func (o *VMInterface) HasBridge() bool {
	if o != nil && o.Bridge.IsSet() {
		return true
	}

	return false
}

// SetBridge gets a reference to the given NullableNestedVMInterface and assigns it to the Bridge field.
func (o *VMInterface) SetBridge(v NestedVMInterface) {
	o.Bridge.Set(&v)
}
// SetBridgeNil sets the value for Bridge to be an explicit nil
func (o *VMInterface) SetBridgeNil() {
	o.Bridge.Set(nil)
}

// UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
func (o *VMInterface) UnsetBridge() {
	o.Bridge.Unset()
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetMtu() int32 {
	if o == nil || IsNil(o.Mtu.Get()) {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetMtuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *VMInterface) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *VMInterface) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *VMInterface) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *VMInterface) UnsetMtu() {
	o.Mtu.Unset()
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress.Get()) {
		var ret string
		return ret
	}
	return *o.MacAddress.Get()
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MacAddress.Get(), o.MacAddress.IsSet()
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VMInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress.IsSet() {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given NullableString and assigns it to the MacAddress field.
func (o *VMInterface) SetMacAddress(v string) {
	o.MacAddress.Set(&v)
}
// SetMacAddressNil sets the value for MacAddress to be an explicit nil
func (o *VMInterface) SetMacAddressNil() {
	o.MacAddress.Set(nil)
}

// UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
func (o *VMInterface) UnsetMacAddress() {
	o.MacAddress.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VMInterface) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VMInterface) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VMInterface) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VMInterface) GetMode() InterfaceMode {
	if o == nil || IsNil(o.Mode) {
		var ret InterfaceMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetModeOk() (*InterfaceMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VMInterface) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given InterfaceMode and assigns it to the Mode field.
func (o *VMInterface) SetMode(v InterfaceMode) {
	o.Mode = &v
}

// GetUntaggedVlan returns the UntaggedVlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetUntaggedVlan() NestedVLAN {
	if o == nil || IsNil(o.UntaggedVlan.Get()) {
		var ret NestedVLAN
		return ret
	}
	return *o.UntaggedVlan.Get()
}

// GetUntaggedVlanOk returns a tuple with the UntaggedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetUntaggedVlanOk() (*NestedVLAN, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntaggedVlan.Get(), o.UntaggedVlan.IsSet()
}

// HasUntaggedVlan returns a boolean if a field has been set.
func (o *VMInterface) HasUntaggedVlan() bool {
	if o != nil && o.UntaggedVlan.IsSet() {
		return true
	}

	return false
}

// SetUntaggedVlan gets a reference to the given NullableNestedVLAN and assigns it to the UntaggedVlan field.
func (o *VMInterface) SetUntaggedVlan(v NestedVLAN) {
	o.UntaggedVlan.Set(&v)
}
// SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil
func (o *VMInterface) SetUntaggedVlanNil() {
	o.UntaggedVlan.Set(nil)
}

// UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
func (o *VMInterface) UnsetUntaggedVlan() {
	o.UntaggedVlan.Unset()
}

// GetTaggedVlans returns the TaggedVlans field value if set, zero value otherwise.
func (o *VMInterface) GetTaggedVlans() []int32 {
	if o == nil || IsNil(o.TaggedVlans) {
		var ret []int32
		return ret
	}
	return o.TaggedVlans
}

// GetTaggedVlansOk returns a tuple with the TaggedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetTaggedVlansOk() ([]int32, bool) {
	if o == nil || IsNil(o.TaggedVlans) {
		return nil, false
	}
	return o.TaggedVlans, true
}

// HasTaggedVlans returns a boolean if a field has been set.
func (o *VMInterface) HasTaggedVlans() bool {
	if o != nil && !IsNil(o.TaggedVlans) {
		return true
	}

	return false
}

// SetTaggedVlans gets a reference to the given []int32 and assigns it to the TaggedVlans field.
func (o *VMInterface) SetTaggedVlans(v []int32) {
	o.TaggedVlans = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VMInterface) GetVrf() NestedVRF {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret NestedVRF
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetVrfOk() (*NestedVRF, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *VMInterface) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableNestedVRF and assigns it to the Vrf field.
func (o *VMInterface) SetVrf(v NestedVRF) {
	o.Vrf.Set(&v)
}
// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *VMInterface) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *VMInterface) UnsetVrf() {
	o.Vrf.Unset()
}

// GetL2vpnTermination returns the L2vpnTermination field value
// If the value is explicit nil, the zero value for NestedL2VPNTermination will be returned
func (o *VMInterface) GetL2vpnTermination() NestedL2VPNTermination {
	if o == nil || o.L2vpnTermination.Get() == nil {
		var ret NestedL2VPNTermination
		return ret
	}

	return *o.L2vpnTermination.Get()
}

// GetL2vpnTerminationOk returns a tuple with the L2vpnTermination field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetL2vpnTerminationOk() (*NestedL2VPNTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2vpnTermination.Get(), o.L2vpnTermination.IsSet()
}

// SetL2vpnTermination sets field value
func (o *VMInterface) SetL2vpnTermination(v NestedL2VPNTermination) {
	o.L2vpnTermination.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VMInterface) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VMInterface) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VMInterface) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VMInterface) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInterface) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VMInterface) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VMInterface) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VMInterface) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VMInterface) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VMInterface) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VMInterface) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VMInterface) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetCountIpaddresses returns the CountIpaddresses field value
func (o *VMInterface) GetCountIpaddresses() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CountIpaddresses
}

// GetCountIpaddressesOk returns a tuple with the CountIpaddresses field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetCountIpaddressesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountIpaddresses, true
}

// SetCountIpaddresses sets field value
func (o *VMInterface) SetCountIpaddresses(v int32) {
	o.CountIpaddresses = v
}

// GetCountFhrpGroups returns the CountFhrpGroups field value
func (o *VMInterface) GetCountFhrpGroups() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CountFhrpGroups
}

// GetCountFhrpGroupsOk returns a tuple with the CountFhrpGroups field value
// and a boolean to check if the value has been set.
func (o *VMInterface) GetCountFhrpGroupsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountFhrpGroups, true
}

// SetCountFhrpGroups sets field value
func (o *VMInterface) SetCountFhrpGroups(v int32) {
	o.CountFhrpGroups = v
}

func (o VMInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	toSerialize["virtual_machine"] = o.VirtualMachine
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Bridge.IsSet() {
		toSerialize["bridge"] = o.Bridge.Get()
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.MacAddress.IsSet() {
		toSerialize["mac_address"] = o.MacAddress.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if o.UntaggedVlan.IsSet() {
		toSerialize["untagged_vlan"] = o.UntaggedVlan.Get()
	}
	if !IsNil(o.TaggedVlans) {
		toSerialize["tagged_vlans"] = o.TaggedVlans
	}
	if o.Vrf.IsSet() {
		toSerialize["vrf"] = o.Vrf.Get()
	}
	toSerialize["l2vpn_termination"] = o.L2vpnTermination.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	// skip: count_ipaddresses is readOnly
	// skip: count_fhrp_groups is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VMInterface) UnmarshalJSON(bytes []byte) (err error) {
	varVMInterface := _VMInterface{}

	if err = json.Unmarshal(bytes, &varVMInterface); err == nil {
		*o = VMInterface(varVMInterface)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "bridge")
		delete(additionalProperties, "mtu")
		delete(additionalProperties, "mac_address")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "untagged_vlan")
		delete(additionalProperties, "tagged_vlans")
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "l2vpn_termination")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "count_ipaddresses")
		delete(additionalProperties, "count_fhrp_groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVMInterface struct {
	value *VMInterface
	isSet bool
}

func (v NullableVMInterface) Get() *VMInterface {
	return v.value
}

func (v *NullableVMInterface) Set(val *VMInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInterface(val *VMInterface) *NullableVMInterface {
	return &NullableVMInterface{value: val, isSet: true}
}

func (v NullableVMInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


