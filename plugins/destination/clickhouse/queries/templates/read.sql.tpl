SELECT {{with .Columns}}{{template "col_names.sql.tpl" .}}{{end}}
FROM {{.Table | sanitize}}
WHERE {{.SourceNameColumn | sanitize}} = @sourceName
ORDER BY {{.SyncTimeColumn | sanitize}};