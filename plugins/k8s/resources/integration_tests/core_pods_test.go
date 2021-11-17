package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationCorePods(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.CorePods(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_core_pods",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("pod-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{

				Count: 1,
				Data: map[string]interface{}{
					"name":                             fmt.Sprintf("pod-%s%s", res.Prefix, res.Suffix),
					"labels":                           nil,
					"annotations":                      nil,
					"owner_references":                 nil,
					"restart_policy":                   "Always",
					"termination_grace_period_seconds": float64(30),
					"dns_policy":                       "None",
					"node_selector":                    nil,
					"automount_service_account_token":  true,
					"host_network":                     false,
					"host_pid":                         false,
					"host_ipc":                         false,
					"image_pull_secrets":               nil,
					"affinity":                         nil,
					"scheduler_name":                   "default-scheduler",
					"qos_class":                        "BestEffort",
					"preemption_policy":                "PreemptLowerPriority",
					"enable_service_links":             true,
					"overhead":                         nil,
					"security_context":                 map[string]interface{}{},
					"share_process_namespace":          false,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "k8s_core_pod_containers",
					ForeignKeyName: "pod_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":                       "example",
							"image":                      "nginx:1.20.1",
							"env_from":                   nil,
							"resources_limits":           nil,
							"resources_requests":         nil,
							"tty":                        false,
							"stdin_once":                 false,
							"stdin":                      false,
							"security_context":           nil,
							"image_pull_policy":          "IfNotPresent",
							"termination_message_policy": "File",
							"termination_message_path":   "/dev/termination-log",
							"lifecycle":                  nil,
							"startup_probe":              nil,
							"readiness_probe":            nil,
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "k8s_core_pod_container_envs",
							ForeignKeyName: "pod_container_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"name":                                   "environment",
									"value":                                  "test",
									"value_from_secret_key_ref_optional":     nil,
									"value_from_config_map_key_ref_optional": nil,
								},
							}},
						},
						{
							Name:           "k8s_core_pod_container_ports",
							ForeignKeyName: "pod_container_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"host_port":      float64(0),
									"container_port": float64(8080),
									"protocol":       "TCP",
								},
							}},
						},
					},
				},
			},
		}
	})
}
