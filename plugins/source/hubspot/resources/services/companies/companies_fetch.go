package companies

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCompanies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	var after string
	for {
		req := c.Companies.BasicApi.GetPage(c.AuthContext(ctx)).Limit(100)
		if len(after) > 0 {
			req.After(after)
		}
		out, _, err := req.Execute()
		if err != nil {
			return err
		}

		res <- out.Results

		page := out.GetPaging()
		next := page.GetNext()
		after = next.GetAfter()
		if after == "" {
			break
		}
	}

	return nil
}
