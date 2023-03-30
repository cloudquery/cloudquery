package ads

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
)

func TestGroupAdLabelsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	ad_group_ad_label.resource_name,
	ad_group_ad_label.ad_group_ad,
	ad_group_ad_label.label
FROM ad_group_ad_label
WHERE ad_group_ad_label.ad_group_ad = "customers/123/adGroupAds/456~789"`

	require.Equal(t, expected,
		gaql.Query(new(resources.AdGroupAdLabel),
			&schema.Resource{Item: &resources.AdGroupAd{ResourceName: `customers/123/adGroupAds/456~789`}},
		),
	)
}

func testAdGroupAdLabels(t *testing.T) map[string][]*services.GoogleAdsRow {
	var label resources.AdGroupAdLabel
	require.NoError(t, faker.FakeObject(&label))
	row := &services.GoogleAdsRow{AdGroupAdLabel: &label}
	return map[string][]*services.GoogleAdsRow{"ad_group_ad_label": {row}}
}
