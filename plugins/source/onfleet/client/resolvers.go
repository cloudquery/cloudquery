package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

// Resolves a unix-timestamp-with-millies field to a timestamp field
func ResolveTimestampField(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}

		var unixTimestampMillis int64

		unixTimestampMillisPtr, ok := data.(*int64)
		if ok {
			unixTimestampMillis = *unixTimestampMillisPtr
		} else {
			unixTimestampMillis, ok = data.(int64)
			if !ok {
				return fmt.Errorf("unexpected type in path %s, wanted int64 or *int64, have '%s'", path, data)
			}
		}

		if unixTimestampMillis == 0 {
			return nil
		}

		return resource.Set(c.Name, time.UnixMilli(unixTimestampMillis))
	}
}

// Sets the `organization_id` field to the OrganizationId in the Cloudquery Client
func ResolveOrganizationId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cqClient := meta.(*Client)
	return r.Set(c.Name, cqClient.OrganizationId)
}
