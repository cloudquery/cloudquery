{{- range .Tables -}}
{{.Name -}}:
  {{- range .Columns}}
  - field_name: {{.Name}}
    data_type: {{.Type}}
  {{- end -}}
{{end }}