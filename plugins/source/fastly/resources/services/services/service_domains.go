package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
			{
				Name:       "service_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ServiceID"),
				PrimaryKey: true,
			},
			{
				Name:       "service_version",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ServiceVersion"),
				PrimaryKey: true,
			},
		},
	}
}
