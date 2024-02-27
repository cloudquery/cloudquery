package client

import (
	"context"
	"fmt"
	"reflect"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/okta/okta-sdk-golang/v4/okta"
	"github.com/thoas/go-funk"
)

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
	transformers.WithPrimaryKeys("Id"),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch field.Type {
	case reflect.TypeOf(okta.NullableTime{}):
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case reflect.TypeOf(okta.NullableString{}):
		return arrow.BinaryTypes.String, nil
	default:
		return nil, nil
	}
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	switch field.Type {
	case reflect.TypeOf(okta.NullableTime{}):
		return resolveNullableTime(path)
	case reflect.TypeOf(okta.NullableString{}):
		return resolveNullableString(path)
	default:
		return transformers.DefaultResolverTransformer(field, path)
	}
}

func resolveNullableTime(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(okta.NullableTime)
		if !ok {
			return fmt.Errorf("unexpected type, want \"okta.NullableTime\", have \"%T\"", data)
		}
		if !ts.IsSet() {
			return resource.Set(c.Name, nil)
		}
		return resource.Set(c.Name, ts.Get())
	}
}

func resolveNullableString(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		str, ok := data.(okta.NullableString)
		if !ok {
			return fmt.Errorf("unexpected type, want \"okta.NullableString\", have \"%T\"", data)
		}
		if !str.IsSet() {
			return resource.Set(c.Name, nil)
		}
		return resource.Set(c.Name, str.Get())
	}
}
