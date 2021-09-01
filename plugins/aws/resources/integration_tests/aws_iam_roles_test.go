package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamRoles(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamRoles(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_roles",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix)})
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_iam_role_policies",
					ForeignKeyName: "role_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
						},
					},
				},
			},
		}
	})
}
