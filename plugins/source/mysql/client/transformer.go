package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

// this is used for tests
type Transformer struct{}

func (*Transformer) TransformBool(v *schema.Bool) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bool
}

func (*Transformer) TransformBytea(v *schema.Bytea) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bytes
}

func (*Transformer) TransformFloat8(v *schema.Float8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Float
}

func (*Transformer) TransformInt8(v *schema.Int8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Int
}

func (*Transformer) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Transformer) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}

	return string(v.Bytes)
}

func (*Transformer) TransformUUID(v *schema.UUID) any {
	if v.Status != schema.Present {
		return nil
	}
	// We need a slice instead of a fixed sized array
	bytes := make([]byte, 16)
	copy(bytes, v.Bytes[:])
	return bytes
}

func (*Transformer) TransformUUIDArray(v *schema.UUIDArray) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}

func (*Transformer) TransformInt8Array(v *schema.Int8Array) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}

func (*Transformer) TransformCIDR(v *schema.CIDR) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Transformer) TransformInet(v *schema.Inet) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Transformer) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Transformer) TransformText(v *schema.Text) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Str
}

func (*Transformer) TransformCIDRArray(v *schema.CIDRArray) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}

func (*Transformer) TransformInetArray(v *schema.InetArray) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}

func (*Transformer) TransformMacaddrArray(v *schema.MacaddrArray) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}

func (*Transformer) TransformTextArray(v *schema.TextArray) any {
	if v.Status != schema.Present {
		return nil
	}

	return v.String()
}
