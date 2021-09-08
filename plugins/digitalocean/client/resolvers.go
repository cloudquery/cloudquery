package client

import (
	"context"
	"strings"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
)

func ResolveResourceTypeFromUrn(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	value := funk.Get(r.Item, "URN", funk.WithAllowZero())
	urn := cast.ToString(value)
	if urn == "" {
		return nil
	}
	parts := strings.Split(urn, ":")
	return r.Set(c.Name, parts[1])
}

func ResolveResourceIdFromUrn(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	value := funk.Get(r.Item, "URN", funk.WithAllowZero())
	urn := cast.ToString(value)
	if urn == "" {
		return nil
	}
	parts := strings.Split(urn, ":")
	if len(parts) < 2 {
		return nil
	}
	return r.Set(c.Name, parts[2])
}
