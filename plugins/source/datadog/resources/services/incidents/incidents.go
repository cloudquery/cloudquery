package incidents

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:      "datadog_incidents",
		Resolver:  fetchIncidents,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV2.IncidentResponseData{}, transformers.WithPrimaryKeys("Id")),
		Columns:   schema.ColumnList{client.AccountNameColumn},
		Relations: schema.Tables{attachments()},
	}
}

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, cancel := c.DDServices.IncidentsAPI.ListIncidentsWithPagination(ctx)
	return client.ConsumePaginatedResponse(resp, cancel, res)
}
