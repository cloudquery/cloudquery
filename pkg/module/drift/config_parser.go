package drift

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type Parser struct {
	p          *hclparse.Parser
	HCLContext *hcl.EvalContext
}

func NewParser(basePath string) *Parser {
	ctx := convert.GetEvalContext(basePath)
	ctx.Variables = make(map[string]cty.Value)
	ctx.Functions["sql"] = function.New(&function.Spec{
		Params: []function.Parameter{
			{
				Name: "sql-expr",
				Type: cty.String,
			},
		},
		Type: function.StaticReturnType(cty.String),
		Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
			if len(args) != 1 {
				return cty.UnknownVal(cty.String), fmt.Errorf("invalid arguments: single expression required")
			}

			return cty.StringVal("${sql:" + args[0].AsString() + "}"), nil
		},
	})

	config.EnvToHCLContext(ctx, config.EnvVarPrefix, os.Environ())

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
		{
			Type: "terraform",
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

func (p *Parser) Decode(body hcl.Body, allowedProvider string, diags hcl.Diagnostics) (*BaseConfig, hcl.Diagnostics) {
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
			if prov == nil {
				continue
			}
			if allowedProvider != "" && prov.Name != allowedProvider {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  `Invalid label`,
					Detail:   `Provider label should be ` + allowedProvider,
					Subject:  &block.DefRange,
				})
				continue
			}

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
				if len(prov.AccountIDs) > 0 {
					diags = append(diags, &hcl.Diagnostic{
						Severity: hcl.DiagError,
						Summary:  `Invalid attribute`,
						Detail:   `account_ids attribute is only valid for non-"*" providers`,
						Subject:  &block.DefRange,
					})
					continue
				}

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

			baseConfig.Providers = append(baseConfig.Providers, prov)
		case "terraform":
			ts, tsDiags := p.decodeTerraformBlock(block, &ctx)
			diags = append(diags, tsDiags...)
			if ts != nil {
				baseConfig.Terraform = ts
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
				Name:     "version", // only valid for non-"*" providers
				Required: false,
			},
			{
				Name:     "ignore_resources",
				Required: false,
			},
			{
				Name:     "check_resources",
				Required: false,
			},
			{
				Name:     "account_ids", // only valid for non-"*" providers
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
			{
				Name:     "filters",
				Required: false,
			},
			{
				Name:     "sets",
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
				Type:       string(iacTerraform),
				LabelNames: nil,
			},
			{
				Type:       string(iacCloudformation),
				LabelNames: nil,
			},
		},
	}

	terraformSourceSchema = &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{
				Name:     "backend",
				Required: false,
			},
			{
				Name:     "files",
				Required: false,
			},
			{
				Name:     "bucket",
				Required: false,
			},
			{
				Name:     "keys",
				Required: false,
			},
			{
				Name:     "region",
				Required: false,
			},
			{
				Name:     "role_arn",
				Required: false,
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
	if attr, ok := content.Attributes["ignore_resources"]; ok {
		var (
			list []string
			err  error
		)
		diags = append(diags, gohcl.DecodeExpression(attr.Expr, ctx, &list)...)
		prov.IgnoreResources, err = parseResourceSelectors(list)
		if err != nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Invalid ignore_resources entry`,
				Detail:   err.Error(),
				Subject:  &attr.Range,
			})
		}
	}
	if attr, ok := content.Attributes["check_resources"]; ok {
		var (
			list []string
			err  error
		)
		diags = append(diags, gohcl.DecodeExpression(attr.Expr, ctx, &list)...)
		prov.CheckResources, err = parseResourceSelectors(list)
		if err != nil {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Invalid check_resources entry`,
				Detail:   err.Error(),
				Subject:  &attr.Range,
			})
		}
	}
	if accountsAttr, ok := content.Attributes["account_ids"]; ok {
		diags = append(diags, gohcl.DecodeExpression(accountsAttr.Expr, ctx, &prov.AccountIDs)...)
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
		IAC: make(map[iacProvider]*IACConfig),
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
	if filtersAttr, ok := content.Attributes["filters"]; ok {
		diags = append(diags, gohcl.DecodeExpression(filtersAttr.Expr, ctx, &res.Filters)...)
	}
	if setsAttr, ok := content.Attributes["sets"]; ok {
		diags = append(diags, gohcl.DecodeExpression(setsAttr.Expr, ctx, &res.Sets)...)
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
					parts := strings.SplitN(v, "=", 2)
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
				res.IAC[iacProvider(iacBlock.Type)] = &ia
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

func (p *Parser) decodeTerraformBlock(b *hcl.Block, ctx *hcl.EvalContext) (*TerraformSourceConfig, hcl.Diagnostics) {
	content, diags := b.Body.Content(terraformSourceSchema)
	if diags.HasErrors() {
		return nil, diags
	}
	ts := &TerraformSourceConfig{}
	if backendAttr, ok := content.Attributes["backend"]; ok {
		diags = append(diags, gohcl.DecodeExpression(backendAttr.Expr, ctx, &ts.Backend)...)
	}
	if filesAttr, ok := content.Attributes["files"]; ok {
		diags = append(diags, gohcl.DecodeExpression(filesAttr.Expr, ctx, &ts.Files)...)
	}
	if bucketAttr, ok := content.Attributes["bucket"]; ok {
		diags = append(diags, gohcl.DecodeExpression(bucketAttr.Expr, ctx, &ts.Bucket)...)
	}
	if keysAttr, ok := content.Attributes["keys"]; ok {
		diags = append(diags, gohcl.DecodeExpression(keysAttr.Expr, ctx, &ts.Keys)...)
	}
	if regionAttr, ok := content.Attributes["region"]; ok {
		diags = append(diags, gohcl.DecodeExpression(regionAttr.Expr, ctx, &ts.Region)...)
	}
	if roleArnAttr, ok := content.Attributes["role_arn"]; ok {
		diags = append(diags, gohcl.DecodeExpression(roleArnAttr.Expr, ctx, &ts.RoleARN)...)
	}
	if diags.HasErrors() {
		return nil, diags
	}

	if err := ts.Validate(); err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  `Invalid terraform config`,
				Detail:   err.Error(),
			},
		}
	}

	return ts, nil
}
