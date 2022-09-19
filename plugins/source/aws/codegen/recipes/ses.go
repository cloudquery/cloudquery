package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SESResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "templates",
			Struct:     &ses.Template{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveSesTemplateArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ses"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("email")`
	}
	return resources
}
