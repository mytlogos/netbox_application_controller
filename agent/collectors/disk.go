package collectors

import (
	"encoding/json"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mytlogos/netbox_application_controller/models"
)

type DiskCollector struct {
}

func (c *DiskCollector) GetName() string {
	return "Disk"
}

func (c *DiskCollector) GetEndpoint() string {
	return "/api/agent/disk"
}

func (c *DiskCollector) getProduct(hostname string) (string, error) {
	cmd := exec.Command("lshw", "-c", "disk", "-json")
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	var lshw []map[string]interface{}
	err = json.Unmarshal(data, &lshw)

	if err != nil {
		// current lshw version on raspberry pi os "02.18.85-0.7" is more or less broken
		// for json, but the way is known for system, so try to fix it manually
		data = regexp.MustCompile(`]\s*$`).ReplaceAll(data, []byte("}]"))
		err = json.Unmarshal(data, &lshw)

		if err != nil {
			return "", err
		}
	}

	// try to get the product if available on an item with hostname as id
	// but do not try to force it
	// only certain devices have a product name
	for _, item := range lshw {
		id := item["id"].(string)

		if strings.EqualFold(id, hostname) {
			product, ok := item["product"]

			if !ok {
				return "", nil
			} else {
				return product.(string), nil
			}
		}
	}

	return "", nil
}

func getLsblkOutput() ([]LsblkDevice, error) {
	// exclude 7 excludes loop devices with major "7:x" and ram "devices" with major "1:x"
	cmd := exec.Command("lsblk", "--exclude", "7,1", "--json", "--output-all", "--tree", "--paths", "--bytes")
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	var result LsblkDevices
	err = json.Unmarshal(output, &result)
	return result.Blockdevices, err
}

func (c *DiskCollector) Collect(uuid string) (interface{}, error) {
	output, err := getLsblkOutput()

	if err != nil {
		return nil, err
	}

	items := []models.StorageItem{}

	for _, outputDisk := range output {
		disk := models.StorageItem{
			SizeBytes:     outputDisk.Size,
			PartitionType: models.PartitionType(outputDisk.Pttype),
			Serial:        outputDisk.Serial,
			DiskType:      models.DiskOther,
			LogicalName:   outputDisk.Name,
		}

		if strings.HasPrefix(outputDisk.Model, "Samsung") {
			disk.Product = strings.TrimPrefix(outputDisk.Model, "Samsung ")
			disk.Manufacturer = "Samsung"
		} else {
			disk.Manufacturer = strings.TrimSpace(outputDisk.Vendor)
			disk.Product = strings.TrimSpace(outputDisk.Model)
		}

		// this assumes that lsblk prints them in correct order
		for i, child := range outputDisk.Children {
			// ignore extended partitions
			if strings.ToLower(child.Parttypename) == "extended" {
				continue
			}
			partition := models.StoragePartition{
				LogicalName: child.Name,
				SizeBytes:   child.Size,
				Position:    i,
			}

			switch strings.ToLower(child.Fstype) {
			case "swap":
				partition.FileSystemType = models.FsSwap
			case "ext4":
				partition.FileSystemType = models.FsExt4
			case "fat":
				partition.FileSystemType = models.FsFat
			case "vfat":
				partition.FileSystemType = models.FsvFat
			case "ntfs":
				partition.FileSystemType = models.FsNtfs
			case "btrfs":
				partition.FileSystemType = models.FsBtrfs
			default:
				partition.FileSystemType = models.FsOther
			}

			// this assumes that either all or none of these properties are set
			if child.Fssize != nil && child.Fsavail != nil && child.Fsuse != nil {
				partition.FsFreeBytes = parseUint64(*child.Fsavail)
				partition.FsTotalBytes = parseUint64(*child.Fssize)
				partition.FsUsedBytes = parseUint64(*child.Fsused)
			} else {
				// if no fs data is available, just assume a full partition?
				partition.FsFreeBytes = 0
				partition.FsTotalBytes = child.Size
				partition.FsUsedBytes = child.Size
			}

			disk.Partitions = append(disk.Partitions, partition)
		}

		items = append(items, disk)
	}

	return models.DiskUpdate{
		BasicUpdate: models.BasicUpdate{
			Uuid: uuid,
		},
		Disk: models.Disk{
			Disks: items,
		},
	}, nil
}

