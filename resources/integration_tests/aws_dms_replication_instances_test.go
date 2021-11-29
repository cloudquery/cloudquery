package integration_tests

import (
	"fmt"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"

	"github.com/cloudquery/cq-provider-aws/resources"
)

func TestIntegrationDmsReplicationInstances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DmsReplicationInstances(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.DmsReplicationInstances().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"allocated_storage":            float64(20),
						"auto_minor_version_upgrade":   true,
						"availability_zone":            "us-east-1a",
						"multi_az":                     false,
						"preferred_maintenance_window": "sun:10:30-sun:14:30",
						"publicly_accessible":          false,
						"class":                        "dms.t2.micro",
						"identifier":                   fmt.Sprintf("dms-replication-instance-%s-%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("dms-replication-instance-%s-%s", res.Prefix, res.Suffix),
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
