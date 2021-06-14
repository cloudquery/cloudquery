package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationIamGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamGroups(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_groups",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("group_name = ?", fmt.Sprintf("aws_iam_group%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					//Data: map[string]interface{}{
					//	"role_name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
					//},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_iam_group_policies",
					ForeignKeyName: "group_id",
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
