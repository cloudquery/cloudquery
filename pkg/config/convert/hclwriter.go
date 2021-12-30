package convert

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func BodyToHCL(ctx *hcl.EvalContext, body hcl.Body) ([]byte, hcl.Diagnostics) {
	f := hclwrite.NewEmptyFile()
	b := f.Body()
	attrs, diags := body.JustAttributes()
	if diags != nil {
		return nil, diags
	}

	for _, v := range attrs {
		val, diags := v.Expr.Value(ctx)
		if diags != nil {
			return nil, diags
		}
		b.SetAttributeValue(v.Name, val)
	}

	newBody, ok := body.(*hclsyntax.Body)
	if !ok {
		return f.Bytes(), nil
	}
	for _, block := range newBody.Blocks {
		newBlock := b.AppendNewBlock(block.Type, block.Labels)
		writeBody(ctx, block.Body, newBlock.Body())
	}

	return f.Bytes(), nil
}

func writeBody(ctx *hcl.EvalContext, body *hclsyntax.Body, hclWriteBlock *hclwrite.Body) hcl.Diagnostics {
	for _, a := range body.Attributes {
		val, diags := a.Expr.Value(ctx)
		if diags != nil {
			return diags
		}
		hclWriteBlock.SetAttributeValue(a.Name, val)
	}
	for _, b := range body.Blocks {
		newBlock := hclWriteBlock.AppendNewBlock(b.Type, b.Labels)
		diags := writeBody(ctx, b.Body, newBlock.Body())
		if diags != nil {
			return diags
		}
	}
	return nil
}
