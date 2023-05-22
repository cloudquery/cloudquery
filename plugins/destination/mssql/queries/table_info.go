package queries

import (
	"database/sql"
)

func AllTables(schemaName string) (query string, params []any) {
	return `SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = @schemaName AND TABLE_TYPE = 'BASE TABLE';`, []any{sql.Named("schemaName", schemaName)}
}

func GetTableSchema(schemaName string, tableName string) (query string, params []any) {
	tableSchema := `SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = @schemaName AND TABLE_NAME = @tableName;`
	return tableSchema, []any{sql.Named("tableName", tableName), sql.Named("schemaName", schemaName)}
}

func GetTablePK(schemaName string, tableName string) (query string, params []any) {
	pks := `SELECT Col.COLUMN_NAME from INFORMATION_SCHEMA.TABLE_CONSTRAINTS Tab, INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE Col
			WHERE Col.CONSTRAINT_NAME = Tab.CONSTRAINT_NAME
			AND Col.TABLE_NAME = Tab.TABLE_NAME
			AND Tab.CONSTRAINT_TYPE = 'PRIMARY KEY'
			AND Col.TABLE_NAME = @tableName
			AND Tab.Table_Schema = @schemaName;`

	return pks, []any{sql.Named("tableName", tableName), sql.Named("schemaName", schemaName)}
}
