package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ResolveProject(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("project_id", client.ProjectId)
}

func ResolveOrganization(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.OrgId)
}

func ResolveFolder(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.FolderId)
}

func AddGcpMetadata(ctx context.Context, c schema.ClientMeta, r *schema.Resource) error {
	return resolveLocation(ctx, c, r)
}

func ResolveResourceId(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	data, err := cast.ToStringE(funk.Get(r.Item, "Id", funk.WithAllowZero()))
	if err != nil {
		return err
	}
	return r.Set(c.Name, data)
}

func resolveLocation(_ context.Context, c schema.ClientMeta, r *schema.Resource) error {
	loc := r.Get("location")
	if loc != nil {
		return nil
	}
	name := r.Get("name")
	if name == nil {
		c.(*Client).Logger().Warn().Str("resource", fmt.Sprintf("%T", r.Item)).Msg("missing name for resource")
		return nil
	}
	nameStr, err := cast.ToStringE(name)
	if err != nil {
		return err
	}
	match := strings.Split(nameStr, "/")
	if len(match) < 3 {
		return nil
	}
	return r.Set("location", match[3])
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
