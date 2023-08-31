package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

// We can't use the default resolver because facebook's time format is not RFC3339.
// '2023-03-05T16:46:23+0200' instead of RFC3339 '2023-03-05T16:46:23+02:00' (missing ':').

func FacebookDatetimeResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		s := funk.Get(r.Item, path, funk.WithAllowZero()).(string)

		if s == "" {
			return r.Set(c.Name, nil)
		}

		t, err := time.Parse("2006-01-02T15:04:05Z0700", s)
		if err != nil {
			return err
		}

		return r.Set(c.Name, t)
	}
}

func ResolveAccountId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).FacebookClient.AdAccountId)
}
