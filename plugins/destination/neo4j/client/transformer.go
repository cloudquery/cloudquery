package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) any {
	if v.Status == schema.Present {
		return v.Bool
	}
	return nil
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	if v.Status == schema.Present {
		return v.Bytes
	}
	return nil
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	if v.Status == schema.Present {
		return v.Float
	}
	return nil
}

func (*Client) TransformInt8(v *schema.Int8) any {
	if v.Status == schema.Present {
		return v.Int
	}
	return nil
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	res := make([]int64, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.Int
	}
	return res
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status == schema.Present {
		return string(v.Bytes)
	}
	return nil
}

func (*Client) TransformText(v *schema.Text) any {
	if v.Status == schema.Present {
		return v.String()
	}
	return nil
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status == schema.Present {
		return v.Time
	}
	return nil
}

func (*Client) TransformUUID(v *schema.UUID) any {
	if v.Status == schema.Present {
		return v.String()
	}
	return nil
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	if v.Status == schema.Present {
		return v.String()
	}
	return nil
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformInet(v *schema.Inet) any {
	if v.Status == schema.Present {
		return v.String()
	}
	return nil
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status == schema.Present {
		return v.String()
	}
	return nil
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}
