package monitors

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func downtimes() *schema.Table {
	return &schema.Table{
		Name:      "datadog_monitor_downtimes",
		Transform: client.TransformWithStruct(&datadogV1.Downtime{}, transformers.WithPrimaryKeys("Id")),
		Resolver:  fetchMonitorDowntimes,
		Columns: schema.ColumnList{
			client.AccountNameColumn,
			{
				Name:       "monitor_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchMonitorDowntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(datadogV1.Monitor)
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.DowntimesAPI.ListMonitorDowntimes(ctx, *p.Id)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
