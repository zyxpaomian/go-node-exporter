package countercol

import (
    "github.com/shirou/gopsutil/cpu"
    "time"    
)

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	return percent[0]
}

func GetCounterData() map[string]int {
    return map[string]int{
        "cputest": int(GetCpuPercent()),
    }
}