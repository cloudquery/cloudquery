package client

import "github.com/cloudquery/plugin-sdk/schema"

func (c *Client) TransformBool(v *schema.Bool) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformBool(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformBool(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformBytea(v *schema.Bytea) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformBytea(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformBytea(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformFloat8(v *schema.Float8) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformFloat8(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformFloat8(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInt8(v *schema.Int8) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformInt8(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformInt8(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformInt8Array(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformInt8Array(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformJSON(v *schema.JSON) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformJSON(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformJSON(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformText(v *schema.Text) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformText(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformText(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformTextArray(v *schema.TextArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformTextArray(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformTextArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformTimestamptz(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformTimestamptz(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformUUID(v *schema.UUID) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformUUID(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformUUID(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformUUIDArray(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformUUIDArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformCIDR(v *schema.CIDR) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformCIDR(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformCIDR(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformCIDRArray(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformCIDRArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInet(v *schema.Inet) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformInet(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformInet(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInetArray(v *schema.InetArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformInetArray(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformInetArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformMacaddr(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformMacaddr(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvTransformer.TransformMacaddrArray(v)
	case FormatTypeJSON:
		return c.jsonTransformer.TransformMacaddrArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}
