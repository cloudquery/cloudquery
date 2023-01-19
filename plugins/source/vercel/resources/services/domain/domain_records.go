package domain

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DomainRecords() *schema.Table {
	return &schema.Table{
		Name:          "vercel_domain_records",
		Resolver:      fetchDomainRecords,
		Transform:     transformers.TransformWithStruct(&vercel.DomainRecord{}, client.SharedTransformers()...),
		Multiplex:     client.TeamMultiplex,
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
