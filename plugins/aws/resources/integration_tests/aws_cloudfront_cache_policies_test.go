package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCloudfrontCachePolicies(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CloudfrontCachePolicies(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {

		return providertest.ResourceIntegrationVerification{
			Name: "aws_cloudfront_cache_policies",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("cache_policy%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"comment":                       "test comment",
					"default_ttl":                   float64(50),
					"max_ttl":                       float64(100),
					"query_strings_behavior":        "whitelist",
					"query_strings_quantity":        float64(1),
					"query_strings":                 []interface{}{"example"},
					"enable_accept_encoding_brotli": false,
				},
			}},
			//Relations: []*providertest.ResourceIntegrationVerification{
			//	{
			//		Name:           "aws_cloudfront_distribution_origins",
			//		ForeignKeyName: "distribution_cq_id",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"id":                  fmt.Sprintf("s3origin%s-%s", res.Prefix, res.Suffix),
			//				"connection_attempts": float64(3),
			//				"connection_timeout":  float64(10),
			//			},
			//		}},
			//	},
			//	{
			//		Name:           "aws_cloudfront_distribution_cache_behaviours",
			//		ForeignKeyName: "distribution_cq_id",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"target_origin_id":       fmt.Sprintf("s3origin%s-%s", res.Prefix, res.Suffix),
			//				"path_pattern":           "/content/immutable/*",
			//				"viewer_protocol_policy": "redirect-to-https",
			//				"allowed_methods":        []interface{}{"HEAD", "GET", "OPTIONS"},
			//			},
			//		}},
			//	},
			//	{
			//		Name:           "aws_cloudfront_distribution_custom_error_responses",
			//		ForeignKeyName: "distribution_cq_id",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"response_code":      "404",
			//				"response_page_path": "/custom_404.html",
			//			},
			//		}},
			//	},
			//},
		}
	})
}
