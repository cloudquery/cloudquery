{{template "base.go.tpl" .}}

func fetch{{.AzureService}}{{.AzureSubService}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{ .AzureService }}.{{ .AzureSubService }}
	{{ range .ListFunctionArgsInit }}
	{{.}}{{ end }}
	response, err := svc.{{ or .ListFunction "ListAll" }}(ctx{{ range .ListFunctionArgs }}, {{.}}{{ end }})
	{{ or .ListHandler `
	if err != nil {
		return errors.WithStack(err)
	}
	
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}
	`}}
	return nil
}