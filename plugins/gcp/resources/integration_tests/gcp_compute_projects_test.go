package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeProjects(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeProjects(), []string{}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeProjects().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data:  map[string]interface{}{},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_project_quotas",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 39,
							Data:  map[string]interface{}{},
						},
					},
				},
			},
		}
	})
}
