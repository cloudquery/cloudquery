package client

import "github.com/cloudquery/plugin-sdk/schema"

func (*Client) transformCSVBool(v *schema.Bool) interface{} {
	return v.String()
}

func (*Client) transformCSVBytea(v *schema.Bytea) interface{} {
	return v.String()
}

func (*Client) transformCSVFloat8(v *schema.Float8) interface{} {
	return v.String()
}

func (*Client) transformCSVInt8(v *schema.Int8) interface{} {
	return v.String()
}

func (*Client) transformCSVInt8Array(v *schema.Int8Array) interface{} {
	return v.String()
}

func (*Client) transformCSVJSON(v *schema.JSON) interface{} {
	return v.String()
}

func (*Client) transformCSVText(v *schema.Text) interface{} {
	return v.String()
}

func (*Client) transformCSVTextArray(v *schema.TextArray) interface{} {
	return v.String()
}

func (*Client) transformCSVTimestamptz(v *schema.Timestamptz) interface{} {
	return v.String()
}

func (*Client) transformCSVUUID(v *schema.UUID) interface{} {
	return v.String()
}

func (*Client) transformCSVUUIDArray(v *schema.UUIDArray) interface{} {
	return v.String()
}

func (*Client) transformCSVCIDR(v *schema.CIDR) interface{} {
	return v.String()
}

func (*Client) transformCSVCIDRArray(v *schema.CIDRArray) interface{} {
	return v.String()
}

func (*Client) transformCSVInet(v *schema.Inet) interface{} {
	return v.String()
}

func (*Client) transformCSVInetArray(v *schema.InetArray) interface{} {
	return v.String()
}

func (*Client) transformCSVMacaddr(v *schema.Macaddr) interface{} {
	return v.String()
}

func (*Client) transformCSVMacaddrArray(v *schema.MacaddrArray) interface{} {
	return v.String()
}
