
{{. | indentToDepth}}- [{{.Name}}]({{.Name}}.md){{ if .IsIncremental}} (Incremental){{ end }}
{{- range $index, $rel := .Relations}}
{{- template "all_tables_entry.md.go.tpl" $rel}}
{{- end}}