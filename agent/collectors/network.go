package collectors

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/Wifx/gonetworkmanager"
	"github.com/jaypipes/ghw"
	"github.com/mytlogos/netbox_application_controller/models"
)

type NetworkCollector struct {
}

func (c *NetworkCollector) GetName() string {
	return "Network"
}

func (c *NetworkCollector) GetEndpoint() string {
	return "/api/agent/network"
}

func (c *NetworkCollector) GetLshwData() ([]LshwNetwork, error) {
	cmd := exec.Command("lshw", "-c", "network", "-json")
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var lshw []LshwNetwork
	err = json.Unmarshal(data, &lshw)

	if err != nil {
		return nil, err
	}
	return lshw, nil
}

func (c *NetworkCollector) GetIPData() ([]IPInfo, error) {
	cmd := exec.Command("ip", "-json", "address")
	data, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var lshw []IPInfo
	err = json.Unmarshal(data, &lshw)

	if err != nil {
		return nil, err
	}
	return lshw, nil
}

func (c *NetworkCollector) GetNmcli() error {
	nm, err := gonetworkmanager.NewNetworkManager()

	if err != nil {
		return err
	}
	devices, err := nm.GetAllDevices()

	if err != nil {
		return err
	}
	for _, d := range devices {
		data, err := d.MarshalJSON()

		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}
	return nil
}

func (c *NetworkCollector) updateMapWithIP(nameInterface map[string]*models.Interface) error {
	data, err := c.GetIPData()

	if err != nil {
		return err
	}
	for _, ipInfo := range data {
		intface, ok := nameInterface[ipInfo.Ifname]

		if !ok {
			log.Printf("interface with name %s not found\n", ipInfo.Ifname)
			continue
		}
		for i, v := range ipInfo.AddrInfo {
			log.Println(intface, i, v)
		}
	}
	return nil
}

func (c *NetworkCollector) updateMapWithNmcli(nameInterface map[string]*models.Interface) error {
	nm, err := gonetworkmanager.NewNetworkManager()

	if err != nil {
		return err
	}

	devices, err := nm.GetAllDevices()

	if err != nil {
		return err
	}
	for _, d := range devices {
		name, err := d.GetPropertyInterface()

		if err != nil {
			return err
		}
		intface, ok := nameInterface[name]

		if !ok {
			log.Printf("interface with name %s not found\n", name)
			continue
		}

		state, err := d.GetPropertyState()

		if err != nil {
			return err
		}

		if state == gonetworkmanager.NmDeviceStateActivated {
			err = c.updateIpAddressesNmcli(d, intface)

			if err != nil {
				return err
			}
		}

		switch state {
		case gonetworkmanager.NmDeviceStateActivated:
			intface.State = models.StateConnected
		case gonetworkmanager.NmDeviceStateDisconnected:
			intface.State = models.StateDisconnected
		default:
			intface.State = models.StateUnknown
		}

		devType, err := d.GetPropertyDeviceType()

		if err != nil {
			return err
		}

		switch devType {
		case gonetworkmanager.NmDeviceTypeBt:
			intface.Type = models.TypeBluetooth
		case gonetworkmanager.NmDeviceTypeEthernet:
			intface.Type = models.TypeEthernet
		case gonetworkmanager.NmDeviceTypeBridge:
			intface.Type = models.TypeBridge
		case gonetworkmanager.NmDeviceTypeVeth:
			intface.Type = models.TypeVirtual
		case gonetworkmanager.NmDeviceTypeTun:
			intface.Type = models.TypeTunnel
		case gonetworkmanager.NmDeviceTypeWifi:
			intface.Type = models.TypeWireless
		}
		mtu, err := d.GetPropertyMtu()

		if err != nil {
			return err
		}

		intface.Mtu = int(mtu)

		x, x1 := d.MarshalJSON()
		d.GetPropertyPhysicalPortId()
		fmt.Println(string(x), x1)
	}
	return nil
}

func (c *NetworkCollector) updateIpAddressesNmcli(d gonetworkmanager.Device, intface *models.Interface) error {
	ipv4Config, err := d.GetPropertyIP4Config()

	if err != nil {
		return err
	}

	iad4, err := ipv4Config.GetPropertyAddressData()

	if err != nil {
		return err
	}
	for _, adddress := range iad4 {
		intface.IPv4 = append(intface.IPv4, models.IPAddress{
			Address: net.ParseIP(adddress.Address),
			MaskLen: adddress.Prefix,
		})
	}

	ipv6Config, err := d.GetPropertyIP6Config()

	if err != nil {
		return err
	}

	iad6, err := ipv6Config.GetPropertyAddressData()

	if err != nil {
		return err
	}
	for _, adddress := range iad6 {
		intface.IPv6 = append(intface.IPv6, models.IPAddress{
			Address: net.ParseIP(adddress.Address),
			MaskLen: adddress.Prefix,
		})
	}
	return nil
}

