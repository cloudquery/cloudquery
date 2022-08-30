// Code generated by codegen; DO NOT EDIT.

package {{.AWSService | ToLower}}

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"{{.TypesImport}}"
{{range .Imports}}	{{.}}
{{end}}
)

func {{.TableFuncName}}() *schema.Table {
	return &schema.Table{{template "table.go.tpl" .Table}}
}

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}

{{template "resolve_parent_defs.go.tpl" .}}
	input := {{.AWSService | ToLower}}.{{.ListMethod}}Input{
{{range .CustomInputs}}{{.}}
{{end}}{{template "resolve_parent_vars.go.tpl" .}}
	}
	paginator := {{.AWSService | ToLower}}.New{{.ListMethod}}Paginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			{{.CustomErrorBlock}}
			return diag.WrapError(err)
		}
		for _, item := range output.{{.PaginatorListName}} {
			do, err := svc.{{.GetMethod}}(ctx, &{{.AWSService | ToLower}}.{{.GetMethod}}Input{
{{range .CustomInputs}}{{.}}
{{end}}{{if not .SkipDescribeParentInputs}}{{template "resolve_parent_vars.go.tpl" .}}{{end}}
{{range $v := .GetAndListOrder}}
	{{$v}}: {{index $.MatchedGetAndListFields $v}},
{{end}}
			})
			if err != nil {
				{{.CustomErrorBlock}}
				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- do.{{.ItemName}}
		}
	}
	return nil
}

{{if .HasTags}}
func resolve{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}Tags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(*types.{{.AWSStructName}})
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}
	out, err := svc.ListTagsFor{{.ItemName}}(ctx, &{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
{{range $v := .GetAndListOrder}}
	{{$v}}: {{index $.MatchedGetAndListFields $v}},
{{end}}
  })
	if err != nil {
		{{.CustomErrorBlock}}
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
{{end}}
