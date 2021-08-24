package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamUsers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamUsers(), []string{"aws_iam_users.tf", "aws_iam_groups.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_users",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"user_name": fmt.Sprintf("user%s%s", res.Prefix, res.Suffix),
				},
			}},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"user_name": fmt.Sprintf("user%s%s", res.Prefix, res.Suffix)})
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_iam_user_policies",
					ForeignKeyName: "user_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"policy_name": fmt.Sprintf("user_policy%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_iam_user_access_keys",
					ForeignKeyName: "user_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
					}},
				},
				{
					Name:           "aws_iam_user_attached_policies",
					ForeignKeyName: "user_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"policy_name": fmt.Sprintf("policy%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_iam_user_groups",
					ForeignKeyName: "user_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"group_name": fmt.Sprintf("aws_iam_group%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
			},
		}
	})
}
