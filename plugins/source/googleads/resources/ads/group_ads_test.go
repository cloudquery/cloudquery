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

func TestGroupAdsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	ad_group_ad.resource_name,
	ad_group_ad.status,
	ad_group_ad.ad_group,
	ad_group_ad.ad.resource_name,
	ad_group_ad.ad.id,
	ad_group_ad.ad.final_urls,
	ad_group_ad.ad.final_app_urls,
	ad_group_ad.ad.final_mobile_urls,
	ad_group_ad.ad.tracking_url_template,
	ad_group_ad.ad.final_url_suffix,
	ad_group_ad.ad.url_custom_parameters,
	ad_group_ad.ad.display_url,
	ad_group_ad.ad.type,
	ad_group_ad.ad.added_by_google_ads,
	ad_group_ad.ad.device_preference,
	ad_group_ad.ad.url_collections,
	ad_group_ad.ad.name,
	ad_group_ad.ad.system_managed_resource_source,
	ad_group_ad.policy_summary.policy_topic_entries,
	ad_group_ad.policy_summary.review_status,
	ad_group_ad.policy_summary.approval_status,
	ad_group_ad.ad_strength,
	ad_group_ad.action_items,
	ad_group_ad.labels
FROM ad_group_ad
WHERE ad_group_ad.ad_group = "customers/123/adGroups/456"`

	require.Equal(t, expected,
		gaql.Query(new(resources.AdGroupAd),
			&schema.Resource{Item: &resources.AdGroup{ResourceName: `customers/123/adGroups/456`}},
			groupAdOptions,
		),
	)
}

func testAdGroupAds(t *testing.T) client.MockedResponses {
	var ad resources.AdGroupAd
	require.NoError(t, faker.FakeObject(&ad))
	responses := client.MockedResponses{"ad_group_ad": {&services.GoogleAdsRow{AdGroupAd: &ad}}}
	maps.Copy(responses, testAdGroupAdLabels(t))
	return responses
}
