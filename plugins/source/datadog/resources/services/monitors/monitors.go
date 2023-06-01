package monitors

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:      "datadog_monitors",
		Resolver:  fetchMonitors,
		Multiplex: client.AccountMultiplex,
		Transform: client.TransformWithStruct(&datadogV1.Monitor{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
			{
				Name:     "deleted",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.NullableResolver("Deleted"),
			},
			{
				Name:     "priority",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: client.NullableResolver("Priority"),
			},
		},

		Relations: []*schema.Table{
			MonitorDowntimes(),
		},
	}
}

func fetchMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.MonitorsAPI.ListMonitors(ctx)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
