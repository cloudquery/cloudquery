package client

import "github.com/cloudquery/plugin-sdk/schema"

func (c *Client) TransformBool(v *schema.Bool) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVBool(v)
	case FormatTypeJSON:
		return c.transformJSONBool(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformBytea(v *schema.Bytea) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVBytea(v)
	case FormatTypeJSON:
		return c.transformJSONBytea(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformFloat8(v *schema.Float8) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVFloat8(v)
	case FormatTypeJSON:
		return c.transformJSONFloat8(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformInt8(v *schema.Int8) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVInt8(v)
	case FormatTypeJSON:
		return c.transformJSONInt8(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVInt8Array(v)
	case FormatTypeJSON:
		return c.transformJSONInt8Array(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformJSON(v *schema.JSON) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVJSON(v)
	case FormatTypeJSON:
		return c.transformJSONJSON(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformText(v *schema.Text) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVText(v)
	case FormatTypeJSON:
		return c.transformJSONText(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformTextArray(v *schema.TextArray) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVTextArray(v)
	case FormatTypeJSON:
		return c.transformJSONTextArray(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVTimestamptz(v)
	case FormatTypeJSON:
		return c.transformJSONTimestamptz(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformUUID(v *schema.UUID) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVUUID(v)
	case FormatTypeJSON:
		return c.transformJSONUUID(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVUUIDArray(v)
	case FormatTypeJSON:
		return c.transformJSONUUIDArray(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformCIDR(v *schema.CIDR) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVCIDR(v)
	case FormatTypeJSON:
		return c.transformJSONCIDR(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVCIDRArray(v)
	case FormatTypeJSON:
		return c.transformJSONCIDRArray(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformInet(v *schema.Inet) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVInet(v)
	case FormatTypeJSON:
		return c.transformJSONInet(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformInetArray(v *schema.InetArray) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVInetArray(v)
	case FormatTypeJSON:
		return c.transformJSONInetArray(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVMacaddr(v)
	case FormatTypeJSON:
		return c.transformJSONMacaddr(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	switch c.csvSpec.Format {
	case FormatTypeCSV:
		return c.transformCSVMacaddrArray(v)
	case FormatTypeJSON:
		return c.transformJSONMacaddrArray(v)
	default:
		panic("unknown format" + c.csvSpec.Format)
	}
}
