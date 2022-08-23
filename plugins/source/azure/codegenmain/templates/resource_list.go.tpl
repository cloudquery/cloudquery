// Auto generated code - DO NOT EDIT.

package {{.AzurePackageName}}

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/pkg/errors"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.AzureService}}{{.AzureSubService}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.AzureService}}{{.AzureSubService}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().{{ .AzureService }}.{{ .AzureSubService }}
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