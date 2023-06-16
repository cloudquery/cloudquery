CREATE PROCEDURE {{sanitizeID .Schema .Name}} @TVP {{sanitizeID .Schema .Type}} READONLY
AS
BEGIN
 SET NOCOUNT ON;
 UPDATE [tgt] WITH (UPDLOCK)
 SET
{{with .Values}}{{template "tvp_assign.sql.tpl" .}}{{end}}
 FROM {{sanitizeID .Schema .Table.Name}} AS [tgt]
 INNER JOIN @TVP AS [src]
 ON
{{with .Table.PrimaryKeys}}{{template "tvp_cmp.sql.tpl" .}}{{end}}
;

INSERT {{sanitizeID .Schema .Table.Name}} (
{{template "col_names.sql.tpl" .Table.Columns.Names}}
) SELECT
{{template "tvp_col_names.sql.tpl" .Table.Columns.Names}}
 FROM @TVP AS [src]
 LEFT JOIN {{sanitizeID .Schema .Table.Name}} AS [tgt] ON (
{{with .Table.PrimaryKeys}}{{template "tvp_cmp.sql.tpl" .}}{{end}}
 ) WHERE (
{{with .Table.PrimaryKeys}}{{template "tvp_null.sql.tpl" .}}{{end}}
);
END;