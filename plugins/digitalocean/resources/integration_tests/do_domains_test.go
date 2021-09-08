package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDomains(t *testing.T) {
	testIntegrationHelper(t, resources.Domains(), []string{"do_domains.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Domains().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s.com", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("%s-%s.com", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_domain_records",
					ForeignKeyName: "domain_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"type": "A",
								"name": "www",
								"data": "192.168.0.11",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"type": "TXT",
								"name": "test",
								"data": fmt.Sprintf("do_record_txt%s-%s", res.Prefix, res.Suffix),
							},
						},
					},
				},
			},
		}
	})
}
