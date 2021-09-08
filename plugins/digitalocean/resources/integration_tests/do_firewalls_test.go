package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-digitalocean/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationFirewalls(t *testing.T) {
	testIntegrationHelper(t, resources.Firewalls(), []string{"do_firewalls.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Firewalls().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("do-firewall-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("do-firewall-%s-%s", res.Prefix, res.Suffix),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "digitalocean_firewall_inbound_rules",
					ForeignKeyName: "firewall_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":          "icmp",
								"port_range":        "0",
								"sources_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":          "tcp",
								"port_range":        "22",
								"sources_addresses": []interface{}{"192.168.1.0/24", "2002:1:2::/48"},
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":          "tcp",
								"port_range":        "80",
								"sources_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":          "tcp",
								"port_range":        "443",
								"sources_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
					},
				},
				{
					Name:           "digitalocean_firewall_outbound_rules",
					ForeignKeyName: "firewall_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":               "icmp",
								"port_range":             "0",
								"destinations_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":               "tcp",
								"port_range":             "53",
								"destinations_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"protocol":               "udp",
								"port_range":             "53",
								"destinations_addresses": []interface{}{"0.0.0.0/0", "::/0"},
							},
						},
					},
				},
			},
		}
	})
}
