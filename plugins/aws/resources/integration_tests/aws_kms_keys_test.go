package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationKmsKeys(t *testing.T) {
	awsTestIntegrationHelper(t, resources.KmsKeys(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.KmsKeys().Name,
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"description": fmt.Sprintf("kms-key-%s%s", res.Prefix, res.Suffix),
					"tags": map[string]interface{}{
						"Type":   "integration_test",
						"TestId": res.Suffix,
					},
				},
			}},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"description": fmt.Sprintf("kms-key-%s%s", res.Prefix, res.Suffix)},
					squirrel.Eq{"deletion_date": nil},
				})
			},
		}
	})
}
