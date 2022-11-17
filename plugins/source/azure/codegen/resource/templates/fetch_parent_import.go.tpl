{{ .Resource.StructPackageName }} "{{ .Resource.StructPackage }}"
{{- with .Parent}}
{{ template "fetch_parent_import.go.tpl" .}}
{{- end -}}
