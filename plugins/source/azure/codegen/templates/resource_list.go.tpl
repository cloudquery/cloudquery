{{template "base.go.tpl" .}}

func fetch{{.AzureService}}{{.AzureSubService}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{ .AzureService }}.{{ .AzureSubService }}
	{{ range .ListFunctionArgsInit }}
	{{.}}{{ end }}
	response, err := svc.{{ or .ListFunction "ListAll" }}(ctx{{ range .ListFunctionArgs }}, {{.}}{{ end }})
	{{ or .ListHandler `
	if err != nil {
		return err
	}
	
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	`}}
	return nil
}

{{if .GetFunction}}
func {{.Table.PreResourceResolver}}(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	svc := meta.(*client.Client).Services().{{ .AzureService }}.{{ .AzureSubService }}
	{{ range .GetFunctionArgsInit }}
	{{.}}{{ end }}
	item, err := svc.{{.GetFunction}}(ctx{{ range .GetFunctionArgs }}, {{.}}{{ end }})
	if err != nil {
		return err
	}
	r.SetItem(item)
	return nil
}
{{end}}