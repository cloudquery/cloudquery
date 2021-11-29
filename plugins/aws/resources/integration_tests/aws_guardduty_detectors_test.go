package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationGuarddutyDetectors(t *testing.T) {
	awsTestIntegrationHelper(t, resources.GuarddutyDetectors(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_guardduty_detectors",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"region": "us-east-1",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("fguardduty-detector-%s%s", res.Prefix, res.Suffix),
							"TestId": res.Suffix,
						},
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"region": "us-east-1"})
			},
		}
	})
}
