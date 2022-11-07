package client

import "github.com/cloudquery/plugin-sdk/schema"

func (c *Client) TransformBool(v *schema.Bool) interface{} {
	return v.String()
}

func (c *Client) TransformBytea(v *schema.Bytea) interface{} {
	return v.String()
}

func (c *Client) TransformFloat8(v *schema.Float8) interface{} {
	return v.String()
}

func (c *Client) TransformInt8(v *schema.Int8) interface{} {
	return v.String()
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	return v.String()
}

func (c *Client) TransformJSON(v *schema.JSON) interface{} {
	return v.String()
}

func (c *Client) TransformText(v *schema.Text) interface{} {
	return v.String()
}

func (c *Client) TransformTextArray(v *schema.TextArray) interface{} {
	return v.String()
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return v.String()
}

func (c *Client) TransformUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	return v.String()
}

func (c *Client) TransformCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	return v.String()
}

func (c *Client) TransformInet(v *schema.Inet) interface{} {
	return v.String()
}

func (c *Client) TransformInetArray(v *schema.InetArray) interface{} {
	return v.String()
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	return v.String()
}