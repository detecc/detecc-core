package main

import (
	"fmt"
	"github.com/detecc/deteccted/plugin"
	cpp "github.com/shirou/gopsutil/cpu"
	mem2 "github.com/shirou/gopsutil/mem"
	"log"
	"time"
)

func init() {
	hwClientPlugin := &HardwareMonitorPlugin{}
	plugin.Register(hwClientPlugin.GetCmdName(), hwClientPlugin)

	go hwClientPlugin.Schedule(LogCpuMemUsage, time.Minute)
}

type HardwareMonitorPlugin struct {
	plugin.Handler
}

func (e HardwareMonitorPlugin) GetCmdName() string {
	return "/get-hw-status"
}

func (e HardwareMonitorPlugin) Execute(args interface{}) (interface{}, error) {
	log.Println(args)
	return GetHardwareInfo()
}

func (e HardwareMonitorPlugin) GetMetadata() plugin.Metadata {
	return plugin.Metadata{Type: plugin.PluginTypeClientServer}
}

// Schedule a function for periodic execution.
func (e HardwareMonitorPlugin) Schedule(f func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		f()
	}
	ticker.Stop()
}

func GetHardwareInfo() (hwInfo map[string]float64, err error) {
	hwInfo = make(map[string]float64)
	cpuUsage, err := cpp.Percent(time.Second, false)
	if err != nil {
		return hwInfo, err
	}

	mem, err := mem2.VirtualMemory()
	if err != nil {
		return hwInfo, err
	}

	hwInfo["cpu"] = cpuUsage[0]
	hwInfo["mem-available"] = float64(mem.Available / 1024 / 1024)
	hwInfo["mem-total"] = float64(mem.Total / 1024 / 1024)
	hwInfo["mem-used"] = float64(mem.Used / 1024 / 1024)
	return hwInfo, err
}
func LogCpuMemUsage() {
	cpuUsage, _ := cpp.Percent(time.Second, false)
	mem, _ := mem2.VirtualMemory()

	fmt.Printf("       CPU Usage: %.2f\n", cpuUsage)
	fmt.Printf("Memory Available: %.2d\n", mem.Available/1024/1024)
	fmt.Printf("    Memory Total: %.2d\n", mem.Total/1024/1024)
	fmt.Printf("     Memory Used: %.2d\n", mem.Used/1024/1024)
}
