package ads

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

var groupCriterionOptions = &gaql.Options{Expand: []string{"QualityInfo", "PositionEstimates"}}

func groupCriteria() *schema.Table {
	return &schema.Table{
		Name:        "googleads_ad_group_criteria",
		Description: client.APIRef + "/AdGroupCriterion",
		Transform: client.TransformWithStruct(new(resources.AdGroupCriterion),
			transformers.WithPrimaryKeys("AdGroup", "ResourceName"),
		),
		Columns:   schema.ColumnList{client.CustomerID, client.IDColumn("CriterionId")},
		Resolver:  client.Fetcher((*services.GoogleAdsRow).GetAdGroupCriterion, groupCriterionOptions),
		Relations: schema.Tables{groupCriterionLabels()},
	}
}
