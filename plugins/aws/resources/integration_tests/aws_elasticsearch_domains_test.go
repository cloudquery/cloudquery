package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationElasticsearchDomains(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ElasticsearchDomains(), []string{"aws_elasticsearch_domains.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_elasticsearch_domains",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("elastic-domain-%.13s", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"elasticsearch_version": "7.10",
						"ebs_volume_type":       "gp2",
						"ebs_volume_size":       float64(10),
						"ebs_enabled":           true,
						"name":                  fmt.Sprintf("elastic-domain-%.13s", res.Suffix),
						"snapshot_options_automated_snapshot_start_hour": float64(23),
					},
				},
			},
		}
	})
}
