package client

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	mssql "github.com/microsoft/go-mssqldb"
)

func (*Client) TransformBool(v *schema.Bool) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bool
}

func (*Client) TransformBytea(v *schema.Bytea) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Bytes
}

func (*Client) TransformFloat8(v *schema.Float8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Float
}

func (*Client) TransformInt8(v *schema.Int8) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Int
}

func (*Client) TransformTimestamptz(v *schema.Timestamptz) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Time
}

func (*Client) TransformJSON(v *schema.JSON) any {
	if v.Status != schema.Present {
		return nil
	}

	return unescape(string(v.Bytes))
}

func (*Client) TransformUUID(v *schema.UUID) any {
	if v.Status != schema.Present {
		return nil
	}
	val, _ := mssql.UniqueIdentifier(v.Bytes).Value()
	return val
}

func (*Client) TransformUUIDArray(v *schema.UUIDArray) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) TransformInt8Array(v *schema.Int8Array) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) TransformCIDR(v *schema.CIDR) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformInet(v *schema.Inet) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformMacaddr(v *schema.Macaddr) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.String()
}

func (*Client) TransformText(v *schema.Text) any {
	if v.Status != schema.Present {
		return nil
	}
	return v.Str
}

func (*Client) TransformCIDRArray(v *schema.CIDRArray) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) TransformInetArray(v *schema.InetArray) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) TransformMacaddrArray(v *schema.MacaddrArray) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) TransformTextArray(v *schema.TextArray) any {
	if v.Status != schema.Present {
		return nil
	}

	result := make([]string, len(v.Elements))
	for i, elem := range v.Elements {
		result[i] = elem.String()
	}

	return strings.Join(result, ",")
}

func (*Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	res := make(schema.CQTypes, len(values))

	for i, v := range values {
		typ := table.Columns[i].Type
		// The array types are written to the db as values, separated by ",".
		// The default implementation will also wrap them with {} that is missing here.
		switch typ {
		case schema.TypeUUIDArray,
			schema.TypeIntArray,
			schema.TypeCIDRArray,
			schema.TypeInetArray,
			schema.TypeMacAddrArray,
			schema.TypeStringArray:
			v = "{" + v.(string) + "}"
		}
		t := schema.NewCqTypeFromValueType(typ)
		if err := t.Set(v); err != nil {
			return nil, fmt.Errorf("failed to convert value %v to type %s: %w", v, table.Columns[i].Type, err)
		}
		res[i] = t
	}
	return res, nil
}
