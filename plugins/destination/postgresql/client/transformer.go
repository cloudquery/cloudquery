package client

import (
	"encoding/json"
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
		return stripNullsFromMarshalledJson(v.Bytes)
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
	// In postgresql, our underlyting type for timestamps is 'timestamp without timezone', so we have to
	// convert it to UTC here (to avoid stripping the timezone information on INSERT).
	return &pgtype.Timestamptz{
		Time:  v.Time.UTC(),
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

// Best effort to strip nulls from json. In case of an error, we fallback to returning the input.
// We only strip the nulls from the values (i.e. never from the keys).
// In some far-fetched case a JSON like this may happen, but we ignore this for now: `{"key\u0000": "v1", "key\u0000\u0000": "v2"}`
func stripNullsFromMarshalledJson(inputJsonMarshaled []byte) []byte {
	var m any
	err := json.Unmarshal(inputJsonMarshaled, &m)
	if err != nil {
		return inputJsonMarshaled
	}

	m = stripNullsFromJsonValue(m)
	outputJsonMarshaled, err := json.Marshal(m)
	if err != nil {
		return inputJsonMarshaled
	}

	return outputJsonMarshaled
}

func stripNullsFromJsonValue(val any) any {
	switch concreteVal := val.(type) {
	case string:
		return stripNulls(concreteVal)
	case map[string]any:
		for k := range concreteVal {
			concreteVal[k] = stripNullsFromJsonValue(concreteVal[k])
		}
		return concreteVal
	case []any:
		for i := range concreteVal {
			concreteVal[i] = stripNullsFromJsonValue(concreteVal[i])
		}
		return concreteVal
	default:
		return val
	}
}
