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

func testAdGroupCriterionLabels(t *testing.T) map[string][]*services.GoogleAdsRow {
	var label resources.AdGroupCriterionLabel
	require.NoError(t, faker.FakeObject(&label))
	row := &services.GoogleAdsRow{AdGroupCriterionLabel: &label}
	return map[string][]*services.GoogleAdsRow{"ad_group_criterion_label": {row}}
}
