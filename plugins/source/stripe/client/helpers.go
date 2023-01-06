package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/transformers"
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
