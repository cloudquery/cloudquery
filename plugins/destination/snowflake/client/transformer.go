package client

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) any {
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	return hex.EncodeToString(v.Bytes)
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) any {
	return v.Int
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]int64, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Int
	}
	return res
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return map[string]any{}
	}
	var res any
	if err := json.Unmarshal(v.Bytes, &res); err != nil {
		panic(err)
	}
	return res
}

func (*Client) TransformText(v *schema.Text) any {
	return stripNulls(v.String())
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = stripNulls(e.String())
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

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}
