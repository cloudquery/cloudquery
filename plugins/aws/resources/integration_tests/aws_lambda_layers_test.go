package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationLambdaLayers(t *testing.T) {
	awsTestIntegrationHelper(t, resources.LambdaLayers(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_lambda_layers",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("layer_name = ?", fmt.Sprintf("lambda_layer%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				//Data: map[string]interface{}{
				//	"tracing_config_mode": "PassThrough",
				//},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_lambda_layer_versions",
					ForeignKeyName: "layer_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						//Data: map[string]interface{}{
						//	"tracing_config_mode": "PassThrough",
						//},
					}},
				},
			},
		}
	})
}
