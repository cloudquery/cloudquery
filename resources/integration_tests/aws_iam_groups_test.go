package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/Masterminds/squirrel"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamGroups(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_groups",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("aws_iam_group%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("aws_iam_group%s%s", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_iam_group_policies",
					ForeignKeyName: "group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							//Data:  map[string]interface{}{},
						},
					},
				},
			},
		}
	})
}
