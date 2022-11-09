package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgtype"
)

func (*Client) TransformBool(v *schema.Bool) interface{} {
	return &pgtype.Bool{
		Bool:   v.Bool,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformBytea(v *schema.Bytea) interface{} {
	return &pgtype.Bytea{
		Bytes:  v.Bytes,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformFloat8(v *schema.Float8) interface{} {
	return &pgtype.Float8{
		Float:  v.Float,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformInt8(v *schema.Int8) interface{} {
	return &pgtype.Int8{
		Int:    v.Int,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	r := pgtype.Int8Array{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Int8{Int: e.Int, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformJSON(v *schema.JSON) interface{} {
	return &pgtype.JSON{
		Bytes:  v.Bytes,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformText(v *schema.Text) interface{} {
	return &pgtype.Text{
		String: v.Str,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformTextArray(v *schema.TextArray) interface{} {
	r := pgtype.TextArray{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: e.Str, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return &pgtype.Timestamptz{
		Time:   v.Time,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformUUID(v *schema.UUID) interface{} {
	return pgtype.UUID{
		Bytes:  v.Bytes,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	r := pgtype.UUIDArray{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.UUID{Bytes: e.Bytes, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformCIDR(v *schema.CIDR) interface{} {
	return &pgtype.CIDR{
		IPNet:  v.IPNet,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.InetArray{}
		for _, e := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Inet{IPNet: e.IPNet, Status: cqStatusToPgStatus[e.Status]})
		}
		r.Status = cqStatusToPgStatus[v.Status]
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	}
	r := pgtype.CIDRArray{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.CIDR{IPNet: e.IPNet, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformInet(v *schema.Inet) interface{} {
	return &pgtype.Inet{
		IPNet:  v.IPNet,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (*Client) TransformInetArray(v *schema.InetArray) interface{} {
	r := pgtype.InetArray{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Inet{IPNet: e.IPNet, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	if c.pgType == pgTypeCockroachDB {
		return &pgtype.Text{
			String: v.String(),
			Status: cqStatusToPgStatus[v.Status],
		}
	}
	return &pgtype.Macaddr{
		Addr:   v.Addr,
		Status: cqStatusToPgStatus[v.Status],
	}
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.TextArray{}
		for _, e := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Text{String: e.String(), Status: cqStatusToPgStatus[e.Status]})
		}
		r.Status = cqStatusToPgStatus[v.Status]
		for _, d := range v.Dimensions {
			r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	}
	r := pgtype.MacaddrArray{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Macaddr{Addr: e.Addr, Status: cqStatusToPgStatus[e.Status]})
	}
	r.Status = cqStatusToPgStatus[v.Status]
	for _, d := range v.Dimensions {
		r.Dimensions = append(r.Dimensions, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}
