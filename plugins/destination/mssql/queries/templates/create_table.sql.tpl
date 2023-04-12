CREATE TABLE {{.Table}} (
{{with .Definitions}}{{template "col_defs.sql.tpl" .}}{{end}}
{{- with .PrimaryKey}}
  {{- if ne (len .Columns) 0}}
    CONSTRAINT {{.Name}} PRIMARY KEY (
      {{template "col_names.sql.tpl" .Columns}}
    )
  {{- end}}
{{- end}}
);