package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/transformers"
)

var _ transformers.NameTransformer = ETagNameTransformer

func ETagNameTransformer(fld reflect.StructField) (string, error) {
	if fld.Name == "ETag" {
		return "etag", nil
	}
	return transformers.DefaultNameTransformer(fld)
}
