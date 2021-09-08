package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDatabases(t *testing.T) {
	testIntegrationHelper(t, resources.Databases(), []string{"do_databases.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Databases().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("do-database-cluster-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do-database-cluster-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_database_users",
					ForeignKeyName: "database_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name": fmt.Sprintf("do-database-user-%s-%s", res.Prefix, res.Suffix),
							},
						},
					},
				},
				{
					Name:           "digitalocean_database_replicas",
					ForeignKeyName: "database_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name": fmt.Sprintf("do-database-replica-%s-%s", res.Prefix, res.Suffix),
							},
						},
					},
				},
				{
					Name:           "digitalocean_database_firewall_rules",
					ForeignKeyName: "database_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"type":  "ip_addr",
								"value": "192.168.1.1",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"type":  "ip_addr",
								"value": "192.0.2.0",
							},
						},
					},
				},
			},
		}
	})
}
