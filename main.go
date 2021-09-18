package main

import (
	"fmt"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-sysinfo"
)

func main() {
	host, err := sysinfo.Host()
	checkErr(err)
	hostInfo := host.Info()
	fmt.Println("arch:", hostInfo.Architecture)
	fmt.Println("boot time:", hostInfo.BootTime)
	if hostInfo.Containerized != nil {
		fmt.Println("containerized:", *hostInfo.Containerized)
	}
	fmt.Println("hostname:", hostInfo.Hostname)
	fmt.Println("ips:")
	for _, ip := range hostInfo.IPs {
		fmt.Println("\t", ip)
	}
	fmt.Println("kernel:", hostInfo.KernelVersion)
	fmt.Println("macs:")
	for _, mac := range hostInfo.MACs {
		fmt.Println("\t", mac)
	}
	fmt.Println("os:", hostInfo.OS.Name, hostInfo.OS.Version)
	fmt.Println("timezone:", hostInfo.Timezone)
	fmt.Println("timezone offset:", hostInfo.TimezoneOffsetSec)
	fmt.Println("id:", hostInfo.UniqueID)
	fmt.Println("uptime:", hostInfo.Uptime())
	fmt.Println()

	memory, err := host.Memory()
	checkErr(err)
	fmt.Println("available mem:", humanize.Bytes(memory.Available))
	fmt.Println("free mem:", humanize.Bytes(memory.Free))
	fmt.Println("total mem:", humanize.Bytes(memory.Total))
	fmt.Println("used mem:", humanize.Bytes(memory.Used))
	fmt.Println("free virt mem:", humanize.Bytes(memory.VirtualFree))
	fmt.Println("total virt mem:", humanize.Bytes(memory.VirtualTotal))
	fmt.Println("used virt mem:", humanize.Bytes(memory.VirtualUsed))
	fmt.Println()

	goInfo := sysinfo.Go()
	fmt.Println("go arch:", goInfo.Arch)
	fmt.Println("go cores:", goInfo.MaxProcs)
	fmt.Println("go os:", goInfo.OS)
	fmt.Println("go version:", goInfo.Version)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
