// Auto generated code - DO NOT EDIT.

package {{.AzureService | ToLower}}

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/pkg/errors"
    {{template "imports.go.tpl" .}}
)

func {{.AzureSubService}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.AzureService}}{{.AzureSubService}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
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