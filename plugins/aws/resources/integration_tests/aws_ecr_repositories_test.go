package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEcrRepositories(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcrRepositories(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ecr_repositories",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                 fmt.Sprintf("ecr_repositories_%s%s", res.Prefix, res.Suffix),
						"image_tag_mutability": "MUTABLE",
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ecr_repositories_%s%s", res.Prefix, res.Suffix)})
			},
		}
	})
}
