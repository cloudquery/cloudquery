// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	{{if not .SkipFetch}}
	"context"
	"google.golang.org/api/iterator"
	{{if .ProtobufImport}}
  pb "{{.ProtobufImport}}"
  {{end}}
	{{end}}
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	{{if not .SkipFetch}}
	{{range .MockImports}}
  "{{.}}"
  {{end}}
	{{end}}
  {{range .Imports}}
  "{{.}}"
  {{end}}
)

func {{.SubService | ToCamel}}() *schema.Table {
    return &schema.Table{{template "table.go.tpl" .Table}}
}

{{if not .SkipFetch}}
func fetch{{.SubService | ToCamel}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.{{.RequestStructName}}{
		{{if .RequestStructFields}}{{.RequestStructFields}}{{end}}
	}
	gcpClient, err := {{.Service}}.{{.NewFunctionName}}(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
  it := gcpClient.{{.ListFunctionName}}(ctx, req)
	for {
    resp, err := it.Next()
    if err == iterator.Done {
            break
    }
    if err != nil {
      return err
    }
		{{if .OutputField}}
			res <- resp.{{.OutputField}}
		{{else}}
			res <- resp
		{{end}}
	}
	return nil
}
{{end}}