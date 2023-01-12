{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}{{else}}, {{end -}}{{. | sanitize}}
{{- end -}}