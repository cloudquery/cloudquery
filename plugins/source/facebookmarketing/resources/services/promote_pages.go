package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PromotePages() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_promote_pages",
		Resolver:  fetchPromotePages,
		Transform: transformers.TransformWithStruct(&rest.PromotePage{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Resolver: client.ResolveAccountId,
				Type:     schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Description: "https://developers.facebook.com/docs/graph-api/reference/page#Reading",
	}
}
