CREATE PROCEDURE {{.Name}} @TVP {{.Type}} READONLY
AS
BEGIN
 SET NOCOUNT ON;
 UPDATE [tgt] WITH (UPDLOCK)
 SET
{{with .Values}}{{template "tvp_assign.sql.tpl" .}}{{end}}
 FROM {{.Table}} AS [tgt]
 INNER JOIN @TVP AS [src]
 ON
{{with .PK}}{{template "tvp_cmp.sql.tpl" .}}{{end}}
;

INSERT {{.Table}} (
{{template "col_names.sql.tpl" .ColumnNames}}
) SELECT
{{template "tvp_col_names.sql.tpl" .ColumnNames}}
 FROM @TVP AS [src]
 LEFT JOIN {{.Table}} AS [tgt] ON (
{{with .PK}}{{template "tvp_cmp.sql.tpl" .}}{{end}}
 ) WHERE (
{{with .PK}}{{template "tvp_null.sql.tpl" .}}{{end}}
);
END;