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

INSERT {{.Table}}
 SELECT * FROM @TVP AS [src]
 WHERE NOT EXISTS (
  SELECT 1 FROM {{.Table}} AS [tgt]
  WHERE (
{{with .PK}}{{template "tvp_cmp.sql.tpl" .}}{{end}}
  )
);
END;