func (c *NetworkCollector) Collect(uuid string) (interface{}, error) {
	network, err := ghw.Network()

	if err != nil {
		return nil, err
	}

	localhostIpv4 := net.ParseIP("127.0.0.1")
	localhostIpv6 := net.ParseIP("::1")

	interfaces := []*models.Interface{
		{
			Name:    "lo",
			Virtual: true,
			Mac:     "00:00:00:00:00:00",
			IPv4: []models.IPAddress{
				{
					Address: localhostIpv4,
					MaskLen: 8,
				},
			},
			IPv6: []models.IPAddress{
				{
					Address: localhostIpv6,
					MaskLen: 128,
				},
			},
			State:        models.StateConnected,
			Type:         models.TypeLoopback,
			Speed:        "unkown",
			Mtu:          65536,
			Capabilities: []models.NicCapability{},
		},
	}
	nameInterface := map[string]*models.Interface{}

	for _, n := range network.NICs {
		capabilities := []models.NicCapability{}
		// TODO: add capabilities
		netInterface := &models.Interface{
			Name:         n.Name,
			Virtual:      n.IsVirtual,
			Mac:          n.MacAddress,
			State:        "",
			IPv4:         nil,
			IPv6:         nil,
			Type:         models.TypeUnknown,
			Speed:        n.Speed,
			Capabilities: capabilities,
		}

		interfaces = append(interfaces, netInterface)
		nameInterface[n.Name] = netInterface
	}
	update := models.NetworkUpdate{
		BasicUpdate: models.BasicUpdate{
			Uuid: uuid,
		},
		Network: models.Network{
			Interfaces: interfaces,
		},
	}

	err = c.updateMapWithNmcli(nameInterface)
	return update, err
}

type LshwNetwork struct {
	Handle        string `json:"handle,omitempty"`
	Product       string `json:"product,omitempty"`
	Vendor        string `json:"vendor,omitempty"`
	Businfo       string `json:"businfo"`
	Version       string `json:"version,omitempty"`
	Width         int    `json:"width,omitempty"`
	Clock         int    `json:"clock,omitempty"`
	Configuration struct {
		Latency         string `json:"latency"`
		Wireless        string `json:"wireless"`
		Firmware        string `json:"firmware"`
		IP              string `json:"ip"`
		Autonegotiation string `json:"autonegotiation"`
		Broadcast       string `json:"broadcast"`
		Driver          string `json:"driver"`
		Driverversion   string `json:"driverversion"`
		Duplex          string `json:"duplex"`
		Link            string `json:"link"`
		Multicast       string `json:"multicast"`
		Port            string `json:"port"`
		Speed           string `json:"speed"`
	} `json:"configuration,omitempty"`
	Capabilities struct {
		One000Bt        string `json:"1000bt"`
		One000BtFd      string `json:"1000bt-fd"`
		One00Bt         string `json:"100bt"`
		One00BtFd       string `json:"100bt-fd"`
		One0Bt          string `json:"10bt"`
		One0BtFd        string `json:"10bt-fd"`
		Autonegotiation string `json:"autonegotiation"`
		Ethernet        bool   `json:"ethernet"`
		Mii             string `json:"mii"`
		Physical        string `json:"physical"`
		Tp              string `json:"tp"`
		Pm              string `json:"pm"`
		Msi             string `json:"msi"`
		Pciexpress      string `json:"pciexpress"`
		Msix            string `json:"msix"`
		BusMaster       string `json:"bus_master"`
		CapList         string `json:"cap_list"`
		Wireless        string `json:"wireless"`
	} `json:"capabilities,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	Capacity    int    `json:"capacity,omitempty"`
	Claimed     bool   `json:"claimed"`
	Class       string `json:"class"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Logicalname string `json:"logicalname"`
	Physid      string `json:"physid"`
	Serial      string `json:"serial"`
	Size        int    `json:"size,omitempty"`
	Units       string `json:"units,omitempty"`
}

type IPInfo struct {
	Ifindex   int      `json:"ifindex"`
	Ifname    string   `json:"ifname"`
	Flags     []string `json:"flags"`
	Mtu       int      `json:"mtu"`
	Qdisc     string   `json:"qdisc"`
	Operstate string   `json:"operstate"`
	Group     string   `json:"group"`
	Txqlen    int      `json:"txqlen,omitempty"`
	LinkType  string   `json:"link_type"`
	Address   string   `json:"address,omitempty"`
	Broadcast string   `json:"broadcast,omitempty"`
	AddrInfo  []struct {
		Family            string `json:"family"`
		Local             string `json:"local"`
		Prefixlen         int    `json:"prefixlen"`
		Scope             string `json:"scope"`
		Label             string `json:"label,omitempty"`
		ValidLifeTime     int64  `json:"valid_life_time"`
		PreferredLifeTime int64  `json:"preferred_life_time"`
	} `json:"addr_info"`
	LinkIndex   int    `json:"link_index,omitempty"`
	Master      string `json:"master,omitempty"`
	LinkNetnsid int    `json:"link_netnsid,omitempty"`
}
