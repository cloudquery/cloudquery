package ads

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

func TestGroupsQuery(t *testing.T) {
	// we need to have this test to check the resulting query using API validator
	const expected = `SELECT
	ad_group.resource_name,
	ad_group.id,
	ad_group.name,
	ad_group.status,
	ad_group.type,
	ad_group.ad_rotation_mode,
	ad_group.base_ad_group,
	ad_group.tracking_url_template,
	ad_group.url_custom_parameters,
	ad_group.campaign,
	ad_group.cpc_bid_micros,
	ad_group.effective_cpc_bid_micros,
	ad_group.cpm_bid_micros,
	ad_group.target_cpa_micros,
	ad_group.cpv_bid_micros,
	ad_group.target_cpm_micros,
	ad_group.target_roas,
	ad_group.percent_cpc_bid_micros,
	ad_group.optimized_targeting_enabled,
	ad_group.display_custom_bid_dimension,
	ad_group.final_url_suffix,
	ad_group.audience_setting.use_audience_grouped,
	ad_group.effective_target_cpa_micros,
	ad_group.effective_target_cpa_source,
	ad_group.effective_target_roas,
	ad_group.effective_target_roas_source,
	ad_group.labels,
	ad_group.excluded_parent_asset_field_types,
	ad_group.excluded_parent_asset_set_types
FROM ad_group`

	require.Equal(t, expected, gaql.Query(new(resources.AdGroup), nil, groupOptions))
}

func testAdGroups(t *testing.T) client.MockedResponses {
	var adGroup resources.AdGroup
	require.NoError(t, faker.FakeObject(&adGroup))
	responses := client.MockedResponses{"ad_group": {&services.GoogleAdsRow{AdGroup: &adGroup}}}
	maps.Copy(responses, testAdGroupLabels(t))
	maps.Copy(responses, testAdGroupCriteria(t))
	maps.Copy(responses, testAdGroupAds(t))
	return responses
}

func TestAdGroups(t *testing.T) {
	client.MockTestHelper(t, Groups(), testAdGroups(t))
}
