package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeInstances(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeInstances(), []string{"gcp_compute_instances.tf", "service-account.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeInstances().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("compute-instance-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                           fmt.Sprintf("compute-instance-%s%s", res.Prefix, res.Suffix),
						"scheduling_on_host_maintenance": "MIGRATE",
						"scheduling_min_node_cpus":       float64(0),
						"scheduling_automatic_restart":   true,
						"satisfies_pzs":                  false,
						"tags_items": []interface{}{
							"test",
						},
						"metadata_items": map[string]interface{}{
							"test":           "test",
							"startup-script": "echo hi > /test.txt",
						},
						"kind":                          "compute#instance",
						"display_device_enable_display": false,
						"deletion_protection":           false,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_instance_disks",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"auto_delete":  true,
								"boot":         true,
								"type":         "PERSISTENT",
								"mode":         "READ_WRITE",
								"disk_size_gb": float64(10),
							},
						},
					},
				},
				{
					Name:           "gcp_compute_instance_network_interfaces",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name": "nic0",
							},
						},
					},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "gcp_compute_instance_network_interface_access_configs",
							ForeignKeyName: "instance_network_interface_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"name":                            "external-nat",
										"instance_network_interface_name": "nic0",
									},
								},
							},
						},
						//todo add gcp_compute_instance_network_interface_alias_ip_ranges relation
					},
				},
				{
					Name:           "gcp_compute_instance_service_accounts",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"scopes": []interface{}{
									"https://www.googleapis.com/auth/cloud-platform",
								},
							},
						},
					},
				},
			},
		}
	})
}
