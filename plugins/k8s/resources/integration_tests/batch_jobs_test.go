package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-k8s/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationBatchJobs(t *testing.T) {
	k8sTestIntegrationHelper(t, resources.BatchJobs(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "k8s_batch_jobs",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("batch-job%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":            fmt.Sprintf("batch-job%s%s", res.Prefix, res.Suffix),
					"annotations":     nil,
					"backoff_limit":   float64(4),
					"manual_selector": false,
					"suspend":         false,
					//"template": map[string]interface{}{
					//	//"metadata": map[string]interface{}{
					//	//	"creationTimestamp": nil,
					//	//	"labels": map[string]interface{}{
					//	//		"job-name":       "batch-jobandriitest",
					//	//		"controller-uid": "e68058f2-e70e-474b-90eb-582a4693f610", // todo make field verification be able to skip unique fields
					//	//	},
					//	//},
					//	"spec": map[string]interface{}{
					//		"containers": []interface{}{
					//			map[string]interface{}{
					//				"name":  "pi",
					//				"image": "perl",
					//				"command": []interface{}{
					//					"perl",
					//					"-Mbignum=bpi",
					//					"-wle",
					//					"print bpi(2000)",
					//				},
					//				"resources":                map[string]interface{}{},
					//				"imagePullPolicy":          "Always",
					//				"terminationMessagePath":   "/dev/termination-log",
					//				"terminationMessagePolicy": "File",
					//			},
					//		},
					//		"dnsPolicy":                     "ClusterFirst",
					//		"restartPolicy":                 "Never",
					//		"schedulerName":                 "default-scheduler",
					//		"securityContext":               map[string]interface{}{},
					//		"enableServiceLinks":            true,
					//		"shareProcessNamespace":         false,
					//		"automountServiceAccountToken":  true,
					//		"terminationGracePeriodSeconds": float64(30),
					//	},
					//},
				}},
			},
		}
	})
}
