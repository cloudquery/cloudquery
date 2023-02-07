package client

import (
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) any {
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	return v.Bytes
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) any {
	return v.Int
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}
	var j map[string]any
	_ = json.Unmarshal(v.Bytes, &j)
	return j
}

func (*Client) TransformText(v *schema.Text) any {
	return v.String()
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Client) TransformUUID(v *schema.UUID) any {
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	return v.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformInet(v *schema.Inet) any {
	return v.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	return v.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}
