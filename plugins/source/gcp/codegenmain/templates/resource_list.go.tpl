// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	{{- if or .ListFunction .GetFunction}}
	"context"
	"github.com/pkg/errors"
	{{- end}}

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

{{if .ListFunction}}
func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := {{.ListFunction}}
		if err != nil {
			return errors.WithStack(err)
		}
    res <- output.{{.OutputField}}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
{{end}}

{{if .GetFunction}}
func {{.Table.PreResourceResolver}}(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	item, err := {{.GetFunction}}
	if err != nil {
		return errors.WithStack(err)
	}
	r.SetItem(item)
	return nil
}
{{end}}