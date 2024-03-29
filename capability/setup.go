package capability

import (
	"github.com/philcantcode/localmapper/system"
	"go.mongodb.org/mongo-driver/bson"
)

func Init() {

	// smurf6 := Capability{
	// 	Type:       "thc-ipv6",
	// 	CCI:        "cci:thc-ipv6:smurf6:default",
	// 	Name:       "Smurf6 Denial of Service (DoS)",
	// 	Desc:       "Attempts to DoS a host.",
	// 	Category:   system.DDOS,
	// 	ResultTags: []string{},
	// 	Command: Command{
	// 		Program: "atk6-smurf6",
	// 		Params: []Param{
	// 			{
	// 				Desc:     "UDP Scan",
	// 				Flag:     "-sU",
	// 				DataType: []system.DataType{system.EMPTY},
	// 				Value:    "",
	// 				Default:  "",
	// 			},
	// 			{
	// 				Desc:     "Run Script",
	// 				Flag:     "--script",
	// 				DataType: []system.DataType{system.STRING},
	// 				Value:    "nbstat.nse",
	// 				Default:  "nbstat.nse",
	// 			},
	// 			{
	// 				Desc:     "Target",
	// 				Flag:     "",
	// 				DataType: []system.DataType{system.CIDR, system.IP},
	// 				Value:    "",
	// 				Default:  "",
	// 			},
	// 			{
	// 				Desc:     "Port 137",
	// 				Flag:     "-p137",
	// 				DataType: []system.DataType{system.EMPTY},
	// 				Value:    "",
	// 				Default:  "",
	// 			},
	// 			{
	// 				Desc:     "XML Output",
	// 				Flag:     "-oX",
	// 				DataType: []system.DataType{system.STRING},
	// 				Value:    "-",
	// 				Default:  "-",
	// 			},
	// 		},
	// 	},
	// }

	searchsploitNMAP := Capability{
		Interpreter: system.Interpreter_SEARCHSPLOIT,
		CCI:         "cci:searchsploit:nmap:json",
		Label:       "Searchsploit Query NMAP Result",
		Description: "Performs a searchsploit query on an XML NMAP result file.",
		Category:    system.Category_RESEARCH,
		ResultTags:  []string{},
		Hidden:      true,
		Command: Command{
			Program: "searchsploit",
			Params: []Param{
				{
					Description: "NMAP XML FIle",
					Flag:        "--nmap",
					DataType:    []system.DataType{system.DataType_FILE_PATH},
					Value:       "",
					Default:     "",
				},
				{
					Description: "JSON Output",
					Flag:        "--json",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
			},
		},
	}

	hydraSNMP := Capability{
		Interpreter: system.Interpreter_UNIVERSAL,
		CCI:         "cci:hydra:bruteforce:snmp",
		Label:       "Brute Force SNMP",
		Description: "Brute force against SNMP using various wordlists.",
		Category:    system.Category_BRUTEFORCE,
		ResultTags:  []string{},
		Hidden:      false,
		Command: Command{
			Program: "hydra",
			Params: []Param{
				{
					Description: "Password File",
					Flag:        "-P",
					DataType:    []system.DataType{system.DataType_FILE_PATH},
					Value:       "",
					Default:     "/localmapper/wordlists/passwords/500-worst-passwords.txt",
					Options: []ParamOpt{
						{
							Label:     "Passwords banned on Twitter",
							Value:     "/localmapper/wordlists/passwords/twitter-banned.txt",
							FileSize:  "2.8 KB",
							RiskLevel: 7,
						},
						{
							Label:     "500 Worst Passwords",
							Value:     "/localmapper/wordlists/passwords/500-worst-passwords.txt",
							FileSize:  "3.5 KB",
							RiskLevel: 7,
						},
						{
							Label:     "John the Ripper Password List",
							Value:     "/localmapper/wordlists/passwords/john.txt",
							FileSize:  "22 KB",
							RiskLevel: 7,
						},
						{
							Label:     "RockYou Password List",
							Value:     "/localmapper/wordlists/passwords/rockyou.txt",
							FileSize:  "134 MB",
							RiskLevel: 7,
						},
					},
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Protocol Type",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_PROTOCOL, system.DataType_STRING},
					Value:       "snmp",
					Default:     "snmp",
				},
			},
		},
	}

	hydraSSH := Capability{
		Interpreter: system.Interpreter_UNIVERSAL,
		CCI:         "cci:hydra:bruteforce:ssh",
		Label:       "Brute Force SSH",
		Description: "Brute force against SSH using various wordlists.",
		Category:    system.Category_BRUTEFORCE,
		Preconditions: []Precondition{
			{
				Label:       "Ports",
				Description: "Check that the SSH port is open",
				DataType:    system.DataType_PORT,
				Values:      []string{"22"},
			},
		},
		ResultTags: []string{},
		Hidden:     false,
		Command: Command{
			Program: "hydra",
			Params: []Param{
				{
					Description: "Username Wordlist",
					Flag:        "-L",
					DataType:    []system.DataType{system.DataType_FILE_PATH},
					Value:       "",
					Default:     "/localmapper/wordlists/usernames/top-usernames-shortlist.txt",
					Options: []ParamOpt{
						{
							Label:     "Top Usernames (shortlist)",
							Value:     "/localmapper/wordlists/usernames/top-usernames-shortlist.txt",
							FileSize:  "118 Bytes",
							RiskLevel: 7,
						},
						{
							Label:     "Top Usernames",
							Value:     "/localmapper/wordlists/usernames/usernames.txt",
							FileSize:  "814 KB",
							RiskLevel: 7,
						},
						{
							Label:     "Top Usernames from 10 Million List",
							Value:     "/localmapper/wordlists/usernames/xato-net-10-million-usernames-dup.txt",
							FileSize:  "5 MB",
							RiskLevel: 7,
						},
						{
							Label:     "10 Million Usernames",
							Value:     "/localmapper/wordlists/usernames/xato-net-10-million-usernames-full.txt",
							FileSize:  "82 MB",
							RiskLevel: 7,
						},
					},
				},
				{
					Description: "Password Wordlist",
					Flag:        "-P",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "",
					Default:     "/localmapper/wordlists/passwords/500-worst-passwords.txt",
					Options: []ParamOpt{
						{
							Label:     "Passwords banned on Twitter",
							Value:     "/localmapper/wordlists/passwords/twitter-banned.txt",
							FileSize:  "2.8 KB",
							RiskLevel: 7,
						},
						{
							Label:     "500 Worst Passwords",
							Value:     "/localmapper/wordlists/passwords/500-worst-passwords.txt",
							FileSize:  "3.5 KB",
							RiskLevel: 7,
						},
						{
							Label:     "John the Ripper Password List",
							Value:     "/localmapper/wordlists/passwords/john.txt",
							FileSize:  "22 KB",
							RiskLevel: 7,
						},
						{
							Label:     "RockYou Password List",
							Value:     "/localmapper/wordlists/passwords/rockyou.txt",
							FileSize:  "134 MB",
							RiskLevel: 7,
						},
					},
				},
				{
					Description: "Number of Parallel Tasks",
					Flag:        "-t",
					DataType:    []system.DataType{system.DataType_INTEGER},
					Value:       "",
					Default:     "4",
				},
				{
					Description: "Verbose",
					Flag:        "-V",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Protocol Type",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_PROTOCOL},
					Value:       "ssh",
					Default:     "ssh",
				},
			},
		},
	}

	nbtScan := Capability{
		Interpreter: system.Interpreter_NBTSCAN,
		CCI:         "cci:kali:nbtscan:default",
		Label:       "NBTScan NetBIOS Scanner",
		Description: "Scan networks for NetBIOS name information.",
		Category:    system.Category_DISCOVERY,
		ResultTags:  []string{},
		Command: Command{
			Program: "nbtscan",
			Params: []Param{
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_IP, system.DataType_CIDR},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Output Style - Comma separated",
					Flag:        "-s",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       ",",
					Default:     ",",
				},
			},
		},
	}

	netBiosScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:nbstat-netbios-script:default",
		Label:       "nbstat NetBIOS",
		Description: "Attempts to retrieve the target's NetBIOS names and MAC address.",
		ResultTags:  []string{"MAC"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "UDP Scan",
					Flag:        "-sU",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Run Script",
					Flag:        "--script",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "nbstat.nse",
					Default:     "nbstat.nse",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Port 137",
					Flag:        "-p137",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	sysDNSScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:sys-dns:default",
		Label:       "System DNS Scan",
		Description: "Use system DNS resolver configured on this host to identify private hostnames.",
		ResultTags:  []string{"HostName"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "Disable Port Scan Flag",
					Flag:        "-sn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "System DNS Flag",
					Flag:        "--system-dns",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	pingSweep := Capability{
		Interpreter: system.Interpreter_NMAP,
		Label:       "Ping Sweep",
		CCI:         "cci:nmap:ping-sweep:default",
		Description: "Perform a discovery Ping Sweep against an IP Range.",
		ResultTags:  []string{"IP"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "Disable Port Scan Flag",
					Flag:        "-sn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "IP CIDR Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	stealthScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:stealth:default",
		Label:       "Stealth Scan",
		Description: "Scan thousands of ports on the target device.",
		ResultTags:  []string{"Ports"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "Stealth Scan Flag",
					Flag:        "-sS",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Disable Ping Flag",
					Flag:        "-Pn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	osIdent := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:os-ident:default",
		Label:       "OS Identification Scan",
		Description: "Attempts to identify the operating system of the host.",
		ResultTags:  []string{"OS", "OSGen", "OSVendor", "OSAccuracy"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "OS Scan Flag",
					Flag:        "-O",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Disable Ping Flag",
					Flag:        "-Pn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				}, {
					Description: "Slows Down Scan",
					Flag:        "--max-rate",
					DataType:    []system.DataType{system.DataType_INTEGER},
					Value:       "100",
					Default:     "100",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	connectScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:tcp-connect:default",
		Label:       "TCP Connect Scan",
		Description: "TCP Connect Scan performs a full connection to the host.",
		ResultTags:  []string{"Ports"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "Connect Scan Flag",
					Flag:        "-sT",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Disable Ping Flag",
					Flag:        "-Pn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "IP Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	arpScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		Label:       "APR Scan",
		CCI:         "cci:nmap:arp:default",
		Description: "Perform a scan of the local network using ARP.",
		ResultTags:  []string{"IP"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "Disable Port Scan Flag",
					Flag:        "-sn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "ARP Flag",
					Flag:        "-PU",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "IP Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	// nmap IP -sU -sS --script smb-os-discovery.nse -p U:137,T:139

	smbScriptScan := Capability{
		Interpreter: system.Interpreter_NMAP,
		Label:       "SMB OS Discovery",
		CCI:         "cci:nmap:smb-os-discovery:default",
		Description: "Attempts to determine the operating system, computer name, domain, workgroup, and current time over the SMB protocol (ports 445 or 139). This is done by starting a session with the anonymous account, in response to a session starting, the server will send back all this information.",
		ResultTags:  []string{"HostName"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "UDP Scan Flag",
					Flag:        "-sU",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Stealth Scan Flag",
					Flag:        "-sS",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Script Flag",
					Flag:        "--script",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "smb-os-discovery.nse",
					Default:     "smb-os-discovery.nse",
				},
				{
					Description: "Port Flag",
					Flag:        "-p",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "U:137,T:139",
					Default:     "U:137,T:139",
				},
				{
					Description: "IP Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	svcDetection := Capability{
		Interpreter: system.Interpreter_NMAP,
		CCI:         "cci:nmap:svc-detection:default",
		Label:       "Service Identification Scan",
		Description: "Attempts to identify the service version of running services the host.",
		ResultTags:  []string{"OS", "OSGen", "OSVendor", "OSAccuracy"},
		Category:    system.Category_DISCOVERY,
		Command: Command{
			Program: "nmap",
			Params: []Param{
				{
					Description: "OS Scan Flag",
					Flag:        "-sV",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				},
				{
					Description: "Disable Ping Flag",
					Flag:        "-Pn",
					DataType:    []system.DataType{system.DataType_EMPTY},
					Value:       "",
					Default:     "",
				}, {
					Description: "Slows Down Scan",
					Flag:        "--max-rate",
					DataType:    []system.DataType{system.DataType_INTEGER},
					Value:       "100",
					Default:     "100",
				},
				{
					Description: "Target",
					Flag:        "",
					DataType:    []system.DataType{system.DataType_CIDR, system.DataType_IP},
					Value:       "",
					Default:     "",
				},
				{
					Description: "XML Output",
					Flag:        "-oX",
					DataType:    []system.DataType{system.DataType_STRING},
					Value:       "-",
					Default:     "-",
				},
			},
		},
	}

	if len(SELECT_Capability(bson.M{}, bson.M{})) == 0 {
		INSERT_Capability(searchsploitNMAP)
		INSERT_Capability(netBiosScan)
		INSERT_Capability(sysDNSScan)
		INSERT_Capability(pingSweep)
		INSERT_Capability(osIdent)
		INSERT_Capability(connectScan)
		INSERT_Capability(stealthScan)
		INSERT_Capability(arpScan)
		INSERT_Capability(smbScriptScan)
		INSERT_Capability(svcDetection)
		INSERT_Capability(nbtScan)
		INSERT_Capability(hydraSNMP)
		INSERT_Capability(hydraSSH)
	}
}
