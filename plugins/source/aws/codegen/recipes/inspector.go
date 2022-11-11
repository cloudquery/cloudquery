package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InspectorResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "findings",
			Struct:      &types.Finding{},
			Description: "https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html",
			SkipFields:  []string{"Arn", "Attributes", "UserAttributes"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "attributes",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTagField("Attributes")`,
					},
					{
						Name:     "user_attributes",
						Type:     schema.TypeJSON,
						Resolver: `client.ResolveTagField("UserAttributes")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "inspector"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("inspector")`
	}
	return resources
}
