package campaigns

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestCampaignsQuery(t *testing.T) {
	// we need to have this test to check the resulting query using API validator
	const expected = `SELECT
	campaign.resource_name,
	campaign.id,
	campaign.name,
	campaign.primary_status,
	campaign.primary_status_reasons,
	campaign.status,
	campaign.serving_status,
	campaign.bidding_strategy_system_status,
	campaign.ad_serving_optimization_status,
	campaign.advertising_channel_type,
	campaign.advertising_channel_sub_type,
	campaign.tracking_url_template,
	campaign.url_custom_parameters,
	campaign.local_services_campaign_settings.category_bids,
	campaign.travel_campaign_settings.travel_account_id,
	campaign.real_time_bidding_setting.opt_in,
	campaign.network_settings.target_google_search,
	campaign.network_settings.target_search_network,
	campaign.network_settings.target_content_network,
	campaign.network_settings.target_partner_search_network,
	campaign.hotel_setting.hotel_center_id,
	campaign.dynamic_search_ads_setting.domain_name,
	campaign.dynamic_search_ads_setting.language_code,
	campaign.dynamic_search_ads_setting.use_supplied_urls_only,
	campaign.dynamic_search_ads_setting.feeds,
	campaign.shopping_setting.merchant_id,
	campaign.shopping_setting.sales_country,
	campaign.shopping_setting.feed_label,
	campaign.shopping_setting.campaign_priority,
	campaign.shopping_setting.enable_local,
	campaign.shopping_setting.use_vehicle_inventory,
	campaign.audience_setting.use_audience_grouped,
	campaign.geo_target_type_setting.positive_geo_target_type,
	campaign.geo_target_type_setting.negative_geo_target_type,
	campaign.local_campaign_setting.location_source_type,
	campaign.app_campaign_setting.bidding_strategy_goal_type,
	campaign.app_campaign_setting.app_id,
	campaign.app_campaign_setting.app_store,
	campaign.labels,
	campaign.experiment_type,
	campaign.base_campaign,
	campaign.campaign_budget,
	campaign.bidding_strategy_type,
	campaign.accessible_bidding_strategy,
	campaign.start_date,
	campaign.campaign_group,
	campaign.end_date,
	campaign.final_url_suffix,
	campaign.frequency_caps,
	campaign.video_brand_safety_suitability,
	campaign.vanity_pharma.vanity_pharma_display_url_mode,
	campaign.vanity_pharma.vanity_pharma_text,
	campaign.selective_optimization.conversion_actions,
	campaign.optimization_goal_setting.optimization_goal_types,
	campaign.tracking_setting.tracking_url,
	campaign.payment_mode,
	campaign.optimization_score,
	campaign.excluded_parent_asset_field_types,
	campaign.excluded_parent_asset_set_types,
	campaign.url_expansion_opt_out,
	campaign.performance_max_upgrade.performance_max_campaign,
	campaign.performance_max_upgrade.pre_upgrade_campaign,
	campaign.performance_max_upgrade.status,
	campaign.hotel_property_asset_set
FROM campaign`

	require.Equal(t, expected, gaql.Query(new(resources.Campaign), nil, campaignOptions))
}

func testCampaigns(t *testing.T) client.MockedResponses {
	var campaign resources.Campaign
	require.NoError(t, faker.FakeObject(&campaign))
	responses := client.MockedResponses{"campaign": {&services.GoogleAdsRow{Campaign: &campaign}}}
	maps.Copy(responses, testCampaignCriteria(t))
	maps.Copy(responses, testCampaignLabels(t))
	return responses
}

func TestCampaigns(t *testing.T) {
	client.MockTestHelper(t, Campaigns(), testCampaigns(t))
}
