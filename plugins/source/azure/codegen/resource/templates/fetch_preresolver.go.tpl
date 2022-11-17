{{if $pre := .Fetcher.PreResolver -}}
func get{{ .SubService | ToCamel }}(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
    c := meta.(*client.Client)
    svc := c.Services().{{ $pre.Service }}

    {{ with $pre -}}
    {{ .BasicStructName | ToLowerCamel}} := resource.Item.(*{{ .StructPackageName }}.{{ .BasicStructName }})
    id, err := arm.ParseResourceID(*{{ .BasicStructName | ToLowerCamel}}.ID)
    if err != nil {
        return err
    }
    {{- end }}

    {{ .StructName | ToLowerCamel }}, err := {{ template "fetch_call.go.tpl" $pre }}
    if err != nil {
        return err
    }

    resource.SetItem({{ .StructName | ToLowerCamel }})
    return nil
}
{{- end -}}