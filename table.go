package dorm

import "github.com/aws/aws-sdk-go/aws"

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
}

type TableOptions struct {
	ForceCreate bool
	ReadUnits   *int64 // if both ReadUnits & WriteUnits is nil, ON_DEMAND billing is used
	WriteUnits  *int64
}

func newTable(awsCfg *aws.Config, opts TableLoadOptions) *Table {
	tbl := &Table{
		aws: newAwsMeta(awsCfg),
		model: &metadata{
			TableName: opts.TableName,
			Raw:       opts.Type,
		},
		options: opts.Options,
	}

	tbl.loadKeyParams()
	tbl.validateExistence()

	return tbl
}
