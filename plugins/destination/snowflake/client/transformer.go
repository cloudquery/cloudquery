package client

import (
	"encoding/hex"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) TransformBool(v *schema.Bool) interface{} {
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) interface{} {
	return hex.EncodeToString(v.Bytes)
}

func (*Client) TransformFloat8(v *schema.Float8) interface{} {
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) interface{} {
	return v.Int
}

func (*Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	// snowIntArray := make([]int64, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowIntArray[i] = e.Int
	// }
	return v.String()
	// return snowIntArray
}

func (*Client) TransformJSON(v *schema.JSON) interface{} {
	return v.String()
	// return nil
	// return v.Bytes
}

func (*Client) TransformText(v *schema.Text) interface{} {
	return v.String()
}

func (*Client) TransformTextArray(v *schema.TextArray) interface{} {
	// snowValue := make([]string, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowValue[i] = e.String()
	// }
	// return snowValue
	// return sf.Array(snowValue)
	return v.String()
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return v.Time
}

func (*Client) TransformUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	// snowValue := make([]string, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowValue[i] = e.String()
	// }
	// return gosnowflake.Array(snowValue)
	return v.String()
	// return snowValue
}

func (*Client) TransformCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	// snowValue := make([]string, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowValue[i] = e.String()
	// }
	// return nil
	// return sf.Array(snowValue)
	return v.String()
}

func (*Client) TransformInet(v *schema.Inet) interface{} {
	return v.String()
}

func (*Client) TransformInetArray(v *schema.InetArray) interface{} {
	// snowValue := make([]string, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowValue[i] = e.String()
	// }
	// return []string{}
	// return sf.Array(snowValue)
	return v.String()
}

func (*Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	// snowValue := make([]string, len(v.Elements))
	// for i, e := range v.Elements {
	// 	snowValue[i] = e.String()
	// }
	// return nil
	return v.String()
}
