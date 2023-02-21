package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func UnixTimeResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		s := funk.Get(r.Item, path, funk.WithAllowZero()).(int64)
		if s == 0 {
			return r.Set(c.Name, nil)
		}

		sSecs := s / 1000
		sNanos := (s % 1000) * 1000000

		t := time.Unix(sSecs, sNanos)
		return r.Set(c.Name, t)
	}
}
