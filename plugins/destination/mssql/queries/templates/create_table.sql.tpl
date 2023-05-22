CREATE TABLE {{sanitizeID .Schema .Table}} (
{{with .Columns}}{{template "col_defs.sql.tpl" .}}{{end}}
{{- with .PrimaryKey}}
  {{- if .Columns }},
  CONSTRAINT {{.Name}} PRIMARY KEY (
{{template "col_names.sql.tpl" .Columns}}
  )
  {{- end}}
{{- end}}
);