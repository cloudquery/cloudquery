package integration_tests

//
//import (
//	"github.com/cloudquery/cq-provider-aws/resources/integration_tests"
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationApigatewayv2ApisTest(t *testing.T) {
//	integration_tests.awsTestIntegrationHelper(t, resources.Apigatewayv2Apis(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_apigatewayv2_apis",
//			ExpectedValues: []providertest.ExpectedValue{{
//				Count: 1,
//				Data: map[string]interface{}{
//					"protocol_type":                        "WEBSOCKET",
//					"api_gateway_managed":                  false,
//					"cors_configuration_allow_credentials": false,
//					"disable_execute_api_endpoint":         false,
//					"disable_schema_validation":            false,
//				},
//			}},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_apigatewayv2_api_authorizers",
//					ForeignKeyName: "api_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"name":                    "example-authorizer",
//							"authorizer_type":         "REQUEST",
//							"enable_simple_responses": false,
//							"identity_source":         []interface{}{"route.request.header.Auth"},
//						},
//					}},
//				},
//				{
//					Name:           "aws_apigatewayv2_api_integrations",
//					ForeignKeyName: "api_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"connection_type": "INTERNET",
//						},
//					}},
//					Relations: []*providertest.ResourceIntegrationVerification{
//						{
//							Name:           "aws_apigatewayv2_api_integration_responses",
//							ForeignKeyName: "api_integration_id",
//							ExpectedValues: []providertest.ExpectedValue{{
//								Count: 1,
//								Data: map[string]interface{}{
//									"integration_response_key": "/200/",
//								},
//							}},
//						},
//					},
//				},
//			},
//		}
//	})
//}
