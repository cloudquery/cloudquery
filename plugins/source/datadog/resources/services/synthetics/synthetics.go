package synthetics

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "datadog_synthetics",
		Resolver:  fetchSynthetics,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.SyntheticsTestDetails{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "public_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PublicId"),
				PrimaryKey: true,
			},
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
