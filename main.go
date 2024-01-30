package main

import (
	"fmt"
	"runtime"
	"syscall"
)

func main() {
	// Method 1: Using runtime.MemStats
	var vmem runtime.MemStats
	runtime.ReadMemStats(&vmem)
	totalMem := vmem.Sys / 1024 / 1024 // convert bytes to megabytes
	memLimit := totalMem / 2
	fmt.Println("Total memory (runtime.MemStats): ", totalMem, "MB")
	fmt.Println("Memory limit (runtime.MemStats): ", memLimit, "MB")

	// Method 2: Using syscall.Sysinfo
	var sysInfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysInfo)
	if err != nil {
		fmt.Println("Error getting system info: ", err)
		return
	}
	totalSysMem := sysInfo.Totalram * uint64(syscall.Getpagesize()) / 1024 / 1024 // convert to megabytes
	fmt.Println("Total memory (syscall.Sysinfo): ", totalSysMem, "MB")
}


