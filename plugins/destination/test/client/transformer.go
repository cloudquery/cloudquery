package client

import "github.com/cloudquery/plugin-sdk/schema"

func (*Client) TransformBool(v *schema.Bool) any {
	return v.String()
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	return v.String()
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	return v.String()
}

func (*Client) TransformInt8(v *schema.Int8) any {
	return v.String()
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	return v.String()
}

func (*Client) TransformJSON(v *schema.JSON) any {
	return v.String()
}

func (*Client) TransformText(v *schema.Text) any {
	return v.String()
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	return v.String()
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	return v.String()
}

func (*Client) TransformUUID(v *schema.UUID) any {
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	return v.String()
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	return v.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	return v.String()
}

func (*Client) TransformInet(v *schema.Inet) any {
	return v.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	return v.String()
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	return v.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	return v.String()
}
