package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationLambdaLayers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.LambdaLayers(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_lambda_layers",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("lambda_layer%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("lambda_layer%s%s", res.Prefix, res.Suffix),
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_lambda_layer_versions",
					ForeignKeyName: "layer_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
					}},
				},
			},
		}
	})
}
