package ads

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

func TestGroupCriterionLabelsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	ad_group_criterion_label.resource_name,
	ad_group_criterion_label.ad_group_criterion,
	ad_group_criterion_label.label
FROM ad_group_criterion_label
WHERE ad_group_criterion_label.ad_group_criterion = "customers/123/adGroupCriteria/456~789"`

	require.Equal(t, expected,
		gaql.Query(new(resources.AdGroupCriterionLabel),
			&schema.Resource{Item: &resources.AdGroupCriterion{ResourceName: `customers/123/adGroupCriteria/456~789`}},
		),
	)
}

func testAdGroupCriterionLabels(t *testing.T) client.MockedResponses {
	var label resources.AdGroupCriterionLabel
	require.NoError(t, faker.FakeObject(&label))
	return client.MockedResponses{"ad_group_criterion_label": {&services.GoogleAdsRow{AdGroupCriterionLabel: &label}}}
}
