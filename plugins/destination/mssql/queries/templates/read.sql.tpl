SELECT
{{with .Columns}}{{template "col_names.sql.tpl" .}}{{end}}
FROM {{.Table}}
WHERE {{.SourceNameColumn}} = @sourceName
ORDER BY {{.SyncTimeColumn}} ASC;