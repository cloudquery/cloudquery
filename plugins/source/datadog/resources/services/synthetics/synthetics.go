package synthetics

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Synthetics() *schema.Table {
	return &schema.Table{
		Name:      "datadog_synthetics",
		Resolver:  fetchSynthetics,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.SyntheticsTestDetails{}),
		Columns: []schema.Column{
			{
				Name:     "account_name",
				Type:     schema.TypeString,
				Resolver: client.ResolveAccountName,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "public_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
