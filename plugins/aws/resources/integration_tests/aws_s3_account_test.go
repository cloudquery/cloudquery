package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationS3Account(t *testing.T) {
	awsTestIntegrationHelper(t, resources.S3Accounts(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_s3_account_config",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"config_exists": false})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"config_exists":           false,
						"block_public_acls":       false,
						"block_public_policy":     false,
						"ignore_public_acls":      false,
						"restrict_public_buckets": false,
					},
				},
			},
		}
	})
}
