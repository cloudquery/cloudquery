package client

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (*Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	res := make(schema.CQTypes, len(values))

	for i, v := range values {
		t := schema.NewCqTypeFromValueType(table.Columns[i].Type)
		if err := t.Set(unwrap(t.Type(), v)); err != nil {
			return nil, fmt.Errorf("failed to convert value %v to type %s: %w", v, table.Columns[i].Type, err)
		}
		res[i] = t
	}
	return res, nil
}

func unwrap(t schema.ValueType, v any) any {
	switch t {
	case schema.TypeByteArray:
		v := v.(**string)
		if *v == nil {
			return nil
		}
		return []byte(**v)

	case schema.TypeJSON:
		v := v.(**string)
		if *v == nil {
			return nil
		}
		return unescape(**v)

	case schema.TypeStringArray,
		schema.TypeMacAddrArray,
		schema.TypeCIDRArray,
		schema.TypeInetArray:
		return *v.(*[]*string)

	case schema.TypeIntArray:
		return *v.(*[]*int64)

	case schema.TypeUUIDArray:
		return *v.(*[]*uuid.UUID)

	default:
		return v
	}
}
