package client

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return schema.TypeTimestamp, nil
	case *durationpb.Duration,
		durationpb.Duration:
		return schema.TypeInt, nil
	case protoreflect.Enum:
		return schema.TypeString, nil
	default:
		return schema.TypeInvalid, nil
	}
}

func resolverTransformer(field reflect.StructField, path string) schema.ColumnResolver {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return ResolveProtoTimestamp(path)
	case *durationpb.Duration,
		durationpb.Duration:
		return ResolveProtoDuration(path)
	case protoreflect.Enum:
		return ResolveProtoEnum(path)
	default:
		return nil
	}
}

var options = []transformers.StructTransformerOption{
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}
