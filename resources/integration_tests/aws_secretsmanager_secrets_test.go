package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecretsmanagerSecrets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.SecretsmanagerSecrets(), []string{"aws_secretsmanager_secrets.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_secretsmanager_secrets",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("secretsmanager-secret-%s%s", res.Prefix, res.Suffix),
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
						"Name":   fmt.Sprintf("secretsmanager-secret-%s%s", res.Prefix, res.Suffix),
					},
				},
			}},
		}
	})
}
