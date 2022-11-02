package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServiceQuotasResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "quotas",
			Struct:     &types.ServiceQuota{},
			SkipFields: []string{"QuotaArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("QuotaArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "servicequotas"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("servicequotas")`
	}
	return resources
}
