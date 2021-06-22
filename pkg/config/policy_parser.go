package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func (p *Parser) DecodePolicies(body hcl.Body, diags hcl.Diagnostics) (*PolicyWrapper, hcl.Diagnostics) {
	policyWrapper := &PolicyWrapper{}
	contentDiags := gohcl.DecodeBody(body, nil, policyWrapper)
	return policyWrapper, append(diags, contentDiags...)
}
