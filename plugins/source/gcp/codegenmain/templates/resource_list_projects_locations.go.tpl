// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.Service | ToCamel}}{{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

func fetch{{.Service | ToCamel}}{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.{{.Service | ToCamel}}.Projects.Locations.{{.SubService | ToCamel}}.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}
    res <- output.{{.SubService | ToCamel}}

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}