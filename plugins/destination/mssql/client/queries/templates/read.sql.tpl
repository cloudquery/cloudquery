SELECT
{{with .Columns}}{{template "col_names.sql.tpl" .}}{{end}}
FROM {{ .Table }}
WHERE {{ .CqSourceNameColumn }} = @sourceName
ORDER BY {{ .CqSyncTimeColumn }} ASC;