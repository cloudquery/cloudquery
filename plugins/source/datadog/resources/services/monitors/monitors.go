package monitors

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Monitors() *schema.Table {
	return &schema.Table{
		Name:      "datadog_monitors",
		Resolver:  fetchMonitors,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.Monitor{}),
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "deleted",
				Type:     schema.TypeTimestamp,
				Resolver: client.NullableResolver("Deleted"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeInt,
				Resolver: client.NullableResolver("Priority"),
			},
		},

		Relations: []*schema.Table{
			MonitorDowntimes(),
		},
	}
}
