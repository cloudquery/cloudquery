SELECT
{{with .Table.Columns.Names}}{{template "col_names.sql.tpl" .}}{{end}}
FROM {{sanitizeID .Schema .Table.Name }}
WHERE {{.SourceNameColumn}} = @sourceName
ORDER BY {{.SyncTimeColumn}} ASC;