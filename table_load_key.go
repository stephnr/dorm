package dorm

import (
	"fmt"
	"reflect"
)

func (tbl *Table) loadKeyParams() {
	typ := reflect.TypeOf(tbl.model.raw)

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
	if tbl.model.hashParamName == "" {
		panic(fmt.Sprintf("dorm: No HASH param was found in (%s) model", tbl.model.tableName))
	}
}

func fetchHashParam(tbl *Table, field reflect.StructField) {
	tableName := tbl.model.tableName
	currVal := tbl.model.hashParamName

	paramTagVal := field.Name
	attrOverride := field.Tag.Get(dormAttrTag)

	if currVal != "" {
		panic(fmt.Sprintf("dorm: Failed to load table (%s): duplicate %s properties found - (%s, %s)", tableName, hashParamTag, currVal, paramTagVal))
	}

	if attrOverride != "" {
		tbl.model.hashParamName = attrOverride
		return
	}

	tbl.model.hashParamName = paramTagVal
}

func fetchRangeParam(tbl *Table, field reflect.StructField) {
	tableName := tbl.model.tableName
	currVal := tbl.model.rangeParamName

	paramTagVal := field.Name
	attrOverride := field.Tag.Get(dormAttrTag)

	if currVal != "" {
		panic(fmt.Sprintf("dorm: Failed to load table (%s): duplicate %s properties found - (%s, %s)", tableName, rangeParamTag, currVal, paramTagVal))
	}

	if attrOverride != "" {
		tbl.model.rangeParamName = attrOverride
		return
	}

	tbl.model.rangeParamName = paramTagVal
}
