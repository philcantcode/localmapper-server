package blueprint

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
