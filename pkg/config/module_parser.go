package config

import (
	"fmt"

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

// DecodeModuleProfile will get the inner part of the module configs with the given name from an hcl.Body, where block identifier is the module name.
func DecodeModuleProfile(body hcl.Body, moduleName string) (map[string]hcl.Body, hcl.Diagnostics) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       moduleName,
				LabelNames: []string{"name"},
			},
		},
	})
	if diags.HasErrors() {
		return nil, diags
	}

	ret := make(map[string]hcl.Body, len(content.Blocks))
	for i := range content.Blocks {
		if _, ok := ret[content.Blocks[i].Labels[0]]; ok {
			return nil, hcl.Diagnostics{
				{
					Severity: hcl.DiagError,
					Summary:  "Duplicate profile name",
					Detail:   fmt.Sprintf("Profile name %q already defined", content.Blocks[i].Labels[0]),
					Subject:  content.Blocks[i].DefRange.Ptr(),
				},
			}
		}

		ret[content.Blocks[i].Labels[0]] = content.Blocks[i].Body
	}
	return ret, nil
}
