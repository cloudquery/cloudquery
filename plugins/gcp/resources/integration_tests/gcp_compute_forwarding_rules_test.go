package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeForwardingRules(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeForwardingRules(), []string{"gcp_compute_forwarding_rules.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeForwardingRules().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("forwarding-rule-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                   fmt.Sprintf("forwarding-rule-%s%s", res.Prefix, res.Suffix),
						"load_balancing_scheme":  "EXTERNAL",
						"is_mirroring_collector": false,
						"ip_protocol":            "TCP",
						"all_ports":              false,
						"allow_global_access":    false,
						"network_tier":           "PREMIUM",
						"port_range":             "80-80",
						"kind":                   "compute#forwardingRule",
						"labels": map[string]interface{}{
							"test": "test",
						},
					},
				},
			},
		}
	})
}
