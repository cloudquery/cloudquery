package business_services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BusinessServices() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_business_services",
		Description: `https://developer.pagerduty.com/api-reference/e67570b9d0e3d-list-business-services`,
		Resolver:    fetchBusinessServices,
		Transform:   transformers.TransformWithStruct(&pagerduty.BusinessService{}, transformers.WithSkipFields("HTMLUrl")),
		Columns: []schema.Column{
			{
				Name:     "dependencies",
				Type:     schema.TypeJSON,
				Resolver: DependenciesResolver,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLUrl"),
			},
		},
	}
}
