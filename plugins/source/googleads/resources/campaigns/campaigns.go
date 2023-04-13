package campaigns

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

var campaignOptions = &gaql.Options{
	Skip: []string{"TargetingSetting"},
	Expand: []string{
		"LocalServicesCampaignSettings",
		"TravelCampaignSettings",
		"RealTimeBiddingSetting",
		"NetworkSettings",
		"HotelSetting",
		"DynamicSearchAdsSetting",
		"ShoppingSetting",
		"AudienceSetting",
		"GeoTargetTypeSetting",
		"LocalCampaignSetting",
		"AppCampaignSetting",
		"VanityPharma",
		"SelectiveOptimization",
		"OptimizationGoalSetting",
		"TrackingSetting",
		"PerformanceMaxUpgrade",
	},
}

func Campaigns() *schema.Table {
	return &schema.Table{
		Name:        "googleads_campaigns",
		Description: client.APIRef + "/Campaign",
		Multiplex:   client.MultiplexByCustomer,
		Transform: client.TransformWithStruct(new(resources.Campaign),
			transformers.WithPrimaryKeys("Id", "ResourceName"),
			transformers.WithSkipFields(campaignOptions.Skip...),
		),
		Columns:   schema.ColumnList{client.CustomerID},
		Resolver:  client.Fetcher((*services.GoogleAdsRow).GetCampaign, campaignOptions),
		Relations: schema.Tables{campaignCriteria(), campaignLabels()},
	}
}
