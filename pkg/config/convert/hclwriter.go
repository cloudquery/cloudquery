package convert

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// This is here because of this - https://github.com/hashicorp/hcl/issues/299

func WriteBody(ctx *hcl.EvalContext, body *hclsyntax.Body, hclWriteBlock *hclwrite.Body) hcl.Diagnostics {
	for _, a := range body.Attributes {
		val, diags := a.Expr.Value(ctx)
		if diags != nil {
			return diags
		}
		hclWriteBlock.SetAttributeValue(a.Name, val)
	}
	for _, b := range body.Blocks {
		newBlock := hclWriteBlock.AppendNewBlock(b.Type, b.Labels)
		diags := WriteBody(ctx, b.Body, newBlock.Body())
		if diags != nil {
			return diags
		}
	}
	return nil
}
