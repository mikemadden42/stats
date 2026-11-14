// Package main displays operating system stats.
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-sysinfo"
)

func main() {
	log.SetFlags(0)

	host, err := sysinfo.Host()
	if err != nil {
		log.Fatal(err)
	}
	hostInfo := host.Info()
	fmt.Println("arch:", hostInfo.Architecture)
	fmt.Println("boot time:", humanize.Time(hostInfo.BootTime))
	if hostInfo.Containerized != nil {
		fmt.Println("containerized:", *hostInfo.Containerized)
	}
	fmt.Println("hostname:", hostInfo.Hostname)
	printList("ips", unique(hostInfo.IPs))
	fmt.Println("kernel:", hostInfo.KernelVersion)
	printList("macs", realMACs(hostInfo.MACs))
	fmt.Println("os:", hostInfo.OS.Name, hostInfo.OS.Version)
	fmt.Println("timezone:", hostInfo.Timezone)
	fmt.Println("timezone offset:", hostInfo.TimezoneOffsetSec)
	fmt.Println("id:", hostInfo.UniqueID)
	fmt.Println("uptime:", hostInfo.Uptime().Truncate(time.Second))
	fmt.Println()

	memory, err := host.Memory()
	if err != nil {
		log.Println("memory:", err)
	} else {
		fmt.Println("available mem:", humanize.IBytes(memory.Available))
		fmt.Println("free mem:", humanize.IBytes(memory.Free))
		fmt.Println("total mem:", humanize.IBytes(memory.Total))
		fmt.Println("used mem:", humanize.IBytes(memory.Used))
		if memory.VirtualTotal > 0 {
			fmt.Println("free virtual mem:", humanize.IBytes(memory.VirtualFree))
			fmt.Println("total virtual mem:", humanize.IBytes(memory.VirtualTotal))
			fmt.Println("used virtual mem:", humanize.IBytes(memory.VirtualUsed))
		}
	}
	fmt.Println()

	goInfo := sysinfo.Go()
	fmt.Println("go arch:", goInfo.Arch)
	fmt.Println("go cores:", goInfo.MaxProcs)
	fmt.Println("go os:", goInfo.OS)
	fmt.Println("go version:", goInfo.Version)
}

// printList prints a labeled, indented list, or nothing if items is empty.
func printList(label string, items []string) {
	if len(items) == 0 {
		return
	}
	fmt.Println(label + ":")
	for _, item := range items {
		fmt.Printf("\t%s\n", item)
	}
}

// unique returns items with duplicates removed, keeping the first occurrence.
func unique(items []string) []string {
	seen := make(map[string]bool, len(items))
	var out []string
	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			out = append(out, item)
		}
	}
	return out
}

// realMACs drops the all-zero and 02:00:00:00:00:00 placeholders that some
// operating systems (such as macOS) report in place of hidden addresses.
func realMACs(macs []string) []string {
	var out []string
	for _, mac := range unique(macs) {
		if mac != "00:00:00:00:00:00" && mac != "02:00:00:00:00:00" {
			out = append(out, mac)
		}
	}
	return out
}
