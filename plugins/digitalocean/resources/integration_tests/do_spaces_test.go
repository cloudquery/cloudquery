package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSpaces(t *testing.T) {
	testIntegrationHelper(t, resources.Spaces(), []string{"do_spaces.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Spaces().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("do-spaces-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do-spaces-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_space_cors",
					ForeignKeyName: "space_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"allowed_origins": []interface{}{fmt.Sprintf("https://www.%s-%s.com", res.Prefix, res.Suffix)},
							},
						},
					},
				},
			},
		}
	})
}
