package {{.Service}}

import (
	"context"
{{- if .StateParamName}}
	"fmt"
	"strconv"
{{end}}

{{if or .HasIDPK .StateParamName}}
	"github.com/apache/arrow/go/v13/arrow"
{{end -}}
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/stripe/stripe-go/v74"
)

func {{.TableName | ToPascal}}() *schema.Table {
    return &schema.Table{
  		Name:        "{{.Plugin}}_{{.TableName}}",
		{{- if .Description}}
      Description: `{{.Description}}`,
    {{- end}}
      Transform:   client.TransformWithStruct(&stripe.{{.StructName}}{},
{{- if .SkipFields}}transformers.WithSkipFields({{.SkipFields | QuoteJoin}}),{{end -}}
{{- if .IgnoreInTests}}transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer({{.IgnoreInTests | QuoteJoin}})),{{end -}}
				),
      Resolver:    fetch{{.TableName | ToPascal}},
{{if .HasIDPK}}
		  Columns: []schema.Column{
				 {
								 Name:       "id",
								 Type:       arrow.BinaryTypes.String,
								 Resolver:   schema.PathResolver("ID"),
								 PrimaryKey: true,
				 },
{{if .StateParamName -}}
				 {
								 Name:           "{{.StateParamName | ToSnake}}",
								 Type:           arrow.FixedWidthTypes.Timestamp_us,
								 Resolver:       schema.PathResolver("{{.StateParamName}}"),
								 IncrementalKey: true,
				 },
{{end -}}
			},
{{if .StateParamName -}}
			IsIncremental: true,
{{end -}}
{{end}}
{{if or .Children .ExtraChildren}}
	  			Relations: []*schema.Table{
				{{- range .Children}}
				{{.TableName | ToPascal}}(),
				{{- end}}
				{{- range .ExtraChildren}}
				{{.}},
				{{- end}}
			},
{{end}}
    }
}

func fetch{{.TableName | ToPascal}}(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)
{{if and (.Parent) (.ListParams)}}
		p := parent.Item.(*stripe.{{.Parent.StructName}})
{{end}}

		lp := &stripe.{{.TableName | ToPascal | Singularize}}ListParams{
{{.ListParams}}
		}

		{{- range .ExpandFields}}
		lp.AddExpand("{{.}}")
		{{- end}}

{{if .StateParamName}}
		const key = "{{.TableName}}"

		if (cl.Backend != nil) {
			value, err := cl.Backend.GetKey(ctx, key)
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}
			if value != "" {
				vi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
				}
				lp.{{.StateParamName}} = &vi
			}
		}
{{end}}

		it := cl.Services.{{.TableName | ToPascal}}.List(lp)
		for it.Next() {
{{if .StateParamName -}}
			data := it.{{.TableName | ToPascal | Singularize}}()
			lp.{{.StateParamName}} = client.MaxInt64(lp.{{.StateParamName}}, &data.{{.StateParamName}})
			res <- data
{{else -}}
			res <- it.{{.TableName | ToPascal | Singularize}}()
{{end -}}
		}

{{if .StateParamName -}}
		err := it.Err()
		if cl.Backend != nil && err == nil && lp.{{.StateParamName}} != nil {
			return cl.Backend.SetKey(ctx, key, strconv.FormatInt(*lp.{{.StateParamName}}, 10))
		}
		return err
{{else -}}
		return it.Err()
{{end -}}
}
