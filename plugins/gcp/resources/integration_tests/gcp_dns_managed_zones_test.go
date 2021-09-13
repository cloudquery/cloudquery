package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDNSManagedZones(t *testing.T) {
	testIntegrationHelper(t, resources.DNSManagedZones(), []string{
		"gcp_dns_managed_zones.tf",
		"network.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.DNSManagedZones().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("managed-zone%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":        fmt.Sprintf("managed-zone%s%s", res.Prefix, res.Suffix),
						"kind":        "dns#managedZone",
						"dns_name":    fmt.Sprintf("example-%s.com.", res.Suffix),
						"visibility":  "public",
						"description": "Example DNS zone",
						"labels": map[string]interface{}{
							"test": "test",
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_dns_managed_zone_dnssec_config_default_key_specs",
					ForeignKeyName: "managed_zone_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"algorithm":  "rsasha256",
								"kind":       "dns#dnsKeySpec",
								"key_length": float64(2048),
								"key_type":   "keySigning",
							},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"algorithm":  "rsasha256",
								"kind":       "dns#dnsKeySpec",
								"key_length": float64(1024),
								"key_type":   "zoneSigning",
							},
						},
					},
				},
			},
		}
	})
}

func TestIntegrationDNSManagedZonesPrivate(t *testing.T) {
	testIntegrationHelper(t, resources.DNSManagedZones(), []string{
		"gcp_dns_managed_zones_private.tf",
		"network.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.DNSManagedZones().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"name": fmt.Sprintf("private-zone%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":        fmt.Sprintf("private-zone%s%s", res.Prefix, res.Suffix),
						"kind":        "dns#managedZone",
						"dns_name":    fmt.Sprintf("example-p-%s.com.", res.Suffix),
						"visibility":  "private",
						"description": "Example DNS zone",
						"labels": map[string]interface{}{
							"test": "test",
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_dns_managed_zone_private_visibility_config_networks",
					ForeignKeyName: "managed_zone_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 2,
							Data: map[string]interface{}{
								"kind": "dns#managedZonePrivateVisibilityConfigNetwork",
							},
						},
					},
				},
				{
					Name:           "gcp_dns_managed_zone_forwarding_config_target_name_servers",
					ForeignKeyName: "managed_zone_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"kind":         "dns#managedZoneForwardingConfigNameServerTarget",
								"ipv4_address": "172.16.1.10",
							},
						},
					},
				},
				{
					Name:           "gcp_dns_managed_zone_forwarding_config_target_name_servers",
					ForeignKeyName: "managed_zone_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"kind":         "dns#managedZoneForwardingConfigNameServerTarget",
								"ipv4_address": "172.16.1.20",
							},
						},
					},
				},
			},
		}
	})
}
