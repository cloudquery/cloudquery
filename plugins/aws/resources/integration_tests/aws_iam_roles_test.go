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
//func TestIntegrationIamRoles(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.IamRoles(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_iam_roles",
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					Data: map[string]interface{}{
//						"role_name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
//					},
//				},
//			},
//			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//				return sq.Where(squirrel.Eq{"role_name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix)})
//			},
//			Relations: []*providertest.ResourceIntegrationVerification{
//				{
//					Name:           "aws_iam_role_policies",
//					ForeignKeyName: "role_id",
//					ExpectedValues: []providertest.ExpectedValue{
//						{
//							Count: 1,
//							//Data:  map[string]interface{}{},
//						},
//					},
//				},
//			},
//		}
//	})
//}
