package customers

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
)

func TestCustomerLabelsQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	const expected = `SELECT
	customer_label.resource_name,
	customer_label.customer,
	customer_label.label
FROM customer_label
WHERE customer_label.customer = "customers/123"`

	require.Equal(t, expected,
		gaql.Query(new(resources.CustomerLabel),
			&schema.Resource{Item: &resources.Customer{ResourceName: `customers/123`}},
		),
	)
}

func testCustomerLabels(t *testing.T) client.MockedResponses {
	var label resources.CustomerLabel
	require.NoError(t, faker.FakeObject(&label))
	return client.MockedResponses{"customer_label": {&services.GoogleAdsRow{CustomerLabel: &label}}}
}