type LsblkPartition struct {
	Name         string    `json:"name"`
	Kname        string    `json:"kname"`
	Path         string    `json:"path"`
	MajMin       string    `json:"maj:min"`
	Fsavail      *string   `json:"fsavail"`
	Fssize       *string   `json:"fssize"`
	Fstype       string    `json:"fstype"`
	Fsused       *string   `json:"fsused"`
	Fsuse        *string   `json:"fsuse%"`
	Fsroots      []string  `json:"fsroots"`
	Fsver        string    `json:"fsver"`
	Mountpoint   *string   `json:"mountpoint"`
	Mountpoints  []*string `json:"mountpoints"`
	Label        *string   `json:"label"`
	UUID         string    `json:"uuid"`
	Ptuuid       string    `json:"ptuuid"`
	Pttype       string    `json:"pttype"`
	Parttype     string    `json:"parttype"`
	Parttypename string    `json:"parttypename"`
	Partlabel    *string   `json:"partlabel"`
	Partuuid     string    `json:"partuuid"`
	Partflags    string    `json:"partflags"`
	Ra           int       `json:"ra"`
	Ro           bool      `json:"ro"`
	Rm           bool      `json:"rm"`
	Hotplug      bool      `json:"hotplug"`
	Size         uint64    `json:"size"`
	State        *string   `json:"state"`
	Owner        string    `json:"owner"`
	Group        string    `json:"group"`
	Mode         string    `json:"mode"`
	Alignment    int       `json:"alignment"`
	MinIo        int       `json:"min-io"`
	OptIo        int       `json:"opt-io"`
	PhySec       int       `json:"phy-sec"`
	LogSec       int       `json:"log-sec"`
	Rota         bool      `json:"rota"`
	Sched        string    `json:"sched"`
	RqSize       int       `json:"rq-size"`
	Type         string    `json:"type"`
	DiscAln      int       `json:"disc-aln"`
	DiscGran     int       `json:"disc-gran"`
	DiscMax      int       `json:"disc-max"`
	DiscZero     bool      `json:"disc-zero"`
	Wsame        int       `json:"wsame"`
	Wwn          string    `json:"wwn"`
	Rand         bool      `json:"rand"`
	Pkname       string    `json:"pkname"`
	Subsystems   string    `json:"subsystems"`
	Zoned        string    `json:"zoned"`
	Dax          bool      `json:"dax"`
}

type LsblkDevice struct {
	Name       string           `json:"name"`
	Kname      string           `json:"kname"`
	Path       string           `json:"path"`
	MajMin     string           `json:"maj:min"`
	Ptuuid     string           `json:"ptuuid"`
	Pttype     string           `json:"pttype"`
	Ra         int              `json:"ra"`
	Ro         bool             `json:"ro"`
	Rm         bool             `json:"rm"`
	Hotplug    bool             `json:"hotplug"`
	Model      string           `json:"model"`
	Serial     string           `json:"serial"`
	Size       uint64           `json:"size"`
	State      string           `json:"state"`
	Owner      string           `json:"owner"`
	Group      string           `json:"group"`
	Mode       string           `json:"mode"`
	Alignment  int              `json:"alignment"`
	MinIo      int              `json:"min-io"`
	OptIo      int              `json:"opt-io"`
	PhySec     int              `json:"phy-sec"`
	LogSec     int              `json:"log-sec"`
	Rota       bool             `json:"rota"`
	Sched      string           `json:"sched"`
	RqSize     int              `json:"rq-size"`
	Type       string           `json:"type"`
	DiscAln    int              `json:"disc-aln"`
	DiscGran   int              `json:"disc-gran"`
	DiscMax    int              `json:"disc-max"`
	DiscZero   bool             `json:"disc-zero"`
	Wsame      int              `json:"wsame"`
	Wwn        string           `json:"wwn"`
	Rand       bool             `json:"rand"`
	Hctl       string           `json:"hctl"`
	Tran       string           `json:"tran"`
	Subsystems string           `json:"subsystems"`
	Rev        string           `json:"rev"`
	Vendor     string           `json:"vendor"`
	Zoned      string           `json:"zoned"`
	Dax        bool             `json:"dax"`
	Children   []LsblkPartition `json:"children"`
}

type LsblkDevices struct {
	Blockdevices []LsblkDevice `json:"blockdevices"`
}
