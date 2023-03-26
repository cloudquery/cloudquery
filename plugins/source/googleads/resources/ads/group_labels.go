package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func groupLabels() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_group_labels",
		Description: client.APIRef + "/AdGroupLabel",
		Transform: client.TransformWithStruct(new(resources.AdGroupLabel),
			transformers.WithPrimaryKeys("AdGroup", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetAdGroupLabel),
	}
}
