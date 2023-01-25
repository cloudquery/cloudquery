package client

import "github.com/cloudquery/plugin-sdk/schema"

func (c *Client) TransformBool(v *schema.Bool) any {
	return c.jsonTransformer.TransformBool(v)
}

func (c *Client) TransformBytea(v *schema.Bytea) any {
	return c.jsonTransformer.TransformBytea(v)
}

func (c *Client) TransformFloat8(v *schema.Float8) any {
	return c.jsonTransformer.TransformFloat8(v)
}

func (c *Client) TransformInt8(v *schema.Int8) any {
	return c.jsonTransformer.TransformInt8(v)
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) any {
	return c.jsonTransformer.TransformInt8Array(v)
}

func (c *Client) TransformJSON(v *schema.JSON) any {
	return c.jsonTransformer.TransformJSON(v)
}

func (c *Client) TransformText(v *schema.Text) any {
	return c.jsonTransformer.TransformText(v)
}

func (c *Client) TransformTextArray(v *schema.TextArray) any {
	return c.jsonTransformer.TransformTextArray(v)
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) any {
	return c.jsonTransformer.TransformTimestamptz(v)
}

func (c *Client) TransformUUID(v *schema.UUID) any {
	return c.jsonTransformer.TransformUUID(v)
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) any {
	return c.jsonTransformer.TransformUUIDArray(v)
}

func (c *Client) TransformCIDR(v *schema.CIDR) any {
	return c.jsonTransformer.TransformCIDR(v)
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) any {
	return c.jsonTransformer.TransformCIDRArray(v)
}

func (c *Client) TransformInet(v *schema.Inet) any {
	return c.jsonTransformer.TransformInet(v)
}

func (c *Client) TransformInetArray(v *schema.InetArray) any {
	return c.jsonTransformer.TransformInetArray(v)
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) any {
	return c.jsonTransformer.TransformMacaddr(v)
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	return c.jsonTransformer.TransformMacaddrArray(v)
}
