package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationAppsStatefulSets(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.AppsStatefulSets(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_apps_stateful_sets",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("stateful-set%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("stateful-set%s%s", res.Prefix, res.Suffix),
					"labels": map[string]interface{}{
						"app":     "local-scheduler",
						"service": "mongodb",
					},
					"selector_match_labels": map[string]interface{}{
						"app":     "local-scheduler",
						"service": "mongodb",
					},
					"service_name":                             "mongodb",
					"pod_management_policy":                    "OrderedReady",
					"update_strategy_type":                     "RollingUpdate",
					"revision_history_limit":                   float64(0),
					"update_strategy_rolling_update_partition": float64(0),
					"min_ready_seconds":                        float64(0),
					"template": map[string]interface{}{
						"metadata": map[string]interface{}{
							"creationTimestamp": nil,
							"labels": map[string]interface{}{
								"app":     "local-scheduler",
								"service": "mongodb",
							},
						},
						"spec": map[string]interface{}{
							"volumes": []interface{}{
								map[string]interface{}{
									"name":     "configdir",
									"emptyDir": map[string]interface{}{},
								},
								map[string]interface{}{
									"name":     "datadir",
									"emptyDir": map[string]interface{}{},
								},
							},
							"containers": []interface{}{
								map[string]interface{}{
									"env": []interface{}{
										map[string]interface{}{
											"name":  "EDGE_PORT",
											"value": "1235",
										},
									},
									"args": []interface{}{
										"--dbpath=/data/db",
										"--port=1235",
										"--bind_ip=0.0.0.0",
									},
									"name":  "mongodb",
									"image": "mongo:bionic",
									"ports": []interface{}{
										map[string]interface{}{
											"protocol":      "TCP",
											"containerPort": float64(1235),
										},
									},
									"command": []interface{}{
										"mongod",
									},
									"resources": map[string]interface{}{},
									"volumeMounts": []interface{}{
										map[string]interface{}{
											"name":             "configdir",
											"mountPath":        "/data/configdb",
											"mountPropagation": "None",
										},
										map[string]interface{}{
											"name":             "datadir",
											"mountPath":        "/data/db",
											"mountPropagation": "None",
										},
									},
									"imagePullPolicy":          "IfNotPresent",
									"terminationMessagePath":   "/dev/termination-log",
									"terminationMessagePolicy": "File",
								},
							},
							"dnsPolicy":                     "ClusterFirst",
							"restartPolicy":                 "Always",
							"schedulerName":                 "default-scheduler",
							"securityContext":               map[string]interface{}{},
							"enableServiceLinks":            true,
							"shareProcessNamespace":         false,
							"automountServiceAccountToken":  true,
							"terminationGracePeriodSeconds": float64(30),
						},
					},
				},
			}},
		}
	})
}
