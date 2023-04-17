package campaigns

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
)

func TestCampaignsCriteriaQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	campaign_criterion.resource_name,
	campaign_criterion.campaign,
	campaign_criterion.criterion_id,
	campaign_criterion.display_name,
	campaign_criterion.bid_modifier,
	campaign_criterion.negative,
	campaign_criterion.type,
	campaign_criterion.status
FROM campaign_criterion
WHERE campaign_criterion.campaign = "customers/123/campaigns/456"`

	require.Equal(t, expected,
		gaql.Query(new(resources.CampaignCriterion),
			&schema.Resource{Item: &resources.Campaign{ResourceName: `customers/123/campaigns/456`}},
		),
	)
}

func testCampaignCriteria(t *testing.T) client.MockedResponses {
	var criterion resources.CampaignCriterion
	require.NoError(t, faker.FakeObject(&criterion))
	return client.MockedResponses{"campaign_criterion": {&services.GoogleAdsRow{CampaignCriterion: &criterion}}}
}
