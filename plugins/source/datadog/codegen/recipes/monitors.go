package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Monitors() []*Resource {
	resources := []*Resource{
		{
			SubService: "monitors",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV1.Monitor),
			SkipFields: []string{"Id", "Deleted", "Priority"},
			ExtraColumns: append(defaultAccountColumnsPK,
				codegen.ColumnDefinition{
					Name:     "id",
					Type:     schema.TypeInt,
					Resolver: `schema.PathResolver("Id")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				codegen.ColumnDefinition{
					Name:     "deleted",
					Type:     schema.TypeTimestamp,
					Resolver: `client.NullableResolver("Deleted")`,
				},
				codegen.ColumnDefinition{
					Name:     "priority",
					Type:     schema.TypeInt,
					Resolver: `client.NullableResolver("Priority")`,
				},
			),
			Relations: []string{"MonitorDowntimes()"},
		},
		{
			SubService:   "monitor_downtimes",
			Struct:       new(datadogV1.Downtime),
			ExtraColumns: defaultAccountColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "monitors"
	}
	return resources
}
