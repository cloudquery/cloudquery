# Source Plugin: {{.PluginName}}
## Tables
{{- range $table := $.Tables }}
{{- template "all_tables_entry.md.go.tpl" $table}}
{{- end }}