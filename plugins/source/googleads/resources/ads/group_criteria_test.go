package ads

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
)

func TestGroupCriteriaQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	ad_group_criterion.resource_name,
	ad_group_criterion.criterion_id,
	ad_group_criterion.display_name,
	ad_group_criterion.status,
	ad_group_criterion.quality_info.quality_score,
	ad_group_criterion.quality_info.creative_quality_score,
	ad_group_criterion.quality_info.post_click_quality_score,
	ad_group_criterion.quality_info.search_predicted_ctr,
	ad_group_criterion.ad_group,
	ad_group_criterion.type,
	ad_group_criterion.negative,
	ad_group_criterion.system_serving_status,
	ad_group_criterion.approval_status,
	ad_group_criterion.disapproval_reasons,
	ad_group_criterion.labels,
	ad_group_criterion.bid_modifier,
	ad_group_criterion.cpc_bid_micros,
	ad_group_criterion.cpm_bid_micros,
	ad_group_criterion.cpv_bid_micros,
	ad_group_criterion.percent_cpc_bid_micros,
	ad_group_criterion.effective_cpc_bid_micros,
	ad_group_criterion.effective_cpm_bid_micros,
	ad_group_criterion.effective_cpv_bid_micros,
	ad_group_criterion.effective_percent_cpc_bid_micros,
	ad_group_criterion.effective_cpc_bid_source,
	ad_group_criterion.effective_cpm_bid_source,
	ad_group_criterion.effective_cpv_bid_source,
	ad_group_criterion.effective_percent_cpc_bid_source,
	ad_group_criterion.position_estimates.first_page_cpc_micros,
	ad_group_criterion.position_estimates.first_position_cpc_micros,
	ad_group_criterion.position_estimates.top_of_page_cpc_micros,
	ad_group_criterion.position_estimates.estimated_add_clicks_at_first_position_cpc,
	ad_group_criterion.position_estimates.estimated_add_cost_at_first_position_cpc,
	ad_group_criterion.final_urls,
	ad_group_criterion.final_mobile_urls,
	ad_group_criterion.final_url_suffix,
	ad_group_criterion.tracking_url_template,
	ad_group_criterion.url_custom_parameters
FROM ad_group_criterion
WHERE ad_group_criterion.ad_group = "customers/123/adGroups/456"`

	require.Equal(t, expected,
		gaql.Query(new(resources.AdGroupCriterion),
			&schema.Resource{Item: &resources.AdGroup{ResourceName: `customers/123/adGroups/456`}},
			groupCriterionOptions,
		),
	)
}

func testAdGroupCriteria(t *testing.T) client.MockedResponses {
	var criterion resources.AdGroupCriterion
	require.NoError(t, faker.FakeObject(&criterion))
	responses := client.MockedResponses{"ad_group_criterion": {&services.GoogleAdsRow{AdGroupCriterion: &criterion}}}
	maps.Copy(responses, testAdGroupCriterionLabels(t))
	return responses
}
