// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"{{.ImportPath}}"
)

func {{.Name | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

{{if not .SkipFetch}}
func fetch{{.Name | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	{{- if .NewFuncHasSubscriptionId}}
  svc, err := {{.Service}}.{{.NewFuncName}}(cl.SubscriptionId, cl.Creds, cl.Options)
  {{- else}}
  svc, err := {{.Service}}.{{.NewFuncName}}(cl.Creds, cl.Options)
  {{- end}}
	if err != nil {
    return err
  }
	{{- if .ListFuncHasSubscriptionId}}
	pager := svc.{{.ListFuncName}}(cl.ResourceGroup, nil)
	{{- else }}
	pager := svc.{{.ListFuncName}}(nil)
	{{- end }}
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
{{end}}