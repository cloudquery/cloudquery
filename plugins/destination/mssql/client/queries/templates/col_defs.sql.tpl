{{- $first := true -}}
{{- range .}}
  {{- if $first}}{{$first = false}}  {{else}},
  {{end -}}
  {{template "col_def.sql.tpl" .}}
{{- end -}}