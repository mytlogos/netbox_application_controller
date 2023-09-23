package models

type BasicUpdate struct {
	Uuid string
}

type DeviceUpdate struct {
	BasicUpdate
	Device
}

type NetworkUpdate struct {
	BasicUpdate
	Network
}

type CpuUpdate struct {
	BasicUpdate
	Cpu
}

type RamUpdate struct {
	BasicUpdate
	Ram
}

type DiskUpdate struct {
	BasicUpdate
	Disk
}

type AppUpdate struct {
	BasicUpdate
	App
}
