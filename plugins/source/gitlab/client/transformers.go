package client

import (
	"net"
	"reflect"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/xanzy/go-gitlab"
)

var options = []transformers.StructTransformerOption{
	transformers.WithUnwrapAllEmbeddedStructs(),
	transformers.WithUnwrapStructFields("Spec", "Status"),
	transformers.WithTypeTransformer(typeTransformer),
}

func TransformWithStruct(t any, opts ...transformers.StructTransformerOption) schema.Transform {
	return transformers.TransformWithStruct(t, append(options, opts...)...)
}

func typeTransformer(field reflect.StructField) (arrow.DataType, error) {
	timestamp := gitlab.ISOTime{}
	ip := net.IP{}
	switch field.Type {
	case reflect.TypeOf(timestamp), reflect.TypeOf(&timestamp):
		return arrow.FixedWidthTypes.Timestamp_us, nil
	case reflect.TypeOf(ip), reflect.TypeOf(&ip):
		return types.ExtensionTypes.Inet, nil
	default:
		return nil, nil
	}
}
