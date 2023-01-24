{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}}
  AND
  {{end -}}
  [tgt].{{.}} IS NULL
{{- end -}}