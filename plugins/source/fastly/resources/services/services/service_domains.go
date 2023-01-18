package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func ServiceDomains() *schema.Table {
	return &schema.Table{
		Name:        "fastly_service_domains",
		Description: `https://developer.fastly.com/reference/api/services/domain/`,
		Resolver:    fetchServiceDomains,
		Transform:   transformers.TransformWithStruct(&fastly.Domain{}),
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServiceVersion"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
