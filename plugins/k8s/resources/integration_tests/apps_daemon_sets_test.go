package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationAppsDaemonSets(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.AppsDaemonSets(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_apps_daemon_sets",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("daemon-set%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("daemon-set%s%s", res.Prefix, res.Suffix),
					"labels": map[string]interface{}{
						"test": "MyExampleApp",
					},
					"selector_match_labels": map[string]interface{}{
						"test": "MyExampleApp",
					},
					"update_strategy_type":   "RollingUpdate",
					"revision_history_limit": float64(10),
					"template": map[string]interface{}{
						"metadata": map[string]interface{}{
							"creationTimestamp": nil,
							"labels": map[string]interface{}{
								"test": "MyExampleApp",
							},
						},
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{
									"name":  "example",
									"image": "nginx:1.7.8",
									"resources": map[string]interface{}{
										"limits": map[string]interface{}{
											"cpu":    "500m",
											"memory": "512Mi",
										},
										"requests": map[string]interface{}{
											"cpu":    "250m",
											"memory": "50Mi",
										},
									},
									"livenessProbe": map[string]interface{}{
										"httpGet": map[string]interface{}{
											"path":   "/nginx_status",
											"port":   float64(80),
											"scheme": "HTTP",
											"httpHeaders": []interface{}{
												map[string]interface{}{
													"name":  "X-Custom-Header",
													"value": "Awesome",
												},
											},
										},
										"periodSeconds":       float64(3),
										"timeoutSeconds":      float64(1),
										"failureThreshold":    float64(3),
										"successThreshold":    float64(1),
										"initialDelaySeconds": float64(3),
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
