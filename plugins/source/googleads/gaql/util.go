package gaql

import (
	"reflect"
	"strings"

	"golang.org/x/exp/slices"
)

func quoted(s string) string {
	return `"` + s + `"`
}

func jsonTag(fld reflect.StructField) (string, bool) {
	tag := strings.Split(fld.Tag.Get("json"), ",")[0]
	return tag, len(tag) > 0 && tag != "-"
}

func trimStrings(prefix string, strs []string) []string {
	if len(prefix) == 0 {
		return slices.Clone(strs)
	}

	res := make([]string, 0, len(strs))
	for _, el := range strs {
		if !strings.HasPrefix(el, prefix) {
			continue
		}

		trimmed := strings.TrimPrefix(el, prefix)
		if len(trimmed) == 0 {
			continue
		}

		res = append(res, trimmed)
	}

	return res
}
