package config

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

func (p *Parser) loadPoliciesFromSource(name string, data []byte, source SourceType) (*PolicyWrapper, hcl.Diagnostics) {
	body, diags := p.loadFromSource(name, data, source)
	if body == nil {
		return nil, diags
	}
	return p.decodePolicies(body, diags)
}

func (p *Parser) LoadPoliciesFromHCL(name string, data []byte) (*PolicyWrapper, hcl.Diagnostics) {
	return p.loadPoliciesFromSource(name, data, SourceHCL)
}

func (p *Parser) LoadPoliciesFromJSON(name string, data []byte) (*PolicyWrapper, hcl.Diagnostics) {
	return p.loadPoliciesFromSource(name, data, SourceJSON)
}

func (p *Parser) decodePolicies(body hcl.Body, diags hcl.Diagnostics) (*PolicyWrapper, hcl.Diagnostics) {
	policyWrapper := &PolicyWrapper{}
	contentDiags := gohcl.DecodeBody(body, nil, policyWrapper)
	return policyWrapper, append(diags, contentDiags...)
}
