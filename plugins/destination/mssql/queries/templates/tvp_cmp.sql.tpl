{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}}
  AND
  {{end -}}
  {{sanitizeID "tgt" .}} = {{sanitizeID "src" .}}
{{- end -}}