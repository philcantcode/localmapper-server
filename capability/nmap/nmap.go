package nmap

import (
	"encoding/xml"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/philcantcode/localmapper/cmdb"
	"github.com/philcantcode/localmapper/system"
	"github.com/philcantcode/localmapper/utils"
)

// Execute takes a list of parameters to execute against NMAP
func Execute(params []string) NmapRun {
	system.Log(fmt.Sprintf("Attempting to run Nmap Command: %s > %v", "nmap", params), true)
	resultByte, err := exec.Command("nmap", params...).CombinedOutput()
	system.Fatal(fmt.Sprintf("Error returned running a command: %s > %v", "nmap", params), err)

	system.Log("Converting from []byte to NmapRun struct", false)

	var nmapRun NmapRun
	err = xml.Unmarshal(resultByte, &nmapRun)
	system.Error("Couldn't unmarshal result from Nmap console", err)

	interpret(nmapRun)

	return nmapRun
}

func interpret(nmapRun NmapRun) {
	// For each host
	for _, host := range nmapRun.Hosts {
		sysTags := []cmdb.EntryTag{}

		ports := cmdb.EntryTag{
			Label:    "Ports",
			Desc:     "Open Ports",
			DataType: system.INTEGER,
			Values:   []string{},
		}

		services := cmdb.EntryTag{
			Label:    "Services",
			Desc:     "Open Services",
			DataType: system.STRING,
			Values:   []string{},
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				if !utils.ArrayContains(strconv.Itoa(port.PortId), ports.Values) {
					ports.Values = append(ports.Values, strconv.Itoa(port.PortId))
				}

				if !utils.ArrayContains(port.Service.Name, services.Values) {
					services.Values = append(services.Values, port.Service.Name)
				}
			}
		}

		vendorTags := cmdb.EntryTag{
			Label:    "MACVendor",
			DataType: system.STRING,
			Values:   []string{},
		}

		osFamily := cmdb.EntryTag{
			Label:    "OS",
			DataType: system.STRING,
			Values:   []string{},
		}

		osGen := cmdb.EntryTag{
			Label:    "OSGen",
			DataType: system.STRING,
			Values:   []string{},
		}

		osAccuracy := cmdb.EntryTag{
			Label:    "OSAccuracy",
			DataType: system.INTEGER,
			Values:   []string{},
		}

		osVendor := cmdb.EntryTag{
			Label:    "OSVendor",
			DataType: system.STRING,
			Values:   []string{},
		}

		osCPE := cmdb.EntryTag{
			Label:    "CPE",
			DataType: system.STRING,
			Values:   []string{},
		}

		if len(host.Os.OsMatches) > 0 {
			match := host.Os.OsMatches[0]

			if match.Accuracy != "" {
				osAccuracy.Values = append(osAccuracy.Values, match.Accuracy)
			}

			for _, osClass := range match.OsClasses {
				if osClass.OsFamily != "" && !utils.ArrayContains(osClass.OsFamily, osFamily.Values) {
					osFamily.Values = append(osFamily.Values, osClass.OsFamily)
				}

				if osClass.OsGen != "" && !utils.ArrayContains(osClass.OsGen, osGen.Values) {
					osGen.Values = append(osGen.Values, osClass.OsGen)
				}

				if osClass.Vendor != "" && !utils.ArrayContains(osClass.Vendor, osVendor.Values) {
					osVendor.Values = append(osVendor.Values, osClass.Vendor)
				}

				for _, cpe := range osClass.CPEs {
					if !utils.ArrayContains(string(cpe), osCPE.Values) {
						osCPE.Values = append(osCPE.Values, string(cpe))
					}
				}
			}
		}

		for _, address := range host.Addresses {
			if address.AddrType == "ipv4" {
				sysTags = append(sysTags, cmdb.EntryTag{
					Label:    "IP",
					DataType: system.IP,
					Values:   []string{address.Addr},
				})
			}

			if address.AddrType == "mac" {
				sysTags = append(sysTags, cmdb.EntryTag{
					Label:    "MAC",
					DataType: system.MAC,
					Values:   []string{address.Addr},
				})
			}

			if address.Vendor != "" {
				vendorTags.Values = append(vendorTags.Values, address.Vendor)
			}
		}

		// Check that they're not empty.
		if len(vendorTags.Values) > 0 {
			sysTags = append(sysTags, vendorTags)
		}

		if len(osFamily.Values) > 0 {
			sysTags = append(sysTags, osFamily)
		}

		if len(osGen.Values) > 0 {
			sysTags = append(sysTags, osGen)
		}

		if len(osVendor.Values) > 0 {
			sysTags = append(sysTags, osVendor)
		}

		if len(osAccuracy.Values) > 0 {
			sysTags = append(sysTags, osAccuracy)
		}

		if len(osCPE.Values) > 0 {
			sysTags = append(sysTags, osCPE)
		}

		if len(ports.Values) > 0 {
			sysTags = append(sysTags, ports)
		}

		if len(services.Values) > 0 {
			sysTags = append(sysTags, services)
		}

		// Hostnames
		if len(host.Hostnames) > 0 {
			hostNameTag := cmdb.EntryTag{
				Label:    "HostName",
				DataType: system.STRING,
				Values:   []string{},
			}

			for _, name := range host.Hostnames {
				hostNameTag.Values = append(hostNameTag.Values, name.Name)
			}

			sysTags = append(sysTags, hostNameTag)
		}

		entry := cmdb.Entry{
			Label:    "Nmap Discovered Device",
			Desc:     "This device was discovered during an Nmap scan: " + nmapRun.Args,
			OSILayer: 0,
			CMDBType: cmdb.PENDING,
			DateSeen: []string{utils.Now()},
			SysTags:  sysTags,
		}

		tag, exists, _ := cmdb.FindSysTag("HostName", entry)

		if exists {
			entry.Label = tag.Values[0]
		}

		cmdb.UpdateOrInsert(entry)
	}
}
