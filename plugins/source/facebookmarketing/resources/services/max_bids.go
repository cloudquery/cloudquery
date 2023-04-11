package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func MaxBids() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_max_bids",
		Resolver:  fetchMaxBids,
		Transform: client.TransformWithStruct(&rest.MaxBid{}),
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
