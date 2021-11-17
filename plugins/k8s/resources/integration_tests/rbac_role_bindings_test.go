package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationRbacRoleBindings(t *testing.T) {
	schema := resources.RbacRoleBindings()
	k8sTestIntegrationHelper(t, schema, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: schema.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("role-binding-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":               fmt.Sprintf("role-binding-%s%s", res.Prefix, res.Suffix),
					"labels":             nil,
					"annotations":        nil,
					"role_ref_name":      "admin",
					"role_ref_kind":      "Role",
					"role_ref_api_group": "rbac.authorization.k8s.io",
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "k8s_rbac_role_binding_subjects",
					ForeignKeyName: "role_binding_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"kind":      "User",
								"api_group": "rbac.authorization.k8s.io",
								"name":      "admin",
							},
						},
					},
				},
			},
		}
	})
}
