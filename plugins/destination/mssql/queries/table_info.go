package queries

import (
	"database/sql"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

func Database() string {
	return `SELECT DB_NAME();`
}

func AllTables(schemaName string) (query string, params []any) {
	return `SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = @schemaName AND TABLE_TYPE = 'BASE TABLE';`, []any{sql.Named("schemaName", schemaName)}
}

func GetTableSchema(schemaName string, databaseName string, table *schema.Table) (query string, params []any) {
	tableSchema := fmt.Sprintf(`SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, CHARACTER_MAXIMUM_LENGTH FROM %s.INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA = @schemaName AND TABLE_NAME = @tableName;`, databaseName)
	return tableSchema, []any{sql.Named("tableName", table.Name), sql.Named("schemaName", schemaName)}
}

func GetTablePK(schemaName string, databaseName string, table *schema.Table) (query string, params []any) {
	pks := fmt.Sprintf(`SELECT Col.COLUMN_NAME from %s.INFORMATION_SCHEMA.TABLE_CONSTRAINTS Tab, %s.INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE Col
			WHERE Col.CONSTRAINT_NAME = Tab.CONSTRAINT_NAME
			AND Col.TABLE_NAME = Tab.TABLE_NAME
			AND Tab.CONSTRAINT_TYPE = 'PRIMARY KEY'
			AND Col.TABLE_NAME = @tableName
			AND Tab.Table_Schema = @schemaName;
	`, databaseName, databaseName)

	return pks, []any{sql.Named("tableName", table.Name), sql.Named("schemaName", schemaName)}
}
