package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgtype"
)


func (c *Client) TransformBool(v *schema.Bool) interface{} {
	return &pgtype.Bool{
		Bool:   v.Bool,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformBytea(v *schema.Bytea) interface{} {
	return &pgtype.Bytea{
		Bytes:  v.Bytes,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformFloat8(v *schema.Float8) interface{} {
	return &pgtype.Float8{
		Float:  v.Float,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformInt8(v *schema.Int8) interface{} {
	return &pgtype.Int8{
		Int:    v.Int,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	r := pgtype.Int8Array{}
	for _, v := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Int8{Int: v.Int, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(v.Status)
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformJSON(v *schema.JSON) interface{} {
	return &pgtype.JSON{
		Bytes:  v.Bytes,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformText(v *schema.Text) interface{} {
	return &pgtype.Text{
		String: v.Str,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformTextArray(v *schema.TextArray) interface{} {
	r := pgtype.TextArray{}
	for _, v := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: v.Str, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(v.Status)
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return &pgtype.Timestamptz{
		Time:   v.Time,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformUUID(v *schema.UUID) interface{} {
	return pgtype.UUID{
		Bytes:  v.Bytes,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	r := pgtype.UUIDArray{}
	for _, v := range v.Elements {
		r.Elements = append(r.Elements, pgtype.UUID{Bytes: v.Bytes, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(v.Status)
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformCIDR(v *schema.CIDR) interface{} {
	return &pgtype.CIDR{
		IPNet:  v.IPNet,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.InetArray{}
		for _, v := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
		}
		r.Status = pgtype.Status(v.Status)
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	} else {
		r := pgtype.CIDRArray{}
		for _, v := range v.Elements {
			r.Elements = append(r.Elements, pgtype.CIDR{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
		}
		r.Status = pgtype.Status(v.Status)
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	}
}

func (c *Client) TransformInet(v *schema.Inet) interface{} {
	return &pgtype.Inet{
		IPNet:  v.IPNet,
		Status: pgtype.Status(v.Status),
	}
}

func (c *Client) TransformInetArray(v *schema.InetArray) interface{} {
	r := pgtype.InetArray{}
	for _, v := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: v.IPNet, Status: pgtype.Status(v.Status)})
	}
	r.Status = pgtype.Status(v.Status)
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	if c.pgType == pgTypeCockroachDB {
		return &pgtype.Text{
			String: v.String(),
			Status: pgtype.Status(v.Status),
		}
	} else {
		return &pgtype.Macaddr{
			Addr:   v.Addr,
			Status: pgtype.Status(v.Status),
		}
	}
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.TextArray{}
		for _, v := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Text{String: v.String(), Status: pgtype.Status(v.Status)})
		}
		r.Status = pgtype.Status(v.Status)
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	} else {
		r := pgtype.MacaddrArray{}
		for _, v := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Macaddr{Addr: v.Addr, Status: pgtype.Status(v.Status)})
		}
		r.Status = pgtype.Status(v.Status)
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	}
}