package synthetics

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "datadog_synthetics",
		Resolver:  fetchSynthetics,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.SyntheticsTestDetails{}, transformers.WithPrimaryKeys("PublicId")),
		Columns: []schema.Column{
			client.AccountNameColumn,
		},
	}
}

func fetchSynthetics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.SyntheticsAPI.ListTests(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetTests()
	return nil
}
