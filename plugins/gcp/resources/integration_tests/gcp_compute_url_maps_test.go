package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeURLMaps(t *testing.T) {
	testIntegrationHelper(t, resources.ComputeURLMaps(), []string{"gcp_compute_url_maps.tf", "network.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeURLMaps().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("urlmap%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                                             fmt.Sprintf("urlmap%s%s", res.Prefix, res.Suffix),
						"cors_policy_allow_credentials":                    false,
						"cors_policy_disabled":                             false,
						"cors_policy_max_age":                              float64(0),
						"fault_injection_policy_abort_http_status":         float64(0),
						"fault_injection_policy_abort_percentage":          float64(0),
						"fault_injection_policy_delay_fixed_delay_nanos":   float64(0),
						"fault_injection_policy_delay_fixed_delay_seconds": float64(0),
						"fault_injection_policy_delay_percentage":          float64(0),
						"max_stream_duration_nanos":                        float64(0),
						"max_stream_duration_seconds":                      float64(0),
						"retry_policy_num_retries":                         float64(0),
						"retry_policy_per_try_timeout_nanos":               float64(0),
						"retry_policy_per_try_timeout_seconds":             float64(0),
						"timeout_nanos":                                    float64(0),
						"timeout_seconds":                                  float64(0),
						"default_url_redirect_https_redirect":              false,
						"default_url_redirect_strip_query":                 false,
						"description":                                      "a description",
						"kind":                                             "compute#urlMap",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_compute_url_map_tests",
					ForeignKeyName: "url_map_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"expected_redirect_response_code": float64(0),
								"host":                            "hi.com",
								"path":                            "/home",
							},
						},
					},
				},
				{
					Name:           "gcp_compute_url_map_path_matchers",
					ForeignKeyName: "url_map_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":                                "mysite",
								"default_url_redirect_strip_query":    false,
								"default_url_redirect_https_redirect": false,
								"default_route_action":                nil,
								"header_action":                       nil,
								"route_rules":                         nil,
								"path_rules": []interface{}{
									map[string]interface{}{
										"paths":   []interface{}{"/home"},
										"service": fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/cq-e2e/global/backendBuckets/static-asset-backend-%s%s", res.Prefix, res.Suffix),
									},
									map[string]interface{}{
										"paths":   []interface{}{"/login"},
										"service": fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/cq-e2e/global/backendServices/url-maps-backend-svc-%s%s", res.Prefix, res.Suffix),
									},
									map[string]interface{}{
										"paths":   []interface{}{"/static"},
										"service": fmt.Sprintf("https://www.googleapis.com/compute/v1/projects/cq-e2e/global/backendBuckets/static-asset-backend-%s%s", res.Prefix, res.Suffix),
									},
								},
							},
						},
					},
				},
				{
					Name:           "gcp_compute_url_map_path_matchers",
					ForeignKeyName: "url_map_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":                                "otherpaths",
								"default_url_redirect_strip_query":    false,
								"default_url_redirect_https_redirect": false,
								"default_route_action":                nil,
								"header_action":                       nil,
								"route_rules":                         nil,
								"path_rules":                          nil,
							},
						},
					},
				},
				{
					Name:           "gcp_compute_url_map_host_rules",
					ForeignKeyName: "url_map_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"path_matcher": "otherpaths",
								"hosts":        []interface{}{"myothersite.com"},
							},
						},
					},
				},
				{
					Name:           "gcp_compute_url_map_host_rules",
					ForeignKeyName: "url_map_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"path_matcher": "mysite",
								"hosts":        []interface{}{"mysite.com"},
							},
						},
					},
				},
			},
		}
	})
}
