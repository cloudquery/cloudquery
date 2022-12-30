package companies

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCompanies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	out, _, err := c.Companies.Get().BatchApi.BatchRead(ctx).Execute()
	if err != nil {
		return err
	}

	res <- out.Results

	return nil
}
