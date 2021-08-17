package integration_tests

//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/Masterminds/squirrel"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationLambdaFunctions(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.LambdaFunctions(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_lambda_functions",
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where("function_name = ?", fmt.Sprintf("test_function_%s%s", res.Prefix, res.Suffix))
//			},
//			ExpectedValues: []providertest.ExpectedValue{{
//				Count: 1,
//				Data: map[string]interface{}{
//					"tracing_config_mode": "PassThrough",
//				},
//			}},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_lambda_function_aliases",
//					ForeignKeyName: "function_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"description": "a sample description",
//						},
//					}},
//				},
//			},
//		}
//	})
//}
