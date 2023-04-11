package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func groupCriterionLabels() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_group_criterion_labels",
		Description: client.APIRef + "/AdGroupCriterionLabel",
		Transform: client.TransformWithStruct(new(resources.AdGroupCriterionLabel),
			transformers.WithPrimaryKeys("AdGroupCriterion", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetAdGroupCriterionLabel),
	}
}
