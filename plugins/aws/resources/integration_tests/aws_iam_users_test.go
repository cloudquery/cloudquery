package integration_tests

//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/Masterminds/squirrel"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationIamUsers(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.IamUsers(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_iam_users",
//			ExpectedValues: []providertest.ExpectedValue{{
//				Count: 1,
//				//Data: map[string]interface{}{
//				//	"tracing_config_mode": "PassThrough",
//				//},
//			}},
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Eq{"user_name": fmt.Sprintf("user%s%s", res.Prefix, res.Suffix)})
//			},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_iam_user_policies",
//					ForeignKeyName: "user_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						//Data: map[string]interface{}{
//						//	"description": "a sample description",
//						//},
//					}},
//				},
//				{
//					Name:           "aws_iam_user_access_keys",
//					ForeignKeyName: "user_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						//Data: map[string]interface{}{
//						//	"description": "a sample description",
//						//},
//					}},
//				},
//				{
//					Name:           "aws_iam_user_attached_policies",
//					ForeignKeyName: "user_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						//Data: map[string]interface{}{
//						//	"description": "a sample description",
//						//},
//					}},
//				},
//				{
//					Name:           "aws_iam_user_groups",
//					ForeignKeyName: "user_id",
//					ExpectedValues: []providertest.ExpectedValue{{
//						Count: 1,
//						//Data: map[string]interface{}{
//						//	"description": "a sample description",
//						//},
//					}},
//				},
//			},
//		}
//	})
//}
