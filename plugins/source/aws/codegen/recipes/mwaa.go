package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func MWAAResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "environments",
			Description:         "https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html",
			Struct:              &types.Environment{},
			SkipFields:          []string{"Arn"},
			PreResourceResolver: "getEnvironment",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "mwaa"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("airflow")`
	}
	return resources
}
