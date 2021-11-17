package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationCoreNamespaces(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.CoreNamespaces(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_core_namespaces",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("namespace%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("namespace%s%s", res.Prefix, res.Suffix),
					"annotations": map[string]interface{}{
						"name": "namespace",
					},
					"labels": map[string]interface{}{
						"mylabel":                     "label-value",
						"kubernetes.io/metadata.name": fmt.Sprintf("namespace%s%s", res.Prefix, res.Suffix),
					},
					"phase":      "Active",
					"conditions": nil,
				}},
			},
		}
	})
}
