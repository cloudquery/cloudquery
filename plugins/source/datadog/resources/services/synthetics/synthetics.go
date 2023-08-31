package synthetics

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "datadog_synthetics",
		Resolver:  fetchSynthetics,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.SyntheticsTestDetails{}, transformers.WithPrimaryKeys("PublicId")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchSynthetics(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.SyntheticsAPI.ListTests(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetTests()
	return nil
}
