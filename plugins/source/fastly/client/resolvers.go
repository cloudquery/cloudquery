package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

func UnixTimeResolver(fieldName string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
		t := funk.Get(r.Item, fieldName, funk.WithAllowZero())
		switch v := t.(type) {
		case int64:
			return r.Set(col.Name, time.Unix(v, 0))
		case int:
			return r.Set(col.Name, time.Unix(int64(v), 0))
		case uint64:
			return r.Set(col.Name, time.Unix(int64(v), 0))
		}
		panic("unknown type for UnixTimeResolver")
	}
}
