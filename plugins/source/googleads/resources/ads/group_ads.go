package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

var groupAdOptions = &gaql.Options{Expand: []string{"Ad", "PolicySummary"}}

func groupAds() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_group_ads",
		Description: client.APIRef + "/AdGroupAd",
		Transform: client.TransformWithStruct(new(resources.AdGroupAd),
			transformers.WithPrimaryKeys("AdGroup", "ResourceName"),
		),
		Columns:   schema.ColumnList{client.CustomerID, client.IDColumn("Ad.Id")},
		Resolver:  client.Fetcher((*services.GoogleAdsRow).GetAdGroupAd, groupAdOptions),
		Relations: schema.Tables{groupAdLabels()},
	}
}
