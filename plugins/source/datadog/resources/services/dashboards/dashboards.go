package dashboards

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Dashboards() *schema.Table {
	return &schema.Table{
		Name:      "datadog_dashboards",
		Resolver:  fetchDashboards,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.DashboardSummaryDefinition{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchDashboards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.DashboardsAPI.ListDashboards(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetDashboards()
	return nil
}
