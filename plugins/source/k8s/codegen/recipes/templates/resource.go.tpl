// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	{{if .GlobalResource}}
		cl := meta.(*client.Client).Client().{{.ServiceFuncName}}().{{.ResourceFuncName}}()
	{{else}}
		cl := meta.(*client.Client).Client().{{.ServiceFuncName}}().{{.ResourceFuncName}}("")
	{{end}}
	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}