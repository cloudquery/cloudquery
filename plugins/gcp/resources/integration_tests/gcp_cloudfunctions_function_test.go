package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCloudfunctionsFunction(t *testing.T) {
	table := resources.CloudfunctionsFunction()
	testIntegrationHelper(t, table, []string{"gcp_cloudfunctions_functions.tf", "helloworld.zip"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"description":         fmt.Sprintf("My function %s%s", res.Prefix, res.Suffix),
						"entry_point":         "HelloHTTP",
						"runtime":             "go113",
						"available_memory_mb": float64(128),
					},
				},
			},
		}
	})
}
