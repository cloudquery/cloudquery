package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeBackendServices(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeBackendServices(), []string{"gcp_compute_forwarding_rules.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeBackendServices().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("backend-svc-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                  fmt.Sprintf("backend-svc-%s%s", res.Prefix, res.Suffix),
						"load_balancing_scheme": "EXTERNAL",
						"cdn_policy_cache_key_policy_include_host":         false,
						"cdn_policy_cache_key_policy_include_protocol":     false,
						"cdn_policy_cache_key_policy_include_query_string": false,
						"cdn_policy_client_ttl":                            float64(0),
						"cdn_policy_default_ttl":                           float64(0),
						"cdn_policy_max_ttl":                               float64(0),
						"cdn_policy_negative_caching":                      false,
						"cdn_policy_request_coalescing":                    false,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_backend_service_backends",
					ForeignKeyName: "backend_service_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"balancing_mode":               "CONNECTION",
								"capacity_scaler":              float64(0),
								"failover":                     false,
								"max_connections":              float64(0),
								"max_connections_per_endpoint": float64(0),
								"max_connections_per_instance": float64(0),
								"max_rate":                     float64(0),
								"max_rate_per_endpoint":        float64(0),
								"max_rate_per_instance":        float64(0),
								"max_utilization":              float64(0),
							},
						},
					},
				},
			},
		}
	})
}
