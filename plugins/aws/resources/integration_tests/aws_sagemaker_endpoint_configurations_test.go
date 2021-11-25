package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSageMakerEndpointConfigurations(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerEndpointConfigurations(), []string{"aws_sagemaker_endpoint_configurations.tf", "aws_sagemaker_models.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_sagemaker_endpoint_configurations",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("sagemaker-endpoint-configuration-%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"TestId": res.Suffix,
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("sagemaker-endpoint-configuration-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_sagemaker_endpoint_configuration_production_variants",
					ForeignKeyName: "endpoint_configuration_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"variant_name":           fmt.Sprintf("variant-%s%s", res.Prefix, res.Suffix),
								"model_name":             fmt.Sprintf("sagemaker-model-%s%s", res.Prefix, res.Suffix),
								"initial_instance_count": float64(1),
								"instance_type":          "ml.t2.medium",
							},
						},
					},
				},
			},
		}
	})
}
