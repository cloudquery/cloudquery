# Table: {{$.Name}}

{{- if $.Title }}
This table shows data for {{$.Title}}.
{{- end }}

{{ $.Description }}
{{ $length := len $.PrimaryKeys -}}
{{ if eq $length 1 }}
The primary key for this table is **{{ index $.PrimaryKeys 0 }}**.
{{ else }}
The composite primary key for this table is ({{ range $index, $pk := $.PrimaryKeys -}}
	{{if $index }}, {{end -}}
		**{{$pk}}**
	{{- end -}}).
{{ end }}
{{- if $.IsIncremental -}}
It supports incremental syncs
{{- $ikLength := len $.IncrementalKeys -}}
{{- if eq $ikLength 1 }} based on the **{{ index $.IncrementalKeys 0 }}** column
{{- else if gt $ikLength 1 }} based on the ({{ range $index, $pk := $.IncrementalKeys -}}
	{{- if $index -}}, {{end -}}
		**{{$pk}}**
	{{- end -}}) columns
{{- end -}}.
{{- end -}}

{{- if or ($.Relations) ($.Parent) }}
## Relations
{{- end }}
{{- if $.Parent }}
This table depends on [{{ $.Parent.Name }}]({{ $.Parent.Name }}.md).
{{- end}}
{{ if $.Relations }}
The following tables depend on {{.Name}}:
{{- range $rel := $.Relations }}
  - [{{ $rel.Name }}]({{ $rel.Name }}.md)
{{- end }}
{{- end }}

## Columns
| Name          | Type          |
| ------------- | ------------- |
{{- range $column := $.Columns }}
|{{$column.Name}}{{if $column.PrimaryKey}} (PK){{end}}{{if $column.IncrementalKey}} (Incremental Key){{end}}|`{{$column.Type}}`|
{{- end }}