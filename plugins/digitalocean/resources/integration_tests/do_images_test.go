package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationImages(t *testing.T) {
	testIntegrationHelper(t, resources.Images(), []string{"do_images.tf", "do_droplets.tf", "do_snapshots.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Images().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Or{
					squirrel.Eq{"name": fmt.Sprintf("do_image%s-%s", res.Prefix, res.Suffix)},
					squirrel.Eq{"name": fmt.Sprintf("do_image_snap%s-%s", res.Prefix, res.Suffix)},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do_image%s-%s", res.Prefix, res.Suffix),
						"type": "custom",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do_image_snap%s-%s", res.Prefix, res.Suffix),
						"type": "snapshot",
					},
				},
			},
		}
	})
}
