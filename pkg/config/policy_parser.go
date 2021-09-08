package config

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

var policyWrapperSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

func (p *Parser) DecodePolicies(body hcl.Body, diags hcl.Diagnostics, basePath string) (*PolicyWrapper, hcl.Diagnostics) {
	policyWrapper := &PolicyWrapper{}
	content, contentDiags := body.Content(policyWrapperSchema)
	diags = append(diags, contentDiags...)
	for _, block := range content.Blocks {
		switch block.Type {
		case "policy":
			policy, policyDiags := p.decodePolicyBlock(block, convert.GetEvalContext(basePath))
			diags = append(diags, policyDiags...)
			if policy != nil {
				policyWrapper.Policies = append(policyWrapper.Policies, policy)
			}
		default:
			panic("unexpected block")
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return policyWrapper, diags
}

var policySchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
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
			Type:       "query",
			LabelNames: []string{"name"},
		},
		{
			Type:       "view",
			LabelNames: []string{"name"},
		},
	},
}

func (p *Parser) decodePolicyBlock(b *hcl.Block, ctx *hcl.EvalContext) (*Policy, hcl.Diagnostics) {
	content, diags := b.Body.Content(policySchema)
	if diags.HasErrors() {
		return nil, diags
	}
	policy := &Policy{Name: b.Labels[0]}
	if descriptionAttr, ok := content.Attributes["description"]; ok {
		diags = append(diags, gohcl.DecodeExpression(descriptionAttr.Expr, nil, &policy.Description)...)
	}
	for _, block := range content.Blocks {
		switch block.Type {
		case "configuration":
			var config Configuration
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &config)...)
			if policy.Config != nil {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  `Duplicate block`,
					Detail:   `There must be at most one block of "configuration" type`,
					Subject:  &block.DefRange,
				})
			}
			policy.Config = &config
		case "policy":
			inner, innerDiags := p.decodePolicyBlock(block, ctx)
			diags = append(diags, innerDiags...)
			policy.Policies = append(policy.Policies, inner)
		case "query":
			var query Query
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
					Detail:   fmt.Sprintf(`Query type value of "%s" is not valid`, query.Type),
					Subject:  &block.DefRange,
				})
			}
			policy.Queries = append(policy.Queries, &query)
		case "view":
			var view View
			view.Name = block.Labels[0]
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &view)...)
			policy.Views = append(policy.Views, &view)
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return policy, diags
}
