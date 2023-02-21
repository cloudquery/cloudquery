CREATE TYPE {{.Type}} AS TABLE (
{{with .Columns}}{{template "col_defs.sql.tpl" .}}{{end}}
);