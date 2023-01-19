CREATE TABLE {{.Table}} (
{{with .Definitions}}{{template "col_defs.sql.tpl" .}}{{end}}
  PRIMARY KEY (
{{with .PrimaryKey}}{{template "col_names.sql.tpl" .}}{{end}}
  )
);