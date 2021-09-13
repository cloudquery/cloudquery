package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

// made to fetch only one
func TestIntegrationResourceManagerProjects(t *testing.T) {
	testIntegrationHelper(t, resources.ResourceManagerProjects(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ResourceManagerProjects().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"state": "ACTIVE"})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"state": "ACTIVE",
					},
				},
			},
		}
	})
}
