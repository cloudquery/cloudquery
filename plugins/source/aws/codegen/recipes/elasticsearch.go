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
			Struct:              new(types.ElasticsearchDomainStatus),
			Description:         "https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_DomainStatus.html",
			PreResourceResolver: "getDomain",
			PKColumns:           []string{"arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				codegen.ColumnDefinition{
					Name:     "authorized_principals",
					Type:     schema.TypeJSON,
					Resolver: `resolveAuthorizedPrincipals`,
				},
				codegen.ColumnDefinition{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveDomainTags`,
				},
			),
		},
		{
			SubService:      "packages",
			Struct:          new(types.PackageDetails),
			Description:     "https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_PackageDetails.html",
			NameTransformer: CreateReplaceTransformer(map[string]string{"package_id": "id"}),
			PKColumns:       []string{"id"},
			ExtraColumns:    defaultRegionalColumnsPK,
		},
		{
			SubService:  "versions",
			Struct:      new(struct{}),
			Description: "https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_ListVersions.html",
			PKColumns:   []string{"version"},
			ExtraColumns: append(defaultRegionalColumnsPK,
				codegen.ColumnDefinition{
					Name:     "version",
					Type:     schema.TypeString,
					Resolver: `resolveVersion`,
				},
				codegen.ColumnDefinition{
					Name:     "instance_types",
					Type:     schema.TypeJSON,
					Resolver: `resolveInstanceTypes`,
				},
			),
		},
		{
			SubService:      "vpc_endpoints",
			Struct:          new(types.VpcEndpoint),
			Description:     "https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_VpcEndpoint.html",
			NameTransformer: CreateReplaceTransformer(map[string]string{"vpc_endpoint_id": "id"}),
			PKColumns:       []string{"id"},
			ExtraColumns:    defaultRegionalColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elasticsearch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("es")`
	}
	return resources
}
