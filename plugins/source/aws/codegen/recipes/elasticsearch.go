package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ElasticsearchResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "domains",
			Struct:              &types.ElasticsearchDomainStatus{},
			PreResourceResolver: "getDomain",
			PKColumns:           []string{"arn"},
			ExtraColumns: append(defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveTags`,
				},
			),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elasticsearch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("es")`
	}
	return resources
}
