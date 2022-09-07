{{if .Parent}}  {{.ChildFieldName | Coalesce .ParentFieldName}}: {{if .ParentFieldName | HasDollar -}}
{{.ParentFieldName | ReplaceDollar (print "r" .NestingLevel) }},
{{else -}}
  r{{.NestingLevel}}.{{.ParentFieldName}},
{{end -}}
{{template "resolve_parent_vars.go.tpl" .Parent}}{{end}}