ALTER TABLE {{.Table}} ADD {{with .Definition}}{{template "col_def.sql.tpl" .}}{{end}};