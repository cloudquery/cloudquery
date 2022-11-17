    data, err := {{ template "fetch_call.go.tpl" . }}
    if err != nil {
        return err
    }
{{- if .ValueField }}
    res <- data.{{ .ValueField }}
{{- else }}
    res <- data
{{- end }}