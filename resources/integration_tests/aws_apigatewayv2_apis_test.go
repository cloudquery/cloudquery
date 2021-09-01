package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationApigatewayv2ApisTest(t *testing.T) {
	const modelSchema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "ExampleModel",
  "type": "object",
  "properties": {
    "id": { "type": "string" }
  }
}
`
	const apiFKName = "api_cq_id"
	awsTestIntegrationHelper(t, resources.Apigatewayv2Apis(), []string{"aws_apigatewayv2_apis.tf", "aws_lambda_functions.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_apigatewayv2_apis",
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name":                                 fmt.Sprintf("apigwv2-api-%s%s", res.Prefix, res.Suffix),
					"protocol_type":                        "WEBSOCKET",
					"api_gateway_managed":                  false,
					"cors_configuration_allow_credentials": false,
					"disable_execute_api_endpoint":         false,
					"disable_schema_validation":            false,
					"tags": map[string]interface{}{
						"Type":   "integration_test",
						"TestId": res.Suffix,
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_apigatewayv2_api_authorizers",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":                    fmt.Sprintf("apigwv2-authorizer-%s%s", res.Prefix, res.Suffix),
							"authorizer_type":         "REQUEST",
							"enable_simple_responses": false,
							"identity_source":         []interface{}{"route.request.header.Auth"},
						},
					}},
				},
				{
					Name:           "aws_apigatewayv2_api_deployments",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"description": fmt.Sprintf("apigwv2-dep-%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "aws_apigatewayv2_api_integrations",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"connection_type":    "INTERNET",
							"integration_type":   "HTTP_PROXY",
							"integration_method": "ANY",
							"integration_uri":    "https://example.com/{proxy}",
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "aws_apigatewayv2_api_integration_responses",
							ForeignKeyName: "api_integration_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"integration_response_key": "/200/",
								},
							}},
						},
					},
				},
				{
					Name:           "aws_apigatewayv2_api_models",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":         fmt.Sprintf("apigwv2model%s", res.Suffix),
							"content_type": "application/json",
							"schema":       modelSchema,
						},
					}},
				},
				{
					Name:           "aws_apigatewayv2_api_routes",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"route_key":           "GET /example/v1/test",
							"api_gateway_managed": false,
							"api_key_required":    false,
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "aws_apigatewayv2_api_route_responses",
							ForeignKeyName: "api_route_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"route_response_key":  "$default",
									"response_models":     nil,
									"response_parameters": nil,
								},
							}},
						},
					},
				},
				{
					Name:           "aws_apigatewayv2_api_stages",
					ForeignKeyName: apiFKName,
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"stage_name": fmt.Sprintf("apigwv2-stage-%s%s", res.Prefix, res.Suffix),
							"tags": map[string]interface{}{
								"Type":   "integration_test",
								"TestId": res.Suffix,
							},
						},
					}},
				},
			},
		}
	})
}
