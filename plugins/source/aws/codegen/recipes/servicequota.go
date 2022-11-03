package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServiceQuotasResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "services",
			Struct:      &types.ServiceInfo{},
			SkipFields:  []string{"ServiceCode", "ServiceName"},
			Description: "https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceInfo.html",
			Multiplex:   `client.ServiceAccountRegionMultiplexer("servicequotas")`,
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "service_code",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ServiceCode")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					}, {
						Name:     "service_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ServiceName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"Quotas()",
			},
		},
		{
			SubService:  "quotas",
			Struct:      &types.ServiceQuota{},
			Description: "https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceQuota.html",
			SkipFields:  []string{"QuotaArn"},
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
	}
	return resources
}
