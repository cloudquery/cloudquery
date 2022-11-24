package client

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) interface{} {
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) interface{} {
	return hex.EncodeToString(v.Bytes)
}

func (*Client) TransformFloat8(v *schema.Float8) interface{} {
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) interface{} {
	return v.Int
}

func (*Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformJSON(v *schema.JSON) interface{} {
	if v.Status != schema.Present {
		return map[string]interface{}{}
	}
	var res interface{}
	if err := json.Unmarshal(v.Bytes, &res); err != nil {
		panic(err)
	}
	return res
}

func (*Client) TransformText(v *schema.Text) interface{} {
	return v.String()
}

func (*Client) TransformTextArray(v *schema.TextArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Client) TransformUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformInet(v *schema.Inet) interface{} {
	return v.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}
