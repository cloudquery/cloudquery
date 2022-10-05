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
			SkipFields:          []string{"DomainId"},
			PreResourceResolver: "getDomain",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "tags",
					Type:     schema.TypeJSON,
					Resolver: `resolveElasticsearchDomainTags`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("DomainId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elasticsearch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("es")`
	}
	return resources
}
