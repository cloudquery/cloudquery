package config

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type Type string

const (
	Hub    Type = "hub"
	Remote Type = "remote"
	Local  Type = "local"
	Inline Type = "inline"
)

type Policy struct {
	Name    string `hcl:"name,label"`
	Type    Type   `hcl:"type"`
	Source  string `hcl:"source"`
	SubPath string `hcl:"sub_path,optional"`
	Version string `hcl:"version,optional"`
}

// DecodePolicyConfigBlock will get the inner part of the module config with the given name from an hcl.Body.
func decodePolicyConfigBlock(block *hcl.Block, ctx *hcl.EvalContext) (*Policy, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	content, _, moreDiags := block.Body.PartialContent(policyConfigBlockSchema)
	diags = append(diags, moreDiags...)
	name := block.Labels[0]
	policy := &Policy{Name: name}

	if attr, exists := content.Attributes["type"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, ctx, &policy.Type)
		diags = append(diags, valDiags...)
		switch policy.Type {
		case Hub, Remote, Local, Inline:
			break
		default:
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Invalid query type`,
				Detail:   fmt.Sprintf(`Policy type value of "%s" is not valid`, policy.Type),
				Subject:  &block.DefRange,
			})
		}
	}

	if attr, exists := content.Attributes["source"]; exists {
		var valDiags hcl.Diagnostics
		if policy.Type == Inline {
			var ctx = convert.GetEvalContext("/")
			valDiags = gohcl.DecodeExpression(attr.Expr, ctx, &policy.Source)
		} else {
			valDiags = gohcl.DecodeExpression(attr.Expr, ctx, &policy.Source)
		}

		diags = append(diags, valDiags...)
	}

	if attr, exists := content.Attributes["sub_path"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, ctx, &policy.SubPath)
		diags = append(diags, valDiags...)
	}

	if attr, exists := content.Attributes["version"]; exists {
		valDiags := gohcl.DecodeExpression(attr.Expr, ctx, &policy.Version)
		diags = append(diags, valDiags...)
	}

	for _, block := range content.Blocks {
		diags = append(diags, &hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Unexpected block type name in policy block",
			Detail:   fmt.Sprintf("The block type name %q is unexpected in policy block.", block.Type),
			Subject:  &block.TypeRange,
		})

	}
	return policy, diags
}

var policyConfigBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name:     "type",
			Required: true,
		},
		{
			Name:     "source",
			Required: true,
		},
		{
			Name: "sub_path",
		},
		{
			Name: "version",
		},
	},
}
