package services

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/netbox"
	"github.com/mytlogos/netbox_application_controller/services/backend"
	"github.com/r3labs/diff/v3"
)

type NetboxConverter struct {
}

func (nc *NetboxConverter) ConvertHost(value *models.Host) *netbox.DeviceWithConfigContext {
	device := netbox.DeviceWithConfigContext{}
	nc.UpdateDevice(value, &device)
	return &device
}

func (nc *NetboxConverter) UpdateDevice(value *models.Host, device *netbox.DeviceWithConfigContext) {
	device.SetName(value.Device.Name)

	customFields := device.GetCustomFields()
	customFields[backend.CustomFieldAgent] = value
	customFields[backend.CustomFieldUuid] = value.Uuid
	device.SetCustomFields(customFields)

	device.SetPlatform(netbox.NestedPlatform{
		Name: value.Device.Platform,
	})
}

func UpdateInterfaces(value *models.Host, deviceId int32, existingInterfaces []netbox.Interface) ([]netbox.Interface, []int32) {
	interfaces := []netbox.Interface{}

	for _, net := range value.Network.Interfaces {
		var foundExisting *netbox.Interface

		for _, existing := range existingInterfaces {
			if existing.MacAddress.Get() != nil && net.Mac == *existing.MacAddress.Get() {
				foundExisting = &existing
				break
			}
			if net.Name == existing.Name {
				foundExisting = &existing
				break
			}
		}

		if foundExisting != nil {
			existingInterfaces = slices.DeleteFunc(existingInterfaces, func(value netbox.Interface) bool {
				return value.Id == foundExisting.Id
			})
			copy := *foundExisting
			updateInterface(*net, &copy)
			differ, err := diff.NewDiffer()
			changed := false

			if err != nil {
				log.Println(err)
				changed = true
			} else {
				// ensure that the types do match
				changes, err := differ.Diff(&copy, foundExisting)

				if err != nil {
					log.Println(err)
					changed = true
				} else {
					changed = len(changes) > 0
				}
			}
			if changed {
				interfaces = append(interfaces, copy)
			}
		} else {
			newInterface := netbox.Interface{}
			updateInterface(*net, &newInterface)
			interfaces = append(interfaces, newInterface)
		}
	}

	removeInterfaceIds := []int32{}
	for _, intface := range existingInterfaces {
		removeInterfaceIds = append(removeInterfaceIds, intface.Id)
	}
	return interfaces, removeInterfaceIds
}

func updateInterface(model models.Interface, target *netbox.Interface) {
	target.Name = model.Name

	if model.Mac == "" {
		target.MacAddress = *netbox.NewNullableString(nil)
	} else {
		model.Mac = strings.ToUpper(model.Mac)
		target.MacAddress = *netbox.NewNullableString(&model.Mac)
	}

	if model.Mtu < 1 {
		model.Mtu = 1
	}

	target.Mtu = *netbox.NewNullableInt32(netbox.PtrInt32(int32(model.Mtu)))
	target.Enabled = netbox.PtrBool(model.State != models.StateDisconnected)
	switch model.Type {
	case models.TypeEthernet:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("100base-tx"),
			Label: netbox.PtrString("100BASE-TX (10/100ME)"),
		}
	case models.TypeLoopback:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("virtual"),
			Label: netbox.PtrString("Virtual"),
		}
	case models.TypeBluetooth:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("ieee802.15.1"),
			Label: netbox.PtrString("IEEE 802.15.1 (Bluetooth)"),
		}
	case models.TypeBridge:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("bridge"),
			Label: netbox.PtrString("Bridge"),
		}
	case models.TypeWireless:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("ieee802.11ax"),
			Label: netbox.PtrString("IEEE 802.11ax"),
		}
	case models.TypeTunnel:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("other"),
			Label: netbox.PtrString("Other"),
		}
	case models.TypeVirtual:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("virtual"),
			Label: netbox.PtrString("Virtual"),
		}
	case models.TypeUnknown:
		target.Type = netbox.InterfaceType{
			Value: netbox.PtrString("other"),
			Label: netbox.PtrString("Other"),
		}
	default:
		log.Printf("unknown type: %s\n", model.Type)
	}
}

