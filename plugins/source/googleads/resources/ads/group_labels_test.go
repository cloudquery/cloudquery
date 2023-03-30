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

func TestGroupLabelsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	ad_group_label.resource_name,
	ad_group_label.ad_group,
	ad_group_label.label
FROM ad_group_label
WHERE ad_group_label.ad_group = "customers/123/adGroups/456"`

	require.Equal(t, expected,
		gaql.Query(new(resources.AdGroupLabel),
			&schema.Resource{Item: &resources.AdGroup{ResourceName: `customers/123/adGroups/456`}},
		),
	)
}

func testAdGroupLabels(t *testing.T) map[string][]*services.GoogleAdsRow {
	var label resources.AdGroupLabel
	require.NoError(t, faker.FakeObject(&label))
	row := &services.GoogleAdsRow{AdGroupLabel: &label}
	return map[string][]*services.GoogleAdsRow{"ad_group_label": {row}}
}
