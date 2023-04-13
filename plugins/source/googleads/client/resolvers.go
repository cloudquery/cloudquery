package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ResolveCustomerID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.CustomerID)
}

func ResolveProtoTimestamp(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(*timestamppb.Timestamp)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"*timestamppb.Timestamp\", have \"%T\"", data)
		}
		return resource.Set(c.Name, ts.AsTime())
	}
}

func ResolveProtoDuration(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		dur, ok := data.(*durationpb.Duration)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"*durationpb.Duration\", have \"%T\"", data)
		}
		return resource.Set(c.Name, dur.AsDuration())
	}
}

func ResolveProtoEnum(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		enum, ok := data.(protoreflect.Enum)
		if !ok {
			return fmt.Errorf("unexpected type, wanted \"protoreflect.Enum\", have \"%T\"", data)
		}
		return resource.Set(c.Name, protoimpl.X.EnumStringOf(enum.Descriptor(), enum.Number()))
	}
}
