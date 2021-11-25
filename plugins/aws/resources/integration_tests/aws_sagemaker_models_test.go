package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSageMakerModels(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerModels(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_sagemaker_models",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("sagemaker-model-%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"TestId": res.Suffix,
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("sagemaker-model-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
