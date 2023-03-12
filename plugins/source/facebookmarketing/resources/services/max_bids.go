package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MaxBids() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_max_bids",
		Resolver:  fetchMaxBids,
		Transform: transformers.TransformWithStruct(&rest.MaxBid{}, client.TransformerOptions()...),
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
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group/max-bid/",
	}
}
