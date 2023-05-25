package client

import (
	"strconv"
	"strings"
)

// https://github.com/snowflakedb/gosnowflake/issues/674
func snowflakeStrToArray(val string) []string {
	val = strings.TrimPrefix(val, "[\n  ")
	val = strings.TrimSuffix(val, "\n]")
	strs := strings.Split(val, ",\n  ")
	for i := range strs {
		if unquoted, err := strconv.Unquote(strs[i]); err == nil {
			strs[i] = unquoted
		}
	}
	return strs
}

// https://github.com/snowflakedb/gosnowflake/issues/674
func snowflakeStrToIntArray(val string) []string {
	val = strings.TrimPrefix(val, "[\n  ")
	val = strings.TrimSuffix(val, "\n]")
	strs := strings.Split(val, ",\n  ")
	for i := range strs {
		strs[i] = strings.ReplaceAll(strs[i], "\\\"", "\"")
		strs[i] = strings.ReplaceAll(strs[i], "\\n", "\n")
	}
	return strs
}
