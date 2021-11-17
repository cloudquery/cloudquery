package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationCoreServices(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.CoreServices(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_core_services",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("service%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{

				Count: 1,
				Data: map[string]interface{}{
					"name":        fmt.Sprintf("service%s%s", res.Prefix, res.Suffix),
					"labels":      nil,
					"annotations": nil,
					"selector": map[string]interface{}{
						"app": "MyApp",
					},
					"ip_family_policy":        "SingleStack",
					"internal_traffic_policy": "Cluster",
				},
			}},
		}
	})
}
