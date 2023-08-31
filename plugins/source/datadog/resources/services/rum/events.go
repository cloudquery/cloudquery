package rum

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Events() *schema.Table {
	return &schema.Table{
		Name:      "datadog_rum_events",
		Title:     "Datadog Real User Monitoring (RUM) Events",
		Resolver:  fetchEvents,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(new(datadogV2.RUMEvent), transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
	}
}

func fetchEvents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	resp, cancel := c.DDServices.RumAPI.ListRUMEventsWithPagination(c.BuildContextV2(ctx))
	return client.ConsumePaginatedResponse(resp, cancel, res)
}
