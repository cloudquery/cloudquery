IF EXISTS (
 SELECT * FROM sys.table_types tt
 INNER JOIN sys.schemas s ON tt.schema_id = s.schema_id
 WHERE s.[name] = @schemaName AND tt.[name] = @typeName
)
DROP TYPE {{.Type}};