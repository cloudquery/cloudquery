{{if .Parent}}
	r{{.NestingLevel}} := parent.Item.({{.Parent.AWSStructName}})
{{if .Parent.Parent}}  parent = parent.Parent
{{template "resolve_parent_defs.go.tpl" .Parent}}
{{end}}
{{end}}