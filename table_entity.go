package dorm

const (
	hashParamTag  = "HASH"
	rangeParamTag = "RANGE"
)

type metadata struct {
	tableName      string
	hashParamName  string
	rangeParamName string
	raw            interface{}
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
