{{if .Parent}}  {{.ParentFieldName}}: r{{.NestingLevel}}.{{.ParentFieldName}},
{{template "resolve_parent_vars.go.tpl" .Parent}}{{end}}