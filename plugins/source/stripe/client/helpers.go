package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func CreateIgnoreInTestsTransformer(fieldNames ...string) transformers.IgnoreInTestsTransformer {
	return func(field reflect.StructField) bool {
		for _, v := range fieldNames {
			if field.Name == v {
				return true
			}
		}
		return false
	}
}

func MaxInt64(a, b *int64) *int64 {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if *a > *b {
		return a
	}
	return b
}
