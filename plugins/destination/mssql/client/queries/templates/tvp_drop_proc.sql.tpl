IF EXISTS (
 SELECT * FROM sys.procedures p
 INNER JOIN sys.schemas s ON p.schema_id = s.schema_id
 WHERE s.[name] = @schemaName AND p.[name] = @procName
)
DROP PROCEDURE {{.Name}};