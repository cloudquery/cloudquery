package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Synthetics() []*Resource {
	resources := []*Resource{
		{
			SubService: "synthetics",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV1.SyntheticsAPITest),
			SkipFields: []string{"PublicId"},
			ExtraColumns: append(defaultAccountColumnsPK, codegen.ColumnDefinition{
				Name:     "public_id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("PublicId")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
		},
		{
			SubService: "global_variables",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV1.SyntheticsGlobalVariable),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumns, codegen.ColumnDefinition{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Id")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "synthetics"
	}
	return resources
}
