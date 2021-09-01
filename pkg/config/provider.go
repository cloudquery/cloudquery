package config

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type Provider struct {
	Name               string   `hcl:"name,label"`
	Alias              string   `hcl:"alias,optional"`
	EnablePartialFetch bool     `hcl:"enable_partial_fetch,optional"`
	Resources          []string `hcl:"resources,optional"`
	Env                []string `hcl:"env,optional"`
	Configuration      hcl.Body `hcl:"configuration,body"`
}

func decodeProviderBlock(block *hcl.Block, existingProviders map[string]bool) (*Provider, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	content, _, moreDiags := block.Body.PartialContent(providerBlockSchema)
	diags = append(diags, moreDiags...)
	name := block.Labels[0]
	provider := &Provider{Name: name, Alias: name}
	if attr, exists := content.Attributes["alias"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, nil, &provider.Alias)
		diags = append(diags, valDiags...)
		if _, ok := existingProviders[provider.Alias]; ok {
			errMsg := fmt.Sprintf("Provider with alias %s for provider %s already exists, give it a different alias.", provider.Alias, name)
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Duplicate Alias",
				Detail:   errMsg,
				Subject:  attr.Range.Ptr(),
			})
		}
		existingProviders[provider.Alias] = true
	} else {
		if _, ok := existingProviders[name]; ok {
			errMsg := fmt.Sprintf("Provider with name %s already exists, use alias in provider configuration block.", name)
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Provider Alias Required",
				Detail:   errMsg,
				Subject:  block.DefRange.Ptr(),
			})
		}
		existingProviders[name] = true
	}

	if attr, exists := content.Attributes["enable_partial_fetch"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, nil, &provider.EnablePartialFetch)
		diags = append(diags, valDiags...)
	}
	if attr, exists := content.Attributes["resources"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, nil, &provider.Resources)
		diags = append(diags, valDiags...)
	}
	if attr, exists := content.Attributes["env"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, nil, &provider.Env)
		diags = append(diags, valDiags...)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "configuration":
			provider.Configuration = block.Body
		default:
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Unexpected block type name in provider block",
				Detail:   fmt.Sprintf("The block type name %q is unexpected in provider block.", block.Type),
				Subject:  &block.TypeRange,
			})
		}
	}
	return provider, diags
}

var providerBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "resources",
		},
		{
			Name: "alias",
		},
		{
			Name: "enable_partial_fetch",
		},
		{
			Name: "env",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		// _All_ of these are reserved for future expansion.
		{Type: "configuration"},
	},
}
