{{- if .Required -}}
{{ .Resource.SubService | Singular | ToLowerCamel }} := parent{{ .Path }}.Item.(*{{ .Resource.StructPackageName }}.{{ .Resource.StructName }})
{{- end -}}
{{ with .Parent }}
    {{ template "fetch_parents.go.tpl" . }}
{{- end -}}
