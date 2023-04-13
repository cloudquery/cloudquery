package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func groupAdLabels() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_group_ad_labels",
		Description: client.APIRef + "/AdGroupAdLabel",
		Transform: client.TransformWithStruct(new(resources.AdGroupAdLabel),
			transformers.WithPrimaryKeys("AdGroupAd", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetAdGroupAdLabel),
	}
}
