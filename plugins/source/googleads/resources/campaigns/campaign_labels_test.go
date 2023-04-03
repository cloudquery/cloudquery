package campaigns

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
)

func TestCampaignLabelsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	campaign_label.resource_name,
	campaign_label.campaign,
	campaign_label.label
FROM campaign_label
WHERE campaign_label.campaign = "customers/123/campaigns/456"`

	require.Equal(t, expected,
		gaql.Query(new(resources.CampaignLabel),
			&schema.Resource{Item: &resources.Campaign{ResourceName: `customers/123/campaigns/456`}},
		),
	)
}

func testCampaignLabels(t *testing.T) map[string][]*services.GoogleAdsRow {
	var label resources.CampaignLabel
	require.NoError(t, faker.FakeObject(&label))
	row := &services.GoogleAdsRow{CampaignLabel: &label}
	return map[string][]*services.GoogleAdsRow{"campaign_label": {row}}
}
