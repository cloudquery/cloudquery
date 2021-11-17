package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationBatchCronJobs(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.BatchCronJobs(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_batch_cron_jobs",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("cron-job%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                          fmt.Sprintf("cron-job%s%s", res.Prefix, res.Suffix),
					"schedule":                      "1 0 * * *",
					"starting_deadline_seconds":     float64(10),
					"concurrency_policy":            "Replace",
					"successful_jobs_history_limit": float64(10),
					"job_template": map[string]interface{}{
						"spec": map[string]interface{}{
							"completions":             float64(1),
							"parallelism":             float64(1),
							"backoffLimit":            float64(2),
							"manualSelector":          false,
							"ttlSecondsAfterFinished": float64(10),
							"template": map[string]interface{}{

								"metadata": map[string]interface{}{
									"creationTimestamp": nil,
								},
								"spec": map[string]interface{}{
									"containers": []interface{}{
										map[string]interface{}{
											"name":  "hello",
											"image": "busybox",
											"command": []interface{}{
												"/bin/sh",
												"-c",
												"date; echo Hello from the Kubernetes cluster",
											},
											"resources":                map[string]interface{}{},
											"imagePullPolicy":          "Never",
											"terminationMessagePath":   "/dev/termination-log",
											"terminationMessagePolicy": "File",
										},
									},
									"dnsPolicy":                     "ClusterFirst",
									"restartPolicy":                 "Never",
									"schedulerName":                 "default-scheduler",
									"securityContext":               map[string]interface{}{},
									"enableServiceLinks":            true,
									"shareProcessNamespace":         false,
									"automountServiceAccountToken":  true,
									"terminationGracePeriodSeconds": float64(30),
								},
							},
						},
						"metadata": map[string]interface{}{
							"creationTimestamp": nil,
						},
					},
				}},
			},
		}
	})
}
