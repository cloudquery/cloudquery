package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (*Client) TransformBool(v *schema.Bool) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bytes
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Int
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	if v.Status != schema.Present {
		return []*int64(nil)
	}

	res := make([]int64, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Int
	}

	return res
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}
	return unescape(string(v.Bytes))
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Addr.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	if v.Status != schema.Present {
		return []*string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Addr.String()
	}

	return res
}

func (*Client) TransformText(v *schema.Text) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Str
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	if v.Status != schema.Present {
		return []*string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Str
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
	if v.Status != schema.Present {
		return nil
	}
	return uuid.UUID(v.Bytes)
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	if v.Status != schema.Present {
		return []*uuid.UUID(nil)
	}

	res := make([]uuid.UUID, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Bytes
	}

	return res
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.IPNet.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	if v.Status != schema.Present {
		return []*string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.IPNet.String()
	}

	return res
}

func (*Client) TransformInet(v *schema.Inet) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.IPNet.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	if v.Status != schema.Present {
		return []*string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.IPNet.String()
	}

	return res
}
