package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElasticbeanstalkEnvironments(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkEnvironments(), []string{"aws_elasticbeanstalk_environments.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elasticbeanstalk_environments",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": firstN(fmt.Sprintf("beanstalk-ee-%s", res.Suffix), 40)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"application_name":                fmt.Sprintf("beanstalk-ea-%s", res.Suffix),
						"tier_name":                       "WebServer",
						"tier_type":                       "Standard",
						"tier_version":                    "1.0",
						"abortable_operation_in_progress": false,
						"description":                     nil,
					},
				},
			},
		}
	})
}
