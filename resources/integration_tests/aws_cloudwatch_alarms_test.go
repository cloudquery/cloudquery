package integration_tests

//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/Masterminds/squirrel"
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationDirectCloudwatchAlarms(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.CloudwatchAlarms(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_cloudwatch_alarms",
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where("alarm_name = ?", fmt.Sprintf("%s-%s", res.Prefix, res.Suffix))
//			},
//			ExpectedValues: []providertest.ExpectedValue{{
//				Count: 1,
//				Data: map[string]interface{}{
//					"actions_enabled": true,
//				},
//			}},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_cloudwatch_alarm_metrics",
//					ForeignKeyName: "alarm_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"metric_id": "m1",
//						},
//					}},
//				},
//				{
//					Name:           "aws_cloudwatch_alarm_metrics",
//					ForeignKeyName: "alarm_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"metric_id": "m2",
//						},
//					}},
//				},
//				{
//					Name:           "aws_cloudwatch_alarm_metrics",
//					ForeignKeyName: "alarm_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						Data: map[string]interface{}{
//							"metric_id": "e1",
//							"label":     "Error Rate",
//						},
//					}},
//				},
//			},
//		}
//	})
//}
