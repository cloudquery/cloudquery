{{if .Parent}}  {{.ChildFieldName | Coalesce .ParentFieldName}}: r{{.NestingLevel}}.{{.ParentFieldName}},
{{template "resolve_parent_vars.go.tpl" .Parent}}{{end}}