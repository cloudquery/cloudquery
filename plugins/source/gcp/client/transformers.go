package client

import (
	"reflect"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/protobuf/types/known/structpb"
)

var toReplace = map[string]string{
	"c_d_n":   "cdn",
	"i_p_":    "ip_",
	"ipv_6":   "ipv6",
	"i_pv_4":  "ipv4",
	"i_pv4":   "ipv4",
	"oauth_2": "oauth2",
	"c_o_r_s": "cors",
	"r_p_o":   "rpo",
}

func replaceTransformer(field reflect.StructField) (string, error) {
	name, err := transformers.DefaultNameTransformer(field)
	if err != nil {
		return "", err
	}
	for k, v := range toReplace {
		name = strings.ReplaceAll(name, k, v)
	}
	return name, nil
}

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

func ignoreInTestsTransformer(field reflect.StructField) bool {
	nonMockables := []any{&structpb.Value{}, &structpb.Struct{}, &aiplatformpb.Model{}, &aiplatformpb.PipelineJob_RuntimeConfig{}}
	for _, nonMockable := range nonMockables {
		if field.Type == reflect.TypeOf(nonMockable) {
			return true
		}
	}

	return false
}

var options = []transformers.StructTransformerOption{
	transformers.WithNameTransformer(replaceTransformer),
	transformers.WithTypeTransformer(typeTransformer),
	transformers.WithResolverTransformer(resolverTransformer),
	transformers.WithIgnoreInTestsTransformer(ignoreInTestsTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}
