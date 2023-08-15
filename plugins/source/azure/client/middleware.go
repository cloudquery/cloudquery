package client

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func LowercaseIDResolver(_ context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	if !meta.(*Client).pluginSpec.NormalizeIDs {
		return nil
	}

	if resource.Table.Columns.Index("id") == -1 {
		return nil
	}

	id := resource.Get("id")
	if arrow.TypeEqual(id.DataType(), arrow.BinaryTypes.String) {
		return id.Set(strings.ToLower(id.String()))
	}

	return nil
}

func ChainRowResolvers(next ...schema.RowResolver) schema.RowResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
		for i := range next {
			if next[i] == nil {
				continue
			}
			if err := next[i](ctx, meta, resource); err != nil {
				return err
			}
		}
		return nil
	}
}
