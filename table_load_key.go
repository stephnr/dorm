package dorm

import (
	"fmt"
	"reflect"
)

func (tbl *Table) loadKeyParams() {
	typ := reflect.TypeOf(tbl.model.Raw)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		switch tag := field.Tag.Get(dormTag); tag {
		case hashParamTag:
			fetchHashParam(tbl, field)
			break
		case rangeParamTag:
			fetchRangeParam(tbl, field)
			break
		}
	}

	// panic if no hash param is defined
	if tbl.model.HashParam != nil && tbl.model.HashParam.Name == "" {
		panic(fmt.Sprintf("dorm: No HASH param was found in (%s) model", tbl.model.TableName))
	}
}

func fetchHashParam(tbl *Table, field reflect.StructField) {
	tableName := tbl.model.TableName
	param := tbl.model.HashParam

	if param == nil {
		param = &paramType{}
	}

	paramTagVal := field.Name
	attrOverride := field.Tag.Get(dormAttrTag)

	if param.Name != "" {
		panic(fmt.Sprintf("dorm: Failed to load table (%s): duplicate %s properties found - (%s, %s)", tableName, hashParamTag, param, paramTagVal))
	}

	if attrOverride != "" {
		param.Name = attrOverride
	} else {
		param.Name = paramTagVal
	}

	param.Type = convertFieldToDynamoType(field)
	tbl.model.HashParam = param
}

func fetchRangeParam(tbl *Table, field reflect.StructField) {
	tableName := tbl.model.TableName
	param := tbl.model.RangeParam

	if param == nil {
		param = &paramType{}
	}

	paramTagVal := field.Name
	attrOverride := field.Tag.Get(dormAttrTag)

	if param == nil && param.Name != "" {
		panic(fmt.Sprintf("dorm: Failed to load table (%s): duplicate %s properties found - (%s, %s)", tableName, rangeParamTag, param, paramTagVal))
	}

	if attrOverride != "" {
		param.Name = attrOverride
	} else {
		param.Name = paramTagVal
	}

	param.Type = convertFieldToDynamoType(field)
	tbl.model.RangeParam = param
}

func convertFieldToDynamoType(field reflect.StructField) dataType {
	switch field.Type.Kind() {
	case reflect.String:
		return dataTypeString

	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return dataTypeNumber
	default:
		return dataTypeUnknown
	}
}
