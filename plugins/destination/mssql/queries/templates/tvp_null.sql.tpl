{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}}
  AND
  {{end -}}
  {{sanitizeID "tgt" .}} IS NULL
{{- end -}}