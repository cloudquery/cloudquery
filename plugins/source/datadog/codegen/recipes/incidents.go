package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Incidents() []*Resource {
	resources := []*Resource{
		{
			SubService: "incidents",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV2.IncidentResponseData),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumnsPK,
				codegen.ColumnDefinition{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Id")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			),
			Relations: []string{"IncidentAttachments()"},
		},
		{
			SubService:   "incident_attachments",
			Struct:       new(datadogV2.IncidentAttachmentData),
			ExtraColumns: defaultAccountColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "incidents"
	}
	return resources
}
