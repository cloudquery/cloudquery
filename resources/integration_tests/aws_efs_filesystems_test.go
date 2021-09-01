package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEfsFilesystems(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EfsFilesystems(), []string{"aws_efs_filesystems.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_efs_filesystems",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":             fmt.Sprintf("efs-%s%s", res.Prefix, res.Suffix),
						"kms_key_id":       nil,
						"encrypted":        false,
						"performance_mode": "generalPurpose",
						"creation_token":   fmt.Sprintf("efs-%s%s", res.Prefix, res.Suffix),
					},
				},
			},
		}
	})
}
