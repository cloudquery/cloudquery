package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) interface{} {
	return v.String()
}

func (*Client) TransformBytea(v *schema.Bytea) interface{} {
	return v.String()
}

func (*Client) TransformFloat8(v *schema.Float8) interface{} {
	return v.String()
}

func (*Client) TransformInt8(v *schema.Int8) interface{} {
	return v.String()
}

func (*Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	return v.String()
}

func (*Client) TransformJSON(v *schema.JSON) interface{} {
	return v.String()
}

func (*Client) TransformText(v *schema.Text) interface{} {
	return stripNulls(v.String())
}

func (*Client) TransformTextArray(v *schema.TextArray) interface{} {
	return stripNulls(v.String())
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return v.String()
}

func (*Client) TransformUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	return v.String()
}

func (*Client) TransformCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	return v.String()
}

func (*Client) TransformInet(v *schema.Inet) interface{} {
	return v.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) interface{} {
	return v.String()
}

func (*Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	return v.String()
}

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}
