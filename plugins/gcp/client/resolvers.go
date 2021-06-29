package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/thoas/go-funk"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
)

func ResolveProject(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("project_id", client.ProjectId)
}

func resolveLocation(_ context.Context, c schema.ClientMeta, r *schema.Resource) error {
	loc := r.Get("location")
	if loc != nil {
		return nil
	}
	name := r.Get("name")
	if name == nil {
		c.Logger().Warn("missing name for resource", "resource", fmt.Sprintf("%T", r.Item))
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

func parseISODate(d string) (*time.Time, error) {
	if d == "" {
		return nil, nil
	}
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation(time.RFC3339, d, location)
	if err != nil {
		return nil, err
	}
	return &date, err
}

func ISODateResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		data, err := cast.ToStringE(funk.Get(r.Item, path, funk.WithAllowZero()))
		if err != nil {
			return err
		}
		date, err := parseISODate(data)
		if err != nil {
			return err
		}
		return r.Set(c.Name, date)
	}
}
