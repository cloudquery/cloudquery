{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}},
  {{end -}}
  {{.Name | sanitize}} {{.Type}}
{{- end -}}