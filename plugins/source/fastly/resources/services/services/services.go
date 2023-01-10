package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "fastly_services",
		Description: `https://developer.fastly.com/reference/api/services/service/`,
		Resolver:    fetchServices,
		Transform:   transformers.TransformWithStruct(&fastly.Service{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			ServiceVersions(),
		},
	}
}
