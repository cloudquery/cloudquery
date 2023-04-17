package client

import (
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/google/uuid"
)

func getValues(arr arrow.Array) []any {
	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.GetOneForMarshal(i)
		}
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
		return []string(nil)
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
		return []string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, el := range v.Elements {
		res[i] = el.Str
	}

	return res
}

var (
	minDateTime64, _ = time.Parse("2006-01-02 15:04:05", "1925-01-01 00:00:00")
	maxDateTime64, _ = time.Parse("2006-01-02 15:04:05", "2262-04-11 23:47:16")
)

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	if v.Time.Before(minDateTime64) || v.Time.After(maxDateTime64) {
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
		return []uuid.UUID(nil)
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
		return []string(nil)
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
		return []string(nil)
	}

	res := make([]string, len(v.Elements))
	for i, e := range v.Elements {
		res[i] = e.IPNet.String()
	}

	return res
}
