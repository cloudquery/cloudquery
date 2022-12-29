package json

import (
	"encoding/json"

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
	return v.Bool
}

func (*Transformer) TransformBytea(v *schema.Bytea) any {
	return v.String()
}

func (*Transformer) TransformFloat8(v *schema.Float8) any {
	return v.Float
}

func (*Transformer) TransformInt8(v *schema.Int8) any {
	return v.Int
}

func (*Transformer) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]int64, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Int
	}
	return res
}

func (*Transformer) TransformJSON(v *schema.JSON) any {
	var res any
	if err := json.Unmarshal(v.Bytes, &res); err != nil {
		panic(err)
	}
	return res
}

func (*Transformer) TransformText(v *schema.Text) any {
	return v.Str
}

func (*Transformer) TransformTextArray(v *schema.TextArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Str
	}
	return res
}

func (*Transformer) TransformTimestamptz(v *schema.Timestamptz) any {
	return v.String()
}

func (*Transformer) TransformUUID(v *schema.UUID) any {
	return v.String()
}

func (*Transformer) TransformUUIDArray(v *schema.UUIDArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Transformer) TransformCIDR(v *schema.CIDR) any {
	return v.String()
}

func (*Transformer) TransformCIDRArray(v *schema.CIDRArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Transformer) TransformInet(v *schema.Inet) any {
	return v.String()
}

func (*Transformer) TransformInetArray(v *schema.InetArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return v.String()
}

func (*Transformer) TransformMacaddr(v *schema.Macaddr) any {
	return v.String()
}

func (*Transformer) TransformMacaddrArray(v *schema.MacaddrArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}
