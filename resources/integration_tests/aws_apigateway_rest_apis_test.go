package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationApigatewayRestApis(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ApigatewayRestApis(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigateway_rest_apis",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"endpoint_configuration_types": []interface{}{"REGIONAL"},
					"api_key_source":               "HEADER",
					"tags": map[string]interface{}{
						"TestId": res.Suffix,
						"Type":   "integration_test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_apigateway_rest_api_deployments",
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"description": "test description",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_authorizers",
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"auth_type":                        "custom",
							"authorizer_result_ttl_in_seconds": float64(500),
							"type":                             "TOKEN",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_resources",
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"path_part": "mydemoresource",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_models",
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"content_type": "application/json",
							"description":  "a JSON schema",
							"name":         "user",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_request_validators",
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"validate_request_parameters": true,
							"validate_request_body":       true,
							"name":                        "example",
						},
					}},
				},
				{
					Name:           "aws_apigateway_rest_api_documentation_parts",
					ForeignKeyName: "rest_api_id",
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
					ForeignKeyName: "rest_api_id",
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
					ForeignKeyName: "rest_api_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"tracing_enabled": false,
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
