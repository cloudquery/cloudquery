package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationFsxBackups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.FsxBackups(), []string{"aws_fsx_backups.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_fsx_backups",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("fsx-%s%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}
