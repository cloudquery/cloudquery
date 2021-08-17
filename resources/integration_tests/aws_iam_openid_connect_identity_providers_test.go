package integration_tests

//
//import (
//	"testing"
//
//	"github.com/cloudquery/cq-provider-aws/resources"
//
//	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
//)
//
//func TestIntegrationIamOpenidConnectIdentityProviders(t *testing.T) {
//	awsTestIntegrationHelper(t, resources.IamOpenidConnectIdentityProviders(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
//		return providertest.ResourceIntegrationVerification{
//			Name: "aws_iam_openid_connect_identity_providers",
//			//Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
//			//	return sq.Where("group_name = ?", fmt.Sprintf("aws_iam_group%s%s", res.Prefix, res.Suffix))
//			//},
//			ExpectedValues: []providertest.ExpectedValue{
//				{
//					Count: 1,
//					//Data: map[string]interface{}{
//					//	"role_name": fmt.Sprintf("%s%s", res.Prefix, res.Suffix),
//					//},
//				},
//			},
//		}
//	})
//}
