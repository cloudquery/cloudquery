package business_services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
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
				Type:     types.ExtensionTypes.JSON,
				Resolver: DependenciesResolver,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLUrl"),
			},
		},
	}
}
