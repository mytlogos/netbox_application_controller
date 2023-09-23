package backend

import "github.com/mytlogos/netbox_application_controller/netbox"

type NetboxHost struct {
	Device         *netbox.DeviceWithConfigContext
	Platform       *netbox.Platform
	Manufacturer   *netbox.Manufacturer
	InventoryItems []*netbox.InventoryItem
	Interfaces     []*netbox.Interface
	DeviceRole     *netbox.DeviceRole
	IPAddresses    []*netbox.IPAddress
	Applications   []*netbox.ApplicationGroup
}

type NetboxOther struct {
	Wlan     []*netbox.WirelessLAN
	IPRanges []*netbox.IPRange
}
