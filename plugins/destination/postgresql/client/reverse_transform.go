package client

import (
	"fmt"
	"net"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/jackc/pgtype"
)

func (*Client) ReverseTransformValues(table *schema.Table, values []interface{}) (schema.CQTypes, error) {
	result := make(schema.CQTypes, len(values))

	for i, value := range values {
		var err error
		result[i], err = reverseTransformValue(table.Columns[i].Type, value)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// To reverse-transform the values from the DB to CQtypes, we try to use `cqtype.Set`.
// `cqtype.Set“ can't convert pgx types, so we handle them separately.
func reverseTransformValue(valueType schema.ValueType, value interface{}) (schema.CQType, error) {
	switch valueType {
	case schema.TypeStringArray:
		pgxTextArray, ok := value.(pgtype.TextArray)
		if ok {
			if len(pgxTextArray.Elements) == 0 {
				return &schema.TextArray{Status: schema.Present}, nil
			}

			elements := make([]schema.Text, len(pgxTextArray.Elements))
			for i := range pgxTextArray.Elements {
				elements[i] = schema.Text{Str: pgxTextArray.Elements[i].String, Status: cqStatusFromPgTypeStatus(pgxTextArray.Elements[i].Status)}
			}

			return &schema.TextArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxTextArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}
	case schema.TypeIntArray:
		pgxIntArray, ok := value.(pgtype.Int8Array)
		if ok {
			if len(pgxIntArray.Elements) == 0 {
				return &schema.Int8Array{Status: schema.Present}, nil
			}

			elements := make([]schema.Int8, len(pgxIntArray.Elements))
			for i := range pgxIntArray.Elements {
				elements[i] = schema.Int8{Int: pgxIntArray.Elements[i].Int, Status: cqStatusFromPgTypeStatus(pgxIntArray.Elements[i].Status)}
			}

			return &schema.Int8Array{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxIntArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}
	case schema.TypeUUIDArray:
		pgxUUIDArray, ok := value.(pgtype.UUIDArray)
		if ok {
			if len(pgxUUIDArray.Elements) == 0 {
				return &schema.UUIDArray{Status: schema.Present}, nil
			}

			elements := make([]schema.UUID, len(pgxUUIDArray.Elements))
			for i := range pgxUUIDArray.Elements {
				elements[i] = schema.UUID{Bytes: pgxUUIDArray.Elements[i].Bytes, Status: cqStatusFromPgTypeStatus(pgxUUIDArray.Elements[i].Status)}
			}

			return &schema.UUIDArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxUUIDArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}

	case schema.TypeInetArray:
		pgxInetArray, ok := value.(pgtype.InetArray)
		if ok {
			if len(pgxInetArray.Elements) == 0 {
				return &schema.InetArray{Status: schema.Present}, nil
			}

			elements := make([]schema.Inet, len(pgxInetArray.Elements))
			for i := range pgxInetArray.Elements {
				elements[i] = schema.Inet{IPNet: pgxInetArray.Elements[i].IPNet, Status: cqStatusFromPgTypeStatus(pgxInetArray.Elements[i].Status)}
			}

			return &schema.InetArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxInetArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}
	case schema.TypeCIDRArray:
		pgxCidrArray, ok := value.(pgtype.CIDRArray)
		if ok {
			if len(pgxCidrArray.Elements) == 0 {
				return &schema.CIDRArray{Status: schema.Present}, nil
			}

			elements := make([]schema.CIDR, len(pgxCidrArray.Elements))
			for i := range pgxCidrArray.Elements {
				elements[i] = schema.CIDR{IPNet: pgxCidrArray.Elements[i].IPNet, Status: cqStatusFromPgTypeStatus(pgxCidrArray.Elements[i].Status)}
			}

			return &schema.CIDRArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxCidrArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}
	case schema.TypeMacAddrArray:
		pgxMacAddrArray, ok := value.(pgtype.MacaddrArray)
		if ok {
			if len(pgxMacAddrArray.Elements) == 0 {
				return &schema.MacaddrArray{Status: schema.Present}, nil
			}

			elements := make([]schema.Macaddr, len(pgxMacAddrArray.Elements))
			for i := range pgxMacAddrArray.Elements {
				elements[i] = schema.Macaddr{Addr: pgxMacAddrArray.Elements[i].Addr, Status: cqStatusFromPgTypeStatus(pgxMacAddrArray.Elements[i].Status)}
			}

			return &schema.MacaddrArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(pgxMacAddrArray.Elements)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}

		// For some reason mac-addr-array sometimes comes back from postgres as a string in the format "{macaddr,macaddr,macaddr}"...
		stringMacAddrArray, ok := value.(string)
		if ok {
			trimmedString := strings.Trim(stringMacAddrArray, "{}")
			stringSlice := strings.Split(trimmedString, ",")

			if len(stringSlice) == 0 {
				return &schema.MacaddrArray{Status: schema.Present}, nil
			}

			elements := make([]schema.Macaddr, len(stringSlice))
			for i := range stringSlice {
				mac, err := net.ParseMAC(stringSlice[i])
				if err != nil {
					return nil, err
				}
				elements[i] = schema.Macaddr{Addr: net.HardwareAddr(mac), Status: schema.Present}
			}

			return &schema.MacaddrArray{
				Elements:   elements,
				Dimensions: []schema.ArrayDimension{{Length: int32(len(stringSlice)), LowerBound: 1}},
				Status:     schema.Present,
			}, nil
		}
	}

	cqtype := schema.NewCqTypeFromValueType(valueType)
	if err := cqtype.Set(value); err != nil {
		return nil, fmt.Errorf("failed to convert value %v to type %s: %w", value, valueType, err)
	}

	return cqtype, nil
}

// Used for testing
func cqStatusFromPgTypeStatus(pgstatus pgtype.Status) schema.Status {
	switch pgstatus {
	case pgtype.Present:
		return schema.Present
	case pgtype.Null:
		return schema.Null
	default:
		return schema.Undefined
	}
}
