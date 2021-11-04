package config

import (
	"github.com/hashicorp/hcl/v2"
)

var moduleWrapperSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

// DecodeModuleConfig will get the inner part of the module config with the given name from an hcl.Body.
func (p *Parser) DecodeModuleConfig(body hcl.Body, moduleName string) (hcl.Body, hcl.Diagnostics) {
	content, _, diags := body.PartialContent(moduleWrapperSchema)
	if diags.HasErrors() {
		return nil, diags
	}

	for i := range content.Blocks {
		if content.Blocks[i].Labels[0] == moduleName {
			return content.Blocks[i].Body, nil
		}
	}

	return nil, nil
}
