package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AppSync() []*Resource {
	resources := []*Resource{
		{
			SubService:  "graphql_apis",
			Struct:      &types.GraphqlApi{},
			Description: "https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html",
			SkipFields:  []string{"Arn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("appsync")`,
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

	for _, r := range resources {
		r.Service = "appsync"
	}
	return resources
}
