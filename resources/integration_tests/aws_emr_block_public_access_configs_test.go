package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEmrBlockPublicAccessConfigs(t *testing.T) {
	table := resources.EmrBlockPublicAccessConfigs()
	awsTestIntegrationHelper(t, table, []string{}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"block_public_security_group_rules": true,
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_emr_block_public_access_config_port_ranges",
					ForeignKeyName: "block_public_access_config_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"min_range": float64(22),
								"max_range": float64(22),
							},
						},
					},
				},
			},
		}
	})
}
