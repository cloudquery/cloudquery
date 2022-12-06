package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

// Called in tests
func (*Client) ReverseTransformValues(table *schema.Table, values []interface{}) (schema.CQTypes, error) {
	result := make(schema.CQTypes, len(values))

	for i, value := range values {
		var err error
		result[i], err = c.reverseTransformValue(table.Columns[i].Type, value)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// By default, we try to use `cqtype.Set` - but some cases need to be handled separately.
func reverseTransformValue(valueType schema.ValueType, value interface{}) (schema.CQType, error) {
	switch valueType {
	case schema.TypeStringArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.TextArray{Status: schema.Present}, nil
			}
			elements := make([]schema.Text, len(stringSlice))
			for i, str := range stringSlice {
				elements[i] = schema.Text{Str: str, Status: schema.Present}
			}

			return &schema.TextArray{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}
	case schema.TypeIntArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.Int8Array{Status: schema.Present}, nil
			}
			elements := make([]schema.Int8, len(stringSlice))
			for i, str := range stringSlice {
				num, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("failed to ParseInt %s", str)
				}
				elements[i] = schema.Int8{Int: num, Status: schema.Present}
			}

			return &schema.Int8Array{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}
	case schema.TypeUUIDArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.UUIDArray{Status: schema.Present}, nil
			}
			elements := make([]schema.UUID, len(stringSlice))
			for i, str := range stringSlice {
				elements[i].Set(str)
			}

			return &schema.UUIDArray{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}
	case schema.TypeInetArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.InetArray{Status: schema.Present}, nil
			}
			elements := make([]schema.Inet, len(stringSlice))
			for i, str := range stringSlice {
				elements[i].Set(str)
			}

			return &schema.InetArray{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}

	case schema.TypeCIDRArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.CIDRArray{Status: schema.Present}, nil
			}
			elements := make([]schema.CIDR, len(stringSlice))
			for i, str := range stringSlice {
				elements[i].Set(str)
			}

			return &schema.CIDRArray{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}
	case schema.TypeMacAddrArray:
		stringSlice, ok := splitSqlArrayToStrings(value)
		if ok {
			if len(stringSlice) == 0 {
				return &schema.MacaddrArray{Status: schema.Present}, nil
			}
			elements := make([]schema.Macaddr, len(stringSlice))
			for i, str := range stringSlice {
				elements[i].Set(str)
			}

			return &schema.MacaddrArray{
				Elements:   elements,
				Status:     schema.Present,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(elements)), LowerBound: 1}},
			}, nil
		}
	}

	cqtype := schema.NewCqTypeFromValueType(valueType)
	if err := cqtype.Set(value); err != nil {
		return nil, fmt.Errorf("failed to convert value %v to type %s: %w", value, valueType, err)
	}

	return cqtype, nil
}

// splitSqlArrayToStrings splits a strings such as '{a,b,c}' into []string{'a', 'b', 'c}
// Receives an interface{} because the value can be either a string or a *string.
// Returns ok if the value was a string or a *string, and false otherwise.
func splitSqlArrayToStrings(val interface{}) ([]string, bool) {
	str, ok := val.(string)
	if !ok {
		var ptr *string
		ptr, ok = val.(*string)
		if ok {
			str = *ptr
		}
	}

	if !ok {
		return nil, false
	}

	str = strings.Trim(str, "{}")
	return strings.Split(str, ","), true
}
