package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/thoas/go-funk"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AccountID)
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Region)
}

func ResolveAWSNamespace(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AutoscalingNamespace)
}

func ResolveAWSPartition(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Partition)
}

func ResolveWAFScope(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).WAFScope)
}

func ResolveLanguageCode(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.LanguageCode)
}

func ResolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return ResolveTagPath("Tags")(ctx, meta, r, c)
}

func ResolveTagPath(fieldPath string) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		val := funk.Get(r.Item, fieldPath, funk.WithAllowZero())
		if val == nil {
			return r.Set(c.Name, map[string]string{}) // can't have nil or the integration test will make a fuss
		}

		return r.Set(c.Name, TagsToMap(val))
	}
}

func ResolveObjectHash(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	hash, err := hashstructure.Hash(r.Item, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}
	return r.Set(c.Name, fmt.Sprint(hash))
}
