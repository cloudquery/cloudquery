# Schema Changes from v1 to v2
This guide summarizes schema changes from CloudQuery v1 to v2. It is automatically generated and
not guaranteed to be complete, but we hope it helps as a starting point and reference when migrating to v2.

Last updated {{$.Date}}.
{{range $table := $.Tables }}
## {{$table.Name}}
{{- if eq $table.Status "removed" }}
{{- if not $table.Comment }}
This table was removed.
{{- end }}
{{- end }}

{{- if eq $table.Status "added" }}
This table was newly added.
{{- end }}
{{- if $table.Comment }}
{{ $table.Comment }}
{{- end }}
{{if and (ne $table.Status "removed") (and (ne $table.Status "moved") (ne $table.Status "renamed")) }}
| Name          | Type          | Status | Comment
| ------------- | ------------- | --------------- | ---------------
{{- range $col := $table.Columns }}
{{- if $col.Status}}
|{{ $col.Name }}|{{ $col.Type }}|{{ $col.Status }}|{{ $col.Comment }}
{{- end}}
{{- end }}
{{- end }}
{{ end }}