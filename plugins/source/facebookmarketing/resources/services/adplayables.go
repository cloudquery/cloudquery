package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Adplayables() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_adplayables",
		Resolver:  fetchAdplayables,
		Transform: transformers.TransformWithStruct(&rest.Adplayable{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
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
