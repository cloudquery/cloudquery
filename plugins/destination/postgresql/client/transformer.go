package client

import (
	"net"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgx/v5/pgtype"
)

func (*Client) TransformBool(v *schema.Bool) any {
	return &pgtype.Bool{
		Bool:  v.Bool,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	if v.Status == schema.Present {
		return v.Bytes
	}
	return nil
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	return &pgtype.Float8{
		Float64: v.Float,
		Valid:   v.Status == schema.Present,
	}
}

func (*Client) TransformInt8(v *schema.Int8) any {
	return &pgtype.Int8{
		Int64: v.Int,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	r := pgtype.FlatArray[pgtype.Int8]{}
	for _, e := range v.Elements {
		r = append(r, pgtype.Int8{Int64: e.Int, Valid: e.Status == schema.Present})
	}
	return &r
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status == schema.Present {
		return v.Bytes
	}
	return nil
}

func (*Client) TransformText(v *schema.Text) any {
	return &pgtype.Text{
		String: stripNulls(v.Str),
		Valid:  v.Status == schema.Present,
	}
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	r := pgtype.FlatArray[pgtype.Text]{}
	for _, e := range v.Elements {
		r = append(r, pgtype.Text{String: stripNulls(e.Str), Valid: e.Status == schema.Present})
	}
	return &r
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	return &pgtype.Timestamptz{
		Time:  v.Time,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformUUID(v *schema.UUID) any {
	return pgtype.UUID{
		Bytes: v.Bytes,
		Valid: v.Status == schema.Present,
	}
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	r := pgtype.FlatArray[pgtype.UUID]{}
	for _, e := range v.Elements {
		r = append(r, pgtype.UUID{Bytes: e.Bytes, Valid: e.Status == schema.Present})
	}
	return &r
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	if v.Status == schema.Present {
		return v.IPNet
	}
	return nil
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	r := pgtype.FlatArray[*net.IPNet]{}
	for _, e := range v.Elements {
		r = append(r, e.IPNet)
	}
	return &r
}

func (*Client) TransformInet(v *schema.Inet) any {
	if v.Status == schema.Present {
		return v.IPNet
	}
	return nil
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	r := pgtype.FlatArray[*net.IPNet]{}
	for _, e := range v.Elements {
		r = append(r, e.IPNet)
	}
	return &r
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status == schema.Present {
		return v.Addr
	}
	return nil
}

func (c *Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	if c.pgType == pgTypeCockroachDB {
		r := pgtype.FlatArray[pgtype.Text]{}
		for _, e := range v.Elements {
			r = append(r, pgtype.Text{String: e.String(), Valid: e.Status == schema.Present})
		}
		return &r
	}
	r := pgtype.FlatArray[net.HardwareAddr]{}
	for _, e := range v.Elements {
		r = append(r, e.Addr)
	}
	return &r
}

func stripNulls(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}
