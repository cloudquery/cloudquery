package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationStorageBuckets(t *testing.T) {
	testIntegrationHelper(t, resources.StorageBuckets(), []string{
		"gcp_storage_buckets.tf",
		"service-account.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.StorageBuckets().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"name": fmt.Sprintf("gcp-storage-buckets-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"location_type":            "multi-region",
						"kind":                     "storage#bucket",
						"id":                       fmt.Sprintf("gcp-storage-buckets-%s-%s", res.Prefix, res.Suffix),
						"billing_requester_pays":   false,
						"default_event_based_hold": false,
						"iam_configuration_bucket_policy_only_enabled":          false,
						"iam_configuration_public_access_prevention":            "inherited",
						"iam_configuration_uniform_bucket_level_access_enabled": false,
						"retention_policy_is_locked":                            false,
						"retention_policy_retention_period":                     float64(0),
						"satisfies_pzs":                                         false,
						"storage_class":                                         "STANDARD",
						"versioning_enabled":                                    false,
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_storage_bucket_lifecycle_rules",
					ForeignKeyName: "bucket_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"action_type":                          "Delete",
								"condition_age":                        float64(3),
								"condition_days_since_custom_time":     float64(0),
								"condition_days_since_noncurrent_time": float64(0),
								"condition_num_newer_versions":         float64(0),
							},
						},
					},
				},
				{
					Name:           "gcp_storage_bucket_cors",
					ForeignKeyName: "bucket_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"response_header": []interface{}{
									"*",
								},
								"origin": []interface{}{
									"http://image-store.com",
								},
								"method":          []interface{}{"GET", "HEAD", "PUT", "POST", "DELETE"},
								"max_age_seconds": float64(3600),
							},
						},
					},
				},
			},
			//todo add 2 more relations checks
		}
	})
}
