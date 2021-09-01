package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayRestApis(), []string{"aws_apigateway_rest_apis.tf", "aws_lambda_functions.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_rest_apis",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"endpoint_configuration_types": []interface{}{"REGIONAL"},
					"api_key_source":               "HEADER",
					"name":                         fmt.Sprintf("apigwv1-api-%s%s", res.Prefix, res.Suffix),
					"version":                      "1.0",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_apigateway_rest_api_deployments",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"description": fmt.Sprintf("apigwv1-dep-%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_authorizers",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":                             fmt.Sprintf("apigwv1-authorizer-%s%s", res.Prefix, res.Suffix),
							"auth_type":                        "custom",
							"authorizer_result_ttl_in_seconds": float64(500),
							"type":                             "TOKEN",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_resources",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"path_part": "gateway_resource_1",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_models",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"content_type": "application/json",
							"description":  "a JSON schema",
							"name":         fmt.Sprintf("apigwv1apimodel%s", res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_request_validators",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"validate_request_parameters": true,
							"validate_request_body":       true,
							"name":                        fmt.Sprintf("apigwv1-req-validation-%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_documentation_parts",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"location_type":   "METHOD",
							"location_method": "GET",
							"location_path":   "/example",
							"properties":      "{\"description\":\"Example description\"}",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_documentation_versions",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"description": "Example description",
							"version":     "example_version",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_stages",
					ForeignKeyName: "rest_api_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"tracing_enabled": false,
								"stage_name":      fmt.Sprintf("apigwv1-stage-%s%s", res.Prefix, res.Suffix),
								"tags": map[string]interface{}{
									"hello":  "world",
									"TestId": res.Suffix,
									"Type":   "integration_test",
								}},
						},
						{
							Count: 1,
							Data: map[string]interface{}{
								"tracing_enabled": false,
								"stage_name":      fmt.Sprintf("apigwv1-stage2-%s%s", res.Prefix, res.Suffix),
								"tags": map[string]interface{}{
									"hello":  "world1",
									"TestId": res.Suffix,
									"Type":   "integration_test",
								}},
						}},
				},
			},
		}
	})
}
