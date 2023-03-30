package client

import (
	"context"
	"reflect"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/thoas/go-funk"
)

var (
	options = []transformers.StructTransformerOption{
		transformers.WithTypeTransformer(timestampTypeTransformer),
		transformers.WithResolverTransformer(timestampResolverTransformer),
	}
)

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}
func timestampTypeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if !strings.HasSuffix(field.Name, "At") {
		return schema.TypeInvalid, nil // fallback
	}

	switch field.Type {
	case reflect.TypeOf(""), reflect.TypeOf(new(string)): // nop
	default:
		return schema.TypeInvalid, nil // fallback
	}

	return schema.TypeTimestamp, nil
}

func timestampResolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if t, _ := timestampTypeTransformer(field); t != schema.TypeTimestamp {
		return transformers.DefaultResolverTransformer(field, path)
	}

	if field.Type.Kind() == reflect.Pointer {
		return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
			val := funk.Get(r.Item, path)
			if val == nil {
				return r.Set(c.Name, nil)
			}

			t, err := time.Parse(time.RFC3339, *(val.(*string)))
			if err != nil {
				return err
			}
			return r.Set(c.Name, t)
		}
	}

	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		val := funk.Get(r.Item, path)
		if val == nil {
			return r.Set(c.Name, nil)
		}

		t, err := time.Parse(time.RFC3339, val.(string))
		if err != nil {
			return err
		}
		return r.Set(c.Name, t)
	}
}
