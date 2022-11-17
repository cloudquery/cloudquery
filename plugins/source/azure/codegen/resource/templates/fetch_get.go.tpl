    {{ .StructName | ToLowerCamel }}, err := {{ template "fetch_call.go.tpl" . }}
    if err != nil {
        return err
    }
    res <- {{ .StructName | ToLowerCamel }}