func UpdateIPAddresses(value *models.Host, interfaces []netbox.Interface, existingIps []netbox.IPAddress) ([]netbox.IPAddress, []int32) {
	addresses := []netbox.IPAddress{}
	interfaceIP := map[int64][]netbox.IPAddress{}

	// map interface id to its assigned ip addresses
	for _, ip := range existingIps {
		if ip.AssignedObjectId.Get() == nil {
			continue
		}
		interfaceId := *ip.AssignedObjectId.Get()
		list, ok := interfaceIP[interfaceId]
		if !ok {
			list = []netbox.IPAddress{}
		}
		interfaceIP[interfaceId] = append(list, ip)
	}

	var objectType = "dcim.interface"

	for _, net := range value.Network.Interfaces {
		if net.IPv4 == nil {
			continue
		}
		var foundExisting *netbox.Interface

		// find netbox interface for current loop item
		for _, netInterface := range interfaces {
			if netInterface.MacAddress.Get() != nil && net.Mac == *netInterface.MacAddress.Get() {
				foundExisting = &netInterface
				break
			}
			if net.Name == netInterface.Name {
				foundExisting = &netInterface
				break
			}
		}

		if foundExisting == nil {
			panic("missing interface - interface should have been created")
		}

		netboxInterfaceIps, ok := interfaceIP[int64(foundExisting.Id)]

		if !ok {
			// if no netbox ip addresses exist, we just add the ons of the current model
			for _, ip := range net.IPv4 {
				addresses = append(addresses, netbox.IPAddress{
					Address: ipv4String(ip),
					Status: &netbox.IPAddressStatus{
						Value: netbox.PtrString("active"),
						Label: netbox.PtrString("Active"),
					},
					AssignedObjectType: *netbox.NewNullableString(&objectType),
					AssignedObjectId:   *netbox.NewNullableInt64(netbox.PtrInt64(int64(foundExisting.Id))),
				})
			}
			continue
		}

		for _, netIP := range net.IPv4 {
			// e.g. 192.168.1.1/24
			ipWithMask := ipv4String(netIP)
			found := false

			// search for netbox ips which match this one, so we do not add duplicates
			for _, netboxIp := range netboxInterfaceIps {
				if netboxIp.Address == ipWithMask {
					existingIps = slices.DeleteFunc(existingIps, func(value netbox.IPAddress) bool {
						return value.Id == netboxIp.Id
					})
					found = true
					break
				}
			}
			// now append a new ip address, where we are sure it is no duplicate
			if !found {
				addresses = append(addresses, netbox.IPAddress{
					Address: ipWithMask,
					Status: &netbox.IPAddressStatus{
						Value: netbox.PtrString("active"),
						Label: netbox.PtrString("Active"),
					},
					AssignedObjectType: *netbox.NewNullableString(&objectType),
					AssignedObjectId:   *netbox.NewNullableInt64(netbox.PtrInt64(int64(foundExisting.Id))),
				})
			}
		}
	}

	removeIpIds := []int32{}
	for _, intface := range existingIps {
		removeIpIds = append(removeIpIds, intface.Id)
	}
	return addresses, removeIpIds
}

func ipv4String(ip models.IPAddress) string {
	return fmt.Sprintf("%s/%d", ip.Address.String(), ip.MaskLen)
}

func UpdateApplicationGroups(app *models.App, deviceId int32, existingAppGroups []netbox.ApplicationGroup) []netbox.ApplicationGroup {
	createGroups := []netbox.ApplicationGroup{}

	for _, ag := range app.Applications {
		found := false
		for _, existing := range existingAppGroups {
			if existing.Name == ag.Name {
				found = true
				break
			}
		}
		if !found {
			createGroups = append(createGroups, netbox.ApplicationGroup{
				Name: ag.Name,
			})
		}
	}
	return createGroups
}

func UpdateApplications(app *models.App, deviceId int32, existingAppGroups []netbox.ApplicationGroup, existingApps []netbox.Application) ([]netbox.Application, []int32) {
	result := []netbox.Application{}
	deleteIds := []int32{}

	for _, group := range app.Applications {
		groupIndex := slices.IndexFunc(existingAppGroups, func(value netbox.ApplicationGroup) bool {
			return value.Name == group.Name
		})

		if groupIndex < 0 {
			log.Printf("could not find application group with name %s even though it should have been created", group.Name)
			continue
		}

		netboxGroup := existingAppGroups[groupIndex]

		for _, appModel := range group.Applications {
			var netboxAppIndex int = -1

			for i, existing := range existingApps {
				if existing.Name == appModel.Name && existing.GetDevice() == deviceId {
					netboxAppIndex = i
					break
				}
			}

			var netboxApp *netbox.Application

			if netboxAppIndex < 0 {
				netboxApp = &netbox.Application{
					Name:   appModel.Name,
					Device: *netbox.NewNullableInt32(&deviceId),
				}
			} else {
				netboxApp = &existingApps[netboxAppIndex]
			}

			netboxApp.RamLimit = netbox.PtrFloat64(float64(appModel.MemoryLimits))
			netboxApp.ApplicationManager = netbox.PtrString(string(appModel.AppManager))
			netboxApp.Group = *netbox.NewNullableInt32(&netboxGroup.Id)
			result = append(result, *netboxApp)
		}
	}
	return result, deleteIds
}

func UpdateApplicationPorts(app *models.App, deviceId int32, existingApps []netbox.Application, existingAppPorts []netbox.ApplicationPort) ([]netbox.ApplicationPort, []int32) {
	results := []netbox.ApplicationPort{}
	deleteIds := []int32{}

	for _, ag := range app.Applications {
		for _, app := range ag.Applications {
			portMapping := map[int][]models.ApplicationPort{}
			for _, port := range app.Ports {
				portMapping[port.Port] = append(portMapping[port.Port], port)
			}

			var appId int32

			for _, existing := range existingApps {
				if existing.Name == app.Name && existing.GetDevice() == deviceId {
					appId = existing.Id
					break
				}
			}

			if appId == 0 {
				log.Printf("could not find existing app with name %s and device id %d\n", app.Name, deviceId)
				continue
			}

			for port, _ := range portMapping {
				found := false
				port := int32(port)

				for _, existing := range existingAppPorts {
					if existing.Port == port && existing.Application == appId {
						found = true
						break
					}
				}
				if !found {
					results = append(results, netbox.ApplicationPort{
						Port:                 port,
						Application:          appId,
						ApplicationProtocols: []int32{},
					})
				}
			}
		}
	}
	return results, deleteIds
}
