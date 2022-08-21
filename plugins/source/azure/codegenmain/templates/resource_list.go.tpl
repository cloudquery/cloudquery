// Auto generated code - DO NOT EDIT.

package {{.AzureService}}

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/pkg/errors"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.AzureService | ToCamel}}{{.AzureSubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.AzureService | ToCamel}}{{.AzureSubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{ .AzureService | ToCamel }}.{{ .AzureSubService | ToCamel }}
	response, err := svc.{{ .ListFunction }}(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}