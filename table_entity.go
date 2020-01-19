package dorm

type dataType string

const (
	hashParamTag  = "HASH"
	rangeParamTag = "RANGE"

	dataTypeBinary  dataType = "B"
	dataTypeNumber  dataType = "N"
	dataTypeString  dataType = "S"
	dataTypeUnknown dataType = ""
)

type metadata struct {
	TableName  string
	HashParam  *paramType
	RangeParam *paramType
	Raw        interface{}
}

type paramType struct {
	Name string
	Type dataType
}

func (en metadata) Key() {
	// typ := reflect.TypeOf(en)
	// val := reflect.ValueOf(en)

	// for i := 0; i < typ.NumField(); i++ {
	// 	field := val.Field(i)
	// 	name := field.Type().Name()

	// 	fmt.Printf("%+v\n", field)
	// 	fmt.Printf("%+v\n", name)
	// }
}
