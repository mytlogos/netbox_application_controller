package models

import (
	"net"
	"time"
)

type Host struct {
	Uuid    string   `diff:"uuid"`
	Agent   Agent    `diff:"agent"`
	Device  *Device  `diff:"device"`
	Network *Network `diff:"network"`
	Cpu     *Cpu     `diff:"cpu"`
	Ram     *Ram     `diff:"ram"`
	Disk    *Disk    `diff:"disk"`
	App     *App     `diff:"app"`
}

type Agent struct {
	DeploymentDate time.Time  `diff:"deploymentdate"`
	LastUpdate     *time.Time `diff:"lastupdate"`
	State          string     `diff:"state"`
	Version        string     `diff:"version"`
}

type Device struct {
	Name            string `diff:"name"`
	HostId          string `diff:"hostid"`
	Arch            string `diff:"arch"`
	KernelVersion   string `diff:"kernel"`
	Platform        string `diff:"platform"`
	PlatformFamily  string `diff:"platformfamily"`
	PlatformVersion string `diff:"platformversion"`
	OS              string `diff:"os"`
	Product         string `diff:"product"`
	Boottime        uint64 `diff:"-"`
	Uptime          uint64 `diff:"-"`
}

type NicCapability struct {
	Name      string `json:"name"`
	IsEnabled bool   `json:"is_enabled"`
	CanEnable bool   `json:"can_enable"`
}

type InterfaceType string

const (
	TypeEthernet  InterfaceType = "ethernet"
	TypeVirtual   InterfaceType = "virtual"
	TypeWireless  InterfaceType = "wireless"
	TypeBluetooth InterfaceType = "bluetooth"
	TypeLoopback  InterfaceType = "loopback"
	TypeBridge    InterfaceType = "bridge"
	TypeTunnel    InterfaceType = "tunnel"
	TypeUnknown   InterfaceType = "unknown"
)

type InterfaceState string

const (
	StateDisabled     InterfaceState = "disabled"
	StateDisconnected InterfaceState = "disconnected"
	StateConnected    InterfaceState = "connected"
	StateUnknown      InterfaceState = "unknown"
)

type IPAddress struct {
	Address net.IP
	MaskLen uint8
}

type Interface struct {
	Name         string
	Virtual      bool
	Mac          string
	IPv4         []IPAddress
	IPv6         []IPAddress
	Mtu          int
	State        InterfaceState
	Type         InterfaceType
	Speed        string
	Capabilities []NicCapability
}

type Network struct {
	Interfaces []*Interface
}

type Cpu struct {
}

type Ram struct {
	TotalSystemBytes uint64
	Slots            []RamSlot
}

type RamSlot struct {
	Index    uint8
	SlotName string
	Content  *RamSlotContent
}

type RamSlotContent struct {
	Description string
	Product     string
	Vendor      string
	SizeInBytes uint64
	Width       uint8
	ClockHertz  uint64
}

type DiskType string

const (
	DiskHDD   DiskType = "hdd"
	DiskSSD   DiskType = "ssd"
	DiskFlash DiskType = "flash"
	DiskOther DiskType = "other"
)

type FileSystemType string

const (
	FsExt4  FileSystemType = "ext4"
	FsBtrfs FileSystemType = "btrfs"
	FsFat   FileSystemType = "fat"
	FsvFat  FileSystemType = "vfat"
	FsNtfs  FileSystemType = "ntfs"
	FsSwap  FileSystemType = "swap"
	FsNone  FileSystemType = "none"
	FsOther FileSystemType = "other"
)

type PartitionType string

const (
	PartitionGpt PartitionType = "gpt"
	PartitionDos PartitionType = "dos"
)

type StoragePartition struct {
	FileSystemType FileSystemType
	LogicalName    string
	Position       int
	SizeBytes      uint64
	FsTotalBytes   uint64 `json:"total"`
	FsFreeBytes    uint64 `json:"free"`
	FsUsedBytes    uint64 `json:"used"`
}

type StorageItem struct {
	Manufacturer  string
	Product       string
	DiskType      DiskType
	LogicalName   string
	Serial        string
	SizeBytes     uint64
	PartitionType PartitionType
	Partitions    []StoragePartition
}

type Disk struct {
	Disks []StorageItem
}

type App struct {
	Applications []ApplicationGroup `diff:"applications"`
}

type ApplicationGroup struct {
	Name           string        `diff:"name"`
	Applications   []Application `diff:"applications"`
	AdditionalInfo interface{}   `diff:"info"`
}

type DockerGroupInfo struct {
	IsCompose   bool   `diff:"isCompose"`
	ComposeFile string `diff:"composeFile"`
	WorkingDir  string `diff:"workinDir"`
}

type ApplicationManager string

const (
	ManagerUnknown    ApplicationManager = "unknown"
	ManagerSystemd    ApplicationManager = "systemd"
	ManagerDocker     ApplicationManager = "docker"
	ManagerPodman     ApplicationManager = "podman"
	ManagerKubernetes ApplicationManager = "kubernetes"
)

type TransportProtocol string

const (
	TransportTcp TransportProtocol = "tcp"
	TransportUdp TransportProtocol = "udp"
)

type ApplicationStatus string

const (
	StatusUnknown ApplicationStatus = "unknown"
	StatusRunning ApplicationStatus = "running"
	StatusStopped ApplicationStatus = "stopped"
	StatusFailed  ApplicationStatus = "failed"
)

type ApplicationPort struct {
	Port              int               `diff:"port"`
	TransportProtocol TransportProtocol `diff:"transport"`
	InterfaceAddress  string            `diff:"interface"`
}

type Application struct {
	Name           string             `diff:"name"`
	AppManager     ApplicationManager `diff:"manager"`
	Status         ApplicationStatus  `diff:"status"`
	Ports          []ApplicationPort  `diff:"ports"`
	CpuLimits      int                `diff:"cpuLimits"`
	MemoryLimits   uint64             `diff:"memoryLimits"`
	AdditionalInfo interface{}        `diff:"info"`
}

type DockerInfo struct {
	ContainerId   string
	ContainerName string `diff:"containerName"`
}
