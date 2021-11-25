package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSageMakerNotebookInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SagemakerNotebookInstances(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_sagemaker_notebook_instances",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":          fmt.Sprintf("sagemaker-%s%s", res.Prefix, res.Suffix),
						"instance_type": "ml.t2.medium",
						"tags": map[string]interface{}{
							"TestId": res.Suffix,
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("sagemaker-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
