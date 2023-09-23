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

type Disk struct {
}

type App struct {
	Applications []ApplicationGroup
}

type ApplicationGroup struct {
	Name           string
	Applications   []Application
	AdditionalInfo interface{}
}

type DockerGroupInfo struct {
	IsCompose   bool
	ComposeFile string
	WorkingDir  string
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
	Port              int
	TransportProtocol TransportProtocol
	InterfaceAddress  string
}

type Application struct {
	Name           string
	AppManager     ApplicationManager
	Status         ApplicationStatus
	Ports          []ApplicationPort
	CpuLimits      int
	MemoryLimits   uint64
	AdditionalInfo interface{}
}

type DockerInfo struct {
	ContainerId   string
	ContainerName string
}