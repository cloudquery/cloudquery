package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpgradeParquetTypeName(t *testing.T) {
	cases := []struct {
		parquetType string
		want        string
	}{
		{"optional int64 (Int(bitWidth=64, isSigned=true))", "int64"},
		{"required int32 (Int(bitWidth=16, isSigned=false))", "required uint16"},
		{"optional int64", "int64"},
		{"optional float", "float"},
		{"optional double", "double"},
		{"required boolean", "required bool"},
		{"optional byte_array (String)", "string"},
		{"optional byte_array (JSON)", "json"},
		{"optional byte_array", "binary"},
		{"optional fixed_len_byte_array(16) (UUID)", "uuid"},
		{"optional fixed_len_byte_array(8)", "binary(8)"},
		{"optional int64 (Timestamp(isAdjustedToUTC=true, timeUnit=microseconds, is_from_converted_type=false, force_set_converted_type=true))", "timestamp (us, UTC)"},
		{"optional int64 (Timestamp(isAdjustedToUTC=false, timeUnit=milliseconds, is_from_converted_type=false, force_set_converted_type=false))", "timestamp (ms)"},
		{"optional int64 (Time(isAdjustedToUTC=true, timeUnit=nanoseconds))", "time (ns, UTC)"},
		{"optional int32 (Date)", "date"},
		{"optional fixed_len_byte_array(16) (Decimal(precision=38, scale=2))", "decimal(38,2)"},
		{"optional group (List) {list: repeated group {element: optional byte_array (String)}}", "list<string>"},
		{"required group (List) {list: repeated group {element: optional group (List) {list: repeated group {element: optional int64 (Int(bitWidth=64, isSigned=true))}}}}", "required list<list<int64>>"},
		{"optional group (List) {element: repeated int32 (Int(bitWidth=32, isSigned=true))}", "list<int32>"},
		{"optional group (Map) {key_value: repeated group {key: required byte_array (String); value: optional byte_array (String)}}", "map<string,string>"},
		{"optional group {name: optional byte_array (String); size: required int64 (Int(bitWidth=64, isSigned=true))}", "struct<name: string, size: int64>"},
		{"optional group {a: optional boolean; b: optional boolean; c: optional boolean; d: optional boolean}", "struct<a: bool, b: bool, c: bool, …>"},
		{"optional group {tags: optional group (List) {list: repeated group {element: optional byte_array (String)}}}", "struct<tags: list<string>>"},
	}
	for _, tc := range cases {
		t.Run(tc.parquetType, func(t *testing.T) {
			got, ok := upgradeParquetTypeName(tc.parquetType)
			require.True(t, ok)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestUpgradeParquetTypeNameKeepsOtherTypes(t *testing.T) {
	for _, destinationType := range []string{
		"",
		"text",
		"text[]",
		"jsonb",
		"TIMESTAMP(MICROS)",
		"BYTE_ARRAY (JSON)",
		"list<item: utf8, nullable>",
		"optional int64 (Unknown)",
		"optional group (List) {list: repeated group {element: optional byte_array (String)}",
		"optional group (Map) {key_value: repeated group {key: required byte_array (String)}}",
	} {
		t.Run(destinationType, func(t *testing.T) {
			_, ok := upgradeParquetTypeName(destinationType)
			require.False(t, ok)
		})
	}
}
