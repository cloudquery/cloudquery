package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"

	"github.com/cloudquery/cloudquery/pkg/config/convert"
)

var policyWrapperSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "policy",
			LabelNames: []string{"name"},
		},
	},
}

func DecodePolicies(body hcl.Body, diags hcl.Diagnostics, basePath string) (Policies, hcl.Diagnostics) {
	policies := Policies{}
	content, contentDiags := body.Content(policyWrapperSchema)
	diags = append(diags, contentDiags...)

	for _, block := range content.Blocks {
		switch block.Type {
		case "policy":
			policy, policyDiags := decodePolicyBlock(block, convert.GetEvalContext(basePath))
			diags = append(diags, policyDiags...)
			if policy != nil {
				policies = append(policies, policy)
			}
		default:
			panic("unexpected block")
		}
	}
	if diags.HasErrors() {
		return nil, diags
	}
	return policies, diags
}

var policySchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "source",
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

func decodePolicyContent(labels []string, content *hcl.BodyContent, ctx *hcl.EvalContext, r *hcl.Range) (*Policy, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	policy := &Policy{Name: labels[0]}
	if descriptionAttr, ok := content.Attributes["description"]; ok {
		diags = append(diags, gohcl.DecodeExpression(descriptionAttr.Expr, ctx, &policy.Description)...)
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

		var data string
		diags = append(diags, gohcl.DecodeExpression(sourceAttr.Expr, ctx, &data)...)
		body, dd := config.NewParser().LoadFromSource("", []byte(data), config.SourceHCL)
		if dd.HasErrors() {
			return nil, dd
		}
		innerContent, contentDiags := body.Content(policySchema)
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
		policy.Views = append(policy.Views, iPolicy.Policies[0].Views...)
		policy.Queries = append(policy.Queries, iPolicy.Policies[0].Queries...)
		policy.Policies = append(policy.Policies, iPolicy.Policies[0].Policies...)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "configuration":
			var cfg Configuration
			diags = append(diags, gohcl.DecodeBody(block.Body, ctx, &cfg)...)
			if policy.Config != nil {
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  `Duplicate block`,
					Detail:   `There must be at most one block of "configuration" type`,
					Subject:  &block.DefRange,
				})
			}
			policy.Config = &cfg
		case "policy":
			inner, innerDiags := decodePolicyBlock(block, ctx)
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

func decodePolicyBlock(b *hcl.Block, ctx *hcl.EvalContext) (*Policy, hcl.Diagnostics) {
	content, diags := b.Body.Content(policySchema)
	if diags.HasErrors() {
		return nil, diags
	}
	return decodePolicyContent(b.Labels, content, ctx, b.TypeRange.Ptr())
}
