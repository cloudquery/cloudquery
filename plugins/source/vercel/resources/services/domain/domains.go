package domain

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:      "vercel_domains",
		Resolver:  fetchDomains,
		Transform: transformers.TransformWithStruct(&vercel.Domain{}, client.SharedTransformers()...),
		Multiplex: client.TeamMultiplex,
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
			DomainRecords(),
		},
	}
}
