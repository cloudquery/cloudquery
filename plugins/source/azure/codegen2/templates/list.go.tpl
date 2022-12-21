// Code generated by codegen2; DO NOT EDIT.
package {{.PackageName}}

import (
	"github.com/cloudquery/plugin-sdk/schema"
	{{- if not .SkipFetch}}
	"context"
	"{{.ImportPath}}"
	{{- end}}
	{{- if not .ChildTable}}
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	{{- end}}
)

{{- if not .ChildTable}}
func {{.Name | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}
{{- else }}
func {{.Name}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}
{{- end}}

{{if not .SkipFetch}}
func fetch{{.Name | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	{{- if .NewFuncHasSubscriptionId}}
  svc, err := {{.Service}}.{{.NewFuncName}}(cl.SubscriptionId, cl.Creds, cl.Options)
  {{- else}}
  svc, err := {{.Service}}.{{.NewFuncName}}(cl.Creds, cl.Options)
  {{- end}}
	if err != nil {
    return err
  }
	pager := svc.{{.ListFuncName}}(nil)
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