package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationRbacRoles(t *testing.T) {
	schema := resources.RbacRoles()
	k8sTestIntegrationHelper(t, schema, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: schema.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("role-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("role-%s%s", res.Prefix, res.Suffix),
					"labels": map[string]interface{}{
						"test": "MyRole",
					},
					"annotations": nil,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "k8s_rbac_role_rules",
					ForeignKeyName: "role_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"verbs":             []interface{}{"list", "get"},
								"api_groups":        []interface{}{"apps"},
								"resources":         []interface{}{"deployments"},
								"resource_names":    nil,
								"non_resource_urls": nil,
							},
						},
					},
				},
			},
		}
	})
}
