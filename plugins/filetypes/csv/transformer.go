package csv

import (
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

type ReverseTransformer struct{
	defaultTransformer destination.DefaultReverseTransformer
}

func (t *ReverseTransformer) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	return t.defaultTransformer.ReverseTransformValues(table, values) 
}

type Transformer struct{}

func (*Transformer) TransformBool(v *schema.Bool) any {
	return v.String()
}

func (*Transformer) TransformBytea(v *schema.Bytea) any {
	return v.String()
}

func (*Transformer) TransformFloat8(v *schema.Float8) any {
	return v.String()
}

func (*Transformer) TransformInt8(v *schema.Int8) any {
	return v.String()
}

func (*Transformer) TransformInt8Array(v *schema.Int8Array) any {
	return v.String()
}

func (*Transformer) TransformJSON(v *schema.JSON) any {
	return v.String()
}

func (*Transformer) TransformText(v *schema.Text) any {
	return v.Str
}

func (*Transformer) TransformTextArray(v *schema.TextArray) any {
	return v.String()
}

func (*Transformer) TransformTimestamptz(v *schema.Timestamptz) any {
	return v.String()
}

func (*Transformer) TransformUUID(v *schema.UUID) any {
	return v.String()
}

func (*Transformer) TransformUUIDArray(v *schema.UUIDArray) any {
	return v.String()
}

func (*Transformer) TransformCIDR(v *schema.CIDR) any {
	return v.String()
}

func (*Transformer) TransformCIDRArray(v *schema.CIDRArray) any {
	return v.String()
}

func (*Transformer) TransformInet(v *schema.Inet) any {
	return v.String()
}

func (*Transformer) TransformInetArray(v *schema.InetArray) any {
	return v.String()
}

func (*Transformer) TransformMacaddr(v *schema.Macaddr) any {
	return v.String()
}

func (*Transformer) TransformMacaddrArray(v *schema.MacaddrArray) any {
	return v.String()
}
