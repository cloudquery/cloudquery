package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func BroadTargetingCategoriess() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_broad_targeting_categories",
		Resolver:  fetchBroadTargetingCategories,
		Transform: transformers.TransformWithStruct(&rest.BroadTargetingCategories{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
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
	}
}
