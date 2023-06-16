CREATE TYPE {{sanitizeID .Schema .Type}} AS TABLE (
{{with .Table.Columns}}{{template "col_defs.sql.tpl" .}}{{end}}
);