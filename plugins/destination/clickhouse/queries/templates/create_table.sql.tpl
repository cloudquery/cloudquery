{{with .Table -}}
CREATE TABLE {{.Name | sanitize}} (
{{with .Columns}}{{template "col_defs.sql.tpl" .}}{{end}}
) ENGINE = MergeTree
{{end -}}
ORDER BY ({{with .SortingKey}}{{template "col_names.sql.tpl" .}}{{end}});