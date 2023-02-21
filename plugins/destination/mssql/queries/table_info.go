package queries

import (
	"database/sql"

	"github.com/cloudquery/plugin-sdk/schema"
)

func GetTableSchema(schemaName string, table *schema.Table) (query string, params []any) {
	// https://stackoverflow.com/a/58995161
	const sqlSelectColumnTypes = `
SELECT
    [name] = c.[name],
    [type] = CASE
		WHEN tp.[name] IN ('varchar', 'char', 'varbinary')
		THEN tp.[name] + '(' + IIF(c.max_length = -1, 'max', CAST (c.max_length AS VARCHAR (25))) + ')'
		WHEN tp.[name] IN ('nvarchar', 'nchar')
		THEN tp.[name] + '(' + IIF(c.max_length = -1, 'max', CAST (c.max_length / 2 AS VARCHAR (25)))+ ')'
		WHEN tp.[name] IN ('decimal', 'numeric')
		THEN tp.[name] + '(' + CAST (c.[precision] AS VARCHAR (25)) + ', ' + CAST (c.[scale] AS VARCHAR (25)) + ')'
		WHEN tp.[name] IN ('datetime2')
		THEN tp.[name] + '(' + CAST (c.[scale] AS VARCHAR (25)) + ')'
		ELSE tp.[name]
    END,
    [nullable] = c.is_nullable
FROM sys.tables t 
JOIN sys.schemas s ON t.schema_id = s.schema_id
JOIN sys.columns c ON t.object_id = c.object_id
JOIN sys.types tp ON c.user_type_id = tp.user_type_id
WHERE s.[name] = @schemaName AND t.[name] = @tableName;`
	return sqlSelectColumnTypes, []any{
		sql.Named("tableName", table.Name),
		sql.Named("schemaName", schemaName),
	}
}

func GetTablePK(schemaName string, table *schema.Table) (query string, params []any) {
	const sqlSelectPKColumns = `SELECT
    [name] = tc.name,
    [order] = ic.key_ordinal
FROM sys.schemas s
    INNER JOIN sys.tables t ON s.schema_id = t.schema_id
    INNER JOIN sys.indexes i ON t.object_id = i.object_id
    INNER JOIN sys.index_columns ic ON i.object_id = ic.object_id AND i.index_id = ic.index_id
    INNER JOIN sys.columns tc ON ic.object_id = tc.object_id AND ic.column_id = tc.column_id
WHERE s.[name] = @schemaName AND t.[name] = @tableName AND i.is_primary_key = 1
ORDER BY ic.key_ordinal`

	return sqlSelectPKColumns, []any{
		sql.Named("tableName", table.Name),
		sql.Named("schemaName", schemaName),
	}
}

func TableExists(schemaName string, table *schema.Table) (query string, params []any) {
	const tableExistsQuery = `SELECT COUNT(*)
FROM sys.tables t
    INNER JOIN sys.schemas s ON t.schema_id = s.schema_id
WHERE s.[name] = @schemaName AND t.[name] = @tableName`

	return tableExistsQuery, []any{
		sql.Named("tableName", table.Name),
		sql.Named("schemaName", schemaName),
	}
}
