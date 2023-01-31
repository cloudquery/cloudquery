package client

import "github.com/cloudquery/plugin-sdk/schema"

func (c *Client) TransformBool(v *schema.Bool) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformBool(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformBool(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformBytea(v *schema.Bytea) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformBytea(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformBytea(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformFloat8(v *schema.Float8) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformFloat8(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformFloat8(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInt8(v *schema.Int8) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformInt8(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformInt8(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformInt8Array(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformInt8Array(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformJSON(v *schema.JSON) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformJSON(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformJSON(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformText(v *schema.Text) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformText(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformText(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformTextArray(v *schema.TextArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformTextArray(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformTextArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformTimestamptz(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformTimestamptz(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformUUID(v *schema.UUID) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformUUID(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformUUID(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformUUIDArray(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformUUIDArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformCIDR(v *schema.CIDR) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformCIDR(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformCIDR(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformCIDRArray(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformCIDRArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInet(v *schema.Inet) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformInet(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformInet(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformInetArray(v *schema.InetArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformInetArray(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformInetArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformMacaddr(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformMacaddr(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.CSVTransformer.TransformMacaddrArray(v)
	case FormatTypeJSON:
		return c.JSONTransformer.TransformMacaddrArray(v)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}
