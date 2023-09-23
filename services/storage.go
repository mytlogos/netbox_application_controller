package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/mytlogos/netbox_application_controller/models"
	"github.com/mytlogos/netbox_application_controller/netbox"
	"github.com/mytlogos/netbox_application_controller/services/backend"
	"golang.org/x/exp/maps"
)

var (
	ErrHostNotFound  = fmt.Errorf("could not find host")
	ErrDuplicateHost = fmt.Errorf("duplicate host")
)

type Storage struct {
	backend  *backend.Backend
	mappings map[string]*models.Host
}

func NewStorage(backend *backend.Backend) *Storage {
	return &Storage{
		backend:  backend,
		mappings: make(map[string]*models.Host),
	}
}

func (s *Storage) Init() error {
	maps.Clear(s.mappings)

	devices, err := s.backend.GetDevices()

	if err != nil {
		return err
	}
	for _, device := range devices {
		hostMap, ok := device.CustomFields[backend.CustomFieldAgent]

		if !ok {
			continue
		}

		// ensure we got a map which can be mapped to a struct
		if _, ok = hostMap.(map[string]interface{}); !ok {
			continue
		}

		data, err := json.Marshal(hostMap)

		if err != nil {
			log.Println(err)
			continue
		}

		var host models.Host
		err = json.Unmarshal(data, &host)

		if err != nil {
			log.Println(err)
			continue
		}
		s.mappings[host.Uuid] = &host
	}
	return nil
}

func (s *Storage) AddAgent(uuid string, value models.Agent) error {
	_, ok := s.mappings[uuid]

	if ok {
		return ErrDuplicateHost
	}
	s.mappings[uuid] = &models.Host{
		Uuid:  uuid,
		Agent: value,
	}
	return nil
}

func (s *Storage) Get(uuid string) (*models.Host, error) {
	data, ok := s.mappings[uuid]

	if !ok {
		return nil, ErrHostNotFound
	}
	return data, nil
}

func (s *Storage) GetAll() ([]*models.Host, error) {
	values := []*models.Host{}
	for _, h := range s.mappings {
		values = append(values, h)
	}
	return values, nil
}

func (s *Storage) getHost(uuid string) (*models.Host, error) {
	value, ok := s.mappings[uuid]

	if !ok {
		return nil, ErrHostNotFound
	}
	return value, nil
}

func (s *Storage) UpdateDevice(uuid string, value *models.Device) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	var copy models.Host = *h
	h.Device = value
	s.MakeUpdate(&copy, h)
	return true, nil
}

func (s *Storage) UpdateNetwork(uuid string, value *models.Network) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	var copy models.Host = *h
	h.Network = value
	s.MakeUpdate(&copy, h)
	return true, nil
}

func (s *Storage) UpdateApp(uuid string, value *models.App) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	var copy models.Host = *h
	h.App = value
	s.MakeUpdate(&copy, h)
	return true, nil
}

func (s *Storage) UpdateCpu(uuid string, value *models.Cpu) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	h.Cpu = value
	return true, nil
}

func (s *Storage) UpdateRam(uuid string, value *models.Ram) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	var copy models.Host = *h
	h.Ram = value
	s.MakeUpdate(&copy, h)
	return true, nil
}

func (s *Storage) UpdateDisk(uuid string, value *models.Disk) (bool, error) {
	h, err := s.getHost(uuid)

	if err != nil {
		return false, err
	}

	h.Disk = value
	return true, nil
}

func (s *Storage) makeDeviceUpdate(value *models.Host) (*netbox.DeviceWithConfigContext, error) {
	platform, err := s.ensurePlatformExists(value)

	if err != nil {
		return nil, err
	}

	device, err := s.backend.GetDeviceByUuid(value.Uuid)

	if err != nil {
		if !errors.Is(err, ErrHostNotFound) {
			conv := NetboxConverter{}
			device = conv.ConvertHost(value)
			device.Platform.Set(netbox.NewNestedPlatform(platform.Id, platform.Url, platform.Display, platform.Name, platform.Slug))
			return s.backend.CreateDevice(device)
		} else {
			return nil, err
		}
	} else {
		conv := NetboxConverter{}
		conv.UpdateDevice(value, device)
		device.Platform.Set(netbox.NewNestedPlatform(platform.Id, platform.Url, platform.Display, platform.Name, platform.Slug))
		return s.backend.UpdateDevice(device.Id, device)
	}
}

func (s *Storage) ensurePlatformExists(value *models.Host) (*netbox.Platform, error) {
	devicePlatformName := fmt.Sprintf("%s %s", value.Device.Platform, value.Device.PlatformVersion)
	modelPlatform := netbox.Platform{
		Name: devicePlatformName,
		Slug: strings.ToLower(fmt.Sprintf("%s-%s", value.Device.Platform, strings.ReplaceAll(value.Device.PlatformVersion, ".", "-"))),
	}

	platforms, err := s.backend.GetPlatforms()

	if err != nil {
		return nil, err
	}

	for _, platform := range platforms {
		if strings.EqualFold(platform.Name, devicePlatformName) {
			modelPlatform = platform
			break
		}
	}

	if modelPlatform.Id == 0 {
		createdPlatform, err := s.backend.CreatePlatform(modelPlatform)

		if err != nil {
			return nil, err
		}

		modelPlatform = *createdPlatform
	}
	return &modelPlatform, nil
}

func (s *Storage) makeNetworkUpdate(value *models.Host, device *netbox.DeviceWithConfigContext) error {
	interfaces, err := s.backend.GetInterfaces(device.Id)

	if err != nil {
		return err
	}

	interfacesToUpdate, removeIds := UpdateInterfaces(value, device.Id, interfaces)

	for _, intface := range interfacesToUpdate {
		if intface.Id == 0 {
			_, err = s.backend.CreateInterface(device.Id, intface)
		} else {
			_, err = s.backend.UpdateInterface(intface)
		}
		if err != nil {
			return err
		}
	}

	for _, id := range removeIds {
		err = s.backend.DeleteInterface(id)

		if err != nil {
			return err
		}
	}

	interfaces, err = s.backend.GetInterfaces(device.Id)

	if err != nil {
		return err
	}
	ipaddresses, err := s.backend.GetIPAddressByDevice(device.Id)

	if err != nil {
		return err
	}

	ipsToUpdate, ipsToRemove := UpdateIPAddresses(value, interfaces, ipaddresses)

	for _, ip := range ipsToUpdate {

		if ip.Id != 0 {
			continue
		}
		_, err = s.backend.CreateIPAddress(ip)

		if err != nil {
			return err
		}
	}

	for _, id := range ipsToRemove {
		err = s.backend.DeleteIP(id)

		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) MakeUpdate(before, after *models.Host) {
	differ := NewDiffer(DifferConfig{
		Backend: s.backend,
	})
	changes, err := differ.DiffDocument(before, after)

	if err != nil {
		log.Println(err)
		return
	}

	// data, err := json.MarshalIndent(changes, "", " ")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	device, err := s.makeDeviceUpdate(after)

	if err != nil {
		panic(err)
	}

	err = s.makeNetworkUpdate(after, device)

	if err != nil {
		panic(err)
	}
	log.Println(changes)
}
