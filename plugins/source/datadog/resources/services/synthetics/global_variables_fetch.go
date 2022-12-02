package synthetics

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlobalVariables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.SyntheticsAPI.ListGlobalVariables(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetVariables()
	return nil
}
