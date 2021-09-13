package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeSubnetworks(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeSubnetworks(), []string{"network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeSubnetworks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("network-subnetwork-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                       fmt.Sprintf("network-subnetwork-%s%s", res.Prefix, res.Suffix),
						"enable_flow_logs":           false,
						"description":                "",
						"private_ip_google_access":   false,
						"private_ipv6_google_access": "DISABLE_GOOGLE_ACCESS",
						"purpose":                    "PRIVATE",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_subnetwork_secondary_ip_ranges",
					ForeignKeyName: "subnetwork_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"range_name":    fmt.Sprintf("range-%s%s", res.Prefix, res.Suffix),
								"ip_cidr_range": "192.168.10.0/24",
							},
						},
					},
				},
			},
		}
	})
}
