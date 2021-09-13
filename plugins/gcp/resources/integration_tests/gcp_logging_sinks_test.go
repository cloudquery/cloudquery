package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationLoggingSinks(t *testing.T) {
	testIntegrationHelper(t, resources.LoggingSinks(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.LoggingSinks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("logging-sink-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":             fmt.Sprintf("logging-sink-%s-%s", res.Prefix, res.Suffix),
						"description":      "a description",
						"include_children": false,
						"disabled":         false,
						"bigquery_options_use_partitioned_tables":             false,
						"bigquery_options_uses_timestamp_column_partitioning": false,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_logging_sink_exclusions",
					ForeignKeyName: "sink_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":        fmt.Sprintf("ex-%s-%s", res.Prefix, res.Suffix),
								"description": "Exclude logs from namespace-1 in k8s",
								"disabled":    false,
								"filter":      "resource.type = k8s_container resource.labels.namespace_name=\"namespace-1\"",
							},
						},
					},
				},
			},
		}
	})
}
