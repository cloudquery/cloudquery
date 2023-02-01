CREATE TABLE {{.Table}} (
{{with .Definitions}}{{template "col_defs.sql.tpl" .}}{{end}}
{{- with .PrimaryKey}}
  CONSTRAINT {{.Name}} PRIMARY KEY (
{{template "col_names.sql.tpl" .Columns}}
{{- end}}
  )
);