package drift

import (
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

type Parser struct {
	p          *hclparse.Parser
	HCLContext *hcl.EvalContext
}

func NewParser(basePath string) *Parser {
	ctx := convert.GetEvalContext(basePath)
	ctx.Variables = make(map[string]cty.Value)

	return &Parser{
		p:          hclparse.NewParser(),
		HCLContext: ctx,
	}
}

var baseSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
	},
}

func (p *Parser) Decode(body hcl.Body, diags hcl.Diagnostics) (*BaseConfig, hcl.Diagnostics) {
	baseConfig := &BaseConfig{}
	content, contentDiags := body.Content(baseSchema)
	diags = append(diags, contentDiags...)

	ctx := *p.HCLContext
	ctx.Variables["resource"] = cty.ObjectVal(map[string]cty.Value{
		// TODO expose the resource struct here, with late binding
		"Value": cty.ObjectVal(map[string]cty.Value{
			"Options": cty.ObjectVal(map[string]cty.Value{
				"PrimaryKeys": cty.ListValEmpty(cty.String),
			}),
		}),
	})

	hasWildcard := false
	for _, block := range content.Blocks {
		switch block.Type {
		case "provider":
			prov, provDiags := p.decodeProviderBlock(block, &ctx)
			diags = append(diags, provDiags...)
			if prov != nil {
				if prov.Name == wildcard {
					if hasWildcard {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  `Duplicate block`,
							Detail:   `There must be at most one block of "*" type provider`,
							Subject:  &block.DefRange,
						})
						continue
					}
					hasWildcard = true
				}
				baseConfig.Providers = append(baseConfig.Providers, prov)
			}
		default:
			panic("unexpected block")
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return baseConfig, diags
}

var (
	providerSchema = &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "resource",
				LabelNames: []string{"name"},
			},
		},
		Attributes: []hcl.AttributeSchema{
			{
				Name:     "source", // only valid for the "*" provider
				Required: false,
			},
			{
				Name:     "version", // only valid for non-"*" providers
				Required: false,
			},
			{
				Name:     "skip_resources",
				Required: false,
			},
		},
	}

/*
	resourceSchema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name:     "identifiers",
				Required: false,
			},
			{
				Name:     "ignore_attributes",
				Required: false,
			},
			{
				Name:     "tf_type",
				Required: false,
			},
			{
				Name:     "tf_name",
				Required: false,
			},
		},
	}
*/
)

func (p *Parser) decodeProviderBlock(b *hcl.Block, ctx *hcl.EvalContext) (*ProviderConfig, hcl.Diagnostics) {
	content, diags := b.Body.Content(providerSchema)
	if diags.HasErrors() {
		return nil, diags
	}
	prov := &ProviderConfig{
		Name:      b.Labels[0],
		Resources: make(map[string]*ResourceConfig),
	}
	// TODO
	//if sourceAttr, ok := content.Attributes["source"]; ok {
	//	diags = append(diags, gohcl.DecodeExpression(sourceAttr.Expr, nil, &prov.Description)...)
	//}
	if versionAttr, ok := content.Attributes["version"]; ok {
		diags = append(diags, gohcl.DecodeExpression(versionAttr.Expr, nil, &prov.Version)...)
	}
	if skipAttr, ok := content.Attributes["skip_resources"]; ok {
		diags = append(diags, gohcl.DecodeExpression(skipAttr.Expr, nil, &prov.SkipResources)...)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			var res ResourceConfig
			if diag := gohcl.DecodeBody(block.Body, ctx, &res); diag.HasErrors() {
				diags = append(diags, diag...)
			} else {
				prov.Resources[block.Labels[0]] = &res
			}
		default:
			panic("unexpected block")
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return prov, diags
}
