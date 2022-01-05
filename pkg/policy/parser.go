package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var policyWrapperSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

var policySchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "source",
		},
		{
			Name: "doc",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "configuration",
		},
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
		{
			Type:       "check",
			LabelNames: []string{"name"},
		},
		{
			Type:       "view",
			LabelNames: []string{"name"},
		},
	},
}

func DecodePolicy(body hcl.Body, diags hcl.Diagnostics, basePath string) (*Policy, hcl.Diagnostics) {
	content, contentDiags := body.Content(policyWrapperSchema)
	diags = append(diags, contentDiags...)
	if len(content.Blocks) > 1 {
		return nil, hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  `Only single root policy block allowed`,
			Detail:   `Only a single policy block is allowed in root level policy`,
			Subject:  &content.MissingItemRange,
		}}
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "policy":
			return DecodePolicyBlock(block, convert.GetEvalContext(basePath))
		default:
			panic("unexpected block")
		}
	}
	return nil, diags
}

func DecodePolicyBlock(b *hcl.Block, ctx *hcl.EvalContext) (*Policy, hcl.Diagnostics) {
	content, diags := b.Body.Content(policySchema)
	if diags.HasErrors() {
		return nil, diags
	}
	return decodePolicyContent(b.Labels, content, ctx, b.TypeRange.Ptr())
}

func decodePolicyContent(labels []string, content *hcl.BodyContent, ctx *hcl.EvalContext, r *hcl.Range) (*Policy, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	p := &Policy{Name: labels[0]}
	if descriptionAttr, ok := content.Attributes["description"]; ok {
		diags = append(diags, gohcl.DecodeExpression(descriptionAttr.Expr, ctx, &p.Description)...)
	}
	if descriptionAttr, ok := content.Attributes["doc"]; ok {
		diags = append(diags, gohcl.DecodeExpression(descriptionAttr.Expr, ctx, &p.Doc)...)
	}

	if sourceAttr, ok := content.Attributes["source"]; ok {
		// Sanity check
		if len(content.Blocks) > 0 {
			return nil, append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Found source with blocks`,
				Detail:   `There must be one of the following: Policy source attribute or blocks`,
				Subject:  r,
			})
		}
		if _, ok := sourceAttr.Expr.(*hclsyntax.FunctionCallExpr); !ok {
			// set source to policy config, it will be loaded later
			if err := gohcl.DecodeExpression(sourceAttr.Expr, ctx, &p.Source); err != nil {
				return nil, err
			}
			return p, nil
		}

		var data string
		if err := gohcl.DecodeExpression(sourceAttr.Expr, ctx, &data); err != nil {
			return nil, err
		}

		f, dd := hclsyntax.ParseConfig([]byte(data), p.Name, hcl.Pos{Byte: 0, Line: 1, Column: 1})
		if dd.HasErrors() {
			return nil, dd
		}
		innerContent, contentDiags := f.Body.Content(policySchema)
		diags = append(diags, contentDiags...)
		if contentDiags.HasErrors() {
			return nil, diags
		}
		iPolicy, decodePolicyDiags := decodePolicyContent([]string{""}, innerContent, ctx, r)
		diags = append(diags, decodePolicyDiags...)
		if decodePolicyDiags.HasErrors() {
			return nil, diags
		}
		if len(iPolicy.Policies) > 1 {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  `Invalid policy source file`,
				Detail:   `Policy source file block should only have a single policy block`,
				Subject:  &sourceAttr.Range,
			})
			return nil, diags
		}
		innerPolicy := iPolicy.Policies[0]
		innerPolicy.Name = p.Name
		return innerPolicy, nil
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "configuration":
			var cfg Configuration
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &cfg)...)
			if p.Config != nil {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  `Duplicate block`,
					Detail:   `There must be at most one block of "configuration" type`,
					Subject:  &block.DefRange,
				})
			}
			p.Config = &cfg
		case "policy":
			inner, innerDiags := DecodePolicyBlock(block, ctx)
			diags = append(diags, innerDiags...)
			p.Policies = append(p.Policies, inner)
		case "check":
			var query Check
			query.Name = block.Labels[0]
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &query)...)
			if query.Type == "" {
				query.Type = AutomaticQuery
			}
			switch query.Type {
			case AutomaticQuery, ManualQuery:
			default:
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  `Invalid query type`,
					Detail:   fmt.Sprintf(`Check type value of "%s" is not valid`, query.Type),
					Subject:  &block.DefRange,
				})
			}
			p.Checks = append(p.Checks, &query)
		case "view":
			var view View
			view.Name = block.Labels[0]
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &view)...)
			p.Views = append(p.Views, &view)
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return p, diags
}
