package config

import (
	"github.com/cloudquery/cloudquery/pkg/config/convert"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func (p *Parser) DecodePolicies(body hcl.Body, diags hcl.Diagnostics, basePath string) (*PolicyWrapper, hcl.Diagnostics) {
	policyWrapper := &PolicyWrapper{}
	contentDiags := gohcl.DecodeBody(body, convert.GetEvalContext(basePath), policyWrapper)
	return policyWrapper, append(diags, contentDiags...)
}
