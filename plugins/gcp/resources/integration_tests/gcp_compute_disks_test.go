package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeDisks(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeDisks(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeDisks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("gcp-compute-disks-disk-%s", res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                      fmt.Sprintf("gcp-compute-disks-disk-%s", res.Suffix),
						"satisfies_pzs":             false,
						"size_gb":                   float64(10),
						"provisioned_iops":          float64(0),
						"physical_block_size_bytes": float64(4096),
						"labels": map[string]interface{}{
							"environment": "dev",
						},
					},
				},
			},
		}
	})
}
