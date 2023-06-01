package incidents

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Incidents() *schema.Table {
	return &schema.Table{
		Name:      "datadog_incidents",
		Resolver:  fetchIncidents,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV2.IncidentResponseData{}, transformers.WithPrimaryKeys("Id")),
		Columns: []schema.Column{
			client.AccountNameColumn,
		},

		Relations: []*schema.Table{
			IncidentAttachments(),
		},
	}
}

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV2(ctx)
	resp, _, err := c.DDServices.IncidentsAPI.ListIncidents(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetData()
	return nil
}
