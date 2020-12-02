package gaugecol

import (
    "github.com/shirou/gopsutil/mem"
)

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetGaugeData() map[string]int {
    return map[string]int{
        "memtest": int(GetMemPercent()),
    }
}