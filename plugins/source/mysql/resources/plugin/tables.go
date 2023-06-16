package plugin

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func getDynamicTables(ctx context.Context, c schema.ClientMeta) (schema.Tables, error) {
	cl := c.(*client.Client)
	return cl.Tables, nil
}
