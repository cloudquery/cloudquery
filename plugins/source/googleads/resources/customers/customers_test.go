package customers

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
	"github.com/stretchr/testify/require"
)

func TestCustomersQuery(t *testing.T) {
	// we need to have this test to check the resulting query
	// using https://developers.google.com/google-ads/api/fields/v13/query_validator
	// TODO: update the link once new API is out
	const expected = `SELECT
	customer.resource_name,
	customer.id,
	customer.descriptive_name,
	customer.currency_code,
	customer.time_zone,
	customer.tracking_url_template,
	customer.final_url_suffix,
	customer.auto_tagging_enabled,
	customer.has_partners_badge,
	customer.manager,
	customer.test_account,
	customer.call_reporting_setting.call_reporting_enabled,
	customer.call_reporting_setting.call_conversion_reporting_enabled,
	customer.call_reporting_setting.call_conversion_action,
	customer.conversion_tracking_setting.conversion_tracking_id,
	customer.conversion_tracking_setting.cross_account_conversion_tracking_id,
	customer.conversion_tracking_setting.accepted_customer_data_terms,
	customer.conversion_tracking_setting.conversion_tracking_status,
	customer.conversion_tracking_setting.enhanced_conversions_for_leads_enabled,
	customer.conversion_tracking_setting.google_ads_conversion_customer,
	customer.remarketing_setting.google_global_site_tag,
	customer.pay_per_conversion_eligibility_failure_reasons,
	customer.optimization_score,
	customer.optimization_score_weight,
	customer.status,
	customer.location_asset_auto_migration_done,
	customer.image_asset_auto_migration_done,
	customer.location_asset_auto_migration_done_date_time,
	customer.image_asset_auto_migration_done_date_time
FROM customer`

	require.Equal(t, expected, gaql.Query(new(resources.Customer), nil, customerOptions))
}

func testCustomers(t *testing.T) map[string][]*services.GoogleAdsRow {
	var customer resources.Customer
	require.NoError(t, faker.FakeObject(&customer))
	row := &services.GoogleAdsRow{Customer: &customer}
	return client.MapsCombine(
		map[string][]*services.GoogleAdsRow{"customer": {row}},
		testCustomerLabels(t),
	)
}

func TestCustomers(t *testing.T) {
	client.MockTestHelper(t, Customers(), testCustomers(t))
}
