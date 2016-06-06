package osgraphql

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/process"
)

// Process process struct wrapper
type Process struct {
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Rss  uint64 `json:"rss"`
	Vms  uint64 `json:"vms"`
	Swap uint64 `json:"swap"`
}

// SystemInfo system info api
type SystemInfo interface {
	GetProcesses() ([]Process, error)
	GetProcessesByName(nameQuery string) ([]Process, error)
	GetCPUInfo() ([]cpu.InfoStat, error)
	GetPartitions(all bool) ([]*disk.UsageStat, error)
}

// NewLocalSystemInfo create a new system info api
func NewLocalSystemInfo() SystemInfo {
	return &localSystemInfo{}
}

type localSystemInfo struct {
}

func (lsi *localSystemInfo) GetProcesses() ([]Process, error) {
	return getProcesses(func(name string) bool { return true })
}

func (lsi *localSystemInfo) GetProcessesByName(nameQuery string) ([]Process, error) {
	return getProcesses(func(name string) bool { return nameQuery == name })
}

func (lsi *localSystemInfo) GetCPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func (lsi *localSystemInfo) GetPartitions(all bool) ([]*disk.UsageStat, error) {
	parts, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}

	usagestats := []*disk.UsageStat{}

	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)

		if err != nil {
			return nil, err
		}

		usagestats = append(usagestats, usage)
	}

	return usagestats, nil
}

func getProcesses(nameMatcher func(name string) bool) ([]Process, error) {

	log.Printf("get processes")

	pids, err := process.Pids()

	if err != nil {
		return nil, err
	}

	procs := []Process{}

	for _, pid := range pids {

		p, err := process.NewProcess(pid)
		if err != nil {
			continue
		}

		name, err := p.Name()
		if err != nil {
			continue
		}

		// if name matcher returns false skip the proc
		if !nameMatcher(name) {
			continue
		}

		memoryInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}

		procs = append(procs, Process{
			Pid:  pid,
			Name: name,
			Rss:  memoryInfo.RSS,
			Vms:  memoryInfo.VMS,
			Swap: memoryInfo.Swap,
		})
	}

	return procs, nil

}
