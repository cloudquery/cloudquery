package config

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type Provider struct {
	Name          string   `hcl:"name,label"`
	Alias         string   `hcl:"alias,optional"`
	Resources     []string `hcl:"resources,optional"`
	Env           []string `hcl:"env,optional"`
	Configuration hcl.Body `hcl:"configuration,body"`
}

func decodeProviderBlock(block *hcl.Block) (*Provider, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	content, _, moreDiags := block.Body.PartialContent(providerBlockSchema)
	diags = append(diags, moreDiags...)

	// Provider names must be localized. Produce an error with a message
	// indicating the action the user can take to fix this message if the local
	// name is not localized.
	name := block.Labels[0]
	provider := &Provider{Name: name}
	if attr, exists := content.Attributes["resources"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, nil, &provider.Resources)
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
	},
	Blocks: []hcl.BlockHeaderSchema{
		// _All_ of these are reserved for future expansion.
		{Type: "configuration"},
	},
}
