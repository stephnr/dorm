package dorm

func (db *DB) LoadTable(tableName string, model interface{}, opts *TableOptions) *Table {
	return newTable(db.awsConfig, TableLoadOptions{
		TableName: tableName,
		Type:      model,
		Options:   opts,
	})
}
