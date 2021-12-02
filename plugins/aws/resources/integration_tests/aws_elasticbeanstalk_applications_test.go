package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElasticbeanstalkApplications(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticbeanstalkApplications(), []string{"aws_elasticbeanstalk_environments.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elasticbeanstalk_applications",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("beanstalk-ea-%s", res.Suffix)})
			},

			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                               fmt.Sprintf("beanstalk-ea-%s", res.Suffix),
						"description":                        "tf-test-desc",
						"max_age_rule_enabled":               false,
						"max_age_rule_delete_source_from_s3": false,
						"max_count_rule_enabled":             false,
						"versions":                           nil,
						"max_count_rule_max_count":           float64(200),
						"max_age_rule_max_age_in_days":       float64(180),
					},
				},
			},
		}
	})
}
