package client

import (
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	switch reflect.New(field.Type).Elem().Interface().(type) {
	case *timestamppb.Timestamp,
		timestamppb.Timestamp:
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case *durationpb.Duration,
		durationpb.Duration:
		return arrow.PrimitiveTypes.Int64, nil
	case protoreflect.Enum:
		return arrow.BinaryTypes.String, nil
	default:
		return nil, nil
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
