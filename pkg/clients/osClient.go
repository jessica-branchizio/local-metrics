package clients

import (
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

type OSClient interface {
	GetCPU() (*load.AvgStat, error)
	GetMem() (*mem.VirtualMemoryStat, error)
}

func (c *Client) GetCPU() (*load.AvgStat, error) {
	return load.Avg()
}

func (c *Client) GetMem() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

type Client struct {
	MemIteration int
	CpuIteration int
	OSClient     *OSClient
}

func NewOSClient(memIteration, cpuIteration int) *Client {
	return &Client{
		MemIteration: memIteration,
		CpuIteration: cpuIteration,
	}
}
