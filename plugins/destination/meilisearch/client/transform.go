package client

import (
	"encoding/json"
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (*Client) TransformBool(v *schema.Bool) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bytes
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.String()
	}
	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Float
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.String()
	}
	return res
}

func (*Client) TransformInet(v *schema.Inet) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]int64, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Int
	}
	return res
}

func (*Client) TransformInt8(v *schema.Int8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Int
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}
	return json.RawMessage(v.Bytes)
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.String()
	}
	return res
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Str
	}
	return res
}

func (*Client) TransformText(v *schema.Text) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Str
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	if v.Status != schema.Present {
		return nil
	}
	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.String()
	}
	return res
}

func (*Client) TransformUUID(v *schema.UUID) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	res := make(schema.CQTypes, len(values))

	for i, v := range values {
		typ := table.Columns[i].Type
		// uuid, cidr, inet, macaddr & string arrays are written as []string
		// The default implementation will write them as coma-separated strings surrounded by "{}".
		switch typ {
		case schema.TypeUUIDArray,
			schema.TypeCIDRArray,
			schema.TypeInetArray,
			schema.TypeMacAddrArray,
			schema.TypeStringArray:
			vTyped := v.([]any)
			parts := make([]string, len(vTyped))
			for idx, el := range vTyped {
				parts[idx] = el.(string)
			}
			v = "{" + strings.Join(parts, ",") + "}"
		}
		t := schema.NewCqTypeFromValueType(typ)
		// update can move value type to another value, so don't worry about err
		// example: str -> bool, and the destination still works fine
		_ = t.Set(v)
		res[i] = t
	}
	return res, nil
}
