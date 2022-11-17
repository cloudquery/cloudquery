{{- if .IsPager -}}
    {{ template "mock_test_pager.go.tpl" . }}
{{- else -}}
    {{- if .IsList -}}
        {{ template "mock_test_list.go.tpl" . }}
    {{- else -}}
        {{ template "mock_test_get.go.tpl" . }}
    {{- end -}}
{{- end -}}