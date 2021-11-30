package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSSMDocuments(t *testing.T) {
	table := resources.SsmDocuments()
	awsTestIntegrationHelper(t, table, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"description":     "Check ip configuration of a Linux instance.",
					"document_format": "JSON",
					"document_type":   "Command",
					"name":            fmt.Sprintf("%sdoc%s", res.Prefix, res.Suffix),
					"account_ids":     []interface{}{"all"},
					"tags": map[string]interface{}{
						"Name":   fmt.Sprintf("%sdoc%s", res.Prefix, res.Suffix),
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
		}
	})
}
