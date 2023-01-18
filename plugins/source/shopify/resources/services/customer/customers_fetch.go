package customer

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCustomers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	cursor := ""
	for {
		ret, cur, err := cl.Services.GetCustomers(ctx, cursor)
		if err != nil {
			return err
		}
		res <- ret.Customers

		if len(ret.Customers) < ret.PageSize || cur == "" {
			return nil
		}

		cursor = cur
	}
}
