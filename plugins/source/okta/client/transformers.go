package client

import (
	"context"
	"fmt"
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/okta/okta-sdk-golang/v3/okta"
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
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}

	return nil, nil
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	if field.Type == reflect.TypeOf(okta.NullableTime{}) {
		return resolveNullableTime(path)
	}

	return transformers.DefaultResolverTransformer(field, path)
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
