package client

import (
	"encoding/json"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) transformJSONBool(v *schema.Bool) interface{} {
	return v.String()
}

func (*Client) transformJSONBytea(v *schema.Bytea) interface{} {
	return v.String()
}

func (*Client) transformJSONFloat8(v *schema.Float8) interface{} {
	return v.Float
}

func (*Client) transformJSONInt8(v *schema.Int8) interface{} {
	return v.Int
}

func (*Client) transformJSONInt8Array(v *schema.Int8Array) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) transformJSONJSON(v *schema.JSON) interface{} {
	var res map[string]interface{}
	if err := json.Unmarshal(v.Bytes, &res); err != nil {
		panic(err)
	}
	return res
}

func (*Client) transformJSONText(v *schema.Text) interface{} {
	return v.String()
}

func (*Client) transformJSONTextArray(v *schema.TextArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) transformJSONTimestamptz(v *schema.Timestamptz) interface{} {
	return v.String()
}

func (*Client) transformJSONUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (*Client) transformJSONUUIDArray(v *schema.UUIDArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) transformJSONCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (*Client) transformJSONCIDRArray(v *schema.CIDRArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) transformJSONInet(v *schema.Inet) interface{} {
	return v.String()
}

func (*Client) transformJSONInetArray(v *schema.InetArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}

func (*Client) transformJSONMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (*Client) transformJSONMacaddrArray(v *schema.MacaddrArray) interface{} {
	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.String()
	}
	return res
}
