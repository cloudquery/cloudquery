package client

import (
	"net"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5/pgtype"
)

func (*Client) TransformBool(v *schema.Bool) interface{} {
	return &pgtype.Bool{
		Bool:   v.Bool,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformBytea(v *schema.Bytea) interface{} {
	return v.Bytes
}

func (*Client) TransformFloat8(v *schema.Float8) interface{} {
	return &pgtype.Float8{
		Float64:  v.Float,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformInt8(v *schema.Int8) interface{} {
	return &pgtype.Int8{
		Int64:    v.Int,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformInt8Array(v *schema.Int8Array) interface{} {
	r := pgtype.Array[pgtype.Int8]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Int8{Int64: e.Int, Valid: e.Status == schema.Present})
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformJSON(v *schema.JSON) interface{} {
	return v.Bytes
}

func (*Client) TransformText(v *schema.Text) interface{} {
	return &pgtype.Text{
		String: stripNulls(v.Str),
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformTextArray(v *schema.TextArray) interface{} {
	r := pgtype.Array[pgtype.Text]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.Text{String: stripNulls(e.Str), Valid: e.Status == schema.Present})
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) interface{} {
	return &pgtype.Timestamptz{
		Time:   v.Time,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformUUID(v *schema.UUID) interface{} {
	return pgtype.UUID{
		Bytes:  v.Bytes,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) interface{} {
	r := pgtype.Array[pgtype.UUID]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, pgtype.UUID{Bytes: e.Bytes, Valid: e.Status == schema.Present})
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformCIDR(v *schema.CIDR) interface{} {
	return v.IPNet
}

func (c *Client) TransformCIDRArray(v *schema.CIDRArray) interface{} {
	r := pgtype.Array[*net.IPNet]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, e.IPNet)
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (*Client) TransformInet(v *schema.Inet) interface{} {
	return v.IPNet
}

func (*Client) TransformInetArray(v *schema.InetArray) interface{} {
	r := pgtype.Array[*net.IPNet]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, e.IPNet)
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func (c *Client) TransformMacaddr(v *schema.Macaddr) interface{} {
	return v.Addr
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) interface{} {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.Array[pgtype.Text]{}
		for _, e := range v.Elements {
			r.Elements = append(r.Elements, pgtype.Text{String: e.String(), Valid: e.Status == schema.Present})
		}
		r.Valid = v.Status == schema.Present
		for _, d := range v.Dimensions {
			r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
		}
		return &r
	}
	r := pgtype.Array[net.HardwareAddr]{}
	for _, e := range v.Elements {
		r.Elements = append(r.Elements, e.Addr)
	}
	r.Valid = v.Status == schema.Present
	for _, d := range v.Dimensions {
		r.Dims = append(r.Dims, pgtype.ArrayDimension{Length: d.Length, LowerBound: d.LowerBound})
	}
	return &r
}

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}
