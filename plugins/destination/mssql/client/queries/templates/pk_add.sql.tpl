ALTER TABLE {{.Table}} ADD CONSTRAINT {{.Name}}
  PRIMARY KEY (
{{with .Columns}}{{template "col_names.sql.tpl" .}}{{end}}
  );