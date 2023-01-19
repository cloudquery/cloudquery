package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) TransformBool(v *schema.Bool) any {
	return c.filetype.TransformBool(v)
}

func (c *Client) TransformBytea(v *schema.Bytea) any {
	return c.filetype.TransformBytea(v)
}

func (c *Client) TransformFloat8(v *schema.Float8) any {
	return c.filetype.TransformFloat8(v)
}

func (c *Client) TransformInt8(v *schema.Int8) any {
	return c.filetype.TransformInt8(v)
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) any {
	return c.filetype.TransformInt8Array(v)
}

func (c *Client) TransformJSON(v *schema.JSON) any {
	return c.filetype.TransformJSON(v)
}

func (c *Client) TransformText(v *schema.Text) any {
	return c.filetype.TransformText(v)
}

func (c *Client) TransformTextArray(v *schema.TextArray) any {
	return c.filetype.TransformTextArray(v)
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) any {
	return c.filetype.TransformTimestamptz(v)
}

func (c *Client) TransformUUID(v *schema.UUID) any {
	return c.filetype.TransformUUID(v)
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) any {
	return c.filetype.TransformUUIDArray(v)
}

func (c *Client) TransformCIDR(v *schema.CIDR) any {
	return c.filetype.TransformCIDR(v)
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) any {
	return c.filetype.TransformCIDRArray(v)
}

func (c *Client) TransformInet(v *schema.Inet) any {
	return c.filetype.TransformInet(v)
}

func (c *Client) TransformInetArray(v *schema.InetArray) any {
	return c.filetype.TransformInetArray(v)
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) any {
	return c.filetype.TransformMacaddr(v)
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	return c.filetype.TransformMacaddrArray(v)
}
