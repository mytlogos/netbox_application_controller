package collectors

import (
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mytlogos/netbox_application_controller/models"
)

type RamCollector struct {
}

func (c *RamCollector) GetName() string {
	return "Ram"
}

func (c *RamCollector) GetEndpoint() string {
	return "/api/agent/ram"
}

func (c *RamCollector) GetLshwData() (uint64, []models.RamSlot, error) {
	cmd := exec.Command("lshw", "-c", "memory", "-json")
	rawData, err := cmd.Output()
	if err != nil {
		return 0, nil, err
	}
	var rawList []json.RawMessage
	err = json.Unmarshal(rawData, &rawList)

	if err != nil {
		return 0, nil, err
	}

	var systemMemory *LshwGenericMemory
	result := []models.RamSlot{}

	for _, rawItem := range rawList {
		var untypedItem map[string]interface{}
		err := json.Unmarshal(rawItem, &untypedItem)

		if err != nil {
			return 0, nil, err
		}

		itemId := untypedItem["id"].(string)
		description := untypedItem["description"].(string)

		if strings.HasPrefix(itemId, "bank:") {

			if description == "[empty]" {
				index, err := strconv.ParseUint(untypedItem["physid"].(string), 10, 64)

				if err != nil {
					return 0, nil, err
				}

				result = append(result, models.RamSlot{
					Index:    uint8(index),
					SlotName: untypedItem["slot"].(string),
				})
			} else {
				var item LshwRamSlot
				err = json.Unmarshal(rawItem, &item)

				if err != nil {
					return 0, nil, err
				}

				index, err := strconv.ParseUint(untypedItem["physid"].(string), 10, 64)

				if err != nil {
					return 0, nil, err
				}

				result = append(result, models.RamSlot{
					Index:    uint8(index),
					SlotName: item.Slot,
					Content: &models.RamSlotContent{
						Description: item.Description,
						Product:     item.Product,
						Vendor:      item.Vendor,
						SizeInBytes: uint64(item.Size),
						Width:       uint8(item.Width),
						ClockHertz:  uint64(item.Clock),
					},
				})
			}
		} else if itemId == "memory" && strings.EqualFold(description, "system memory") {
			if systemMemory != nil {
				log.Println("duplicate system memory item - ignoring")
				continue
			}
			var genericItem LshwGenericMemory
			err = json.Unmarshal(rawItem, &genericItem)

			if err != nil {
				return 0, nil, err
			}
			systemMemory = &genericItem
		}
	}
	var systemMemoryBytes uint64 = 0

	if systemMemory != nil {
		systemMemoryBytes = uint64(systemMemory.Size)
	}
	return systemMemoryBytes, result, nil
}

func (c *RamCollector) Collect(uuid string) (interface{}, error) {
	systemTotalBytes, slots, err := c.GetLshwData()

	if err != nil {
		return nil, err
	}

	return models.RamUpdate{
		BasicUpdate: models.BasicUpdate{
			Uuid: uuid,
		},
		Ram: models.Ram{
			TotalSystemBytes: systemTotalBytes,
			Slots:            slots,
		},
	}, nil
}

type LshwGenericMemory struct {
	ID          string `json:"id"`
	Class       string `json:"class"`
	Claimed     bool   `json:"claimed,omitempty"`
	Description string `json:"description"`
	Physid      string `json:"physid"`
	Units       string `json:"units,omitempty"`
	Size        int    `json:"size,omitempty"`
}

type LshwRamSlot struct {
	ID          string `json:"id"`
	Class       string `json:"class"`
	Claimed     bool   `json:"claimed"`
	Handle      string `json:"handle"`
	Description string `json:"description"`
	Product     string `json:"product"`
	Vendor      string `json:"vendor"`
	Physid      string `json:"physid"`
	Serial      string `json:"serial"`
	Slot        string `json:"slot"`
	Units       string `json:"units"`
	Size        int64  `json:"size"`
	Width       int    `json:"width"`
	Clock       int    `json:"clock"`
}
