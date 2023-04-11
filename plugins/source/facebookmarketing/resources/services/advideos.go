package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Advideos() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_advideos",
		Resolver:  fetchAdvideos,
		Transform: client.TransformWithStruct(&rest.Advideo{}, transformers.WithPrimaryKeys("Id")),
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
		Description: "https://developers.facebook.com/docs/graph-api/reference/video#Reading",
	}
}
