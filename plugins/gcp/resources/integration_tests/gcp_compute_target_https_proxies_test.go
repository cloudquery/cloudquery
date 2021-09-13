package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeTargetHTTPSProxies(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeTargetHTTPSProxies(), []string{
		"gcp_compute_target_https_proxies.tf",
		"gcp_compute_ssl_certificates.tf",
		"gcp_compute_ssl_policies.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeTargetHTTPSProxies().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("https-proxy-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":          fmt.Sprintf("https-proxy-%s%s", res.Prefix, res.Suffix),
						"proxy_bind":    false,
						"kind":          "compute#targetHttpsProxy",
						"quic_override": "NONE",
					},
				},
			},
		}
	})
}
