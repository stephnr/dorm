package dorm

import (
	"github.com/aws/aws-sdk-go/aws"
)

const (
	dormTag     = "dorm"
	dormAttrTag = "dormav"
)

type Table struct {
	aws     *awsMeta
	model   *metadata
	options *TableOptions
}

type TableLoadOptions struct {
	TableName string
	Type      interface{}
	Options   *TableOptions
	AwsConfig *aws.Config
}

type TableOptions struct {
	ForceCreate bool
	ReadUnits   *int64 // if both ReadUnits & WriteUnits is nil, ON_DEMAND billing is used
	WriteUnits  *int64
}

func Load(opts TableLoadOptions) *Table {
	return newTable(opts)
}

func newTable(opts TableLoadOptions) *Table {
	tbl := &Table{
		aws: newAwsMeta(opts.AwsConfig),
		model: &metadata{
			tableName: opts.TableName,
			raw:       opts.Type,
		},
		options: opts.Options,
	}

	tbl.loadKeyParams()
	tbl.validateExistence()

	return tbl
}
