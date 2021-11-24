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
						"replication_instance_class": "dms.t2.micro",
						"replication_instance_id":    fmt.Sprintf("dms-replication-instance-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
		}
	})
}
