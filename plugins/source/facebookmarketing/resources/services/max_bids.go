package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func MaxBids() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_max_bids",
		Resolver:  fetchMaxBids,
		Transform: client.TransformWithStruct(&rest.MaxBid{}),
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Resolver:   client.ResolveAccountId,
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-campaign-group/max-bid/",
	}
}
