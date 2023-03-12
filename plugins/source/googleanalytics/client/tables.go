package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var _ source.GetTables = GetTables

func GetTables(_ context.Context, meta schema.ClientMeta) (schema.Tables, error) {
	c := meta.(*Client)

	res := make(schema.Tables, len(c.reports))
	for i, r := range c.reports {
		res[i] = r.table(c.PropertyID)
	}

	return res, nil
}
