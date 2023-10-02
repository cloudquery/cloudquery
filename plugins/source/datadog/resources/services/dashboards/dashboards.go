package dashboards

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Dashboards() *schema.Table {
	return &schema.Table{
		Name:      "datadog_dashboards",
		Resolver:  fetchDashboards,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.DashboardSummaryDefinition{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchDashboards(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, cancel := c.DDServices.DashboardsAPI.ListDashboardsWithPagination(ctx)
	return client.ConsumePaginatedResponse(resp, cancel, res)
}
