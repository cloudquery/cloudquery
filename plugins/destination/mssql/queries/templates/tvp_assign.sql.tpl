{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}},
  {{end -}}
  [tgt].{{.}} = [src].{{.}}
{{- end -}}