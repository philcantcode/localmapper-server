package core

import (
	"github.com/philcantcode/localmapper/utils"
)

type Capability struct {
	ID            int
	Command       Command
	Type          string
	Name          string
	Desc          string
	DisplayFields []string
}

type Command struct {
	Program string
	Params  []Param
}

type Param struct {
	Flag  string
	Value string

	MetaType    DataType
	MetaDefault string
	MetaInfo    string
}

type Proposition struct {
	ID          int
	Type        string
	Date        string
	Description string
	Proposition PropositionItem
	Correction  PropositionItem
	Status      int
	User        int
}

type PropositionItem struct {
	Name     string
	Value    string
	DataType DataType
	Options  []string
}

/* OSI Layer:
   7 - Firewalls, IDS
   6 -
   5 -
   4 - Firewalls (some)
   3 - Routers, L3 switches
   2 - L2 switches, Bridges
   1 - Hubs, repeaters, modems
*/
type CMDBItem struct {
	ID          int
	OSILayer    int      // 1 - 7
	DateSeen    []string //[] Array of dates seen
	Description string
	StatusTags  map[string]string // [Stopped, Running] etc
	UserTags    map[string]string // [Project-X, Bob's Server] etc
	InfoTags    map[string]string // [IP: xxx, MAC: xxx, URL: xxx] etc
}

type DataType int

const (
	String DataType = iota
	IP
	IPRange
	MAC
	Integer
	Decimal
	Bool
	None
)

func ReverseDataTypeLookup(datType DataType) string {
	switch datType {
	case 0:
		return "String"
	case 1:
		return "IP"
	case 2:
		return "IPRange"
	case 3:
		return "MAC"
	case 4:
		return "Integer"
	case 5:
		return "Decimal"
	case 6:
		return "Bool"
	case 7:
		return "None"
	default:
		utils.ErrorForceFatal("Couldn't do a reverse lookup for DataType (definitions)")
	}

	return "nil"
}

func ParamsToArray(params []Param) []string {
	var paramArr []string

	for _, param := range params {
		// If the flag is NOT empty, add the flag
		if param.Flag != "" {
			paramArr = append(paramArr, param.Flag)
		}

		// If the MetaType is NOT 'none' and the value is NOT empty, add the value
		if param.MetaType != None && param.Value != "" {
			paramArr = append(paramArr, param.Value)
		}

		// If the MetaType and Value are empty, use the default
		if param.MetaType != None && param.Value == "" && param.MetaDefault != "" {
			paramArr = append(paramArr, param.MetaDefault)
		}

	}

	return paramArr
}
