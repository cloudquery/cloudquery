{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}},
  {{end -}}
  {{sanitizeID "tgt" .}} = {{sanitizeID "src" .}}
{{- end -}}