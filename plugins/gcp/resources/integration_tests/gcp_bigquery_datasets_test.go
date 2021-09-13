package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationBigqueryDatasets(t *testing.T) {
	testIntegrationHelper(t, resources.BigqueryDatasets(), []string{"gcp_bigquery_datasets.tf", "service-account.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.BigqueryDatasets().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"friendly_name": fmt.Sprintf("bigquery_dataset_%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"friendly_name": fmt.Sprintf("bigquery_dataset_%s%s", res.Prefix, res.Suffix),
					},
				},
			},
		}
	})
}
