{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}},
  {{end -}}
  [src].{{.}}
{{- end -}}