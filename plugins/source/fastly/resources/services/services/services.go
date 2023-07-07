package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			ServiceVersions(),
		},
	}
}
