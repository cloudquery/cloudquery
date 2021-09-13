package integration_tests

import (
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeDiskTypes(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeDiskTypes(), []string{"gcp_compute_disks.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeDiskTypes().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"id": "us-east1/pd-balanced"})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"id":                   "us-east1/pd-balanced",
						"description":          "Balanced Persistent Disk",
						"valid_disk_size":      "10GB-65536GB",
						"default_disk_size_gb": float64(100),
					},
				},
			},
		}
	})
}
