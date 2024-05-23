package clients

import (
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type OSClient interface {
	GetCPU() (*load.AvgStat, error)
	GetMem() (*mem.VirtualMemoryStat, error)
}

func (moc *MacOSClient) GetCPU() (*load.AvgStat, error) {
	return load.Avg()
}

func (moc *MacOSClient) GetMem() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

type MacOSClient struct{}

func (loc *LinuxOSClient) GetCPU() (*load.AvgStat, error) {
	return load.Avg()
}

func (loc *LinuxOSClient) GetMem() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

type LinuxOSClient struct{}

func NewMacOSClient() OSClient {
	return &MacOSClient{}
}

func NewLinuxOSClient() OSClient {
	return &LinuxOSClient{}
}
