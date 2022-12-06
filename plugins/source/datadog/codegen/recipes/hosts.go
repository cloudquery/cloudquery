package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Hosts() []*Resource {
	resources := []*Resource{
		{
			SubService: "hosts",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV1.Host),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumnsPK,
				codegen.ColumnDefinition{
					Name:     "id",
					Type:     schema.TypeInt,
					Resolver: `schema.PathResolver("Id")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "hosts"
	}
	return resources
}
