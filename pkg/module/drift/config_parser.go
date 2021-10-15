package drift

import (
	"fmt"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/go-version"
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

type placeholder string

const (
	placeholderResourceKey             placeholder = "resourceKey"
	placeholderResourceName            placeholder = "resourceName"
	placeholderResourceColumnNames     placeholder = "resourceColumnNames"
	placeholderResourceOptsPrimaryKeys placeholder = "resourceOptionsPrimaryKeys"
)

func makePlaceholder(varName placeholder) cty.Value {
	return cty.StringVal("${" + string(varName) + "}")
}

func replacePlaceholderInSlice(varName placeholder, value, subject []string) []string {
	plc := "${" + string(varName) + "}"
	newSubj := make([]string, 0, len(subject))
	for i := range subject {
		if subject[i] == plc {
			newSubj = append(newSubj, value...)
		} else {
			newSubj = append(newSubj, subject[i])
		}
	}
	return newSubj
}

func (p *Parser) Decode(body hcl.Body, diags hcl.Diagnostics) (*BaseConfig, hcl.Diagnostics) {
	baseConfig := &BaseConfig{}
	content, contentDiags := body.Content(baseSchema)
	diags = append(diags, contentDiags...)

	ctx := *p.HCLContext
	ctx.Variables["resource"] = cty.ObjectVal(map[string]cty.Value{
		// FIXME expose the resource struct here, with late binding?
		"Key": makePlaceholder(placeholderResourceKey),
		"Value": cty.ObjectVal(map[string]cty.Value{
			"Name":        makePlaceholder(placeholderResourceName),
			"ColumnNames": cty.ListVal([]cty.Value{makePlaceholder(placeholderResourceColumnNames)}),
			"Options": cty.ObjectVal(map[string]cty.Value{
				"PrimaryKeys": cty.ListVal([]cty.Value{makePlaceholder(placeholderResourceOptsPrimaryKeys)}),
			}),
		}),
	})

	for _, block := range content.Blocks {
		switch block.Type {
		case "provider":
			prov, provDiags := p.decodeProviderBlock(block, &ctx)
			diags = append(diags, provDiags...)
			if prov != nil {
				if prov.Name == wildcard {
					if baseConfig.WildProvider != nil {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  `Duplicate block`,
							Detail:   `There must be at most one block of "*" type provider`,
							Subject:  &block.DefRange,
						})
						continue
					}
					if prov.Version != "" {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  `Invalid attribute`,
							Detail:   `version attribute is only valid for non-"*" providers`,
							Subject:  &block.DefRange,
						})
						continue
					}

					// TODO handle source?

					baseConfig.WildProvider = prov
					continue
				}

				if prov.Version != "" {
					var err error
					prov.versionConstraints, err = version.NewConstraint(prov.Version)
					if err != nil {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  `Invalid attribute`,
							Detail:   fmt.Sprintf(`version attribute is invalid: %v`, err),
							Subject:  &block.DefRange,
						})
						continue
					}
				}

				if prov.Source != "" {
					diags = append(diags, &hcl.Diagnostic{
						Severity: hcl.DiagError,
						Summary:  `Invalid attribute`,
						Detail:   `source attribute is only valid for "*" providers`,
						Subject:  &block.DefRange,
					})
					continue
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

	diags = p.interpret(baseConfig)

	return baseConfig, diags
}

// interpret iterates over every provider/resource and replaces missing values with the ones in wildprovider/wildresource
func (p *Parser) interpret(cfg *BaseConfig) hcl.Diagnostics {
	for _, prov := range cfg.Providers {
		prov.applyWildProvider(cfg.WildProvider)

		for _, res := range prov.Resources {
			res.applyWildResource(prov.WildResource)
			if cfg.WildProvider != nil {
				res.applyWildResource(cfg.WildProvider.WildResource)
			}
		}
	}
	return nil
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

	resourceSchema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name:     "identifiers",
				Required: false,
			},
			{
				Name:     "ignore_identifiers",
				Required: false,
			},
			{
				Name:     "attributes",
				Required: false,
			},
			{
				Name:     "ignore_attributes",
				Required: false,
			},
			{
				Name:     "deep",
				Required: false,
			},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "iac",
				LabelNames: nil,
			},
		},
	}

	iacSchema = &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "terraform",
				LabelNames: nil,
			},
			{
				Type:       "cloudformation",
				LabelNames: nil,
			},
		},
	}
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
	if versionAttr, ok := content.Attributes["version"]; ok {
		diags = append(diags, gohcl.DecodeExpression(versionAttr.Expr, ctx, &prov.Version)...)
	}
	if skipAttr, ok := content.Attributes["skip_resources"]; ok {
		diags = append(diags, gohcl.DecodeExpression(skipAttr.Expr, ctx, &prov.SkipResources)...)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			res, resDiags := p.decodeResourceBlock(block, ctx.NewChild())
			if resDiags.HasErrors() {
				diags = append(diags, resDiags...)
				continue
			}
			res.defRange = &block.DefRange
			if block.Labels[0] == wildcard {
				prov.WildResource = res
			} else {
				prov.Resources[block.Labels[0]] = res
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

func (p *Parser) decodeResourceBlock(b *hcl.Block, ctx *hcl.EvalContext) (*ResourceConfig, hcl.Diagnostics) {
	content, diags := b.Body.Content(resourceSchema)
	if diags.HasErrors() {
		return nil, diags
	}
	res := &ResourceConfig{
		IAC: make(map[string]*IACConfig),
	}
	if idAttr, ok := content.Attributes["identifiers"]; ok {
		diags = append(diags, gohcl.DecodeExpression(idAttr.Expr, ctx, &res.Identifiers)...)
	}
	if ignoreIdAttr, ok := content.Attributes["ignore_identifiers"]; ok {
		diags = append(diags, gohcl.DecodeExpression(ignoreIdAttr.Expr, ctx, &res.IgnoreIdentifiers)...)
	}
	if attr, ok := content.Attributes["attributes"]; ok {
		diags = append(diags, gohcl.DecodeExpression(attr.Expr, ctx, &res.Attributes)...)
	}
	if ignoreAttr, ok := content.Attributes["ignore_attributes"]; ok {
		diags = append(diags, gohcl.DecodeExpression(ignoreAttr.Expr, ctx, &res.IgnoreAttributes)...)
	}
	if deepAttr, ok := content.Attributes["deep"]; ok {
		diags = append(diags, gohcl.DecodeExpression(deepAttr.Expr, ctx, &res.Deep)...)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "iac":
			iacContent, diags := block.Body.Content(iacSchema)
			if diags.HasErrors() {
				return nil, diags
			}
			for _, iacBlock := range iacContent.Blocks {
				var ia IACConfig
				if diag := gohcl.DecodeBody(iacBlock.Body, ctx.NewChild(), &ia); diag.HasErrors() {
					diags = append(diags, diag...)
					continue
				}
				ia.defRange = &block.DefRange
				ia.attributeMap = make(map[string]string, len(ia.AttributeMap))

				for _, v := range ia.AttributeMap {
					parts := strings.Split(v, "=")
					if len(parts) != 2 {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Summary:  `Invalid attribute_map entry`,
							Detail:   `attribute_map entry should have a "cloud_attribute=iac_attribute" format`,
							Subject:  &block.DefRange,
						})
						continue
					}
					ia.attributeMap[parts[0]] = parts[1]
				}
				res.IAC[iacBlock.Type] = &ia
			}
			if diags.HasErrors() {
				return nil, diags
			}
		default:
			panic("unexpected block")
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return res, diags
}
