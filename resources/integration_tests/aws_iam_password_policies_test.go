package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIamPasswordPolicies(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamPasswordPolicies(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_password_policies",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"policy_exists": false})

			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"allow_users_to_change_password": false,
						"expire_passwords":               false,
						"hard_expiry":                    nil,
						"max_password_age":               nil,
						"minimum_password_length":        nil,
						"password_reuse_prevention":      nil,
						"require_lowercase_characters":   false,
						"require_numbers":                false,
						"require_symbols":                false,
						"require_uppercase_characters":   false,
						"policy_exists":                  false,
					},
				},
			},
		}
	})
}
